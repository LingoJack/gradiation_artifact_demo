package query

import (
	"encoding/json"
	"time"
)

// ProductSkusDto 商品SKU表 数据传输对象
type ProductSkusDto struct {
	Id              *uint64    `json:"id"`                //
	ProductId       uint64     `json:"product_id"`        // 商品ID
	SkuCode         string     `json:"sku_code"`          // SKU编码
	SpecValues      *string    `json:"spec_values"`       // 规格值（颜色、尺寸等）
	Price           float64    `json:"price"`             // 价格
	Stock           *int       `json:"stock"`             // 库存
	Image           *string    `json:"image"`             // SKU图片
	Status          *int8      `json:"status"`            // 状态：0-禁用 1-正常
	CreatedAt       *time.Time `json:"created_at"`        //
	UpdatedAt       *time.Time `json:"updated_at"`        //
	DeletedAt       *time.Time `json:"deleted_at"`        //
	ProductIdList   []uint64   `json:"product_id_list"`   // 商品ID IN 查询
	SkuCodeFuzzy    string     `json:"sku_code_fuzzy"`    // SKU编码 模糊查询
	SkuCodeList     []string   `json:"sku_code_list"`     // SKU编码 IN 查询
	SpecValuesFuzzy *string    `json:"spec_values_fuzzy"` // 规格值（颜色、尺寸等） 模糊查询
	ImageFuzzy      *string    `json:"image_fuzzy"`       // SKU图片 模糊查询
	CreatedAtStart  *time.Time `json:"created_at_start"`  //  开始时间
	CreatedAtEnd    *time.Time `json:"created_at_end"`    //  结束时间
	UpdatedAtStart  *time.Time `json:"updated_at_start"`  //  开始时间
	UpdatedAtEnd    *time.Time `json:"updated_at_end"`    //  结束时间
	DeletedAtStart  *time.Time `json:"deleted_at_start"`  //  开始时间
	DeletedAtEnd    *time.Time `json:"deleted_at_end"`    //  结束时间
	OrderBy         string     `json:"order_by"`          // 排序字段
	PageOffset      int        `json:"page_offset"`       // 分页偏移量
	PageSize        int        `json:"page_size"`         // 每页数量
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *ProductSkusDto) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *ProductSkusDto) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// ProductSkusDtoBuilder 用于构建 ProductSkusDto 实例的 Builder
type ProductSkusDtoBuilder struct {
	instance *ProductSkusDto
}

// NewProductSkusDtoBuilder 创建一个新的 ProductSkusDtoBuilder 实例
// 返回:
//   - *ProductSkusDtoBuilder: Builder 实例，用于链式调用
func NewProductSkusDtoBuilder() *ProductSkusDtoBuilder {
	return &ProductSkusDtoBuilder{
		instance: &ProductSkusDto{},
	}
}

// WithProductId 设置 product_id 字段
// 参数:
//   - productId: 商品ID
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithProductId(productId uint64) *ProductSkusDtoBuilder {
	b.instance.ProductId = productId
	return b
}

// WithSkuCode 设置 sku_code 字段
// 参数:
//   - skuCode: SKU编码
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithSkuCode(skuCode string) *ProductSkusDtoBuilder {
	b.instance.SkuCode = skuCode
	return b
}

// WithSpecValues 设置 spec_values 字段
// 参数:
//   - specValues: 规格值（颜色、尺寸等）
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithSpecValues(specValues *string) *ProductSkusDtoBuilder {
	b.instance.SpecValues = specValues
	return b
}

// WithSpecValuesValue 设置 spec_values 字段（便捷方法，自动转换为指针）
// 参数:
//   - specValues: 规格值（颜色、尺寸等）
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithSpecValuesValue(specValues string) *ProductSkusDtoBuilder {
	b.instance.SpecValues = &specValues
	return b
}

// WithPrice 设置 price 字段
// 参数:
//   - price: 价格
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithPrice(price float64) *ProductSkusDtoBuilder {
	b.instance.Price = price
	return b
}

// WithStock 设置 stock 字段
// 参数:
//   - stock: 库存
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithStock(stock *int) *ProductSkusDtoBuilder {
	b.instance.Stock = stock
	return b
}

// WithStockValue 设置 stock 字段（便捷方法，自动转换为指针）
// 参数:
//   - stock: 库存
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithStockValue(stock int) *ProductSkusDtoBuilder {
	b.instance.Stock = &stock
	return b
}

// WithImage 设置 image 字段
// 参数:
//   - image: SKU图片
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithImage(image *string) *ProductSkusDtoBuilder {
	b.instance.Image = image
	return b
}

// WithImageValue 设置 image 字段（便捷方法，自动转换为指针）
// 参数:
//   - image: SKU图片
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithImageValue(image string) *ProductSkusDtoBuilder {
	b.instance.Image = &image
	return b
}

// WithStatus 设置 status 字段
// 参数:
//   - status: 状态：0-禁用 1-正常
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithStatus(status *int8) *ProductSkusDtoBuilder {
	b.instance.Status = status
	return b
}

// WithStatusValue 设置 status 字段（便捷方法，自动转换为指针）
// 参数:
//   - status: 状态：0-禁用 1-正常
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithStatusValue(status int8) *ProductSkusDtoBuilder {
	b.instance.Status = &status
	return b
}

// WithCreatedAt 设置 created_at 字段
// 参数:
//   - createdAt:
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithCreatedAt(createdAt *time.Time) *ProductSkusDtoBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

// WithCreatedAtValue 设置 created_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - createdAt:
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithCreatedAtValue(createdAt time.Time) *ProductSkusDtoBuilder {
	b.instance.CreatedAt = &createdAt
	return b
}

// WithUpdatedAt 设置 updated_at 字段
// 参数:
//   - updatedAt:
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithUpdatedAt(updatedAt *time.Time) *ProductSkusDtoBuilder {
	b.instance.UpdatedAt = updatedAt
	return b
}

// WithUpdatedAtValue 设置 updated_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - updatedAt:
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithUpdatedAtValue(updatedAt time.Time) *ProductSkusDtoBuilder {
	b.instance.UpdatedAt = &updatedAt
	return b
}

// WithDeletedAt 设置 deleted_at 字段
// 参数:
//   - deletedAt:
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithDeletedAt(deletedAt *time.Time) *ProductSkusDtoBuilder {
	b.instance.DeletedAt = deletedAt
	return b
}

// WithDeletedAtValue 设置 deleted_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - deletedAt:
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithDeletedAtValue(deletedAt time.Time) *ProductSkusDtoBuilder {
	b.instance.DeletedAt = &deletedAt
	return b
}

// WithProductIdList 设置 product_idList 字段
// 参数:
//   - productIdList: 商品ID IN 查询
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithProductIdList(productIdList []uint64) *ProductSkusDtoBuilder {
	b.instance.ProductIdList = productIdList
	return b
}

// WithSkuCodeFuzzy 设置 sku_code_fuzzy 字段
// 参数:
//   - skuCodeFuzzy: SKU编码 模糊查询
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithSkuCodeFuzzy(skuCodeFuzzy string) *ProductSkusDtoBuilder {
	b.instance.SkuCodeFuzzy = skuCodeFuzzy
	return b
}

// WithSkuCodeList 设置 sku_codeList 字段
// 参数:
//   - skuCodeList: SKU编码 IN 查询
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithSkuCodeList(skuCodeList []string) *ProductSkusDtoBuilder {
	b.instance.SkuCodeList = skuCodeList
	return b
}

// WithSpecValuesFuzzy 设置 spec_values_fuzzy 字段
// 参数:
//   - specValuesFuzzy: 规格值（颜色、尺寸等） 模糊查询
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithSpecValuesFuzzy(specValuesFuzzy *string) *ProductSkusDtoBuilder {
	b.instance.SpecValuesFuzzy = specValuesFuzzy
	return b
}

// WithImageFuzzy 设置 image_fuzzy 字段
// 参数:
//   - imageFuzzy: SKU图片 模糊查询
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithImageFuzzy(imageFuzzy *string) *ProductSkusDtoBuilder {
	b.instance.ImageFuzzy = imageFuzzy
	return b
}

// WithCreatedAtStart 设置 created_atStart 字段
// 参数:
//   - createdAtStart:  开始时间
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithCreatedAtStart(createdAtStart *time.Time) *ProductSkusDtoBuilder {
	b.instance.CreatedAtStart = createdAtStart
	return b
}

// WithCreatedAtEnd 设置 created_atEnd 字段
// 参数:
//   - createdAtEnd:  结束时间
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithCreatedAtEnd(createdAtEnd *time.Time) *ProductSkusDtoBuilder {
	b.instance.CreatedAtEnd = createdAtEnd
	return b
}

// WithUpdatedAtStart 设置 updated_atStart 字段
// 参数:
//   - updatedAtStart:  开始时间
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithUpdatedAtStart(updatedAtStart *time.Time) *ProductSkusDtoBuilder {
	b.instance.UpdatedAtStart = updatedAtStart
	return b
}

// WithUpdatedAtEnd 设置 updated_atEnd 字段
// 参数:
//   - updatedAtEnd:  结束时间
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithUpdatedAtEnd(updatedAtEnd *time.Time) *ProductSkusDtoBuilder {
	b.instance.UpdatedAtEnd = updatedAtEnd
	return b
}

// WithDeletedAtStart 设置 deleted_atStart 字段
// 参数:
//   - deletedAtStart:  开始时间
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithDeletedAtStart(deletedAtStart *time.Time) *ProductSkusDtoBuilder {
	b.instance.DeletedAtStart = deletedAtStart
	return b
}

// WithDeletedAtEnd 设置 deleted_atEnd 字段
// 参数:
//   - deletedAtEnd:  结束时间
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithDeletedAtEnd(deletedAtEnd *time.Time) *ProductSkusDtoBuilder {
	b.instance.DeletedAtEnd = deletedAtEnd
	return b
}

// WithOrderBy 设置 orderBy 字段
// 参数:
//   - orderBy: 排序字段
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithOrderBy(orderBy string) *ProductSkusDtoBuilder {
	b.instance.OrderBy = orderBy
	return b
}

// WithPageOffset 设置 pageOffset 字段
// 参数:
//   - pageOffset: 分页偏移量
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithPageOffset(pageOffset int) *ProductSkusDtoBuilder {
	b.instance.PageOffset = pageOffset
	return b
}

// WithPageSize 设置 pageSize 字段
// 参数:
//   - pageSize: 每页数量
//
// 返回:
//   - *ProductSkusDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *ProductSkusDtoBuilder) WithPageSize(pageSize int) *ProductSkusDtoBuilder {
	b.instance.PageSize = pageSize
	return b
}

// Build 构建并返回 ProductSkusDto 实例
// 返回:
//   - *ProductSkusDto: 构建完成的实例
func (b *ProductSkusDtoBuilder) Build() *ProductSkusDto {
	return b.instance
}
