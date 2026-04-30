package model

import (
	"encoding/json"
	"time"
)

// UserFavorites 用户收藏表
type UserFavorites struct {
	Id        *uint64    `gorm:"column:id;type:bigint(20) UNSIGNED;primaryKey;autoIncrement;comment:;" json:"id"`
	UserId    uint64     `gorm:"column:user_id;type:bigint(20) UNSIGNED;comment:用户ID;not null" json:"user_id"`
	ProductId uint64     `gorm:"column:product_id;type:bigint(20) UNSIGNED;comment:商品ID;not null" json:"product_id"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:;" json:"created_at"`
}

// TableName 返回表名
func (t *UserFavorites) TableName() string {
	return "user_favorites"
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *UserFavorites) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *UserFavorites) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// UserFavoritesBuilder 用于构建 UserFavorites 实例的 Builder
type UserFavoritesBuilder struct {
	instance *UserFavorites
}

// NewUserFavoritesBuilder 创建一个新的 UserFavoritesBuilder 实例
// 返回:
//   - *UserFavoritesBuilder: Builder 实例，用于链式调用
func NewUserFavoritesBuilder() *UserFavoritesBuilder {
	return &UserFavoritesBuilder{
		instance: &UserFavorites{},
	}
}

// WithUserId 设置 user_id 字段
// 参数:
//   - userId: 用户ID
//
// 返回:
//   - *UserFavoritesBuilder: 返回 Builder 实例，支持链式调用
func (b *UserFavoritesBuilder) WithUserId(userId uint64) *UserFavoritesBuilder {
	b.instance.UserId = userId
	return b
}

// WithProductId 设置 product_id 字段
// 参数:
//   - productId: 商品ID
//
// 返回:
//   - *UserFavoritesBuilder: 返回 Builder 实例，支持链式调用
func (b *UserFavoritesBuilder) WithProductId(productId uint64) *UserFavoritesBuilder {
	b.instance.ProductId = productId
	return b
}

// WithCreatedAt 设置 created_at 字段
// 参数:
//   - createdAt:
//
// 返回:
//   - *UserFavoritesBuilder: 返回 Builder 实例，支持链式调用
func (b *UserFavoritesBuilder) WithCreatedAt(createdAt *time.Time) *UserFavoritesBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

// WithCreatedAtValue 设置 created_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - createdAt:
//
// 返回:
//   - *UserFavoritesBuilder: 返回 Builder 实例，支持链式调用
func (b *UserFavoritesBuilder) WithCreatedAtValue(createdAt time.Time) *UserFavoritesBuilder {
	b.instance.CreatedAt = &createdAt
	return b
}

// Build 构建并返回 UserFavorites 实例
// 返回:
//   - *UserFavorites: 构建完成的实例
func (b *UserFavoritesBuilder) Build() *UserFavorites {
	return b.instance
}
