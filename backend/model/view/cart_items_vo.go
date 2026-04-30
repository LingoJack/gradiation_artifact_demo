package view

import (
	"encoding/json"
	"time"
)

// CartItemsVo 购物车表 视图对象
type CartItemsVo struct {
	Id        *uint64    `json:"id,omitempty"`         //
	UserId    uint64     `json:"user_id,omitempty"`    // 用户ID
	ProductId uint64     `json:"product_id,omitempty"` // 商品ID
	SkuId     *uint64    `json:"sku_id,omitempty"`     // SKU ID
	Quantity  int        `json:"quantity,omitempty"`   // 数量
	Selected  *int8      `json:"selected,omitempty"`   // 是否选中：0-否 1-是
	CreatedAt *time.Time `json:"created_at,omitempty"` //
	UpdatedAt *time.Time `json:"updated_at,omitempty"` //
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *CartItemsVo) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *CartItemsVo) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}
