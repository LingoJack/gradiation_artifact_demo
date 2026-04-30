package model

import (
	"encoding/json"
	"time"
)

// OrderItems 订单项表
type OrderItems struct {
	Id            *uint64    `gorm:"column:id;type:bigint(20) UNSIGNED;primaryKey;autoIncrement;comment:;" json:"id"`
	OrderId       uint64     `gorm:"column:order_id;type:bigint(20) UNSIGNED;comment:订单ID;not null" json:"order_id"`
	ProductId     uint64     `gorm:"column:product_id;type:bigint(20) UNSIGNED;comment:商品ID;not null" json:"product_id"`
	SkuId         *uint64    `gorm:"column:sku_id;type:bigint(20) UNSIGNED;comment:SKU ID;" json:"sku_id"`
	ProductName   string     `gorm:"column:product_name;type:varchar(200);comment:商品名称;not null" json:"product_name"`
	SkuSpecValues *string    `gorm:"column:sku_spec_values;type:json;comment:SKU规格值;" json:"sku_spec_values"`
	ProductImage  *string    `gorm:"column:product_image;type:varchar(500);comment:商品图片;" json:"product_image"`
	Price         float64    `gorm:"column:price;type:decimal(10,2);comment:单价;not null" json:"price"`
	Quantity      int        `gorm:"column:quantity;type:int(11);comment:数量;not null" json:"quantity"`
	TotalAmount   float64    `gorm:"column:total_amount;type:decimal(10,2);comment:小计金额;not null" json:"total_amount"`
	CreatedAt     *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:;" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:;" json:"updated_at"`
}

// TableName 返回表名
func (t *OrderItems) TableName() string {
	return "order_items"
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *OrderItems) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *OrderItems) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// OrderItemsBuilder 用于构建 OrderItems 实例的 Builder
type OrderItemsBuilder struct {
	instance *OrderItems
}

// NewOrderItemsBuilder 创建一个新的 OrderItemsBuilder 实例
// 返回:
//   - *OrderItemsBuilder: Builder 实例，用于链式调用
func NewOrderItemsBuilder() *OrderItemsBuilder {
	return &OrderItemsBuilder{
		instance: &OrderItems{},
	}
}

// WithOrderId 设置 order_id 字段
// 参数:
//   - orderId: 订单ID
//
// 返回:
//   - *OrderItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsBuilder) WithOrderId(orderId uint64) *OrderItemsBuilder {
	b.instance.OrderId = orderId
	return b
}

// WithProductId 设置 product_id 字段
// 参数:
//   - productId: 商品ID
//
// 返回:
//   - *OrderItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsBuilder) WithProductId(productId uint64) *OrderItemsBuilder {
	b.instance.ProductId = productId
	return b
}

// WithSkuId 设置 sku_id 字段
// 参数:
//   - skuId: SKU ID
//
// 返回:
//   - *OrderItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsBuilder) WithSkuId(skuId *uint64) *OrderItemsBuilder {
	b.instance.SkuId = skuId
	return b
}

// WithSkuIdValue 设置 sku_id 字段（便捷方法，自动转换为指针）
// 参数:
//   - skuId: SKU ID
//
// 返回:
//   - *OrderItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsBuilder) WithSkuIdValue(skuId uint64) *OrderItemsBuilder {
	b.instance.SkuId = &skuId
	return b
}

// WithProductName 设置 product_name 字段
// 参数:
//   - productName: 商品名称
//
// 返回:
//   - *OrderItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsBuilder) WithProductName(productName string) *OrderItemsBuilder {
	b.instance.ProductName = productName
	return b
}

// WithSkuSpecValues 设置 sku_spec_values 字段
// 参数:
//   - skuSpecValues: SKU规格值
//
// 返回:
//   - *OrderItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsBuilder) WithSkuSpecValues(skuSpecValues *string) *OrderItemsBuilder {
	b.instance.SkuSpecValues = skuSpecValues
	return b
}

// WithSkuSpecValuesValue 设置 sku_spec_values 字段（便捷方法，自动转换为指针）
// 参数:
//   - skuSpecValues: SKU规格值
//
// 返回:
//   - *OrderItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsBuilder) WithSkuSpecValuesValue(skuSpecValues string) *OrderItemsBuilder {
	b.instance.SkuSpecValues = &skuSpecValues
	return b
}

// WithProductImage 设置 product_image 字段
// 参数:
//   - productImage: 商品图片
//
// 返回:
//   - *OrderItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsBuilder) WithProductImage(productImage *string) *OrderItemsBuilder {
	b.instance.ProductImage = productImage
	return b
}

// WithProductImageValue 设置 product_image 字段（便捷方法，自动转换为指针）
// 参数:
//   - productImage: 商品图片
//
// 返回:
//   - *OrderItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsBuilder) WithProductImageValue(productImage string) *OrderItemsBuilder {
	b.instance.ProductImage = &productImage
	return b
}

// WithPrice 设置 price 字段
// 参数:
//   - price: 单价
//
// 返回:
//   - *OrderItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsBuilder) WithPrice(price float64) *OrderItemsBuilder {
	b.instance.Price = price
	return b
}

// WithQuantity 设置 quantity 字段
// 参数:
//   - quantity: 数量
//
// 返回:
//   - *OrderItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsBuilder) WithQuantity(quantity int) *OrderItemsBuilder {
	b.instance.Quantity = quantity
	return b
}

// WithTotalAmount 设置 total_amount 字段
// 参数:
//   - totalAmount: 小计金额
//
// 返回:
//   - *OrderItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsBuilder) WithTotalAmount(totalAmount float64) *OrderItemsBuilder {
	b.instance.TotalAmount = totalAmount
	return b
}

// WithCreatedAt 设置 created_at 字段
// 参数:
//   - createdAt:
//
// 返回:
//   - *OrderItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsBuilder) WithCreatedAt(createdAt *time.Time) *OrderItemsBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

// WithCreatedAtValue 设置 created_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - createdAt:
//
// 返回:
//   - *OrderItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsBuilder) WithCreatedAtValue(createdAt time.Time) *OrderItemsBuilder {
	b.instance.CreatedAt = &createdAt
	return b
}

// WithUpdatedAt 设置 updated_at 字段
// 参数:
//   - updatedAt:
//
// 返回:
//   - *OrderItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsBuilder) WithUpdatedAt(updatedAt *time.Time) *OrderItemsBuilder {
	b.instance.UpdatedAt = updatedAt
	return b
}

// WithUpdatedAtValue 设置 updated_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - updatedAt:
//
// 返回:
//   - *OrderItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsBuilder) WithUpdatedAtValue(updatedAt time.Time) *OrderItemsBuilder {
	b.instance.UpdatedAt = &updatedAt
	return b
}

// Build 构建并返回 OrderItems 实例
// 返回:
//   - *OrderItems: 构建完成的实例
func (b *OrderItemsBuilder) Build() *OrderItems {
	return b.instance
}
