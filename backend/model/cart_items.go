package model

import (
	"encoding/json"
	"time"
)

// CartItems 购物车表
type CartItems struct {
	Id        *uint64    `gorm:"column:id;type:bigint(20) UNSIGNED;primaryKey;autoIncrement;comment:;" json:"id"`
	UserId    uint64     `gorm:"column:user_id;type:bigint(20) UNSIGNED;comment:用户ID;not null" json:"user_id"`
	ProductId uint64     `gorm:"column:product_id;type:bigint(20) UNSIGNED;comment:商品ID;not null" json:"product_id"`
	SkuId     *uint64    `gorm:"column:sku_id;type:bigint(20) UNSIGNED;comment:SKU ID;" json:"sku_id"`
	Quantity  int        `gorm:"column:quantity;type:int(11);default:;comment:数量;not null" json:"quantity"`
	Selected  *int8      `gorm:"column:selected;type:tinyint(4);default:;comment:是否选中：0-否 1-是;" json:"selected"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:;" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:;" json:"updated_at"`
}

// TableName 返回表名
func (t *CartItems) TableName() string {
	return "cart_items"
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *CartItems) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *CartItems) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// CartItemsBuilder 用于构建 CartItems 实例的 Builder
type CartItemsBuilder struct {
	instance *CartItems
}

// NewCartItemsBuilder 创建一个新的 CartItemsBuilder 实例
// 返回:
//   - *CartItemsBuilder: Builder 实例，用于链式调用
func NewCartItemsBuilder() *CartItemsBuilder {
	return &CartItemsBuilder{
		instance: &CartItems{},
	}
}

// WithUserId 设置 user_id 字段
// 参数:
//   - userId: 用户ID
//
// 返回:
//   - *CartItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsBuilder) WithUserId(userId uint64) *CartItemsBuilder {
	b.instance.UserId = userId
	return b
}

// WithProductId 设置 product_id 字段
// 参数:
//   - productId: 商品ID
//
// 返回:
//   - *CartItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsBuilder) WithProductId(productId uint64) *CartItemsBuilder {
	b.instance.ProductId = productId
	return b
}

// WithSkuId 设置 sku_id 字段
// 参数:
//   - skuId: SKU ID
//
// 返回:
//   - *CartItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsBuilder) WithSkuId(skuId *uint64) *CartItemsBuilder {
	b.instance.SkuId = skuId
	return b
}

// WithSkuIdValue 设置 sku_id 字段（便捷方法，自动转换为指针）
// 参数:
//   - skuId: SKU ID
//
// 返回:
//   - *CartItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsBuilder) WithSkuIdValue(skuId uint64) *CartItemsBuilder {
	b.instance.SkuId = &skuId
	return b
}

// WithQuantity 设置 quantity 字段
// 参数:
//   - quantity: 数量
//
// 返回:
//   - *CartItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsBuilder) WithQuantity(quantity int) *CartItemsBuilder {
	b.instance.Quantity = quantity
	return b
}

// WithSelected 设置 selected 字段
// 参数:
//   - selected: 是否选中：0-否 1-是
//
// 返回:
//   - *CartItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsBuilder) WithSelected(selected *int8) *CartItemsBuilder {
	b.instance.Selected = selected
	return b
}

// WithSelectedValue 设置 selected 字段（便捷方法，自动转换为指针）
// 参数:
//   - selected: 是否选中：0-否 1-是
//
// 返回:
//   - *CartItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsBuilder) WithSelectedValue(selected int8) *CartItemsBuilder {
	b.instance.Selected = &selected
	return b
}

// WithCreatedAt 设置 created_at 字段
// 参数:
//   - createdAt:
//
// 返回:
//   - *CartItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsBuilder) WithCreatedAt(createdAt *time.Time) *CartItemsBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

// WithCreatedAtValue 设置 created_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - createdAt:
//
// 返回:
//   - *CartItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsBuilder) WithCreatedAtValue(createdAt time.Time) *CartItemsBuilder {
	b.instance.CreatedAt = &createdAt
	return b
}

// WithUpdatedAt 设置 updated_at 字段
// 参数:
//   - updatedAt:
//
// 返回:
//   - *CartItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsBuilder) WithUpdatedAt(updatedAt *time.Time) *CartItemsBuilder {
	b.instance.UpdatedAt = updatedAt
	return b
}

// WithUpdatedAtValue 设置 updated_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - updatedAt:
//
// 返回:
//   - *CartItemsBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsBuilder) WithUpdatedAtValue(updatedAt time.Time) *CartItemsBuilder {
	b.instance.UpdatedAt = &updatedAt
	return b
}

// Build 构建并返回 CartItems 实例
// 返回:
//   - *CartItems: 构建完成的实例
func (b *CartItemsBuilder) Build() *CartItems {
	return b.instance
}
