package query

import (
	"encoding/json"
	"time"
)

// SearchHistoriesDto 搜索历史表 数据传输对象
type SearchHistoriesDto struct {
	Id             *uint64     `json:"id"`               //
	UserId         uint64      `json:"user_id"`          // 用户ID
	Keyword        string      `json:"keyword"`          // 搜索关键词
	CreatedAt      *time.Time  `json:"created_at"`       //
	UserIdList     []uint64    `json:"user_id_list"`     // 用户ID IN 查询
	KeywordFuzzy   string      `json:"keyword_fuzzy"`    // 搜索关键词 模糊查询
	CreatedAtStart *time.Time  `json:"created_at_start"` //  开始时间
	CreatedAtEnd   *time.Time  `json:"created_at_end"`   //  结束时间
	CreatedAtList  []time.Time `json:"created_at_list"`  //  IN 查询
	OrderBy        string      `json:"order_by"`         // 排序字段
	PageOffset     int         `json:"page_offset"`      // 分页偏移量
	PageSize       int         `json:"page_size"`        // 每页数量
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *SearchHistoriesDto) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *SearchHistoriesDto) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// SearchHistoriesDtoBuilder 用于构建 SearchHistoriesDto 实例的 Builder
type SearchHistoriesDtoBuilder struct {
	instance *SearchHistoriesDto
}

// NewSearchHistoriesDtoBuilder 创建一个新的 SearchHistoriesDtoBuilder 实例
// 返回:
//   - *SearchHistoriesDtoBuilder: Builder 实例，用于链式调用
func NewSearchHistoriesDtoBuilder() *SearchHistoriesDtoBuilder {
	return &SearchHistoriesDtoBuilder{
		instance: &SearchHistoriesDto{},
	}
}

// WithUserId 设置 user_id 字段
// 参数:
//   - userId: 用户ID
//
// 返回:
//   - *SearchHistoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *SearchHistoriesDtoBuilder) WithUserId(userId uint64) *SearchHistoriesDtoBuilder {
	b.instance.UserId = userId
	return b
}

// WithKeyword 设置 keyword 字段
// 参数:
//   - keyword: 搜索关键词
//
// 返回:
//   - *SearchHistoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *SearchHistoriesDtoBuilder) WithKeyword(keyword string) *SearchHistoriesDtoBuilder {
	b.instance.Keyword = keyword
	return b
}

// WithCreatedAt 设置 created_at 字段
// 参数:
//   - createdAt:
//
// 返回:
//   - *SearchHistoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *SearchHistoriesDtoBuilder) WithCreatedAt(createdAt *time.Time) *SearchHistoriesDtoBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

// WithCreatedAtValue 设置 created_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - createdAt:
//
// 返回:
//   - *SearchHistoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *SearchHistoriesDtoBuilder) WithCreatedAtValue(createdAt time.Time) *SearchHistoriesDtoBuilder {
	b.instance.CreatedAt = &createdAt
	return b
}

// WithUserIdList 设置 user_idList 字段
// 参数:
//   - userIdList: 用户ID IN 查询
//
// 返回:
//   - *SearchHistoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *SearchHistoriesDtoBuilder) WithUserIdList(userIdList []uint64) *SearchHistoriesDtoBuilder {
	b.instance.UserIdList = userIdList
	return b
}

// WithKeywordFuzzy 设置 keyword_fuzzy 字段
// 参数:
//   - keywordFuzzy: 搜索关键词 模糊查询
//
// 返回:
//   - *SearchHistoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *SearchHistoriesDtoBuilder) WithKeywordFuzzy(keywordFuzzy string) *SearchHistoriesDtoBuilder {
	b.instance.KeywordFuzzy = keywordFuzzy
	return b
}

// WithCreatedAtStart 设置 created_atStart 字段
// 参数:
//   - createdAtStart:  开始时间
//
// 返回:
//   - *SearchHistoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *SearchHistoriesDtoBuilder) WithCreatedAtStart(createdAtStart *time.Time) *SearchHistoriesDtoBuilder {
	b.instance.CreatedAtStart = createdAtStart
	return b
}

// WithCreatedAtEnd 设置 created_atEnd 字段
// 参数:
//   - createdAtEnd:  结束时间
//
// 返回:
//   - *SearchHistoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *SearchHistoriesDtoBuilder) WithCreatedAtEnd(createdAtEnd *time.Time) *SearchHistoriesDtoBuilder {
	b.instance.CreatedAtEnd = createdAtEnd
	return b
}

// WithCreatedAtList 设置 created_atList 字段
// 参数:
//   - createdAtList:  IN 查询
//
// 返回:
//   - *SearchHistoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *SearchHistoriesDtoBuilder) WithCreatedAtList(createdAtList []time.Time) *SearchHistoriesDtoBuilder {
	b.instance.CreatedAtList = createdAtList
	return b
}

// WithOrderBy 设置 orderBy 字段
// 参数:
//   - orderBy: 排序字段
//
// 返回:
//   - *SearchHistoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *SearchHistoriesDtoBuilder) WithOrderBy(orderBy string) *SearchHistoriesDtoBuilder {
	b.instance.OrderBy = orderBy
	return b
}

// WithPageOffset 设置 pageOffset 字段
// 参数:
//   - pageOffset: 分页偏移量
//
// 返回:
//   - *SearchHistoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *SearchHistoriesDtoBuilder) WithPageOffset(pageOffset int) *SearchHistoriesDtoBuilder {
	b.instance.PageOffset = pageOffset
	return b
}

// WithPageSize 设置 pageSize 字段
// 参数:
//   - pageSize: 每页数量
//
// 返回:
//   - *SearchHistoriesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *SearchHistoriesDtoBuilder) WithPageSize(pageSize int) *SearchHistoriesDtoBuilder {
	b.instance.PageSize = pageSize
	return b
}

// Build 构建并返回 SearchHistoriesDto 实例
// 返回:
//   - *SearchHistoriesDto: 构建完成的实例
func (b *SearchHistoriesDtoBuilder) Build() *SearchHistoriesDto {
	return b.instance
}
