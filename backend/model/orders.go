package model

import (
	"encoding/json"
	"time"
)

// Orders 订单表
type Orders struct {
	Id              *uint64    `gorm:"column:id;type:bigint(20) UNSIGNED;primaryKey;autoIncrement;comment:;" json:"id"`
	OrderNo         string     `gorm:"column:order_no;type:varchar(50);comment:订单号;not null" json:"order_no"`
	UserId          uint64     `gorm:"column:user_id;type:bigint(20) UNSIGNED;comment:用户ID;not null" json:"user_id"`
	TotalAmount     float64    `gorm:"column:total_amount;type:decimal(10,2);comment:订单总金额;not null" json:"total_amount"`
	PayAmount       float64    `gorm:"column:pay_amount;type:decimal(10,2);comment:实付金额;not null" json:"pay_amount"`
	Status          string     `gorm:"column:status;type:varchar(20);comment:订单状态;not null" json:"status"`
	ReceiverName    string     `gorm:"column:receiver_name;type:varchar(50);comment:收货人姓名;not null" json:"receiver_name"`
	ReceiverPhone   string     `gorm:"column:receiver_phone;type:varchar(20);comment:收货人电话;not null" json:"receiver_phone"`
	ReceiverAddress string     `gorm:"column:receiver_address;type:varchar(300);comment:收货地址;not null" json:"receiver_address"`
	Remark          *string    `gorm:"column:remark;type:varchar(500);comment:备注;" json:"remark"`
	PayTime         *time.Time `gorm:"column:pay_time;type:timestamp;comment:支付时间;" json:"pay_time"`
	DeliveryTime    *time.Time `gorm:"column:delivery_time;type:timestamp;comment:发货时间;" json:"delivery_time"`
	ReceiveTime     *time.Time `gorm:"column:receive_time;type:timestamp;comment:收货时间;" json:"receive_time"`
	CreatedAt       *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:;" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:;" json:"updated_at"`
	DeletedAt       *time.Time `gorm:"column:deleted_at;type:timestamp;default:;comment:;" json:"deleted_at"`
}

// TableName 返回表名
func (t *Orders) TableName() string {
	return "orders"
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *Orders) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *Orders) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// OrdersBuilder 用于构建 Orders 实例的 Builder
type OrdersBuilder struct {
	instance *Orders
}

// NewOrdersBuilder 创建一个新的 OrdersBuilder 实例
// 返回:
//   - *OrdersBuilder: Builder 实例，用于链式调用
func NewOrdersBuilder() *OrdersBuilder {
	return &OrdersBuilder{
		instance: &Orders{},
	}
}

// WithOrderNo 设置 order_no 字段
// 参数:
//   - orderNo: 订单号
//
// 返回:
//   - *OrdersBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersBuilder) WithOrderNo(orderNo string) *OrdersBuilder {
	b.instance.OrderNo = orderNo
	return b
}

// WithUserId 设置 user_id 字段
// 参数:
//   - userId: 用户ID
//
// 返回:
//   - *OrdersBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersBuilder) WithUserId(userId uint64) *OrdersBuilder {
	b.instance.UserId = userId
	return b
}

// WithTotalAmount 设置 total_amount 字段
// 参数:
//   - totalAmount: 订单总金额
//
// 返回:
//   - *OrdersBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersBuilder) WithTotalAmount(totalAmount float64) *OrdersBuilder {
	b.instance.TotalAmount = totalAmount
	return b
}

// WithPayAmount 设置 pay_amount 字段
// 参数:
//   - payAmount: 实付金额
//
// 返回:
//   - *OrdersBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersBuilder) WithPayAmount(payAmount float64) *OrdersBuilder {
	b.instance.PayAmount = payAmount
	return b
}

// WithStatus 设置 status 字段
// 参数:
//   - status: 订单状态
//
// 返回:
//   - *OrdersBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersBuilder) WithStatus(status string) *OrdersBuilder {
	b.instance.Status = status
	return b
}

// WithReceiverName 设置 receiver_name 字段
// 参数:
//   - receiverName: 收货人姓名
//
// 返回:
//   - *OrdersBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersBuilder) WithReceiverName(receiverName string) *OrdersBuilder {
	b.instance.ReceiverName = receiverName
	return b
}

// WithReceiverPhone 设置 receiver_phone 字段
// 参数:
//   - receiverPhone: 收货人电话
//
// 返回:
//   - *OrdersBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersBuilder) WithReceiverPhone(receiverPhone string) *OrdersBuilder {
	b.instance.ReceiverPhone = receiverPhone
	return b
}

// WithReceiverAddress 设置 receiver_address 字段
// 参数:
//   - receiverAddress: 收货地址
//
// 返回:
//   - *OrdersBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersBuilder) WithReceiverAddress(receiverAddress string) *OrdersBuilder {
	b.instance.ReceiverAddress = receiverAddress
	return b
}

// WithRemark 设置 remark 字段
// 参数:
//   - remark: 备注
//
// 返回:
//   - *OrdersBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersBuilder) WithRemark(remark *string) *OrdersBuilder {
	b.instance.Remark = remark
	return b
}

// WithRemarkValue 设置 remark 字段（便捷方法，自动转换为指针）
// 参数:
//   - remark: 备注
//
// 返回:
//   - *OrdersBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersBuilder) WithRemarkValue(remark string) *OrdersBuilder {
	b.instance.Remark = &remark
	return b
}

// WithPayTime 设置 pay_time 字段
// 参数:
//   - payTime: 支付时间
//
// 返回:
//   - *OrdersBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersBuilder) WithPayTime(payTime *time.Time) *OrdersBuilder {
	b.instance.PayTime = payTime
	return b
}

// WithPayTimeValue 设置 pay_time 字段（便捷方法，自动转换为指针）
// 参数:
//   - payTime: 支付时间
//
// 返回:
//   - *OrdersBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersBuilder) WithPayTimeValue(payTime time.Time) *OrdersBuilder {
	b.instance.PayTime = &payTime
	return b
}

// WithDeliveryTime 设置 delivery_time 字段
// 参数:
//   - deliveryTime: 发货时间
//
// 返回:
//   - *OrdersBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersBuilder) WithDeliveryTime(deliveryTime *time.Time) *OrdersBuilder {
	b.instance.DeliveryTime = deliveryTime
	return b
}

// WithDeliveryTimeValue 设置 delivery_time 字段（便捷方法，自动转换为指针）
// 参数:
//   - deliveryTime: 发货时间
//
// 返回:
//   - *OrdersBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersBuilder) WithDeliveryTimeValue(deliveryTime time.Time) *OrdersBuilder {
	b.instance.DeliveryTime = &deliveryTime
	return b
}

// WithReceiveTime 设置 receive_time 字段
// 参数:
//   - receiveTime: 收货时间
//
// 返回:
//   - *OrdersBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersBuilder) WithReceiveTime(receiveTime *time.Time) *OrdersBuilder {
	b.instance.ReceiveTime = receiveTime
	return b
}

// WithReceiveTimeValue 设置 receive_time 字段（便捷方法，自动转换为指针）
// 参数:
//   - receiveTime: 收货时间
//
// 返回:
//   - *OrdersBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersBuilder) WithReceiveTimeValue(receiveTime time.Time) *OrdersBuilder {
	b.instance.ReceiveTime = &receiveTime
	return b
}

// WithCreatedAt 设置 created_at 字段
// 参数:
//   - createdAt:
//
// 返回:
//   - *OrdersBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersBuilder) WithCreatedAt(createdAt *time.Time) *OrdersBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

// WithCreatedAtValue 设置 created_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - createdAt:
//
// 返回:
//   - *OrdersBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersBuilder) WithCreatedAtValue(createdAt time.Time) *OrdersBuilder {
	b.instance.CreatedAt = &createdAt
	return b
}

// WithUpdatedAt 设置 updated_at 字段
// 参数:
//   - updatedAt:
//
// 返回:
//   - *OrdersBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersBuilder) WithUpdatedAt(updatedAt *time.Time) *OrdersBuilder {
	b.instance.UpdatedAt = updatedAt
	return b
}

// WithUpdatedAtValue 设置 updated_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - updatedAt:
//
// 返回:
//   - *OrdersBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersBuilder) WithUpdatedAtValue(updatedAt time.Time) *OrdersBuilder {
	b.instance.UpdatedAt = &updatedAt
	return b
}

// WithDeletedAt 设置 deleted_at 字段
// 参数:
//   - deletedAt:
//
// 返回:
//   - *OrdersBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersBuilder) WithDeletedAt(deletedAt *time.Time) *OrdersBuilder {
	b.instance.DeletedAt = deletedAt
	return b
}

// WithDeletedAtValue 设置 deleted_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - deletedAt:
//
// 返回:
//   - *OrdersBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersBuilder) WithDeletedAtValue(deletedAt time.Time) *OrdersBuilder {
	b.instance.DeletedAt = &deletedAt
	return b
}

// Build 构建并返回 Orders 实例
// 返回:
//   - *Orders: 构建完成的实例
func (b *OrdersBuilder) Build() *Orders {
	return b.instance
}
