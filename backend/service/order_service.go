package service

import (
	"errors"
	"time"

	"github.com/lingojack/taobao_clone/model"
	"gorm.io/gorm"
)

type OrderService struct {
	db *gorm.DB
}

func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{db: db}
}

type OrderItemInput struct {
	ProductID uint64 `json:"productId" validate:"required"`
	SkuID     uint64 `json:"skuId"`
	Quantity  int    `json:"quantity" validate:"required,min=1"`
}

type CreateOrderRequest struct {
	AddressID uint64          `json:"addressId" validate:"required"`
	CouponID  uint64          `json:"couponId"`
	Remark    string          `json:"remark"`
	Items     []OrderItemInput `json:"items" validate:"required,min=1"`
}

type OrderWithItems struct {
	model.Orders
	Items []model.OrderItems   `json:"items"`
}

func (s *OrderService) CreateOrder(userID uint64, req *CreateOrderRequest) (*model.Orders, error) {
	// 验证地址
	var address model.UserAddresses
	if err := s.db.Where("id = ? AND user_id = ?", req.AddressID, userID).First(&address).Error; err != nil {
		return nil, errors.New("收货地址不存在")
	}

	// 计算总金额
	var totalAmount float64
	var orderItems []model.OrderItems

	for _, item := range req.Items {
		var product model.Products
		if err := s.db.First(&product, item.ProductID).Error; err != nil {
			return nil, errors.New("商品不存在")
		}

		price := product.Price
		var skuSpecValues *string
		var skuId *uint64

		// 如果有 SKU，使用 SKU 价格
		if item.SkuID > 0 {
			var sku model.ProductSkus
			if err := s.db.First(&sku, item.SkuID).Error; err != nil {
				return nil, errors.New("商品规格不存在")
			}
			price = sku.Price
			skuSpecValues = sku.SpecValues
			skuId = sku.Id
		}

		amount := price * float64(item.Quantity)
		totalAmount += amount

		orderItems = append(orderItems, model.OrderItems{
			ProductId:     item.ProductID,
			ProductName:   product.Name,
			ProductImage:  product.MainImage,
			SkuId:         skuId,
			SkuSpecValues: skuSpecValues,
			Price:         price,
			Quantity:      item.Quantity,
			TotalAmount:   amount,
		})
	}

	// 应用优惠券
	discount := float64(0)
	if req.CouponID > 0 {
		var uc model.UserCoupons
		if err := s.db.Where("id = ? AND user_id = ? AND status = ?", req.CouponID, userID, "unused").First(&uc).Error; err != nil {
			return nil, errors.New("优惠券不可用")
		}
		var coupon model.Coupons
		if err := s.db.First(&coupon, uc.CouponId).Error; err == nil {
			if totalAmount >= coupon.MinSpend {
				discount = coupon.Discount
				s.db.Model(&uc).Updates(map[string]interface{}{"status": "used"})
			}
		}
	}

	payAmount := totalAmount - discount
	if payAmount < 0 {
		payAmount = 0
	}

	orderNo := generateOrderNo()
	remark := req.Remark
	order := &model.Orders{
		OrderNo:         orderNo,
		UserId:          userID,
		TotalAmount:     totalAmount,
		PayAmount:       payAmount,
		Status:          "pending",
		ReceiverName:    address.ReceiverName,
		ReceiverPhone:   address.ReceiverPhone,
		ReceiverAddress: address.Province + address.City + address.District + address.DetailAddress,
		Remark:          &remark,
	}

	// 事务创建订单
	tx := s.db.Begin()
	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for i := range orderItems {
		orderItems[i].OrderId = *order.Id
	}
	if err := tx.Create(&orderItems).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 删除购物车中已下单的商品
	tx.Where("user_id = ? AND selected = 1", userID).Delete(&model.CartItems{})

	tx.Commit()
	return order, nil
}

func (s *OrderService) GetOrders(userID uint64, status string, page, pageSize int) ([]OrderWithItems, int64, error) {
	var total int64
	query := s.db.Model(&model.Orders{}).Where("user_id = ?", userID)
	if status != "" && status != "all" {
		query = query.Where("status = ?", status)
	}
	query.Count(&total)

	var orders []model.Orders
	offset := (page - 1) * pageSize
	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	result := make([]OrderWithItems, 0, len(orders))
	for _, order := range orders {
		item := OrderWithItems{Orders: order}

		var items []model.OrderItems
		s.db.Where("order_id = ?", *order.Id).Find(&items)
		item.Items = items

		result = append(result, item)
	}

	return result, total, nil
}

func (s *OrderService) GetOrderDetail(userID uint64, orderID uint64) (*OrderWithItems, error) {
	var order model.Orders
	if err := s.db.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		return nil, errors.New("订单不存在")
	}

	result := &OrderWithItems{Orders: order}

	var items []model.OrderItems
	s.db.Where("order_id = ?", *order.Id).Find(&items)
	result.Items = items

	return result, nil
}

func (s *OrderService) CancelOrder(userID uint64, orderID uint64) error {
	var order model.Orders
	if err := s.db.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		return errors.New("订单不存在")
	}

	if order.Status != "pending" {
		return errors.New("只能取消待付款订单")
	}

	return s.db.Model(&order).Update("status", "cancelled").Error
}

func (s *OrderService) PayOrder(userID uint64, orderID uint64) error {
	var order model.Orders
	if err := s.db.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		return errors.New("订单不存在")
	}

	if order.Status != "pending" {
		return errors.New("订单状态不可支付")
	}

	now := time.Now()
	return s.db.Model(&order).Updates(map[string]interface{}{
		"status":   "paid",
		"pay_time": now,
	}).Error
}

func (s *OrderService) ConfirmReceive(userID uint64, orderID uint64) error {
	var order model.Orders
	if err := s.db.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		return errors.New("订单不存在")
	}

	if order.Status != "shipped" {
		return errors.New("只能确认已发货的订单")
	}

	now := time.Now()
	return s.db.Model(&order).Updates(map[string]interface{}{
		"status":       "completed",
		"receive_time": now,
	}).Error
}

func (s *OrderService) DeleteOrder(userID uint64, orderID uint64) error {
	var order model.Orders
	if err := s.db.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		return errors.New("订单不存在")
	}

	if order.Status != "completed" && order.Status != "cancelled" {
		return errors.New("只能删除已完成或已取消的订单")
	}

	return s.db.Delete(&order).Error
}

func generateOrderNo() string {
	return "ORD" + time.Now().Format("20060102150405")
}
