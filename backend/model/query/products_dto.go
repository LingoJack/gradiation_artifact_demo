package query

import (
	"encoding/json"
	"time"
)

// ProductsDto 商品表 数据传输对象
type ProductsDto struct {
	Id               *uint64    `json:"id"`                //
	CategoryId       uint64     `json:"category_id"`       // 分类ID
	Name             string     `json:"name"`              // 商品名称
	Description      *string    `json:"description"`       // 商品描述
	MainImage        *string    `json:"main_image"`        // 主图URL
	Images           *string    `json:"images"`            // 商品图片列表
	Price            float64    `json:"price"`             // 价格
	OriginalPrice    *float64   `json:"original_price"`    // 原价
	Stock            *int       `json:"stock"`             // 库存
	Sales            *int       `json:"sales"`             // 销量
	Status           *int8      `json:"status"`            // 状态：0-下架 1-上架
	SortOrder        *int       `json:"sort_order"`        // 排序
	CreatedAt        *time.Time `json:"created_at"`        //
	UpdatedAt        *time.Time `json:"updated_at"`        //
	DeletedAt        *time.Time `json:"deleted_at"`        //
	CategoryIdList   []uint64   `json:"category_id_list"`  // 分类ID IN 查询
	NameFuzzy        string     `json:"name_fuzzy"`        // 商品名称 模糊查询
	DescriptionFuzzy *string    `json:"description_fuzzy"` // 商品描述 模糊查询
	MainImageFuzzy   *string    `json:"main_image_fuzzy"`  // 主图URL 模糊查询
	ImagesFuzzy      *string    `json:"images_fuzzy"`      // 商品图片列表 模糊查询
	SalesList        []*int     `json:"sales_list"`        // 销量 IN 查询
	StatusList       []*int8    `json:"status_list"`       // 状态：0-下架 1-上架 IN 查询
	SortOrderList    []*int     `json:"sort_order_list"`   // 排序 IN 查询
	CreatedAtStart   *time.Time `json:"created_at_start"`  //  开始时间
	CreatedAtEnd     *time.Time `json:"created_at_end"`    //  结束时间
	UpdatedAtStart   *time.Time `json:"updated_at_start"`  //  开始时间
	UpdatedAtEnd     *time.Time `json:"updated_at_end"`    //  结束时间
	DeletedAtStart   *time.Time `json:"deleted_at_start"`  //  开始时间
	DeletedAtEnd     *time.Time `json:"deleted_at_end"`    //  结束时间
	OrderBy          string     `json:"order_by"`          // 排序字段
	PageOffset       int        `json:"page_offset"`       // 分页偏移量
	PageSize         int        `json:"page_size"`         // 每页数量
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *ProductsDto) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *ProductsDto) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// ProductsDtoBuilder 用于构建 ProductsDto 实例的 Builder
type ProductsDtoBuilder struct {
	instance *ProductsDto
}

// NewProductsDtoBuilder 创建一个新的 ProductsDtoBuilder 实例
// 返回:
//   - *ProductsDtoBuilder: Builder 实例，用于链式调用
func NewProductsDtoBuilder() *ProductsDtoBuilder {
	return &ProductsDtoBuilder{
		instance: &ProductsDto{},
	}
}

// WithCategoryId 设置 category_id 字段
// 参数:
//   - categoryId: 分类ID
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithCategoryId(categoryId uint64) *ProductsDtoBuilder {
	b.instance.CategoryId = categoryId
	return b
}

// WithName 设置 name 字段
// 参数:
//   - name: 商品名称
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithName(name string) *ProductsDtoBuilder {
	b.instance.Name = name
	return b
}

// WithDescription 设置 description 字段
// 参数:
//   - description: 商品描述
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithDescription(description *string) *ProductsDtoBuilder {
	b.instance.Description = description
	return b
}

// WithDescriptionValue 设置 description 字段（便捷方法，自动转换为指针）
// 参数:
//   - description: 商品描述
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithDescriptionValue(description string) *ProductsDtoBuilder {
	b.instance.Description = &description
	return b
}

// WithMainImage 设置 main_image 字段
// 参数:
//   - mainImage: 主图URL
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithMainImage(mainImage *string) *ProductsDtoBuilder {
	b.instance.MainImage = mainImage
	return b
}

// WithMainImageValue 设置 main_image 字段（便捷方法，自动转换为指针）
// 参数:
//   - mainImage: 主图URL
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithMainImageValue(mainImage string) *ProductsDtoBuilder {
	b.instance.MainImage = &mainImage
	return b
}

// WithImages 设置 images 字段
// 参数:
//   - images: 商品图片列表
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithImages(images *string) *ProductsDtoBuilder {
	b.instance.Images = images
	return b
}

// WithImagesValue 设置 images 字段（便捷方法，自动转换为指针）
// 参数:
//   - images: 商品图片列表
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithImagesValue(images string) *ProductsDtoBuilder {
	b.instance.Images = &images
	return b
}

// WithPrice 设置 price 字段
// 参数:
//   - price: 价格
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithPrice(price float64) *ProductsDtoBuilder {
	b.instance.Price = price
	return b
}

// WithOriginalPrice 设置 original_price 字段
// 参数:
//   - originalPrice: 原价
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithOriginalPrice(originalPrice *float64) *ProductsDtoBuilder {
	b.instance.OriginalPrice = originalPrice
	return b
}

// WithOriginalPriceValue 设置 original_price 字段（便捷方法，自动转换为指针）
// 参数:
//   - originalPrice: 原价
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithOriginalPriceValue(originalPrice float64) *ProductsDtoBuilder {
	b.instance.OriginalPrice = &originalPrice
	return b
}

// WithStock 设置 stock 字段
// 参数:
//   - stock: 库存
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithStock(stock *int) *ProductsDtoBuilder {
	b.instance.Stock = stock
	return b
}

// WithStockValue 设置 stock 字段（便捷方法，自动转换为指针）
// 参数:
//   - stock: 库存
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithStockValue(stock int) *ProductsDtoBuilder {
	b.instance.Stock = &stock
	return b
}

// WithSales 设置 sales 字段
// 参数:
//   - sales: 销量
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithSales(sales *int) *ProductsDtoBuilder {
	b.instance.Sales = sales
	return b
}

// WithSalesValue 设置 sales 字段（便捷方法，自动转换为指针）
// 参数:
//   - sales: 销量
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithSalesValue(sales int) *ProductsDtoBuilder {
	b.instance.Sales = &sales
	return b
}

// WithStatus 设置 status 字段
// 参数:
//   - status: 状态：0-下架 1-上架
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithStatus(status *int8) *ProductsDtoBuilder {
	b.instance.Status = status
	return b
}

// WithStatusValue 设置 status 字段（便捷方法，自动转换为指针）
// 参数:
//   - status: 状态：0-下架 1-上架
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithStatusValue(status int8) *ProductsDtoBuilder {
	b.instance.Status = &status
	return b
}

// WithSortOrder 设置 sort_order 字段
// 参数:
//   - sortOrder: 排序
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithSortOrder(sortOrder *int) *ProductsDtoBuilder {
	b.instance.SortOrder = sortOrder
	return b
}

// WithSortOrderValue 设置 sort_order 字段（便捷方法，自动转换为指针）
// 参数:
//   - sortOrder: 排序
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithSortOrderValue(sortOrder int) *ProductsDtoBuilder {
	b.instance.SortOrder = &sortOrder
	return b
}

// WithCreatedAt 设置 created_at 字段
// 参数:
//   - createdAt:
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithCreatedAt(createdAt *time.Time) *ProductsDtoBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

// WithCreatedAtValue 设置 created_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - createdAt:
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithCreatedAtValue(createdAt time.Time) *ProductsDtoBuilder {
	b.instance.CreatedAt = &createdAt
	return b
}

// WithUpdatedAt 设置 updated_at 字段
// 参数:
//   - updatedAt:
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithUpdatedAt(updatedAt *time.Time) *ProductsDtoBuilder {
	b.instance.UpdatedAt = updatedAt
	return b
}

// WithUpdatedAtValue 设置 updated_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - updatedAt:
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithUpdatedAtValue(updatedAt time.Time) *ProductsDtoBuilder {
	b.instance.UpdatedAt = &updatedAt
	return b
}

// WithDeletedAt 设置 deleted_at 字段
// 参数:
//   - deletedAt:
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithDeletedAt(deletedAt *time.Time) *ProductsDtoBuilder {
	b.instance.DeletedAt = deletedAt
	return b
}

// WithDeletedAtValue 设置 deleted_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - deletedAt:
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithDeletedAtValue(deletedAt time.Time) *ProductsDtoBuilder {
	b.instance.DeletedAt = &deletedAt
	return b
}

// WithCategoryIdList 设置 category_idList 字段
// 参数:
//   - categoryIdList: 分类ID IN 查询
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithCategoryIdList(categoryIdList []uint64) *ProductsDtoBuilder {
	b.instance.CategoryIdList = categoryIdList
	return b
}

// WithNameFuzzy 设置 name_fuzzy 字段
// 参数:
//   - nameFuzzy: 商品名称 模糊查询
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithNameFuzzy(nameFuzzy string) *ProductsDtoBuilder {
	b.instance.NameFuzzy = nameFuzzy
	return b
}

// WithDescriptionFuzzy 设置 description_fuzzy 字段
// 参数:
//   - descriptionFuzzy: 商品描述 模糊查询
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithDescriptionFuzzy(descriptionFuzzy *string) *ProductsDtoBuilder {
	b.instance.DescriptionFuzzy = descriptionFuzzy
	return b
}

// WithMainImageFuzzy 设置 main_image_fuzzy 字段
// 参数:
//   - mainImageFuzzy: 主图URL 模糊查询
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithMainImageFuzzy(mainImageFuzzy *string) *ProductsDtoBuilder {
	b.instance.MainImageFuzzy = mainImageFuzzy
	return b
}

// WithImagesFuzzy 设置 images_fuzzy 字段
// 参数:
//   - imagesFuzzy: 商品图片列表 模糊查询
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithImagesFuzzy(imagesFuzzy *string) *ProductsDtoBuilder {
	b.instance.ImagesFuzzy = imagesFuzzy
	return b
}

// WithSalesList 设置 salesList 字段
// 参数:
//   - salesList: 销量 IN 查询
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithSalesList(salesList []*int) *ProductsDtoBuilder {
	b.instance.SalesList = salesList
	return b
}

// WithStatusList 设置 statusList 字段
// 参数:
//   - statusList: 状态：0-下架 1-上架 IN 查询
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithStatusList(statusList []*int8) *ProductsDtoBuilder {
	b.instance.StatusList = statusList
	return b
}

// WithSortOrderList 设置 sort_orderList 字段
// 参数:
//   - sortOrderList: 排序 IN 查询
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithSortOrderList(sortOrderList []*int) *ProductsDtoBuilder {
	b.instance.SortOrderList = sortOrderList
	return b
}

// WithCreatedAtStart 设置 created_atStart 字段
// 参数:
//   - createdAtStart:  开始时间
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithCreatedAtStart(createdAtStart *time.Time) *ProductsDtoBuilder {
	b.instance.CreatedAtStart = createdAtStart
	return b
}

// WithCreatedAtEnd 设置 created_atEnd 字段
// 参数:
//   - createdAtEnd:  结束时间
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithCreatedAtEnd(createdAtEnd *time.Time) *ProductsDtoBuilder {
	b.instance.CreatedAtEnd = createdAtEnd
	return b
}

// WithUpdatedAtStart 设置 updated_atStart 字段
// 参数:
//   - updatedAtStart:  开始时间
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithUpdatedAtStart(updatedAtStart *time.Time) *ProductsDtoBuilder {
	b.instance.UpdatedAtStart = updatedAtStart
	return b
}

// WithUpdatedAtEnd 设置 updated_atEnd 字段
// 参数:
//   - updatedAtEnd:  结束时间
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithUpdatedAtEnd(updatedAtEnd *time.Time) *ProductsDtoBuilder {
	b.instance.UpdatedAtEnd = updatedAtEnd
	return b
}

// WithDeletedAtStart 设置 deleted_atStart 字段
// 参数:
//   - deletedAtStart:  开始时间
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithDeletedAtStart(deletedAtStart *time.Time) *ProductsDtoBuilder {
	b.instance.DeletedAtStart = deletedAtStart
	return b
}

// WithDeletedAtEnd 设置 deleted_atEnd 字段
// 参数:
//   - deletedAtEnd:  结束时间
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithDeletedAtEnd(deletedAtEnd *time.Time) *ProductsDtoBuilder {
	b.instance.DeletedAtEnd = deletedAtEnd
	return b
}

// WithOrderBy 设置 orderBy 字段
// 参数:
//   - orderBy: 排序字段
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithOrderBy(orderBy string) *ProductsDtoBuilder {
	b.instance.OrderBy = orderBy
	return b
}

// WithPageOffset 设置 pageOffset 字段
// 参数:
//   - pageOffset: 分页偏移量
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithPageOffset(pageOffset int) *ProductsDtoBuilder {
	b.instance.PageOffset = pageOffset
	return b
}

// WithPageSize 设置 pageSize 字段
// 参数:
//   - pageSize: 每页数量
//
// 返回:
//   - *ProductsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductsDtoBuilder) WithPageSize(pageSize int) *ProductsDtoBuilder {
	b.instance.PageSize = pageSize
	return b
}

// Build 构建并返回 ProductsDto 实例
// 返回:
//   - *ProductsDto: 构建完成的实例
func (b *ProductsDtoBuilder) Build() *ProductsDto {
	return b.instance
}
