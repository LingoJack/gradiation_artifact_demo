package model

import (
	"encoding/json"
	"time"
)

// Categories 商品分类表
type Categories struct {
	Id        *uint64    `gorm:"column:id;type:bigint(20) UNSIGNED;primaryKey;autoIncrement;comment:;" json:"id"`
	Name      string     `gorm:"column:name;type:varchar(50);comment:分类名称;not null" json:"name"`
	ParentId  *uint64    `gorm:"column:parent_id;type:bigint(20) UNSIGNED;default:;comment:父分类ID（0表示顶级分类）;" json:"parent_id"`
	Icon      *string    `gorm:"column:icon;type:varchar(500);comment:分类图标;" json:"icon"`
	SortOrder *int       `gorm:"column:sort_order;type:int(11);default:;comment:排序;" json:"sort_order"`
	Status    *int8      `gorm:"column:status;type:tinyint(4);default:;comment:状态：0-禁用 1-正常;" json:"status"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:;" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:;" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:timestamp;default:;comment:;" json:"deleted_at"`
}

// TableName 返回表名
func (t *Categories) TableName() string {
	return "categories"
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *Categories) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *Categories) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// CategoriesBuilder 用于构建 Categories 实例的 Builder
type CategoriesBuilder struct {
	instance *Categories
}

// NewCategoriesBuilder 创建一个新的 CategoriesBuilder 实例
// 返回:
//   - *CategoriesBuilder: Builder 实例，用于链式调用
func NewCategoriesBuilder() *CategoriesBuilder {
	return &CategoriesBuilder{
		instance: &Categories{},
	}
}

// WithName 设置 name 字段
// 参数:
//   - name: 分类名称
//
// 返回:
//   - *CategoriesBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesBuilder) WithName(name string) *CategoriesBuilder {
	b.instance.Name = name
	return b
}

// WithParentId 设置 parent_id 字段
// 参数:
//   - parentId: 父分类ID（0表示顶级分类）
//
// 返回:
//   - *CategoriesBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesBuilder) WithParentId(parentId *uint64) *CategoriesBuilder {
	b.instance.ParentId = parentId
	return b
}

// WithParentIdValue 设置 parent_id 字段（便捷方法，自动转换为指针）
// 参数:
//   - parentId: 父分类ID（0表示顶级分类）
//
// 返回:
//   - *CategoriesBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesBuilder) WithParentIdValue(parentId uint64) *CategoriesBuilder {
	b.instance.ParentId = &parentId
	return b
}

// WithIcon 设置 icon 字段
// 参数:
//   - icon: 分类图标
//
// 返回:
//   - *CategoriesBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesBuilder) WithIcon(icon *string) *CategoriesBuilder {
	b.instance.Icon = icon
	return b
}

// WithIconValue 设置 icon 字段（便捷方法，自动转换为指针）
// 参数:
//   - icon: 分类图标
//
// 返回:
//   - *CategoriesBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesBuilder) WithIconValue(icon string) *CategoriesBuilder {
	b.instance.Icon = &icon
	return b
}

// WithSortOrder 设置 sort_order 字段
// 参数:
//   - sortOrder: 排序
//
// 返回:
//   - *CategoriesBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesBuilder) WithSortOrder(sortOrder *int) *CategoriesBuilder {
	b.instance.SortOrder = sortOrder
	return b
}

// WithSortOrderValue 设置 sort_order 字段（便捷方法，自动转换为指针）
// 参数:
//   - sortOrder: 排序
//
// 返回:
//   - *CategoriesBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesBuilder) WithSortOrderValue(sortOrder int) *CategoriesBuilder {
	b.instance.SortOrder = &sortOrder
	return b
}

// WithStatus 设置 status 字段
// 参数:
//   - status: 状态：0-禁用 1-正常
//
// 返回:
//   - *CategoriesBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesBuilder) WithStatus(status *int8) *CategoriesBuilder {
	b.instance.Status = status
	return b
}

// WithStatusValue 设置 status 字段（便捷方法，自动转换为指针）
// 参数:
//   - status: 状态：0-禁用 1-正常
//
// 返回:
//   - *CategoriesBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesBuilder) WithStatusValue(status int8) *CategoriesBuilder {
	b.instance.Status = &status
	return b
}

// WithCreatedAt 设置 created_at 字段
// 参数:
//   - createdAt:
//
// 返回:
//   - *CategoriesBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesBuilder) WithCreatedAt(createdAt *time.Time) *CategoriesBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

// WithCreatedAtValue 设置 created_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - createdAt:
//
// 返回:
//   - *CategoriesBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesBuilder) WithCreatedAtValue(createdAt time.Time) *CategoriesBuilder {
	b.instance.CreatedAt = &createdAt
	return b
}

// WithUpdatedAt 设置 updated_at 字段
// 参数:
//   - updatedAt:
//
// 返回:
//   - *CategoriesBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesBuilder) WithUpdatedAt(updatedAt *time.Time) *CategoriesBuilder {
	b.instance.UpdatedAt = updatedAt
	return b
}

// WithUpdatedAtValue 设置 updated_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - updatedAt:
//
// 返回:
//   - *CategoriesBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesBuilder) WithUpdatedAtValue(updatedAt time.Time) *CategoriesBuilder {
	b.instance.UpdatedAt = &updatedAt
	return b
}

// WithDeletedAt 设置 deleted_at 字段
// 参数:
//   - deletedAt:
//
// 返回:
//   - *CategoriesBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesBuilder) WithDeletedAt(deletedAt *time.Time) *CategoriesBuilder {
	b.instance.DeletedAt = deletedAt
	return b
}

// WithDeletedAtValue 设置 deleted_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - deletedAt:
//
// 返回:
//   - *CategoriesBuilder: 返回 Builder 实例，支持链式调用
func (b *CategoriesBuilder) WithDeletedAtValue(deletedAt time.Time) *CategoriesBuilder {
	b.instance.DeletedAt = &deletedAt
	return b
}

// Build 构建并返回 Categories 实例
// 返回:
//   - *Categories: 构建完成的实例
func (b *CategoriesBuilder) Build() *Categories {
	return b.instance
}
