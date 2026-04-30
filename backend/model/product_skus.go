package model

import (
	"encoding/json"
	"time"
)

// ProductSkus 商品SKU表
type ProductSkus struct {
	Id         *uint64    `gorm:"column:id;type:bigint(20) UNSIGNED;primaryKey;autoIncrement;comment:;" json:"id"`
	ProductId  uint64     `gorm:"column:product_id;type:bigint(20) UNSIGNED;comment:商品ID;not null" json:"product_id"`
	SkuCode    string     `gorm:"column:sku_code;type:varchar(100);comment:SKU编码;not null" json:"sku_code"`
	SpecValues *string    `gorm:"column:spec_values;type:json;comment:规格值（颜色、尺寸等）;" json:"spec_values"`
	Price      float64    `gorm:"column:price;type:decimal(10,2);comment:价格;not null" json:"price"`
	Stock      *int       `gorm:"column:stock;type:int(11);default:;comment:库存;" json:"stock"`
	Image      *string    `gorm:"column:image;type:varchar(500);comment:SKU图片;" json:"image"`
	Status     *int8      `gorm:"column:status;type:tinyint(4);default:;comment:状态：0-禁用 1-正常;" json:"status"`
	CreatedAt  *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:;" json:"created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:;" json:"updated_at"`
	DeletedAt  *time.Time `gorm:"column:deleted_at;type:timestamp;default:;comment:;" json:"deleted_at"`
}

// TableName 返回表名
func (t *ProductSkus) TableName() string {
	return "product_skus"
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *ProductSkus) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *ProductSkus) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// ProductSkusBuilder 用于构建 ProductSkus 实例的 Builder
type ProductSkusBuilder struct {
	instance *ProductSkus
}

// NewProductSkusBuilder 创建一个新的 ProductSkusBuilder 实例
// 返回:
//   - *ProductSkusBuilder: Builder 实例，用于链式调用
func NewProductSkusBuilder() *ProductSkusBuilder {
	return &ProductSkusBuilder{
		instance: &ProductSkus{},
	}
}

// WithProductId 设置 product_id 字段
// 参数:
//   - productId: 商品ID
//
// 返回:
//   - *ProductSkusBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusBuilder) WithProductId(productId uint64) *ProductSkusBuilder {
	b.instance.ProductId = productId
	return b
}

// WithSkuCode 设置 sku_code 字段
// 参数:
//   - skuCode: SKU编码
//
// 返回:
//   - *ProductSkusBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusBuilder) WithSkuCode(skuCode string) *ProductSkusBuilder {
	b.instance.SkuCode = skuCode
	return b
}

// WithSpecValues 设置 spec_values 字段
// 参数:
//   - specValues: 规格值（颜色、尺寸等）
//
// 返回:
//   - *ProductSkusBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusBuilder) WithSpecValues(specValues *string) *ProductSkusBuilder {
	b.instance.SpecValues = specValues
	return b
}

// WithSpecValuesValue 设置 spec_values 字段（便捷方法，自动转换为指针）
// 参数:
//   - specValues: 规格值（颜色、尺寸等）
//
// 返回:
//   - *ProductSkusBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusBuilder) WithSpecValuesValue(specValues string) *ProductSkusBuilder {
	b.instance.SpecValues = &specValues
	return b
}

// WithPrice 设置 price 字段
// 参数:
//   - price: 价格
//
// 返回:
//   - *ProductSkusBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusBuilder) WithPrice(price float64) *ProductSkusBuilder {
	b.instance.Price = price
	return b
}

// WithStock 设置 stock 字段
// 参数:
//   - stock: 库存
//
// 返回:
//   - *ProductSkusBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusBuilder) WithStock(stock *int) *ProductSkusBuilder {
	b.instance.Stock = stock
	return b
}

// WithStockValue 设置 stock 字段（便捷方法，自动转换为指针）
// 参数:
//   - stock: 库存
//
// 返回:
//   - *ProductSkusBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusBuilder) WithStockValue(stock int) *ProductSkusBuilder {
	b.instance.Stock = &stock
	return b
}

// WithImage 设置 image 字段
// 参数:
//   - image: SKU图片
//
// 返回:
//   - *ProductSkusBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusBuilder) WithImage(image *string) *ProductSkusBuilder {
	b.instance.Image = image
	return b
}

// WithImageValue 设置 image 字段（便捷方法，自动转换为指针）
// 参数:
//   - image: SKU图片
//
// 返回:
//   - *ProductSkusBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusBuilder) WithImageValue(image string) *ProductSkusBuilder {
	b.instance.Image = &image
	return b
}

// WithStatus 设置 status 字段
// 参数:
//   - status: 状态：0-禁用 1-正常
//
// 返回:
//   - *ProductSkusBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusBuilder) WithStatus(status *int8) *ProductSkusBuilder {
	b.instance.Status = status
	return b
}

// WithStatusValue 设置 status 字段（便捷方法，自动转换为指针）
// 参数:
//   - status: 状态：0-禁用 1-正常
//
// 返回:
//   - *ProductSkusBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusBuilder) WithStatusValue(status int8) *ProductSkusBuilder {
	b.instance.Status = &status
	return b
}

// WithCreatedAt 设置 created_at 字段
// 参数:
//   - createdAt:
//
// 返回:
//   - *ProductSkusBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusBuilder) WithCreatedAt(createdAt *time.Time) *ProductSkusBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

// WithCreatedAtValue 设置 created_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - createdAt:
//
// 返回:
//   - *ProductSkusBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusBuilder) WithCreatedAtValue(createdAt time.Time) *ProductSkusBuilder {
	b.instance.CreatedAt = &createdAt
	return b
}

// WithUpdatedAt 设置 updated_at 字段
// 参数:
//   - updatedAt:
//
// 返回:
//   - *ProductSkusBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusBuilder) WithUpdatedAt(updatedAt *time.Time) *ProductSkusBuilder {
	b.instance.UpdatedAt = updatedAt
	return b
}

// WithUpdatedAtValue 设置 updated_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - updatedAt:
//
// 返回:
//   - *ProductSkusBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusBuilder) WithUpdatedAtValue(updatedAt time.Time) *ProductSkusBuilder {
	b.instance.UpdatedAt = &updatedAt
	return b
}

// WithDeletedAt 设置 deleted_at 字段
// 参数:
//   - deletedAt:
//
// 返回:
//   - *ProductSkusBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusBuilder) WithDeletedAt(deletedAt *time.Time) *ProductSkusBuilder {
	b.instance.DeletedAt = deletedAt
	return b
}

// WithDeletedAtValue 设置 deleted_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - deletedAt:
//
// 返回:
//   - *ProductSkusBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusBuilder) WithDeletedAtValue(deletedAt time.Time) *ProductSkusBuilder {
	b.instance.DeletedAt = &deletedAt
	return b
}

// Build 构建并返回 ProductSkus 实例
// 返回:
//   - *ProductSkus: 构建完成的实例
func (b *ProductSkusBuilder) Build() *ProductSkus {
	return b.instance
}
