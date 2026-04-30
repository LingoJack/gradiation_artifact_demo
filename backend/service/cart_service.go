package service

import (
	"errors"

	"github.com/lingojack/taobao_clone/model"
	"gorm.io/gorm"
)

type CartService struct {
	db *gorm.DB
}

func NewCartService(db *gorm.DB) *CartService {
	return &CartService{db: db}
}

type ProductBrief struct {
	ID        uint64  `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	MainImage string  `json:"mainImage"`
	Stock     int     `json:"stock"`
}

type SkuBrief struct {
	ID         uint64  `json:"id"`
	Price      float64 `json:"price"`
	Stock      int     `json:"stock"`
	SpecValues string  `json:"specValues"`
	Image      string  `json:"image"`
}

type CartItemWithProduct struct {
	model.CartItems
	Product *ProductBrief `json:"product"`
	Sku     *SkuBrief     `json:"sku,omitempty"`
}

func (s *CartService) GetCart(userID uint64) ([]CartItemWithProduct, error) {
	var items []model.CartItems
	if err := s.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}

	result := make([]CartItemWithProduct, 0, len(items))
	for _, item := range items {
		cartItem := CartItemWithProduct{CartItems: item}

		// 加载商品信息
		var product model.Products
		if err := s.db.Select("id, name, price, main_image, stock").First(&product, item.ProductId).Error; err == nil {
			mainImage := ""
			if product.MainImage != nil {
				mainImage = *product.MainImage
			}
			stock := 0
			if product.Stock != nil {
				stock = *product.Stock
			}
			cartItem.Product = &ProductBrief{
				ID:        *product.Id,
				Name:      product.Name,
				Price:     product.Price,
				MainImage: mainImage,
				Stock:     stock,
			}
		}

		// 加载 SKU 信息
		if item.SkuId != nil && *item.SkuId > 0 {
			var sku model.ProductSkus
			if err := s.db.First(&sku, *item.SkuId).Error; err == nil {
				specValues := ""
				if sku.SpecValues != nil {
					specValues = *sku.SpecValues
				}
				image := ""
				if sku.Image != nil {
					image = *sku.Image
				}
				stock := 0
				if sku.Stock != nil {
					stock = *sku.Stock
				}
				cartItem.Sku = &SkuBrief{
					ID:         *sku.Id,
					Price:      sku.Price,
					Stock:      stock,
					SpecValues: specValues,
					Image:      image,
				}
			}
		}

		result = append(result, cartItem)
	}

	return result, nil
}

type AddCartRequest struct {
	ProductID uint64 `json:"productId" validate:"required"`
	SkuID     uint64 `json:"skuId"`
	Quantity  int    `json:"quantity" validate:"required,min=1"`
}

func (s *CartService) AddItem(userID uint64, req *AddCartRequest) (*model.CartItems, error) {
	// 检查商品是否存在
	var product model.Products
	if err := s.db.First(&product, req.ProductID).Error; err != nil {
		return nil, errors.New("商品不存在")
	}

	// 检查是否已在购物车中
	var existing model.CartItems
	query := s.db.Where("user_id = ? AND product_id = ?", userID, req.ProductID)
	if req.SkuID > 0 {
		query = query.Where("sku_id = ?", req.SkuID)
	} else {
		query = query.Where("sku_id IS NULL OR sku_id = 0")
	}

	if err := query.First(&existing).Error; err == nil {
		// 已存在，更新数量
		s.db.Model(&existing).Update("quantity", existing.Quantity+req.Quantity)
		s.db.First(&existing, *existing.Id)
		return &existing, nil
	}

	// 新建购物车项
	selected := int8(1)
	skuId := req.SkuID
	cartItem := &model.CartItems{
		UserId:    userID,
		ProductId: req.ProductID,
		SkuId:     &skuId,
		Quantity:  req.Quantity,
		Selected:  &selected,
	}

	if err := s.db.Create(cartItem).Error; err != nil {
		return nil, err
	}

	return cartItem, nil
}

func (s *CartService) UpdateQuantity(userID uint64, itemID uint64, quantity int) (*model.CartItems, error) {
	var item model.CartItems
	if err := s.db.Where("id = ? AND user_id = ?", itemID, userID).First(&item).Error; err != nil {
		return nil, errors.New("购物车项不存在")
	}

	if quantity <= 0 {
		return nil, errors.New("数量必须大于0")
	}

	if err := s.db.Model(&item).Update("quantity", quantity).Error; err != nil {
		return nil, err
	}

	s.db.First(&item, *item.Id)
	return &item, nil
}

func (s *CartService) UpdateSelected(userID uint64, itemIDs []uint64, selected bool) error {
	sel := int8(0)
	if selected {
		sel = 1
	}
	return s.db.Model(&model.CartItems{}).
		Where("id IN ? AND user_id = ?", itemIDs, userID).
		Update("selected", sel).Error
}

func (s *CartService) SelectAll(userID uint64, selected bool) error {
	sel := int8(0)
	if selected {
		sel = 1
	}
	return s.db.Model(&model.CartItems{}).
		Where("user_id = ?", userID).
		Update("selected", sel).Error
}

func (s *CartService) RemoveItem(userID uint64, itemID uint64) error {
	result := s.db.Where("id = ? AND user_id = ?", itemID, userID).Delete(&model.CartItems{})
	if result.RowsAffected == 0 {
		return errors.New("购物车项不存在")
	}
	return nil
}

func (s *CartService) ClearCart(userID uint64) error {
	return s.db.Where("user_id = ?", userID).Delete(&model.CartItems{}).Error
}

func (s *CartService) GetSelectedItems(userID uint64) ([]CartItemWithProduct, error) {
	var items []model.CartItems
	if err := s.db.Where("user_id = ? AND selected = 1", userID).Find(&items).Error; err != nil {
		return nil, err
	}

	result := make([]CartItemWithProduct, 0, len(items))
	for _, item := range items {
		cartItem := CartItemWithProduct{CartItems: item}

		var product model.Products
		if err := s.db.First(&product, item.ProductId).Error; err == nil {
			mainImage := ""
			if product.MainImage != nil {
				mainImage = *product.MainImage
			}
			stock := 0
			if product.Stock != nil {
				stock = *product.Stock
			}
			cartItem.Product = &ProductBrief{
				ID:        *product.Id,
				Name:      product.Name,
				Price:     product.Price,
				MainImage: mainImage,
				Stock:     stock,
			}
		}

		if item.SkuId != nil && *item.SkuId > 0 {
			var sku model.ProductSkus
			if err := s.db.First(&sku, *item.SkuId).Error; err == nil {
				specValues := ""
				if sku.SpecValues != nil {
					specValues = *sku.SpecValues
				}
				image := ""
				if sku.Image != nil {
					image = *sku.Image
				}
				stock := 0
				if sku.Stock != nil {
					stock = *sku.Stock
				}
				cartItem.Sku = &SkuBrief{
					ID:         *sku.Id,
					Price:      sku.Price,
					Stock:      stock,
					SpecValues: specValues,
					Image:      image,
				}
			}
		}

		result = append(result, cartItem)
	}

	return result, nil
}
