package view

import (
	"encoding/json"
	"time"
)

// UsersVo 用户表 视图对象
type UsersVo struct {
	Id        *uint64    `json:"id,omitempty"`         //
	Username  string     `json:"username,omitempty"`   // 用户名
	Password  string     `json:"password,omitempty"`   // 密码（加密）
	Nickname  *string    `json:"nickname,omitempty"`   // 昵称
	Avatar    *string    `json:"avatar,omitempty"`     // 头像URL
	Phone     *string    `json:"phone,omitempty"`      // 手机号
	Email     *string    `json:"email,omitempty"`      // 邮箱
	Gender    *int8      `json:"gender,omitempty"`     // 性别：0-未知 1-男 2-女
	Birthday  *time.Time `json:"birthday,omitempty"`   // 生日
	Status    *int8      `json:"status,omitempty"`     // 状态：0-禁用 1-正常
	CreatedAt *time.Time `json:"created_at,omitempty"` //
	UpdatedAt *time.Time `json:"updated_at,omitempty"` //
	DeletedAt *time.Time `json:"deleted_at,omitempty"` //
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *UsersVo) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *UsersVo) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}
