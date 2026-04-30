package query

import (
	"encoding/json"
	"time"
)

// BannersDto Banner表 数据传输对象
type BannersDto struct {
	Id             *uint64    `json:"id"`               //
	Title          *string    `json:"title"`            // 标题
	ImageUrl       string     `json:"image_url"`        // 图片URL
	LinkUrl        *string    `json:"link_url"`         // 跳转链接
	SortOrder      *int       `json:"sort_order"`       // 排序
	Status         *int8      `json:"status"`           // 状态：0-禁用 1-正常
	StartTime      *time.Time `json:"start_time"`       // 开始时间
	EndTime        *time.Time `json:"end_time"`         // 结束时间
	CreatedAt      *time.Time `json:"created_at"`       //
	UpdatedAt      *time.Time `json:"updated_at"`       //
	DeletedAt      *time.Time `json:"deleted_at"`       //
	TitleFuzzy     *string    `json:"title_fuzzy"`      // 标题 模糊查询
	ImageUrlFuzzy  string     `json:"image_url_fuzzy"`  // 图片URL 模糊查询
	LinkUrlFuzzy   *string    `json:"link_url_fuzzy"`   // 跳转链接 模糊查询
	SortOrderList  []*int     `json:"sort_order_list"`  // 排序 IN 查询
	StatusList     []*int8    `json:"status_list"`      // 状态：0-禁用 1-正常 IN 查询
	StartTimeStart *time.Time `json:"start_time_start"` // 开始时间 开始时间
	StartTimeEnd   *time.Time `json:"start_time_end"`   // 开始时间 结束时间
	EndTimeStart   *time.Time `json:"end_time_start"`   // 结束时间 开始时间
	EndTimeEnd     *time.Time `json:"end_time_end"`     // 结束时间 结束时间
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
func (t *BannersDto) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *BannersDto) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// BannersDtoBuilder 用于构建 BannersDto 实例的 Builder
type BannersDtoBuilder struct {
	instance *BannersDto
}

// NewBannersDtoBuilder 创建一个新的 BannersDtoBuilder 实例
// 返回:
//   - *BannersDtoBuilder: Builder 实例，用于链式调用
func NewBannersDtoBuilder() *BannersDtoBuilder {
	return &BannersDtoBuilder{
		instance: &BannersDto{},
	}
}

// WithTitle 设置 title 字段
// 参数:
//   - title: 标题
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithTitle(title *string) *BannersDtoBuilder {
	b.instance.Title = title
	return b
}

// WithTitleValue 设置 title 字段（便捷方法，自动转换为指针）
// 参数:
//   - title: 标题
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithTitleValue(title string) *BannersDtoBuilder {
	b.instance.Title = &title
	return b
}

// WithImageUrl 设置 image_url 字段
// 参数:
//   - imageUrl: 图片URL
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithImageUrl(imageUrl string) *BannersDtoBuilder {
	b.instance.ImageUrl = imageUrl
	return b
}

// WithLinkUrl 设置 link_url 字段
// 参数:
//   - linkUrl: 跳转链接
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithLinkUrl(linkUrl *string) *BannersDtoBuilder {
	b.instance.LinkUrl = linkUrl
	return b
}

// WithLinkUrlValue 设置 link_url 字段（便捷方法，自动转换为指针）
// 参数:
//   - linkUrl: 跳转链接
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithLinkUrlValue(linkUrl string) *BannersDtoBuilder {
	b.instance.LinkUrl = &linkUrl
	return b
}

// WithSortOrder 设置 sort_order 字段
// 参数:
//   - sortOrder: 排序
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithSortOrder(sortOrder *int) *BannersDtoBuilder {
	b.instance.SortOrder = sortOrder
	return b
}

// WithSortOrderValue 设置 sort_order 字段（便捷方法，自动转换为指针）
// 参数:
//   - sortOrder: 排序
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithSortOrderValue(sortOrder int) *BannersDtoBuilder {
	b.instance.SortOrder = &sortOrder
	return b
}

// WithStatus 设置 status 字段
// 参数:
//   - status: 状态：0-禁用 1-正常
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithStatus(status *int8) *BannersDtoBuilder {
	b.instance.Status = status
	return b
}

// WithStatusValue 设置 status 字段（便捷方法，自动转换为指针）
// 参数:
//   - status: 状态：0-禁用 1-正常
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithStatusValue(status int8) *BannersDtoBuilder {
	b.instance.Status = &status
	return b
}

// WithStartTime 设置 start_time 字段
// 参数:
//   - startTime: 开始时间
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithStartTime(startTime *time.Time) *BannersDtoBuilder {
	b.instance.StartTime = startTime
	return b
}

// WithStartTimeValue 设置 start_time 字段（便捷方法，自动转换为指针）
// 参数:
//   - startTime: 开始时间
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithStartTimeValue(startTime time.Time) *BannersDtoBuilder {
	b.instance.StartTime = &startTime
	return b
}

// WithEndTime 设置 end_time 字段
// 参数:
//   - endTime: 结束时间
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithEndTime(endTime *time.Time) *BannersDtoBuilder {
	b.instance.EndTime = endTime
	return b
}

// WithEndTimeValue 设置 end_time 字段（便捷方法，自动转换为指针）
// 参数:
//   - endTime: 结束时间
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithEndTimeValue(endTime time.Time) *BannersDtoBuilder {
	b.instance.EndTime = &endTime
	return b
}

// WithCreatedAt 设置 created_at 字段
// 参数:
//   - createdAt:
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithCreatedAt(createdAt *time.Time) *BannersDtoBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

// WithCreatedAtValue 设置 created_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - createdAt:
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithCreatedAtValue(createdAt time.Time) *BannersDtoBuilder {
	b.instance.CreatedAt = &createdAt
	return b
}

// WithUpdatedAt 设置 updated_at 字段
// 参数:
//   - updatedAt:
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithUpdatedAt(updatedAt *time.Time) *BannersDtoBuilder {
	b.instance.UpdatedAt = updatedAt
	return b
}

// WithUpdatedAtValue 设置 updated_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - updatedAt:
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithUpdatedAtValue(updatedAt time.Time) *BannersDtoBuilder {
	b.instance.UpdatedAt = &updatedAt
	return b
}

// WithDeletedAt 设置 deleted_at 字段
// 参数:
//   - deletedAt:
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithDeletedAt(deletedAt *time.Time) *BannersDtoBuilder {
	b.instance.DeletedAt = deletedAt
	return b
}

// WithDeletedAtValue 设置 deleted_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - deletedAt:
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithDeletedAtValue(deletedAt time.Time) *BannersDtoBuilder {
	b.instance.DeletedAt = &deletedAt
	return b
}

// WithTitleFuzzy 设置 title_fuzzy 字段
// 参数:
//   - titleFuzzy: 标题 模糊查询
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithTitleFuzzy(titleFuzzy *string) *BannersDtoBuilder {
	b.instance.TitleFuzzy = titleFuzzy
	return b
}

// WithImageUrlFuzzy 设置 image_url_fuzzy 字段
// 参数:
//   - imageUrlFuzzy: 图片URL 模糊查询
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithImageUrlFuzzy(imageUrlFuzzy string) *BannersDtoBuilder {
	b.instance.ImageUrlFuzzy = imageUrlFuzzy
	return b
}

// WithLinkUrlFuzzy 设置 link_url_fuzzy 字段
// 参数:
//   - linkUrlFuzzy: 跳转链接 模糊查询
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithLinkUrlFuzzy(linkUrlFuzzy *string) *BannersDtoBuilder {
	b.instance.LinkUrlFuzzy = linkUrlFuzzy
	return b
}

// WithSortOrderList 设置 sort_orderList 字段
// 参数:
//   - sortOrderList: 排序 IN 查询
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithSortOrderList(sortOrderList []*int) *BannersDtoBuilder {
	b.instance.SortOrderList = sortOrderList
	return b
}

// WithStatusList 设置 statusList 字段
// 参数:
//   - statusList: 状态：0-禁用 1-正常 IN 查询
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithStatusList(statusList []*int8) *BannersDtoBuilder {
	b.instance.StatusList = statusList
	return b
}

// WithStartTimeStart 设置 start_timeStart 字段
// 参数:
//   - startTimeStart: 开始时间 开始时间
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithStartTimeStart(startTimeStart *time.Time) *BannersDtoBuilder {
	b.instance.StartTimeStart = startTimeStart
	return b
}

// WithStartTimeEnd 设置 start_timeEnd 字段
// 参数:
//   - startTimeEnd: 开始时间 结束时间
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithStartTimeEnd(startTimeEnd *time.Time) *BannersDtoBuilder {
	b.instance.StartTimeEnd = startTimeEnd
	return b
}

// WithEndTimeStart 设置 end_timeStart 字段
// 参数:
//   - endTimeStart: 结束时间 开始时间
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithEndTimeStart(endTimeStart *time.Time) *BannersDtoBuilder {
	b.instance.EndTimeStart = endTimeStart
	return b
}

// WithEndTimeEnd 设置 end_timeEnd 字段
// 参数:
//   - endTimeEnd: 结束时间 结束时间
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithEndTimeEnd(endTimeEnd *time.Time) *BannersDtoBuilder {
	b.instance.EndTimeEnd = endTimeEnd
	return b
}

// WithCreatedAtStart 设置 created_atStart 字段
// 参数:
//   - createdAtStart:  开始时间
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithCreatedAtStart(createdAtStart *time.Time) *BannersDtoBuilder {
	b.instance.CreatedAtStart = createdAtStart
	return b
}

// WithCreatedAtEnd 设置 created_atEnd 字段
// 参数:
//   - createdAtEnd:  结束时间
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithCreatedAtEnd(createdAtEnd *time.Time) *BannersDtoBuilder {
	b.instance.CreatedAtEnd = createdAtEnd
	return b
}

// WithUpdatedAtStart 设置 updated_atStart 字段
// 参数:
//   - updatedAtStart:  开始时间
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithUpdatedAtStart(updatedAtStart *time.Time) *BannersDtoBuilder {
	b.instance.UpdatedAtStart = updatedAtStart
	return b
}

// WithUpdatedAtEnd 设置 updated_atEnd 字段
// 参数:
//   - updatedAtEnd:  结束时间
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithUpdatedAtEnd(updatedAtEnd *time.Time) *BannersDtoBuilder {
	b.instance.UpdatedAtEnd = updatedAtEnd
	return b
}

// WithDeletedAtStart 设置 deleted_atStart 字段
// 参数:
//   - deletedAtStart:  开始时间
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithDeletedAtStart(deletedAtStart *time.Time) *BannersDtoBuilder {
	b.instance.DeletedAtStart = deletedAtStart
	return b
}

// WithDeletedAtEnd 设置 deleted_atEnd 字段
// 参数:
//   - deletedAtEnd:  结束时间
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithDeletedAtEnd(deletedAtEnd *time.Time) *BannersDtoBuilder {
	b.instance.DeletedAtEnd = deletedAtEnd
	return b
}

// WithOrderBy 设置 orderBy 字段
// 参数:
//   - orderBy: 排序字段
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithOrderBy(orderBy string) *BannersDtoBuilder {
	b.instance.OrderBy = orderBy
	return b
}

// WithPageOffset 设置 pageOffset 字段
// 参数:
//   - pageOffset: 分页偏移量
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithPageOffset(pageOffset int) *BannersDtoBuilder {
	b.instance.PageOffset = pageOffset
	return b
}

// WithPageSize 设置 pageSize 字段
// 参数:
//   - pageSize: 每页数量
//
// 返回:
//   - *BannersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *BannersDtoBuilder) WithPageSize(pageSize int) *BannersDtoBuilder {
	b.instance.PageSize = pageSize
	return b
}

// Build 构建并返回 BannersDto 实例
// 返回:
//   - *BannersDto: 构建完成的实例
func (b *BannersDtoBuilder) Build() *BannersDto {
	return b.instance
}
