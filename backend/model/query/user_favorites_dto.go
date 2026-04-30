package query

import (
	"encoding/json"
	"time"
)

// UserFavoritesDto 用户收藏表 数据传输对象
type UserFavoritesDto struct {
	Id             *uint64    `json:"id"`               //
	UserId         uint64     `json:"user_id"`          // 用户ID
	ProductId      uint64     `json:"product_id"`       // 商品ID
	CreatedAt      *time.Time `json:"created_at"`       //
	UserIdList     []uint64   `json:"user_id_list"`     // 用户ID IN 查询
	ProductIdList  []uint64   `json:"product_id_list"`  // 商品ID IN 查询
	CreatedAtStart *time.Time `json:"created_at_start"` //  开始时间
	CreatedAtEnd   *time.Time `json:"created_at_end"`   //  结束时间
	OrderBy        string     `json:"order_by"`         // 排序字段
	PageOffset     int        `json:"page_offset"`      // 分页偏移量
	PageSize       int        `json:"page_size"`        // 每页数量
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *UserFavoritesDto) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *UserFavoritesDto) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// UserFavoritesDtoBuilder 用于构建 UserFavoritesDto 实例的 Builder
type UserFavoritesDtoBuilder struct {
	instance *UserFavoritesDto
}

// NewUserFavoritesDtoBuilder 创建一个新的 UserFavoritesDtoBuilder 实例
// 返回:
//   - *UserFavoritesDtoBuilder: Builder 实例，用于链式调用
func NewUserFavoritesDtoBuilder() *UserFavoritesDtoBuilder {
	return &UserFavoritesDtoBuilder{
		instance: &UserFavoritesDto{},
	}
}

// WithUserId 设置 user_id 字段
// 参数:
//   - userId: 用户ID
//
// 返回:
//   - *UserFavoritesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserFavoritesDtoBuilder) WithUserId(userId uint64) *UserFavoritesDtoBuilder {
	b.instance.UserId = userId
	return b
}

// WithProductId 设置 product_id 字段
// 参数:
//   - productId: 商品ID
//
// 返回:
//   - *UserFavoritesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserFavoritesDtoBuilder) WithProductId(productId uint64) *UserFavoritesDtoBuilder {
	b.instance.ProductId = productId
	return b
}

// WithCreatedAt 设置 created_at 字段
// 参数:
//   - createdAt:
//
// 返回:
//   - *UserFavoritesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserFavoritesDtoBuilder) WithCreatedAt(createdAt *time.Time) *UserFavoritesDtoBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

// WithCreatedAtValue 设置 created_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - createdAt:
//
// 返回:
//   - *UserFavoritesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserFavoritesDtoBuilder) WithCreatedAtValue(createdAt time.Time) *UserFavoritesDtoBuilder {
	b.instance.CreatedAt = &createdAt
	return b
}

// WithUserIdList 设置 user_idList 字段
// 参数:
//   - userIdList: 用户ID IN 查询
//
// 返回:
//   - *UserFavoritesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserFavoritesDtoBuilder) WithUserIdList(userIdList []uint64) *UserFavoritesDtoBuilder {
	b.instance.UserIdList = userIdList
	return b
}

// WithProductIdList 设置 product_idList 字段
// 参数:
//   - productIdList: 商品ID IN 查询
//
// 返回:
//   - *UserFavoritesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserFavoritesDtoBuilder) WithProductIdList(productIdList []uint64) *UserFavoritesDtoBuilder {
	b.instance.ProductIdList = productIdList
	return b
}

// WithCreatedAtStart 设置 created_atStart 字段
// 参数:
//   - createdAtStart:  开始时间
//
// 返回:
//   - *UserFavoritesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserFavoritesDtoBuilder) WithCreatedAtStart(createdAtStart *time.Time) *UserFavoritesDtoBuilder {
	b.instance.CreatedAtStart = createdAtStart
	return b
}

// WithCreatedAtEnd 设置 created_atEnd 字段
// 参数:
//   - createdAtEnd:  结束时间
//
// 返回:
//   - *UserFavoritesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserFavoritesDtoBuilder) WithCreatedAtEnd(createdAtEnd *time.Time) *UserFavoritesDtoBuilder {
	b.instance.CreatedAtEnd = createdAtEnd
	return b
}

// WithOrderBy 设置 orderBy 字段
// 参数:
//   - orderBy: 排序字段
//
// 返回:
//   - *UserFavoritesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserFavoritesDtoBuilder) WithOrderBy(orderBy string) *UserFavoritesDtoBuilder {
	b.instance.OrderBy = orderBy
	return b
}

// WithPageOffset 设置 pageOffset 字段
// 参数:
//   - pageOffset: 分页偏移量
//
// 返回:
//   - *UserFavoritesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserFavoritesDtoBuilder) WithPageOffset(pageOffset int) *UserFavoritesDtoBuilder {
	b.instance.PageOffset = pageOffset
	return b
}

// WithPageSize 设置 pageSize 字段
// 参数:
//   - pageSize: 每页数量
//
// 返回:
//   - *UserFavoritesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserFavoritesDtoBuilder) WithPageSize(pageSize int) *UserFavoritesDtoBuilder {
	b.instance.PageSize = pageSize
	return b
}

// Build 构建并返回 UserFavoritesDto 实例
// 返回:
//   - *UserFavoritesDto: 构建完成的实例
func (b *UserFavoritesDtoBuilder) Build() *UserFavoritesDto {
	return b.instance
}
