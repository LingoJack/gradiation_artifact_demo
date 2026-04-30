package model

import (
	"encoding/json"
	"time"
)

// Products 商品表
type Products struct {
	Id            *uint64    `gorm:"column:id;type:bigint(20) UNSIGNED;primaryKey;autoIncrement;comment:;" json:"id"`
	CategoryId    uint64     `gorm:"column:category_id;type:bigint(20) UNSIGNED;comment:分类ID;not null" json:"category_id"`
	Name          string     `gorm:"column:name;type:varchar(200);comment:商品名称;not null" json:"name"`
	Description   *string    `gorm:"column:description;type:text;comment:商品描述;" json:"description"`
	MainImage     *string    `gorm:"column:main_image;type:varchar(500);comment:主图URL;" json:"main_image"`
	Images        *string    `gorm:"column:images;type:json;comment:商品图片列表;" json:"images"`
	Price         float64    `gorm:"column:price;type:decimal(10,2);comment:价格;not null" json:"price"`
	OriginalPrice *float64   `gorm:"column:original_price;type:decimal(10,2);comment:原价;" json:"original_price"`
	Stock         *int       `gorm:"column:stock;type:int(11);default:;comment:库存;" json:"stock"`
	Sales         *int       `gorm:"column:sales;type:int(11);default:;comment:销量;" json:"sales"`
	Status        *int8      `gorm:"column:status;type:tinyint(4);default:;comment:状态：0-下架 1-上架;" json:"status"`
	SortOrder     *int       `gorm:"column:sort_order;type:int(11);default:;comment:排序;" json:"sort_order"`
	CreatedAt     *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:;" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:;" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at;type:timestamp;default:;comment:;" json:"deleted_at"`
}

// TableName 返回表名
func (t *Products) TableName() string {
	return "products"
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *Products) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *Products) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// ProductsBuilder 用于构建 Products 实例的 Builder
type ProductsBuilder struct {
	instance *Products
}

// NewProductsBuilder 创建一个新的 ProductsBuilder 实例
// 返回:
//   - *ProductsBuilder: Builder 实例，用于链式调用
func NewProductsBuilder() *ProductsBuilder {
	return &ProductsBuilder{
		instance: &Products{},
	}
}

// WithCategoryId 设置 category_id 字段
// 参数:
//   - categoryId: 分类ID
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithCategoryId(categoryId uint64) *ProductsBuilder {
	b.instance.CategoryId = categoryId
	return b
}

// WithName 设置 name 字段
// 参数:
//   - name: 商品名称
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithName(name string) *ProductsBuilder {
	b.instance.Name = name
	return b
}

// WithDescription 设置 description 字段
// 参数:
//   - description: 商品描述
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithDescription(description *string) *ProductsBuilder {
	b.instance.Description = description
	return b
}

// WithDescriptionValue 设置 description 字段（便捷方法，自动转换为指针）
// 参数:
//   - description: 商品描述
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithDescriptionValue(description string) *ProductsBuilder {
	b.instance.Description = &description
	return b
}

// WithMainImage 设置 main_image 字段
// 参数:
//   - mainImage: 主图URL
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithMainImage(mainImage *string) *ProductsBuilder {
	b.instance.MainImage = mainImage
	return b
}

// WithMainImageValue 设置 main_image 字段（便捷方法，自动转换为指针）
// 参数:
//   - mainImage: 主图URL
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithMainImageValue(mainImage string) *ProductsBuilder {
	b.instance.MainImage = &mainImage
	return b
}

// WithImages 设置 images 字段
// 参数:
//   - images: 商品图片列表
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithImages(images *string) *ProductsBuilder {
	b.instance.Images = images
	return b
}

// WithImagesValue 设置 images 字段（便捷方法，自动转换为指针）
// 参数:
//   - images: 商品图片列表
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithImagesValue(images string) *ProductsBuilder {
	b.instance.Images = &images
	return b
}

// WithPrice 设置 price 字段
// 参数:
//   - price: 价格
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithPrice(price float64) *ProductsBuilder {
	b.instance.Price = price
	return b
}

// WithOriginalPrice 设置 original_price 字段
// 参数:
//   - originalPrice: 原价
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithOriginalPrice(originalPrice *float64) *ProductsBuilder {
	b.instance.OriginalPrice = originalPrice
	return b
}

// WithOriginalPriceValue 设置 original_price 字段（便捷方法，自动转换为指针）
// 参数:
//   - originalPrice: 原价
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithOriginalPriceValue(originalPrice float64) *ProductsBuilder {
	b.instance.OriginalPrice = &originalPrice
	return b
}

// WithStock 设置 stock 字段
// 参数:
//   - stock: 库存
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithStock(stock *int) *ProductsBuilder {
	b.instance.Stock = stock
	return b
}

// WithStockValue 设置 stock 字段（便捷方法，自动转换为指针）
// 参数:
//   - stock: 库存
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithStockValue(stock int) *ProductsBuilder {
	b.instance.Stock = &stock
	return b
}

// WithSales 设置 sales 字段
// 参数:
//   - sales: 销量
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithSales(sales *int) *ProductsBuilder {
	b.instance.Sales = sales
	return b
}

// WithSalesValue 设置 sales 字段（便捷方法，自动转换为指针）
// 参数:
//   - sales: 销量
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithSalesValue(sales int) *ProductsBuilder {
	b.instance.Sales = &sales
	return b
}

// WithStatus 设置 status 字段
// 参数:
//   - status: 状态：0-下架 1-上架
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithStatus(status *int8) *ProductsBuilder {
	b.instance.Status = status
	return b
}

// WithStatusValue 设置 status 字段（便捷方法，自动转换为指针）
// 参数:
//   - status: 状态：0-下架 1-上架
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithStatusValue(status int8) *ProductsBuilder {
	b.instance.Status = &status
	return b
}

// WithSortOrder 设置 sort_order 字段
// 参数:
//   - sortOrder: 排序
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithSortOrder(sortOrder *int) *ProductsBuilder {
	b.instance.SortOrder = sortOrder
	return b
}

// WithSortOrderValue 设置 sort_order 字段（便捷方法，自动转换为指针）
// 参数:
//   - sortOrder: 排序
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithSortOrderValue(sortOrder int) *ProductsBuilder {
	b.instance.SortOrder = &sortOrder
	return b
}

// WithCreatedAt 设置 created_at 字段
// 参数:
//   - createdAt:
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithCreatedAt(createdAt *time.Time) *ProductsBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

// WithCreatedAtValue 设置 created_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - createdAt:
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithCreatedAtValue(createdAt time.Time) *ProductsBuilder {
	b.instance.CreatedAt = &createdAt
	return b
}

// WithUpdatedAt 设置 updated_at 字段
// 参数:
//   - updatedAt:
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithUpdatedAt(updatedAt *time.Time) *ProductsBuilder {
	b.instance.UpdatedAt = updatedAt
	return b
}

// WithUpdatedAtValue 设置 updated_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - updatedAt:
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithUpdatedAtValue(updatedAt time.Time) *ProductsBuilder {
	b.instance.UpdatedAt = &updatedAt
	return b
}

// WithDeletedAt 设置 deleted_at 字段
// 参数:
//   - deletedAt:
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithDeletedAt(deletedAt *time.Time) *ProductsBuilder {
	b.instance.DeletedAt = deletedAt
	return b
}

// WithDeletedAtValue 设置 deleted_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - deletedAt:
//
// 返回:
//   - *ProductsBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsBuilder) WithDeletedAtValue(deletedAt time.Time) *ProductsBuilder {
	b.instance.DeletedAt = &deletedAt
	return b
}

// Build 构建并返回 Products 实例
// 返回:
//   - *Products: 构建完成的实例
func (b *ProductsBuilder) Build() *Products {
	return b.instance
}
