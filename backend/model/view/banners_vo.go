package view

import (
	"encoding/json"
	"time"
)

// BannersVo Banner表 视图对象
type BannersVo struct {
	Id        *uint64    `json:"id,omitempty"`         //
	Title     *string    `json:"title,omitempty"`      // 标题
	ImageUrl  string     `json:"image_url,omitempty"`  // 图片URL
	LinkUrl   *string    `json:"link_url,omitempty"`   // 跳转链接
	SortOrder *int       `json:"sort_order,omitempty"` // 排序
	Status    *int8      `json:"status,omitempty"`     // 状态：0-禁用 1-正常
	StartTime *time.Time `json:"start_time,omitempty"` // 开始时间
	EndTime   *time.Time `json:"end_time,omitempty"`   // 结束时间
	CreatedAt *time.Time `json:"created_at,omitempty"` //
	UpdatedAt *time.Time `json:"updated_at,omitempty"` //
	DeletedAt *time.Time `json:"deleted_at,omitempty"` //
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *BannersVo) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *BannersVo) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}
