package query

import (
	"encoding/json"
	"time"
)

// CategoriesDto 商品分类表 数据传输对象
type CategoriesDto struct {
	Id             *uint64    `json:"id"`               //
	Name           string     `json:"name"`             // 分类名称
	ParentId       *uint64    `json:"parent_id"`        // 父分类ID（0表示顶级分类）
	Icon           *string    `json:"icon"`             // 分类图标
	SortOrder      *int       `json:"sort_order"`       // 排序
	Status         *int8      `json:"status"`           // 状态：0-禁用 1-正常
	CreatedAt      *time.Time `json:"created_at"`       //
	UpdatedAt      *time.Time `json:"updated_at"`       //
	DeletedAt      *time.Time `json:"deleted_at"`       //
	NameFuzzy      string     `json:"name_fuzzy"`       // 分类名称 模糊查询
	ParentIdList   []*uint64  `json:"parent_id_list"`   // 父分类ID（0表示顶级分类） IN 查询
	IconFuzzy      *string    `json:"icon_fuzzy"`       // 分类图标 模糊查询
	SortOrderList  []*int     `json:"sort_order_list"`  // 排序 IN 查询
	CreatedAtStart *time.Time `json:"created_at_start"` //  开始时间
	CreatedAtEnd   *time.Time `json:"created_at_end"`   //  结束时间
	UpdatedAtStart *time.Time `json:"updated_at_start"` //  开始时间
	UpdatedAtEnd   *time.Time `json:"updated_at_end"`   //  结束时间
	DeletedAtStart *time.Time `json:"deleted_at_start"` //  开始时间
	DeletedAtEnd   *time.Time `json:"deleted_at_end"`   //  结束时间
	OrderBy        string     `json:"order_by"`         // 排序字段
	PageOffset     int        `json:"page_offset"`      // 分页偏移量
	PageSize       int        `json:"page_size"`        // 每页数量
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *CategoriesDto) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *CategoriesDto) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// CategoriesDtoBuilder 用于构建 CategoriesDto 实例的 Builder
type CategoriesDtoBuilder struct {
	instance *CategoriesDto
}

// NewCategoriesDtoBuilder 创建一个新的 CategoriesDtoBuilder 实例
// 返回:
//   - *CategoriesDtoBuilder: Builder 实例，用于链式调用
func NewCategoriesDtoBuilder() *CategoriesDtoBuilder {
	return &CategoriesDtoBuilder{
		instance: &CategoriesDto{},
	}
}

// WithName 设置 name 字段
// 参数:
//   - name: 分类名称
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithName(name string) *CategoriesDtoBuilder {
	b.instance.Name = name
	return b
}

// WithParentId 设置 parent_id 字段
// 参数:
//   - parentId: 父分类ID（0表示顶级分类）
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithParentId(parentId *uint64) *CategoriesDtoBuilder {
	b.instance.ParentId = parentId
	return b
}

// WithParentIdValue 设置 parent_id 字段（便捷方法，自动转换为指针）
// 参数:
//   - parentId: 父分类ID（0表示顶级分类）
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithParentIdValue(parentId uint64) *CategoriesDtoBuilder {
	b.instance.ParentId = &parentId
	return b
}

// WithIcon 设置 icon 字段
// 参数:
//   - icon: 分类图标
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithIcon(icon *string) *CategoriesDtoBuilder {
	b.instance.Icon = icon
	return b
}

// WithIconValue 设置 icon 字段（便捷方法，自动转换为指针）
// 参数:
//   - icon: 分类图标
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithIconValue(icon string) *CategoriesDtoBuilder {
	b.instance.Icon = &icon
	return b
}

// WithSortOrder 设置 sort_order 字段
// 参数:
//   - sortOrder: 排序
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithSortOrder(sortOrder *int) *CategoriesDtoBuilder {
	b.instance.SortOrder = sortOrder
	return b
}

// WithSortOrderValue 设置 sort_order 字段（便捷方法，自动转换为指针）
// 参数:
//   - sortOrder: 排序
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithSortOrderValue(sortOrder int) *CategoriesDtoBuilder {
	b.instance.SortOrder = &sortOrder
	return b
}

// WithStatus 设置 status 字段
// 参数:
//   - status: 状态：0-禁用 1-正常
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithStatus(status *int8) *CategoriesDtoBuilder {
	b.instance.Status = status
	return b
}

// WithStatusValue 设置 status 字段（便捷方法，自动转换为指针）
// 参数:
//   - status: 状态：0-禁用 1-正常
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithStatusValue(status int8) *CategoriesDtoBuilder {
	b.instance.Status = &status
	return b
}

// WithCreatedAt 设置 created_at 字段
// 参数:
//   - createdAt:
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithCreatedAt(createdAt *time.Time) *CategoriesDtoBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

// WithCreatedAtValue 设置 created_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - createdAt:
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithCreatedAtValue(createdAt time.Time) *CategoriesDtoBuilder {
	b.instance.CreatedAt = &createdAt
	return b
}

// WithUpdatedAt 设置 updated_at 字段
// 参数:
//   - updatedAt:
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithUpdatedAt(updatedAt *time.Time) *CategoriesDtoBuilder {
	b.instance.UpdatedAt = updatedAt
	return b
}

// WithUpdatedAtValue 设置 updated_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - updatedAt:
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithUpdatedAtValue(updatedAt time.Time) *CategoriesDtoBuilder {
	b.instance.UpdatedAt = &updatedAt
	return b
}

// WithDeletedAt 设置 deleted_at 字段
// 参数:
//   - deletedAt:
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithDeletedAt(deletedAt *time.Time) *CategoriesDtoBuilder {
	b.instance.DeletedAt = deletedAt
	return b
}

// WithDeletedAtValue 设置 deleted_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - deletedAt:
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithDeletedAtValue(deletedAt time.Time) *CategoriesDtoBuilder {
	b.instance.DeletedAt = &deletedAt
	return b
}

// WithNameFuzzy 设置 name_fuzzy 字段
// 参数:
//   - nameFuzzy: 分类名称 模糊查询
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithNameFuzzy(nameFuzzy string) *CategoriesDtoBuilder {
	b.instance.NameFuzzy = nameFuzzy
	return b
}

// WithParentIdList 设置 parent_idList 字段
// 参数:
//   - parentIdList: 父分类ID（0表示顶级分类） IN 查询
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithParentIdList(parentIdList []*uint64) *CategoriesDtoBuilder {
	b.instance.ParentIdList = parentIdList
	return b
}

// WithIconFuzzy 设置 icon_fuzzy 字段
// 参数:
//   - iconFuzzy: 分类图标 模糊查询
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithIconFuzzy(iconFuzzy *string) *CategoriesDtoBuilder {
	b.instance.IconFuzzy = iconFuzzy
	return b
}

// WithSortOrderList 设置 sort_orderList 字段
// 参数:
//   - sortOrderList: 排序 IN 查询
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithSortOrderList(sortOrderList []*int) *CategoriesDtoBuilder {
	b.instance.SortOrderList = sortOrderList
	return b
}

// WithCreatedAtStart 设置 created_atStart 字段
// 参数:
//   - createdAtStart:  开始时间
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithCreatedAtStart(createdAtStart *time.Time) *CategoriesDtoBuilder {
	b.instance.CreatedAtStart = createdAtStart
	return b
}

// WithCreatedAtEnd 设置 created_atEnd 字段
// 参数:
//   - createdAtEnd:  结束时间
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithCreatedAtEnd(createdAtEnd *time.Time) *CategoriesDtoBuilder {
	b.instance.CreatedAtEnd = createdAtEnd
	return b
}

// WithUpdatedAtStart 设置 updated_atStart 字段
// 参数:
//   - updatedAtStart:  开始时间
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithUpdatedAtStart(updatedAtStart *time.Time) *CategoriesDtoBuilder {
	b.instance.UpdatedAtStart = updatedAtStart
	return b
}

// WithUpdatedAtEnd 设置 updated_atEnd 字段
// 参数:
//   - updatedAtEnd:  结束时间
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithUpdatedAtEnd(updatedAtEnd *time.Time) *CategoriesDtoBuilder {
	b.instance.UpdatedAtEnd = updatedAtEnd
	return b
}

// WithDeletedAtStart 设置 deleted_atStart 字段
// 参数:
//   - deletedAtStart:  开始时间
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithDeletedAtStart(deletedAtStart *time.Time) *CategoriesDtoBuilder {
	b.instance.DeletedAtStart = deletedAtStart
	return b
}

// WithDeletedAtEnd 设置 deleted_atEnd 字段
// 参数:
//   - deletedAtEnd:  结束时间
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithDeletedAtEnd(deletedAtEnd *time.Time) *CategoriesDtoBuilder {
	b.instance.DeletedAtEnd = deletedAtEnd
	return b
}

// WithOrderBy 设置 orderBy 字段
// 参数:
//   - orderBy: 排序字段
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithOrderBy(orderBy string) *CategoriesDtoBuilder {
	b.instance.OrderBy = orderBy
	return b
}

// WithPageOffset 设置 pageOffset 字段
// 参数:
//   - pageOffset: 分页偏移量
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithPageOffset(pageOffset int) *CategoriesDtoBuilder {
	b.instance.PageOffset = pageOffset
	return b
}

// WithPageSize 设置 pageSize 字段
// 参数:
//   - pageSize: 每页数量
//
// 返回:
//   - *CategoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesDtoBuilder) WithPageSize(pageSize int) *CategoriesDtoBuilder {
	b.instance.PageSize = pageSize
	return b
}

// Build 构建并返回 CategoriesDto 实例
// 返回:
//   - *CategoriesDto: 构建完成的实例
func (b *CategoriesDtoBuilder) Build() *CategoriesDto {
	return b.instance
}
