package query

import (
	"encoding/json"
	"time"
)

// OrderItemsDto 订单项表 数据传输对象
type OrderItemsDto struct {
	Id                 *uint64    `json:"id"`                    //
	OrderId            uint64     `json:"order_id"`              // 订单ID
	ProductId          uint64     `json:"product_id"`            // 商品ID
	SkuId              *uint64    `json:"sku_id"`                // SKU ID
	ProductName        string     `json:"product_name"`          // 商品名称
	SkuSpecValues      *string    `json:"sku_spec_values"`       // SKU规格值
	ProductImage       *string    `json:"product_image"`         // 商品图片
	Price              float64    `json:"price"`                 // 单价
	Quantity           int        `json:"quantity"`              // 数量
	TotalAmount        float64    `json:"total_amount"`          // 小计金额
	CreatedAt          *time.Time `json:"created_at"`            //
	UpdatedAt          *time.Time `json:"updated_at"`            //
	OrderIdList        []uint64   `json:"order_id_list"`         // 订单ID IN 查询
	ProductIdList      []uint64   `json:"product_id_list"`       // 商品ID IN 查询
	ProductNameFuzzy   string     `json:"product_name_fuzzy"`    // 商品名称 模糊查询
	SkuSpecValuesFuzzy *string    `json:"sku_spec_values_fuzzy"` // SKU规格值 模糊查询
	ProductImageFuzzy  *string    `json:"product_image_fuzzy"`   // 商品图片 模糊查询
	CreatedAtStart     *time.Time `json:"created_at_start"`      //  开始时间
	CreatedAtEnd       *time.Time `json:"created_at_end"`        //  结束时间
	UpdatedAtStart     *time.Time `json:"updated_at_start"`      //  开始时间
	UpdatedAtEnd       *time.Time `json:"updated_at_end"`        //  结束时间
	OrderBy            string     `json:"order_by"`              // 排序字段
	PageOffset         int        `json:"page_offset"`           // 分页偏移量
	PageSize           int        `json:"page_size"`             // 每页数量
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *OrderItemsDto) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *OrderItemsDto) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// OrderItemsDtoBuilder 用于构建 OrderItemsDto 实例的 Builder
type OrderItemsDtoBuilder struct {
	instance *OrderItemsDto
}

// NewOrderItemsDtoBuilder 创建一个新的 OrderItemsDtoBuilder 实例
// 返回:
//   - *OrderItemsDtoBuilder: Builder 实例，用于链式调用
func NewOrderItemsDtoBuilder() *OrderItemsDtoBuilder {
	return &OrderItemsDtoBuilder{
		instance: &OrderItemsDto{},
	}
}

// WithOrderId 设置 order_id 字段
// 参数:
//   - orderId: 订单ID
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithOrderId(orderId uint64) *OrderItemsDtoBuilder {
	b.instance.OrderId = orderId
	return b
}

// WithProductId 设置 product_id 字段
// 参数:
//   - productId: 商品ID
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithProductId(productId uint64) *OrderItemsDtoBuilder {
	b.instance.ProductId = productId
	return b
}

// WithSkuId 设置 sku_id 字段
// 参数:
//   - skuId: SKU ID
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithSkuId(skuId *uint64) *OrderItemsDtoBuilder {
	b.instance.SkuId = skuId
	return b
}

// WithSkuIdValue 设置 sku_id 字段（便捷方法，自动转换为指针）
// 参数:
//   - skuId: SKU ID
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithSkuIdValue(skuId uint64) *OrderItemsDtoBuilder {
	b.instance.SkuId = &skuId
	return b
}

// WithProductName 设置 product_name 字段
// 参数:
//   - productName: 商品名称
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithProductName(productName string) *OrderItemsDtoBuilder {
	b.instance.ProductName = productName
	return b
}

// WithSkuSpecValues 设置 sku_spec_values 字段
// 参数:
//   - skuSpecValues: SKU规格值
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithSkuSpecValues(skuSpecValues *string) *OrderItemsDtoBuilder {
	b.instance.SkuSpecValues = skuSpecValues
	return b
}

// WithSkuSpecValuesValue 设置 sku_spec_values 字段（便捷方法，自动转换为指针）
// 参数:
//   - skuSpecValues: SKU规格值
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithSkuSpecValuesValue(skuSpecValues string) *OrderItemsDtoBuilder {
	b.instance.SkuSpecValues = &skuSpecValues
	return b
}

// WithProductImage 设置 product_image 字段
// 参数:
//   - productImage: 商品图片
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithProductImage(productImage *string) *OrderItemsDtoBuilder {
	b.instance.ProductImage = productImage
	return b
}

// WithProductImageValue 设置 product_image 字段（便捷方法，自动转换为指针）
// 参数:
//   - productImage: 商品图片
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithProductImageValue(productImage string) *OrderItemsDtoBuilder {
	b.instance.ProductImage = &productImage
	return b
}

// WithPrice 设置 price 字段
// 参数:
//   - price: 单价
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithPrice(price float64) *OrderItemsDtoBuilder {
	b.instance.Price = price
	return b
}

// WithQuantity 设置 quantity 字段
// 参数:
//   - quantity: 数量
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithQuantity(quantity int) *OrderItemsDtoBuilder {
	b.instance.Quantity = quantity
	return b
}

// WithTotalAmount 设置 total_amount 字段
// 参数:
//   - totalAmount: 小计金额
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithTotalAmount(totalAmount float64) *OrderItemsDtoBuilder {
	b.instance.TotalAmount = totalAmount
	return b
}

// WithCreatedAt 设置 created_at 字段
// 参数:
//   - createdAt:
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithCreatedAt(createdAt *time.Time) *OrderItemsDtoBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

// WithCreatedAtValue 设置 created_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - createdAt:
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithCreatedAtValue(createdAt time.Time) *OrderItemsDtoBuilder {
	b.instance.CreatedAt = &createdAt
	return b
}

// WithUpdatedAt 设置 updated_at 字段
// 参数:
//   - updatedAt:
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithUpdatedAt(updatedAt *time.Time) *OrderItemsDtoBuilder {
	b.instance.UpdatedAt = updatedAt
	return b
}

// WithUpdatedAtValue 设置 updated_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - updatedAt:
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithUpdatedAtValue(updatedAt time.Time) *OrderItemsDtoBuilder {
	b.instance.UpdatedAt = &updatedAt
	return b
}

// WithOrderIdList 设置 order_idList 字段
// 参数:
//   - orderIdList: 订单ID IN 查询
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithOrderIdList(orderIdList []uint64) *OrderItemsDtoBuilder {
	b.instance.OrderIdList = orderIdList
	return b
}

// WithProductIdList 设置 product_idList 字段
// 参数:
//   - productIdList: 商品ID IN 查询
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithProductIdList(productIdList []uint64) *OrderItemsDtoBuilder {
	b.instance.ProductIdList = productIdList
	return b
}

// WithProductNameFuzzy 设置 product_name_fuzzy 字段
// 参数:
//   - productNameFuzzy: 商品名称 模糊查询
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithProductNameFuzzy(productNameFuzzy string) *OrderItemsDtoBuilder {
	b.instance.ProductNameFuzzy = productNameFuzzy
	return b
}

// WithSkuSpecValuesFuzzy 设置 sku_spec_values_fuzzy 字段
// 参数:
//   - skuSpecValuesFuzzy: SKU规格值 模糊查询
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithSkuSpecValuesFuzzy(skuSpecValuesFuzzy *string) *OrderItemsDtoBuilder {
	b.instance.SkuSpecValuesFuzzy = skuSpecValuesFuzzy
	return b
}

// WithProductImageFuzzy 设置 product_image_fuzzy 字段
// 参数:
//   - productImageFuzzy: 商品图片 模糊查询
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithProductImageFuzzy(productImageFuzzy *string) *OrderItemsDtoBuilder {
	b.instance.ProductImageFuzzy = productImageFuzzy
	return b
}

// WithCreatedAtStart 设置 created_atStart 字段
// 参数:
//   - createdAtStart:  开始时间
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithCreatedAtStart(createdAtStart *time.Time) *OrderItemsDtoBuilder {
	b.instance.CreatedAtStart = createdAtStart
	return b
}

// WithCreatedAtEnd 设置 created_atEnd 字段
// 参数:
//   - createdAtEnd:  结束时间
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithCreatedAtEnd(createdAtEnd *time.Time) *OrderItemsDtoBuilder {
	b.instance.CreatedAtEnd = createdAtEnd
	return b
}

// WithUpdatedAtStart 设置 updated_atStart 字段
// 参数:
//   - updatedAtStart:  开始时间
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithUpdatedAtStart(updatedAtStart *time.Time) *OrderItemsDtoBuilder {
	b.instance.UpdatedAtStart = updatedAtStart
	return b
}

// WithUpdatedAtEnd 设置 updated_atEnd 字段
// 参数:
//   - updatedAtEnd:  结束时间
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithUpdatedAtEnd(updatedAtEnd *time.Time) *OrderItemsDtoBuilder {
	b.instance.UpdatedAtEnd = updatedAtEnd
	return b
}

// WithOrderBy 设置 orderBy 字段
// 参数:
//   - orderBy: 排序字段
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithOrderBy(orderBy string) *OrderItemsDtoBuilder {
	b.instance.OrderBy = orderBy
	return b
}

// WithPageOffset 设置 pageOffset 字段
// 参数:
//   - pageOffset: 分页偏移量
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithPageOffset(pageOffset int) *OrderItemsDtoBuilder {
	b.instance.PageOffset = pageOffset
	return b
}

// WithPageSize 设置 pageSize 字段
// 参数:
//   - pageSize: 每页数量
//
// 返回:
//   - *OrderItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrderItemsDtoBuilder) WithPageSize(pageSize int) *OrderItemsDtoBuilder {
	b.instance.PageSize = pageSize
	return b
}

// Build 构建并返回 OrderItemsDto 实例
// 返回:
//   - *OrderItemsDto: 构建完成的实例
func (b *OrderItemsDtoBuilder) Build() *OrderItemsDto {
	return b.instance
}
