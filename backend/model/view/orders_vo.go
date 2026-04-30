package view

import (
	"encoding/json"
	"time"
)

// OrdersVo 订单表 视图对象
type OrdersVo struct {
	Id              *uint64    `json:"id,omitempty"`               //
	OrderNo         string     `json:"order_no,omitempty"`         // 订单号
	UserId          uint64     `json:"user_id,omitempty"`          // 用户ID
	TotalAmount     float64    `json:"total_amount,omitempty"`     // 订单总金额
	PayAmount       float64    `json:"pay_amount,omitempty"`       // 实付金额
	Status          string     `json:"status,omitempty"`           // 订单状态
	ReceiverName    string     `json:"receiver_name,omitempty"`    // 收货人姓名
	ReceiverPhone   string     `json:"receiver_phone,omitempty"`   // 收货人电话
	ReceiverAddress string     `json:"receiver_address,omitempty"` // 收货地址
	Remark          *string    `json:"remark,omitempty"`           // 备注
	PayTime         *time.Time `json:"pay_time,omitempty"`         // 支付时间
	DeliveryTime    *time.Time `json:"delivery_time,omitempty"`    // 发货时间
	ReceiveTime     *time.Time `json:"receive_time,omitempty"`     // 收货时间
	CreatedAt       *time.Time `json:"created_at,omitempty"`       //
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`       //
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`       //
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *OrdersVo) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *OrdersVo) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}
