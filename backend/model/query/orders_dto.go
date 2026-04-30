package query

import (
	"encoding/json"
	"time"
)

// OrdersDto 订单表 数据传输对象
type OrdersDto struct {
	Id                   *uint64     `json:"id"`                     //
	OrderNo              string      `json:"order_no"`               // 订单号
	UserId               uint64      `json:"user_id"`                // 用户ID
	TotalAmount          float64     `json:"total_amount"`           // 订单总金额
	PayAmount            float64     `json:"pay_amount"`             // 实付金额
	Status               string      `json:"status"`                 // 订单状态
	ReceiverName         string      `json:"receiver_name"`          // 收货人姓名
	ReceiverPhone        string      `json:"receiver_phone"`         // 收货人电话
	ReceiverAddress      string      `json:"receiver_address"`       // 收货地址
	Remark               *string     `json:"remark"`                 // 备注
	PayTime              *time.Time  `json:"pay_time"`               // 支付时间
	DeliveryTime         *time.Time  `json:"delivery_time"`          // 发货时间
	ReceiveTime          *time.Time  `json:"receive_time"`           // 收货时间
	CreatedAt            *time.Time  `json:"created_at"`             //
	UpdatedAt            *time.Time  `json:"updated_at"`             //
	DeletedAt            *time.Time  `json:"deleted_at"`             //
	OrderNoFuzzy         string      `json:"order_no_fuzzy"`         // 订单号 模糊查询
	OrderNoList          []string    `json:"order_no_list"`          // 订单号 IN 查询
	UserIdList           []uint64    `json:"user_id_list"`           // 用户ID IN 查询
	StatusFuzzy          string      `json:"status_fuzzy"`           // 订单状态 模糊查询
	StatusList           []string    `json:"status_list"`            // 订单状态 IN 查询
	ReceiverNameFuzzy    string      `json:"receiver_name_fuzzy"`    // 收货人姓名 模糊查询
	ReceiverPhoneFuzzy   string      `json:"receiver_phone_fuzzy"`   // 收货人电话 模糊查询
	ReceiverAddressFuzzy string      `json:"receiver_address_fuzzy"` // 收货地址 模糊查询
	RemarkFuzzy          *string     `json:"remark_fuzzy"`           // 备注 模糊查询
	PayTimeStart         *time.Time  `json:"pay_time_start"`         // 支付时间 开始时间
	PayTimeEnd           *time.Time  `json:"pay_time_end"`           // 支付时间 结束时间
	DeliveryTimeStart    *time.Time  `json:"delivery_time_start"`    // 发货时间 开始时间
	DeliveryTimeEnd      *time.Time  `json:"delivery_time_end"`      // 发货时间 结束时间
	ReceiveTimeStart     *time.Time  `json:"receive_time_start"`     // 收货时间 开始时间
	ReceiveTimeEnd       *time.Time  `json:"receive_time_end"`       // 收货时间 结束时间
	CreatedAtStart       *time.Time  `json:"created_at_start"`       //  开始时间
	CreatedAtEnd         *time.Time  `json:"created_at_end"`         //  结束时间
	CreatedAtList        []time.Time `json:"created_at_list"`        //  IN 查询
	UpdatedAtStart       *time.Time  `json:"updated_at_start"`       //  开始时间
	UpdatedAtEnd         *time.Time  `json:"updated_at_end"`         //  结束时间
	DeletedAtStart       *time.Time  `json:"deleted_at_start"`       //  开始时间
	DeletedAtEnd         *time.Time  `json:"deleted_at_end"`         //  结束时间
	OrderBy              string      `json:"order_by"`               // 排序字段
	PageOffset           int         `json:"page_offset"`            // 分页偏移量
	PageSize             int         `json:"page_size"`              // 每页数量
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *OrdersDto) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *OrdersDto) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// OrdersDtoBuilder 用于构建 OrdersDto 实例的 Builder
type OrdersDtoBuilder struct {
	instance *OrdersDto
}

// NewOrdersDtoBuilder 创建一个新的 OrdersDtoBuilder 实例
// 返回:
//   - *OrdersDtoBuilder: Builder 实例，用于链式调用
func NewOrdersDtoBuilder() *OrdersDtoBuilder {
	return &OrdersDtoBuilder{
		instance: &OrdersDto{},
	}
}

// WithOrderNo 设置 order_no 字段
// 参数:
//   - orderNo: 订单号
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithOrderNo(orderNo string) *OrdersDtoBuilder {
	b.instance.OrderNo = orderNo
	return b
}

// WithUserId 设置 user_id 字段
// 参数:
//   - userId: 用户ID
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithUserId(userId uint64) *OrdersDtoBuilder {
	b.instance.UserId = userId
	return b
}

// WithTotalAmount 设置 total_amount 字段
// 参数:
//   - totalAmount: 订单总金额
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithTotalAmount(totalAmount float64) *OrdersDtoBuilder {
	b.instance.TotalAmount = totalAmount
	return b
}

// WithPayAmount 设置 pay_amount 字段
// 参数:
//   - payAmount: 实付金额
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithPayAmount(payAmount float64) *OrdersDtoBuilder {
	b.instance.PayAmount = payAmount
	return b
}

// WithStatus 设置 status 字段
// 参数:
//   - status: 订单状态
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithStatus(status string) *OrdersDtoBuilder {
	b.instance.Status = status
	return b
}

// WithReceiverName 设置 receiver_name 字段
// 参数:
//   - receiverName: 收货人姓名
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithReceiverName(receiverName string) *OrdersDtoBuilder {
	b.instance.ReceiverName = receiverName
	return b
}

// WithReceiverPhone 设置 receiver_phone 字段
// 参数:
//   - receiverPhone: 收货人电话
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithReceiverPhone(receiverPhone string) *OrdersDtoBuilder {
	b.instance.ReceiverPhone = receiverPhone
	return b
}

// WithReceiverAddress 设置 receiver_address 字段
// 参数:
//   - receiverAddress: 收货地址
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithReceiverAddress(receiverAddress string) *OrdersDtoBuilder {
	b.instance.ReceiverAddress = receiverAddress
	return b
}

// WithRemark 设置 remark 字段
// 参数:
//   - remark: 备注
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithRemark(remark *string) *OrdersDtoBuilder {
	b.instance.Remark = remark
	return b
}

// WithRemarkValue 设置 remark 字段（便捷方法，自动转换为指针）
// 参数:
//   - remark: 备注
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithRemarkValue(remark string) *OrdersDtoBuilder {
	b.instance.Remark = &remark
	return b
}

// WithPayTime 设置 pay_time 字段
// 参数:
//   - payTime: 支付时间
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithPayTime(payTime *time.Time) *OrdersDtoBuilder {
	b.instance.PayTime = payTime
	return b
}

// WithPayTimeValue 设置 pay_time 字段（便捷方法，自动转换为指针）
// 参数:
//   - payTime: 支付时间
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithPayTimeValue(payTime time.Time) *OrdersDtoBuilder {
	b.instance.PayTime = &payTime
	return b
}

// WithDeliveryTime 设置 delivery_time 字段
// 参数:
//   - deliveryTime: 发货时间
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithDeliveryTime(deliveryTime *time.Time) *OrdersDtoBuilder {
	b.instance.DeliveryTime = deliveryTime
	return b
}

// WithDeliveryTimeValue 设置 delivery_time 字段（便捷方法，自动转换为指针）
// 参数:
//   - deliveryTime: 发货时间
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithDeliveryTimeValue(deliveryTime time.Time) *OrdersDtoBuilder {
	b.instance.DeliveryTime = &deliveryTime
	return b
}

// WithReceiveTime 设置 receive_time 字段
// 参数:
//   - receiveTime: 收货时间
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithReceiveTime(receiveTime *time.Time) *OrdersDtoBuilder {
	b.instance.ReceiveTime = receiveTime
	return b
}

// WithReceiveTimeValue 设置 receive_time 字段（便捷方法，自动转换为指针）
// 参数:
//   - receiveTime: 收货时间
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithReceiveTimeValue(receiveTime time.Time) *OrdersDtoBuilder {
	b.instance.ReceiveTime = &receiveTime
	return b
}

// WithCreatedAt 设置 created_at 字段
// 参数:
//   - createdAt:
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithCreatedAt(createdAt *time.Time) *OrdersDtoBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

// WithCreatedAtValue 设置 created_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - createdAt:
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithCreatedAtValue(createdAt time.Time) *OrdersDtoBuilder {
	b.instance.CreatedAt = &createdAt
	return b
}

// WithUpdatedAt 设置 updated_at 字段
// 参数:
//   - updatedAt:
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithUpdatedAt(updatedAt *time.Time) *OrdersDtoBuilder {
	b.instance.UpdatedAt = updatedAt
	return b
}

// WithUpdatedAtValue 设置 updated_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - updatedAt:
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithUpdatedAtValue(updatedAt time.Time) *OrdersDtoBuilder {
	b.instance.UpdatedAt = &updatedAt
	return b
}

// WithDeletedAt 设置 deleted_at 字段
// 参数:
//   - deletedAt:
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithDeletedAt(deletedAt *time.Time) *OrdersDtoBuilder {
	b.instance.DeletedAt = deletedAt
	return b
}

// WithDeletedAtValue 设置 deleted_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - deletedAt:
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithDeletedAtValue(deletedAt time.Time) *OrdersDtoBuilder {
	b.instance.DeletedAt = &deletedAt
	return b
}

// WithOrderNoFuzzy 设置 order_no_fuzzy 字段
// 参数:
//   - orderNoFuzzy: 订单号 模糊查询
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithOrderNoFuzzy(orderNoFuzzy string) *OrdersDtoBuilder {
	b.instance.OrderNoFuzzy = orderNoFuzzy
	return b
}

// WithOrderNoList 设置 order_noList 字段
// 参数:
//   - orderNoList: 订单号 IN 查询
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithOrderNoList(orderNoList []string) *OrdersDtoBuilder {
	b.instance.OrderNoList = orderNoList
	return b
}

// WithUserIdList 设置 user_idList 字段
// 参数:
//   - userIdList: 用户ID IN 查询
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithUserIdList(userIdList []uint64) *OrdersDtoBuilder {
	b.instance.UserIdList = userIdList
	return b
}

// WithStatusFuzzy 设置 status_fuzzy 字段
// 参数:
//   - statusFuzzy: 订单状态 模糊查询
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithStatusFuzzy(statusFuzzy string) *OrdersDtoBuilder {
	b.instance.StatusFuzzy = statusFuzzy
	return b
}

// WithStatusList 设置 statusList 字段
// 参数:
//   - statusList: 订单状态 IN 查询
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithStatusList(statusList []string) *OrdersDtoBuilder {
	b.instance.StatusList = statusList
	return b
}

// WithReceiverNameFuzzy 设置 receiver_name_fuzzy 字段
// 参数:
//   - receiverNameFuzzy: 收货人姓名 模糊查询
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithReceiverNameFuzzy(receiverNameFuzzy string) *OrdersDtoBuilder {
	b.instance.ReceiverNameFuzzy = receiverNameFuzzy
	return b
}

// WithReceiverPhoneFuzzy 设置 receiver_phone_fuzzy 字段
// 参数:
//   - receiverPhoneFuzzy: 收货人电话 模糊查询
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithReceiverPhoneFuzzy(receiverPhoneFuzzy string) *OrdersDtoBuilder {
	b.instance.ReceiverPhoneFuzzy = receiverPhoneFuzzy
	return b
}

// WithReceiverAddressFuzzy 设置 receiver_address_fuzzy 字段
// 参数:
//   - receiverAddressFuzzy: 收货地址 模糊查询
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithReceiverAddressFuzzy(receiverAddressFuzzy string) *OrdersDtoBuilder {
	b.instance.ReceiverAddressFuzzy = receiverAddressFuzzy
	return b
}

// WithRemarkFuzzy 设置 remark_fuzzy 字段
// 参数:
//   - remarkFuzzy: 备注 模糊查询
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithRemarkFuzzy(remarkFuzzy *string) *OrdersDtoBuilder {
	b.instance.RemarkFuzzy = remarkFuzzy
	return b
}

// WithPayTimeStart 设置 pay_timeStart 字段
// 参数:
//   - payTimeStart: 支付时间 开始时间
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithPayTimeStart(payTimeStart *time.Time) *OrdersDtoBuilder {
	b.instance.PayTimeStart = payTimeStart
	return b
}

// WithPayTimeEnd 设置 pay_timeEnd 字段
// 参数:
//   - payTimeEnd: 支付时间 结束时间
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithPayTimeEnd(payTimeEnd *time.Time) *OrdersDtoBuilder {
	b.instance.PayTimeEnd = payTimeEnd
	return b
}

// WithDeliveryTimeStart 设置 delivery_timeStart 字段
// 参数:
//   - deliveryTimeStart: 发货时间 开始时间
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithDeliveryTimeStart(deliveryTimeStart *time.Time) *OrdersDtoBuilder {
	b.instance.DeliveryTimeStart = deliveryTimeStart
	return b
}

// WithDeliveryTimeEnd 设置 delivery_timeEnd 字段
// 参数:
//   - deliveryTimeEnd: 发货时间 结束时间
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithDeliveryTimeEnd(deliveryTimeEnd *time.Time) *OrdersDtoBuilder {
	b.instance.DeliveryTimeEnd = deliveryTimeEnd
	return b
}

// WithReceiveTimeStart 设置 receive_timeStart 字段
// 参数:
//   - receiveTimeStart: 收货时间 开始时间
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithReceiveTimeStart(receiveTimeStart *time.Time) *OrdersDtoBuilder {
	b.instance.ReceiveTimeStart = receiveTimeStart
	return b
}

// WithReceiveTimeEnd 设置 receive_timeEnd 字段
// 参数:
//   - receiveTimeEnd: 收货时间 结束时间
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithReceiveTimeEnd(receiveTimeEnd *time.Time) *OrdersDtoBuilder {
	b.instance.ReceiveTimeEnd = receiveTimeEnd
	return b
}

// WithCreatedAtStart 设置 created_atStart 字段
// 参数:
//   - createdAtStart:  开始时间
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithCreatedAtStart(createdAtStart *time.Time) *OrdersDtoBuilder {
	b.instance.CreatedAtStart = createdAtStart
	return b
}

// WithCreatedAtEnd 设置 created_atEnd 字段
// 参数:
//   - createdAtEnd:  结束时间
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithCreatedAtEnd(createdAtEnd *time.Time) *OrdersDtoBuilder {
	b.instance.CreatedAtEnd = createdAtEnd
	return b
}

// WithCreatedAtList 设置 created_atList 字段
// 参数:
//   - createdAtList:  IN 查询
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithCreatedAtList(createdAtList []time.Time) *OrdersDtoBuilder {
	b.instance.CreatedAtList = createdAtList
	return b
}

// WithUpdatedAtStart 设置 updated_atStart 字段
// 参数:
//   - updatedAtStart:  开始时间
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithUpdatedAtStart(updatedAtStart *time.Time) *OrdersDtoBuilder {
	b.instance.UpdatedAtStart = updatedAtStart
	return b
}

// WithUpdatedAtEnd 设置 updated_atEnd 字段
// 参数:
//   - updatedAtEnd:  结束时间
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithUpdatedAtEnd(updatedAtEnd *time.Time) *OrdersDtoBuilder {
	b.instance.UpdatedAtEnd = updatedAtEnd
	return b
}

// WithDeletedAtStart 设置 deleted_atStart 字段
// 参数:
//   - deletedAtStart:  开始时间
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithDeletedAtStart(deletedAtStart *time.Time) *OrdersDtoBuilder {
	b.instance.DeletedAtStart = deletedAtStart
	return b
}

// WithDeletedAtEnd 设置 deleted_atEnd 字段
// 参数:
//   - deletedAtEnd:  结束时间
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithDeletedAtEnd(deletedAtEnd *time.Time) *OrdersDtoBuilder {
	b.instance.DeletedAtEnd = deletedAtEnd
	return b
}

// WithOrderBy 设置 orderBy 字段
// 参数:
//   - orderBy: 排序字段
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithOrderBy(orderBy string) *OrdersDtoBuilder {
	b.instance.OrderBy = orderBy
	return b
}

// WithPageOffset 设置 pageOffset 字段
// 参数:
//   - pageOffset: 分页偏移量
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithPageOffset(pageOffset int) *OrdersDtoBuilder {
	b.instance.PageOffset = pageOffset
	return b
}

// WithPageSize 设置 pageSize 字段
// 参数:
//   - pageSize: 每页数量
//
// 返回:
//   - *OrdersDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *OrdersDtoBuilder) WithPageSize(pageSize int) *OrdersDtoBuilder {
	b.instance.PageSize = pageSize
	return b
}

// Build 构建并返回 OrdersDto 实例
// 返回:
//   - *OrdersDto: 构建完成的实例
func (b *OrdersDtoBuilder) Build() *OrdersDto {
	return b.instance
}
