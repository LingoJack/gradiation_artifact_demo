package view

import (
	"encoding/json"
	"time"
)

// UserAddressesVo 用户地址表 视图对象
type UserAddressesVo struct {
	Id            *uint64    `json:"id,omitempty"`             //
	UserId        uint64     `json:"user_id,omitempty"`        // 用户ID
	ReceiverName  string     `json:"receiver_name,omitempty"`  // 收货人姓名
	ReceiverPhone string     `json:"receiver_phone,omitempty"` // 收货人电话
	Province      string     `json:"province,omitempty"`       // 省
	City          string     `json:"city,omitempty"`           // 市
	District      string     `json:"district,omitempty"`       // 区
	DetailAddress string     `json:"detail_address,omitempty"` // 详细地址
	IsDefault     *int8      `json:"is_default,omitempty"`     // 是否默认：0-否 1-是
	CreatedAt     *time.Time `json:"created_at,omitempty"`     //
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`     //
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`     //
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *UserAddressesVo) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *UserAddressesVo) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}
