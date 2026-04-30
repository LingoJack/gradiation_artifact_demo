package query

import (
	"encoding/json"
	"time"
)

// UsersDto 用户表 数据传输对象
type UsersDto struct {
	Id             *uint64    `json:"id"`               //
	Username       string     `json:"username"`         // 用户名
	Password       string     `json:"password"`         // 密码（加密）
	Nickname       *string    `json:"nickname"`         // 昵称
	Avatar         *string    `json:"avatar"`           // 头像URL
	Phone          *string    `json:"phone"`            // 手机号
	Email          *string    `json:"email"`            // 邮箱
	Gender         *int8      `json:"gender"`           // 性别：0-未知 1-男 2-女
	Birthday       *time.Time `json:"birthday"`         // 生日
	Status         *int8      `json:"status"`           // 状态：0-禁用 1-正常
	CreatedAt      *time.Time `json:"created_at"`       //
	UpdatedAt      *time.Time `json:"updated_at"`       //
	DeletedAt      *time.Time `json:"deleted_at"`       //
	UsernameFuzzy  string     `json:"username_fuzzy"`   // 用户名 模糊查询
	UsernameList   []string   `json:"username_list"`    // 用户名 IN 查询
	PasswordFuzzy  string     `json:"password_fuzzy"`   // 密码（加密） 模糊查询
	NicknameFuzzy  *string    `json:"nickname_fuzzy"`   // 昵称 模糊查询
	AvatarFuzzy    *string    `json:"avatar_fuzzy"`     // 头像URL 模糊查询
	PhoneFuzzy     *string    `json:"phone_fuzzy"`      // 手机号 模糊查询
	PhoneList      []string   `json:"phone_list"`       // 手机号 IN 查询
	EmailFuzzy     *string    `json:"email_fuzzy"`      // 邮箱 模糊查询
	EmailList      []string   `json:"email_list"`       // 邮箱 IN 查询
	BirthdayStart  *time.Time `json:"birthday_start"`   // 生日 开始时间
	BirthdayEnd    *time.Time `json:"birthday_end"`     // 生日 结束时间
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
func (t *UsersDto) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *UsersDto) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// UsersDtoBuilder 用于构建 UsersDto 实例的 Builder
type UsersDtoBuilder struct {
	instance *UsersDto
}

// NewUsersDtoBuilder 创建一个新的 UsersDtoBuilder 实例
// 返回:
//   - *UsersDtoBuilder: Builder 实例，用于链式调用
func NewUsersDtoBuilder() *UsersDtoBuilder {
	return &UsersDtoBuilder{
		instance: &UsersDto{},
	}
}

// WithUsername 设置 username 字段
// 参数:
//   - username: 用户名
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithUsername(username string) *UsersDtoBuilder {
	b.instance.Username = username
	return b
}

// WithPassword 设置 password 字段
// 参数:
//   - password: 密码（加密）
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithPassword(password string) *UsersDtoBuilder {
	b.instance.Password = password
	return b
}

// WithNickname 设置 nickname 字段
// 参数:
//   - nickname: 昵称
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithNickname(nickname *string) *UsersDtoBuilder {
	b.instance.Nickname = nickname
	return b
}

// WithNicknameValue 设置 nickname 字段（便捷方法，自动转换为指针）
// 参数:
//   - nickname: 昵称
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithNicknameValue(nickname string) *UsersDtoBuilder {
	b.instance.Nickname = &nickname
	return b
}

// WithAvatar 设置 avatar 字段
// 参数:
//   - avatar: 头像URL
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithAvatar(avatar *string) *UsersDtoBuilder {
	b.instance.Avatar = avatar
	return b
}

// WithAvatarValue 设置 avatar 字段（便捷方法，自动转换为指针）
// 参数:
//   - avatar: 头像URL
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithAvatarValue(avatar string) *UsersDtoBuilder {
	b.instance.Avatar = &avatar
	return b
}

// WithPhone 设置 phone 字段
// 参数:
//   - phone: 手机号
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithPhone(phone *string) *UsersDtoBuilder {
	b.instance.Phone = phone
	return b
}

// WithPhoneValue 设置 phone 字段（便捷方法，自动转换为指针）
// 参数:
//   - phone: 手机号
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithPhoneValue(phone string) *UsersDtoBuilder {
	b.instance.Phone = &phone
	return b
}

// WithEmail 设置 email 字段
// 参数:
//   - email: 邮箱
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithEmail(email *string) *UsersDtoBuilder {
	b.instance.Email = email
	return b
}

// WithEmailValue 设置 email 字段（便捷方法，自动转换为指针）
// 参数:
//   - email: 邮箱
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithEmailValue(email string) *UsersDtoBuilder {
	b.instance.Email = &email
	return b
}

// WithGender 设置 gender 字段
// 参数:
//   - gender: 性别：0-未知 1-男 2-女
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithGender(gender *int8) *UsersDtoBuilder {
	b.instance.Gender = gender
	return b
}

// WithGenderValue 设置 gender 字段（便捷方法，自动转换为指针）
// 参数:
//   - gender: 性别：0-未知 1-男 2-女
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithGenderValue(gender int8) *UsersDtoBuilder {
	b.instance.Gender = &gender
	return b
}

// WithBirthday 设置 birthday 字段
// 参数:
//   - birthday: 生日
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithBirthday(birthday *time.Time) *UsersDtoBuilder {
	b.instance.Birthday = birthday
	return b
}

// WithBirthdayValue 设置 birthday 字段（便捷方法，自动转换为指针）
// 参数:
//   - birthday: 生日
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithBirthdayValue(birthday time.Time) *UsersDtoBuilder {
	b.instance.Birthday = &birthday
	return b
}

// WithStatus 设置 status 字段
// 参数:
//   - status: 状态：0-禁用 1-正常
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithStatus(status *int8) *UsersDtoBuilder {
	b.instance.Status = status
	return b
}

// WithStatusValue 设置 status 字段（便捷方法，自动转换为指针）
// 参数:
//   - status: 状态：0-禁用 1-正常
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithStatusValue(status int8) *UsersDtoBuilder {
	b.instance.Status = &status
	return b
}

// WithCreatedAt 设置 created_at 字段
// 参数:
//   - createdAt:
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithCreatedAt(createdAt *time.Time) *UsersDtoBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

// WithCreatedAtValue 设置 created_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - createdAt:
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithCreatedAtValue(createdAt time.Time) *UsersDtoBuilder {
	b.instance.CreatedAt = &createdAt
	return b
}

// WithUpdatedAt 设置 updated_at 字段
// 参数:
//   - updatedAt:
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithUpdatedAt(updatedAt *time.Time) *UsersDtoBuilder {
	b.instance.UpdatedAt = updatedAt
	return b
}

// WithUpdatedAtValue 设置 updated_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - updatedAt:
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithUpdatedAtValue(updatedAt time.Time) *UsersDtoBuilder {
	b.instance.UpdatedAt = &updatedAt
	return b
}

// WithDeletedAt 设置 deleted_at 字段
// 参数:
//   - deletedAt:
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithDeletedAt(deletedAt *time.Time) *UsersDtoBuilder {
	b.instance.DeletedAt = deletedAt
	return b
}

// WithDeletedAtValue 设置 deleted_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - deletedAt:
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithDeletedAtValue(deletedAt time.Time) *UsersDtoBuilder {
	b.instance.DeletedAt = &deletedAt
	return b
}

// WithUsernameFuzzy 设置 username_fuzzy 字段
// 参数:
//   - usernameFuzzy: 用户名 模糊查询
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithUsernameFuzzy(usernameFuzzy string) *UsersDtoBuilder {
	b.instance.UsernameFuzzy = usernameFuzzy
	return b
}

// WithUsernameList 设置 usernameList 字段
// 参数:
//   - usernameList: 用户名 IN 查询
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithUsernameList(usernameList []string) *UsersDtoBuilder {
	b.instance.UsernameList = usernameList
	return b
}

// WithPasswordFuzzy 设置 password_fuzzy 字段
// 参数:
//   - passwordFuzzy: 密码（加密） 模糊查询
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithPasswordFuzzy(passwordFuzzy string) *UsersDtoBuilder {
	b.instance.PasswordFuzzy = passwordFuzzy
	return b
}

// WithNicknameFuzzy 设置 nickname_fuzzy 字段
// 参数:
//   - nicknameFuzzy: 昵称 模糊查询
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithNicknameFuzzy(nicknameFuzzy *string) *UsersDtoBuilder {
	b.instance.NicknameFuzzy = nicknameFuzzy
	return b
}

// WithAvatarFuzzy 设置 avatar_fuzzy 字段
// 参数:
//   - avatarFuzzy: 头像URL 模糊查询
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithAvatarFuzzy(avatarFuzzy *string) *UsersDtoBuilder {
	b.instance.AvatarFuzzy = avatarFuzzy
	return b
}

// WithPhoneFuzzy 设置 phone_fuzzy 字段
// 参数:
//   - phoneFuzzy: 手机号 模糊查询
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithPhoneFuzzy(phoneFuzzy *string) *UsersDtoBuilder {
	b.instance.PhoneFuzzy = phoneFuzzy
	return b
}

// WithPhoneList 设置 phoneList 字段
// 参数:
//   - phoneList: 手机号 IN 查询
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithPhoneList(phoneList []string) *UsersDtoBuilder {
	b.instance.PhoneList = phoneList
	return b
}

// WithEmailFuzzy 设置 email_fuzzy 字段
// 参数:
//   - emailFuzzy: 邮箱 模糊查询
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithEmailFuzzy(emailFuzzy *string) *UsersDtoBuilder {
	b.instance.EmailFuzzy = emailFuzzy
	return b
}

// WithEmailList 设置 emailList 字段
// 参数:
//   - emailList: 邮箱 IN 查询
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithEmailList(emailList []string) *UsersDtoBuilder {
	b.instance.EmailList = emailList
	return b
}

// WithBirthdayStart 设置 birthdayStart 字段
// 参数:
//   - birthdayStart: 生日 开始时间
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithBirthdayStart(birthdayStart *time.Time) *UsersDtoBuilder {
	b.instance.BirthdayStart = birthdayStart
	return b
}

// WithBirthdayEnd 设置 birthdayEnd 字段
// 参数:
//   - birthdayEnd: 生日 结束时间
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithBirthdayEnd(birthdayEnd *time.Time) *UsersDtoBuilder {
	b.instance.BirthdayEnd = birthdayEnd
	return b
}

// WithCreatedAtStart 设置 created_atStart 字段
// 参数:
//   - createdAtStart:  开始时间
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithCreatedAtStart(createdAtStart *time.Time) *UsersDtoBuilder {
	b.instance.CreatedAtStart = createdAtStart
	return b
}

// WithCreatedAtEnd 设置 created_atEnd 字段
// 参数:
//   - createdAtEnd:  结束时间
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithCreatedAtEnd(createdAtEnd *time.Time) *UsersDtoBuilder {
	b.instance.CreatedAtEnd = createdAtEnd
	return b
}

// WithUpdatedAtStart 设置 updated_atStart 字段
// 参数:
//   - updatedAtStart:  开始时间
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithUpdatedAtStart(updatedAtStart *time.Time) *UsersDtoBuilder {
	b.instance.UpdatedAtStart = updatedAtStart
	return b
}

// WithUpdatedAtEnd 设置 updated_atEnd 字段
// 参数:
//   - updatedAtEnd:  结束时间
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithUpdatedAtEnd(updatedAtEnd *time.Time) *UsersDtoBuilder {
	b.instance.UpdatedAtEnd = updatedAtEnd
	return b
}

// WithDeletedAtStart 设置 deleted_atStart 字段
// 参数:
//   - deletedAtStart:  开始时间
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithDeletedAtStart(deletedAtStart *time.Time) *UsersDtoBuilder {
	b.instance.DeletedAtStart = deletedAtStart
	return b
}

// WithDeletedAtEnd 设置 deleted_atEnd 字段
// 参数:
//   - deletedAtEnd:  结束时间
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithDeletedAtEnd(deletedAtEnd *time.Time) *UsersDtoBuilder {
	b.instance.DeletedAtEnd = deletedAtEnd
	return b
}

// WithOrderBy 设置 orderBy 字段
// 参数:
//   - orderBy: 排序字段
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithOrderBy(orderBy string) *UsersDtoBuilder {
	b.instance.OrderBy = orderBy
	return b
}

// WithPageOffset 设置 pageOffset 字段
// 参数:
//   - pageOffset: 分页偏移量
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithPageOffset(pageOffset int) *UsersDtoBuilder {
	b.instance.PageOffset = pageOffset
	return b
}

// WithPageSize 设置 pageSize 字段
// 参数:
//   - pageSize: 每页数量
//
// 返回:
//   - *UsersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UsersDtoBuilder) WithPageSize(pageSize int) *UsersDtoBuilder {
	b.instance.PageSize = pageSize
	return b
}

// Build 构建并返回 UsersDto 实例
// 返回:
//   - *UsersDto: 构建完成的实例
func (b *UsersDtoBuilder) Build() *UsersDto {
	return b.instance
}
