package view

import (
	"encoding/json"
	"time"
)

// OrderItemsVo 订单项表 视图对象
type OrderItemsVo struct {
	Id            *uint64    `json:"id,omitempty"`              //
	OrderId       uint64     `json:"order_id,omitempty"`        // 订单ID
	ProductId     uint64     `json:"product_id,omitempty"`      // 商品ID
	SkuId         *uint64    `json:"sku_id,omitempty"`          // SKU ID
	ProductName   string     `json:"product_name,omitempty"`    // 商品名称
	SkuSpecValues *string    `json:"sku_spec_values,omitempty"` // SKU规格值
	ProductImage  *string    `json:"product_image,omitempty"`   // 商品图片
	Price         float64    `json:"price,omitempty"`           // 单价
	Quantity      int        `json:"quantity,omitempty"`        // 数量
	TotalAmount   float64    `json:"total_amount,omitempty"`    // 小计金额
	CreatedAt     *time.Time `json:"created_at,omitempty"`      //
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`      //
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *OrderItemsVo) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *OrderItemsVo) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}
