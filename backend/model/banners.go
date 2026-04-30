package model

import (
	"encoding/json"
	"time"
)

// Banners Banner表
type Banners struct {
	Id        *uint64    `gorm:"column:id;type:bigint(20) UNSIGNED;primaryKey;autoIncrement;comment:;" json:"id"`
	Title     *string    `gorm:"column:title;type:varchar(100);comment:标题;" json:"title"`
	ImageUrl  string     `gorm:"column:image_url;type:varchar(500);comment:图片URL;not null" json:"image_url"`
	LinkUrl   *string    `gorm:"column:link_url;type:varchar(500);comment:跳转链接;" json:"link_url"`
	SortOrder *int       `gorm:"column:sort_order;type:int(11);default:;comment:排序;" json:"sort_order"`
	Status    *int8      `gorm:"column:status;type:tinyint(4);default:;comment:状态：0-禁用 1-正常;" json:"status"`
	StartTime *time.Time `gorm:"column:start_time;type:timestamp;comment:开始时间;" json:"start_time"`
	EndTime   *time.Time `gorm:"column:end_time;type:timestamp;comment:结束时间;" json:"end_time"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:;" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:;" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:timestamp;default:;comment:;" json:"deleted_at"`
}

// TableName 返回表名
func (t *Banners) TableName() string {
	return "banners"
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *Banners) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *Banners) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// BannersBuilder 用于构建 Banners 实例的 Builder
type BannersBuilder struct {
	instance *Banners
}

// NewBannersBuilder 创建一个新的 BannersBuilder 实例
// 返回:
//   - *BannersBuilder: Builder 实例，用于链式调用
func NewBannersBuilder() *BannersBuilder {
	return &BannersBuilder{
		instance: &Banners{},
	}
}

// WithTitle 设置 title 字段
// 参数:
//   - title: 标题
//
// 返回:
//   - *BannersBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersBuilder) WithTitle(title *string) *BannersBuilder {
	b.instance.Title = title
	return b
}

// WithTitleValue 设置 title 字段（便捷方法，自动转换为指针）
// 参数:
//   - title: 标题
//
// 返回:
//   - *BannersBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersBuilder) WithTitleValue(title string) *BannersBuilder {
	b.instance.Title = &title
	return b
}

// WithImageUrl 设置 image_url 字段
// 参数:
//   - imageUrl: 图片URL
//
// 返回:
//   - *BannersBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersBuilder) WithImageUrl(imageUrl string) *BannersBuilder {
	b.instance.ImageUrl = imageUrl
	return b
}

// WithLinkUrl 设置 link_url 字段
// 参数:
//   - linkUrl: 跳转链接
//
// 返回:
//   - *BannersBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersBuilder) WithLinkUrl(linkUrl *string) *BannersBuilder {
	b.instance.LinkUrl = linkUrl
	return b
}

// WithLinkUrlValue 设置 link_url 字段（便捷方法，自动转换为指针）
// 参数:
//   - linkUrl: 跳转链接
//
// 返回:
//   - *BannersBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersBuilder) WithLinkUrlValue(linkUrl string) *BannersBuilder {
	b.instance.LinkUrl = &linkUrl
	return b
}

// WithSortOrder 设置 sort_order 字段
// 参数:
//   - sortOrder: 排序
//
// 返回:
//   - *BannersBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersBuilder) WithSortOrder(sortOrder *int) *BannersBuilder {
	b.instance.SortOrder = sortOrder
	return b
}

// WithSortOrderValue 设置 sort_order 字段（便捷方法，自动转换为指针）
// 参数:
//   - sortOrder: 排序
//
// 返回:
//   - *BannersBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersBuilder) WithSortOrderValue(sortOrder int) *BannersBuilder {
	b.instance.SortOrder = &sortOrder
	return b
}

// WithStatus 设置 status 字段
// 参数:
//   - status: 状态：0-禁用 1-正常
//
// 返回:
//   - *BannersBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersBuilder) WithStatus(status *int8) *BannersBuilder {
	b.instance.Status = status
	return b
}

// WithStatusValue 设置 status 字段（便捷方法，自动转换为指针）
// 参数:
//   - status: 状态：0-禁用 1-正常
//
// 返回:
//   - *BannersBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersBuilder) WithStatusValue(status int8) *BannersBuilder {
	b.instance.Status = &status
	return b
}

// WithStartTime 设置 start_time 字段
// 参数:
//   - startTime: 开始时间
//
// 返回:
//   - *BannersBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersBuilder) WithStartTime(startTime *time.Time) *BannersBuilder {
	b.instance.StartTime = startTime
	return b
}

// WithStartTimeValue 设置 start_time 字段（便捷方法，自动转换为指针）
// 参数:
//   - startTime: 开始时间
//
// 返回:
//   - *BannersBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersBuilder) WithStartTimeValue(startTime time.Time) *BannersBuilder {
	b.instance.StartTime = &startTime
	return b
}

// WithEndTime 设置 end_time 字段
// 参数:
//   - endTime: 结束时间
//
// 返回:
//   - *BannersBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersBuilder) WithEndTime(endTime *time.Time) *BannersBuilder {
	b.instance.EndTime = endTime
	return b
}

// WithEndTimeValue 设置 end_time 字段（便捷方法，自动转换为指针）
// 参数:
//   - endTime: 结束时间
//
// 返回:
//   - *BannersBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersBuilder) WithEndTimeValue(endTime time.Time) *BannersBuilder {
	b.instance.EndTime = &endTime
	return b
}

// WithCreatedAt 设置 created_at 字段
// 参数:
//   - createdAt:
//
// 返回:
//   - *BannersBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersBuilder) WithCreatedAt(createdAt *time.Time) *BannersBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

// WithCreatedAtValue 设置 created_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - createdAt:
//
// 返回:
//   - *BannersBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersBuilder) WithCreatedAtValue(createdAt time.Time) *BannersBuilder {
	b.instance.CreatedAt = &createdAt
	return b
}

// WithUpdatedAt 设置 updated_at 字段
// 参数:
//   - updatedAt:
//
// 返回:
//   - *BannersBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersBuilder) WithUpdatedAt(updatedAt *time.Time) *BannersBuilder {
	b.instance.UpdatedAt = updatedAt
	return b
}

// WithUpdatedAtValue 设置 updated_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - updatedAt:
//
// 返回:
//   - *BannersBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersBuilder) WithUpdatedAtValue(updatedAt time.Time) *BannersBuilder {
	b.instance.UpdatedAt = &updatedAt
	return b
}

// WithDeletedAt 设置 deleted_at 字段
// 参数:
//   - deletedAt:
//
// 返回:
//   - *BannersBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersBuilder) WithDeletedAt(deletedAt *time.Time) *BannersBuilder {
	b.instance.DeletedAt = deletedAt
	return b
}

// WithDeletedAtValue 设置 deleted_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - deletedAt:
//
// 返回:
//   - *BannersBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersBuilder) WithDeletedAtValue(deletedAt time.Time) *BannersBuilder {
	b.instance.DeletedAt = &deletedAt
	return b
}

// Build 构建并返回 Banners 实例
// 返回:
//   - *Banners: 构建完成的实例
func (b *BannersBuilder) Build() *Banners {
	return b.instance
}
