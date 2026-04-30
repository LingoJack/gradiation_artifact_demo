package query

import (
	"encoding/json"
	"time"
)

// CartItemsDto 购物车表 数据传输对象
type CartItemsDto struct {
	Id             *uint64    `json:"id"`               //
	UserId         uint64     `json:"user_id"`          // 用户ID
	ProductId      uint64     `json:"product_id"`       // 商品ID
	SkuId          *uint64    `json:"sku_id"`           // SKU ID
	Quantity       int        `json:"quantity"`         // 数量
	Selected       *int8      `json:"selected"`         // 是否选中：0-否 1-是
	CreatedAt      *time.Time `json:"created_at"`       //
	UpdatedAt      *time.Time `json:"updated_at"`       //
	UserIdList     []uint64   `json:"user_id_list"`     // 用户ID IN 查询
	ProductIdList  []uint64   `json:"product_id_list"`  // 商品ID IN 查询
	SkuIdList      []*uint64  `json:"sku_id_list"`      // SKU ID IN 查询
	CreatedAtStart *time.Time `json:"created_at_start"` //  开始时间
	CreatedAtEnd   *time.Time `json:"created_at_end"`   //  结束时间
	UpdatedAtStart *time.Time `json:"updated_at_start"` //  开始时间
	UpdatedAtEnd   *time.Time `json:"updated_at_end"`   //  结束时间
	OrderBy        string     `json:"order_by"`         // 排序字段
	PageOffset     int        `json:"page_offset"`      // 分页偏移量
	PageSize       int        `json:"page_size"`        // 每页数量
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *CartItemsDto) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *CartItemsDto) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// CartItemsDtoBuilder 用于构建 CartItemsDto 实例的 Builder
type CartItemsDtoBuilder struct {
	instance *CartItemsDto
}

// NewCartItemsDtoBuilder 创建一个新的 CartItemsDtoBuilder 实例
// 返回:
//   - *CartItemsDtoBuilder: Builder 实例，用于链式调用
func NewCartItemsDtoBuilder() *CartItemsDtoBuilder {
	return &CartItemsDtoBuilder{
		instance: &CartItemsDto{},
	}
}

// WithUserId 设置 user_id 字段
// 参数:
//   - userId: 用户ID
//
// 返回:
//   - *CartItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsDtoBuilder) WithUserId(userId uint64) *CartItemsDtoBuilder {
	b.instance.UserId = userId
	return b
}

// WithProductId 设置 product_id 字段
// 参数:
//   - productId: 商品ID
//
// 返回:
//   - *CartItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsDtoBuilder) WithProductId(productId uint64) *CartItemsDtoBuilder {
	b.instance.ProductId = productId
	return b
}

// WithSkuId 设置 sku_id 字段
// 参数:
//   - skuId: SKU ID
//
// 返回:
//   - *CartItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsDtoBuilder) WithSkuId(skuId *uint64) *CartItemsDtoBuilder {
	b.instance.SkuId = skuId
	return b
}

// WithSkuIdValue 设置 sku_id 字段（便捷方法，自动转换为指针）
// 参数:
//   - skuId: SKU ID
//
// 返回:
//   - *CartItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsDtoBuilder) WithSkuIdValue(skuId uint64) *CartItemsDtoBuilder {
	b.instance.SkuId = &skuId
	return b
}

// WithQuantity 设置 quantity 字段
// 参数:
//   - quantity: 数量
//
// 返回:
//   - *CartItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsDtoBuilder) WithQuantity(quantity int) *CartItemsDtoBuilder {
	b.instance.Quantity = quantity
	return b
}

// WithSelected 设置 selected 字段
// 参数:
//   - selected: 是否选中：0-否 1-是
//
// 返回:
//   - *CartItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsDtoBuilder) WithSelected(selected *int8) *CartItemsDtoBuilder {
	b.instance.Selected = selected
	return b
}

// WithSelectedValue 设置 selected 字段（便捷方法，自动转换为指针）
// 参数:
//   - selected: 是否选中：0-否 1-是
//
// 返回:
//   - *CartItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsDtoBuilder) WithSelectedValue(selected int8) *CartItemsDtoBuilder {
	b.instance.Selected = &selected
	return b
}

// WithCreatedAt 设置 created_at 字段
// 参数:
//   - createdAt:
//
// 返回:
//   - *CartItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsDtoBuilder) WithCreatedAt(createdAt *time.Time) *CartItemsDtoBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

// WithCreatedAtValue 设置 created_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - createdAt:
//
// 返回:
//   - *CartItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsDtoBuilder) WithCreatedAtValue(createdAt time.Time) *CartItemsDtoBuilder {
	b.instance.CreatedAt = &createdAt
	return b
}

// WithUpdatedAt 设置 updated_at 字段
// 参数:
//   - updatedAt:
//
// 返回:
//   - *CartItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsDtoBuilder) WithUpdatedAt(updatedAt *time.Time) *CartItemsDtoBuilder {
	b.instance.UpdatedAt = updatedAt
	return b
}

// WithUpdatedAtValue 设置 updated_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - updatedAt:
//
// 返回:
//   - *CartItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsDtoBuilder) WithUpdatedAtValue(updatedAt time.Time) *CartItemsDtoBuilder {
	b.instance.UpdatedAt = &updatedAt
	return b
}

// WithUserIdList 设置 user_idList 字段
// 参数:
//   - userIdList: 用户ID IN 查询
//
// 返回:
//   - *CartItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsDtoBuilder) WithUserIdList(userIdList []uint64) *CartItemsDtoBuilder {
	b.instance.UserIdList = userIdList
	return b
}

// WithProductIdList 设置 product_idList 字段
// 参数:
//   - productIdList: 商品ID IN 查询
//
// 返回:
//   - *CartItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsDtoBuilder) WithProductIdList(productIdList []uint64) *CartItemsDtoBuilder {
	b.instance.ProductIdList = productIdList
	return b
}

// WithSkuIdList 设置 sku_idList 字段
// 参数:
//   - skuIdList: SKU ID IN 查询
//
// 返回:
//   - *CartItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsDtoBuilder) WithSkuIdList(skuIdList []*uint64) *CartItemsDtoBuilder {
	b.instance.SkuIdList = skuIdList
	return b
}

// WithCreatedAtStart 设置 created_atStart 字段
// 参数:
//   - createdAtStart:  开始时间
//
// 返回:
//   - *CartItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsDtoBuilder) WithCreatedAtStart(createdAtStart *time.Time) *CartItemsDtoBuilder {
	b.instance.CreatedAtStart = createdAtStart
	return b
}

// WithCreatedAtEnd 设置 created_atEnd 字段
// 参数:
//   - createdAtEnd:  结束时间
//
// 返回:
//   - *CartItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsDtoBuilder) WithCreatedAtEnd(createdAtEnd *time.Time) *CartItemsDtoBuilder {
	b.instance.CreatedAtEnd = createdAtEnd
	return b
}

// WithUpdatedAtStart 设置 updated_atStart 字段
// 参数:
//   - updatedAtStart:  开始时间
//
// 返回:
//   - *CartItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsDtoBuilder) WithUpdatedAtStart(updatedAtStart *time.Time) *CartItemsDtoBuilder {
	b.instance.UpdatedAtStart = updatedAtStart
	return b
}

// WithUpdatedAtEnd 设置 updated_atEnd 字段
// 参数:
//   - updatedAtEnd:  结束时间
//
// 返回:
//   - *CartItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsDtoBuilder) WithUpdatedAtEnd(updatedAtEnd *time.Time) *CartItemsDtoBuilder {
	b.instance.UpdatedAtEnd = updatedAtEnd
	return b
}

// WithOrderBy 设置 orderBy 字段
// 参数:
//   - orderBy: 排序字段
//
// 返回:
//   - *CartItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsDtoBuilder) WithOrderBy(orderBy string) *CartItemsDtoBuilder {
	b.instance.OrderBy = orderBy
	return b
}

// WithPageOffset 设置 pageOffset 字段
// 参数:
//   - pageOffset: 分页偏移量
//
// 返回:
//   - *CartItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsDtoBuilder) WithPageOffset(pageOffset int) *CartItemsDtoBuilder {
	b.instance.PageOffset = pageOffset
	return b
}

// WithPageSize 设置 pageSize 字段
// 参数:
//   - pageSize: 每页数量
//
// 返回:
//   - *CartItemsDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CartItemsDtoBuilder) WithPageSize(pageSize int) *CartItemsDtoBuilder {
	b.instance.PageSize = pageSize
	return b
}

// Build 构建并返回 CartItemsDto 实例
// 返回:
//   - *CartItemsDto: 构建完成的实例
func (b *CartItemsDtoBuilder) Build() *CartItemsDto {
	return b.instance
}
