package view

import (
	"encoding/json"
	"time"
)

// ProductsVo 商品表 视图对象
type ProductsVo struct {
	Id            *uint64    `json:"id,omitempty"`             //
	CategoryId    uint64     `json:"category_id,omitempty"`    // 分类ID
	Name          string     `json:"name,omitempty"`           // 商品名称
	Description   *string    `json:"description,omitempty"`    // 商品描述
	MainImage     *string    `json:"main_image,omitempty"`     // 主图URL
	Images        *string    `json:"images,omitempty"`         // 商品图片列表
	Price         float64    `json:"price,omitempty"`          // 价格
	OriginalPrice *float64   `json:"original_price,omitempty"` // 原价
	Stock         *int       `json:"stock,omitempty"`          // 库存
	Sales         *int       `json:"sales,omitempty"`          // 销量
	Status        *int8      `json:"status,omitempty"`         // 状态：0-下架 1-上架
	SortOrder     *int       `json:"sort_order,omitempty"`     // 排序
	CreatedAt     *time.Time `json:"created_at,omitempty"`     //
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`     //
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`     //
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *ProductsVo) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *ProductsVo) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}
