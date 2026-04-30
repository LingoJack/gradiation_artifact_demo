package model

import (
	"encoding/json"
	"time"
)

// SearchHistories 搜索历史表
type SearchHistories struct {
	Id        *uint64    `gorm:"column:id;type:bigint(20) UNSIGNED;primaryKey;autoIncrement;comment:;" json:"id"`
	UserId    uint64     `gorm:"column:user_id;type:bigint(20) UNSIGNED;comment:用户ID;not null" json:"user_id"`
	Keyword   string     `gorm:"column:keyword;type:varchar(200);comment:搜索关键词;not null" json:"keyword"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:;" json:"created_at"`
}

// TableName 返回表名
func (t *SearchHistories) TableName() string {
	return "search_histories"
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *SearchHistories) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *SearchHistories) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// SearchHistoriesBuilder 用于构建 SearchHistories 实例的 Builder
type SearchHistoriesBuilder struct {
	instance *SearchHistories
}

// NewSearchHistoriesBuilder 创建一个新的 SearchHistoriesBuilder 实例
// 返回:
//   - *SearchHistoriesBuilder: Builder 实例，用于链式调用
func NewSearchHistoriesBuilder() *SearchHistoriesBuilder {
	return &SearchHistoriesBuilder{
		instance: &SearchHistories{},
	}
}

// WithUserId 设置 user_id 字段
// 参数:
//   - userId: 用户ID
//
// 返回:
//   - *SearchHistoriesBuilder: 返回 Builder 实例，支持链式调用
func (b *SearchHistoriesBuilder) WithUserId(userId uint64) *SearchHistoriesBuilder {
	b.instance.UserId = userId
	return b
}

// WithKeyword 设置 keyword 字段
// 参数:
//   - keyword: 搜索关键词
//
// 返回:
//   - *SearchHistoriesBuilder: 返回 Builder 实例，支持链式调用
func (b *SearchHistoriesBuilder) WithKeyword(keyword string) *SearchHistoriesBuilder {
	b.instance.Keyword = keyword
	return b
}

// WithCreatedAt 设置 created_at 字段
// 参数:
//   - createdAt:
//
// 返回:
//   - *SearchHistoriesBuilder: 返回 Builder 实例，支持链式调用
func (b *SearchHistoriesBuilder) WithCreatedAt(createdAt *time.Time) *SearchHistoriesBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

// WithCreatedAtValue 设置 created_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - createdAt:
//
// 返回:
//   - *SearchHistoriesBuilder: 返回 Builder 实例，支持链式调用
func (b *SearchHistoriesBuilder) WithCreatedAtValue(createdAt time.Time) *SearchHistoriesBuilder {
	b.instance.CreatedAt = &createdAt
	return b
}

// Build 构建并返回 SearchHistories 实例
// 返回:
//   - *SearchHistories: 构建完成的实例
func (b *SearchHistoriesBuilder) Build() *SearchHistories {
	return b.instance
}
