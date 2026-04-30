package view

import (
	"encoding/json"
	"time"
)

// CategoriesVo 商品分类表 视图对象
type CategoriesVo struct {
	Id        *uint64    `json:"id,omitempty"`         //
	Name      string     `json:"name,omitempty"`       // 分类名称
	ParentId  *uint64    `json:"parent_id,omitempty"`  // 父分类ID（0表示顶级分类）
	Icon      *string    `json:"icon,omitempty"`       // 分类图标
	SortOrder *int       `json:"sort_order,omitempty"` // 排序
	Status    *int8      `json:"status,omitempty"`     // 状态：0-禁用 1-正常
	CreatedAt *time.Time `json:"created_at,omitempty"` //
	UpdatedAt *time.Time `json:"updated_at,omitempty"` //
	DeletedAt *time.Time `json:"deleted_at,omitempty"` //
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *CategoriesVo) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *CategoriesVo) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}
