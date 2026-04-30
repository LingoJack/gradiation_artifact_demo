package view

import (
	"encoding/json"
	"time"
)

// ProductSkusVo 商品SKU表 视图对象
type ProductSkusVo struct {
	Id         *uint64    `json:"id,omitempty"`          //
	ProductId  uint64     `json:"product_id,omitempty"`  // 商品ID
	SkuCode    string     `json:"sku_code,omitempty"`    // SKU编码
	SpecValues *string    `json:"spec_values,omitempty"` // 规格值（颜色、尺寸等）
	Price      float64    `json:"price,omitempty"`       // 价格
	Stock      *int       `json:"stock,omitempty"`       // 库存
	Image      *string    `json:"image,omitempty"`       // SKU图片
	Status     *int8      `json:"status,omitempty"`      // 状态：0-禁用 1-正常
	CreatedAt  *time.Time `json:"created_at,omitempty"`  //
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`  //
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`  //
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *ProductSkusVo) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *ProductSkusVo) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}
