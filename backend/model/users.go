package model

import (
	"encoding/json"
	"time"
)

// Users 用户表
type Users struct {
	Id        *uint64    `gorm:"column:id;type:bigint(20) UNSIGNED;primaryKey;autoIncrement;comment:;" json:"id"`
	Username  string     `gorm:"column:username;type:varchar(50);comment:用户名;not null" json:"username"`
	Password  string     `gorm:"column:password;type:varchar(255);comment:密码（加密）;not null" json:"password"`
	Nickname  *string    `gorm:"column:nickname;type:varchar(100);comment:昵称;" json:"nickname"`
	Avatar    *string    `gorm:"column:avatar;type:varchar(500);comment:头像URL;" json:"avatar"`
	Phone     *string    `gorm:"column:phone;type:varchar(20);comment:手机号;" json:"phone"`
	Email     *string    `gorm:"column:email;type:varchar(100);comment:邮箱;" json:"email"`
	Gender    *int8      `gorm:"column:gender;type:tinyint(4);default:;comment:性别：0-未知 1-男 2-女;" json:"gender"`
	Birthday  *time.Time `gorm:"column:birthday;type:date;comment:生日;" json:"birthday"`
	Status    *int8      `gorm:"column:status;type:tinyint(4);default:;comment:状态：0-禁用 1-正常;" json:"status"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:;" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:;" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:timestamp;default:;comment:;" json:"deleted_at"`
}

// TableName 返回表名
func (t *Users) TableName() string {
	return "users"
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *Users) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *Users) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// UsersBuilder 用于构建 Users 实例的 Builder
type UsersBuilder struct {
	instance *Users
}

// NewUsersBuilder 创建一个新的 UsersBuilder 实例
// 返回:
//   - *UsersBuilder: Builder 实例，用于链式调用
func NewUsersBuilder() *UsersBuilder {
	return &UsersBuilder{
		instance: &Users{},
	}
}

// WithUsername 设置 username 字段
// 参数:
//   - username: 用户名
//
// 返回:
//   - *UsersBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersBuilder) WithUsername(username string) *UsersBuilder {
	b.instance.Username = username
	return b
}

// WithPassword 设置 password 字段
// 参数:
//   - password: 密码（加密）
//
// 返回:
//   - *UsersBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersBuilder) WithPassword(password string) *UsersBuilder {
	b.instance.Password = password
	return b
}

// WithNickname 设置 nickname 字段
// 参数:
//   - nickname: 昵称
//
// 返回:
//   - *UsersBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersBuilder) WithNickname(nickname *string) *UsersBuilder {
	b.instance.Nickname = nickname
	return b
}

// WithNicknameValue 设置 nickname 字段（便捷方法，自动转换为指针）
// 参数:
//   - nickname: 昵称
//
// 返回:
//   - *UsersBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersBuilder) WithNicknameValue(nickname string) *UsersBuilder {
	b.instance.Nickname = &nickname
	return b
}

// WithAvatar 设置 avatar 字段
// 参数:
//   - avatar: 头像URL
//
// 返回:
//   - *UsersBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersBuilder) WithAvatar(avatar *string) *UsersBuilder {
	b.instance.Avatar = avatar
	return b
}

// WithAvatarValue 设置 avatar 字段（便捷方法，自动转换为指针）
// 参数:
//   - avatar: 头像URL
//
// 返回:
//   - *UsersBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersBuilder) WithAvatarValue(avatar string) *UsersBuilder {
	b.instance.Avatar = &avatar
	return b
}

// WithPhone 设置 phone 字段
// 参数:
//   - phone: 手机号
//
// 返回:
//   - *UsersBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersBuilder) WithPhone(phone *string) *UsersBuilder {
	b.instance.Phone = phone
	return b
}

// WithPhoneValue 设置 phone 字段（便捷方法，自动转换为指针）
// 参数:
//   - phone: 手机号
//
// 返回:
//   - *UsersBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersBuilder) WithPhoneValue(phone string) *UsersBuilder {
	b.instance.Phone = &phone
	return b
}

// WithEmail 设置 email 字段
// 参数:
//   - email: 邮箱
//
// 返回:
//   - *UsersBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersBuilder) WithEmail(email *string) *UsersBuilder {
	b.instance.Email = email
	return b
}

// WithEmailValue 设置 email 字段（便捷方法，自动转换为指针）
// 参数:
//   - email: 邮箱
//
// 返回:
//   - *UsersBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersBuilder) WithEmailValue(email string) *UsersBuilder {
	b.instance.Email = &email
	return b
}

// WithGender 设置 gender 字段
// 参数:
//   - gender: 性别：0-未知 1-男 2-女
//
// 返回:
//   - *UsersBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersBuilder) WithGender(gender *int8) *UsersBuilder {
	b.instance.Gender = gender
	return b
}

// WithGenderValue 设置 gender 字段（便捷方法，自动转换为指针）
// 参数:
//   - gender: 性别：0-未知 1-男 2-女
//
// 返回:
//   - *UsersBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersBuilder) WithGenderValue(gender int8) *UsersBuilder {
	b.instance.Gender = &gender
	return b
}

// WithBirthday 设置 birthday 字段
// 参数:
//   - birthday: 生日
//
// 返回:
//   - *UsersBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersBuilder) WithBirthday(birthday *time.Time) *UsersBuilder {
	b.instance.Birthday = birthday
	return b
}

// WithBirthdayValue 设置 birthday 字段（便捷方法，自动转换为指针）
// 参数:
//   - birthday: 生日
//
// 返回:
//   - *UsersBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersBuilder) WithBirthdayValue(birthday time.Time) *UsersBuilder {
	b.instance.Birthday = &birthday
	return b
}

// WithStatus 设置 status 字段
// 参数:
//   - status: 状态：0-禁用 1-正常
//
// 返回:
//   - *UsersBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersBuilder) WithStatus(status *int8) *UsersBuilder {
	b.instance.Status = status
	return b
}

// WithStatusValue 设置 status 字段（便捷方法，自动转换为指针）
// 参数:
//   - status: 状态：0-禁用 1-正常
//
// 返回:
//   - *UsersBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersBuilder) WithStatusValue(status int8) *UsersBuilder {
	b.instance.Status = &status
	return b
}

// WithCreatedAt 设置 created_at 字段
// 参数:
//   - createdAt:
//
// 返回:
//   - *UsersBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersBuilder) WithCreatedAt(createdAt *time.Time) *UsersBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

// WithCreatedAtValue 设置 created_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - createdAt:
//
// 返回:
//   - *UsersBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersBuilder) WithCreatedAtValue(createdAt time.Time) *UsersBuilder {
	b.instance.CreatedAt = &createdAt
	return b
}

// WithUpdatedAt 设置 updated_at 字段
// 参数:
//   - updatedAt:
//
// 返回:
//   - *UsersBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersBuilder) WithUpdatedAt(updatedAt *time.Time) *UsersBuilder {
	b.instance.UpdatedAt = updatedAt
	return b
}

// WithUpdatedAtValue 设置 updated_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - updatedAt:
//
// 返回:
//   - *UsersBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersBuilder) WithUpdatedAtValue(updatedAt time.Time) *UsersBuilder {
	b.instance.UpdatedAt = &updatedAt
	return b
}

// WithDeletedAt 设置 deleted_at 字段
// 参数:
//   - deletedAt:
//
// 返回:
//   - *UsersBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersBuilder) WithDeletedAt(deletedAt *time.Time) *UsersBuilder {
	b.instance.DeletedAt = deletedAt
	return b
}

// WithDeletedAtValue 设置 deleted_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - deletedAt:
//
// 返回:
//   - *UsersBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersBuilder) WithDeletedAtValue(deletedAt time.Time) *UsersBuilder {
	b.instance.DeletedAt = &deletedAt
	return b
}

// Build 构建并返回 Users 实例
// 返回:
//   - *Users: 构建完成的实例
func (b *UsersBuilder) Build() *Users {
	return b.instance
}
