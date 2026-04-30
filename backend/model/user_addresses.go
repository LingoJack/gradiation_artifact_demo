package model

import (
	"encoding/json"
	"time"
)

// UserAddresses 用户地址表
type UserAddresses struct {
	Id            *uint64    `gorm:"column:id;type:bigint(20) UNSIGNED;primaryKey;autoIncrement;comment:;" json:"id"`
	UserId        uint64     `gorm:"column:user_id;type:bigint(20) UNSIGNED;comment:用户ID;not null" json:"user_id"`
	ReceiverName  string     `gorm:"column:receiver_name;type:varchar(50);comment:收货人姓名;not null" json:"receiver_name"`
	ReceiverPhone string     `gorm:"column:receiver_phone;type:varchar(20);comment:收货人电话;not null" json:"receiver_phone"`
	Province      string     `gorm:"column:province;type:varchar(50);comment:省;not null" json:"province"`
	City          string     `gorm:"column:city;type:varchar(50);comment:市;not null" json:"city"`
	District      string     `gorm:"column:district;type:varchar(50);comment:区;not null" json:"district"`
	DetailAddress string     `gorm:"column:detail_address;type:varchar(200);comment:详细地址;not null" json:"detail_address"`
	IsDefault     *int8      `gorm:"column:is_default;type:tinyint(4);default:;comment:是否默认：0-否 1-是;" json:"is_default"`
	CreatedAt     *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:;" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:;" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at;type:timestamp;default:;comment:;" json:"deleted_at"`
}

// TableName 返回表名
func (t *UserAddresses) TableName() string {
	return "user_addresses"
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *UserAddresses) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *UserAddresses) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// UserAddressesBuilder 用于构建 UserAddresses 实例的 Builder
type UserAddressesBuilder struct {
	instance *UserAddresses
}

// NewUserAddressesBuilder 创建一个新的 UserAddressesBuilder 实例
// 返回:
//   - *UserAddressesBuilder: Builder 实例，用于链式调用
func NewUserAddressesBuilder() *UserAddressesBuilder {
	return &UserAddressesBuilder{
		instance: &UserAddresses{},
	}
}

// WithUserId 设置 user_id 字段
// 参数:
//   - userId: 用户ID
//
// 返回:
//   - *UserAddressesBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesBuilder) WithUserId(userId uint64) *UserAddressesBuilder {
	b.instance.UserId = userId
	return b
}

// WithReceiverName 设置 receiver_name 字段
// 参数:
//   - receiverName: 收货人姓名
//
// 返回:
//   - *UserAddressesBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesBuilder) WithReceiverName(receiverName string) *UserAddressesBuilder {
	b.instance.ReceiverName = receiverName
	return b
}

// WithReceiverPhone 设置 receiver_phone 字段
// 参数:
//   - receiverPhone: 收货人电话
//
// 返回:
//   - *UserAddressesBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesBuilder) WithReceiverPhone(receiverPhone string) *UserAddressesBuilder {
	b.instance.ReceiverPhone = receiverPhone
	return b
}

// WithProvince 设置 province 字段
// 参数:
//   - province: 省
//
// 返回:
//   - *UserAddressesBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesBuilder) WithProvince(province string) *UserAddressesBuilder {
	b.instance.Province = province
	return b
}

// WithCity 设置 city 字段
// 参数:
//   - city: 市
//
// 返回:
//   - *UserAddressesBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesBuilder) WithCity(city string) *UserAddressesBuilder {
	b.instance.City = city
	return b
}

// WithDistrict 设置 district 字段
// 参数:
//   - district: 区
//
// 返回:
//   - *UserAddressesBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesBuilder) WithDistrict(district string) *UserAddressesBuilder {
	b.instance.District = district
	return b
}

// WithDetailAddress 设置 detail_address 字段
// 参数:
//   - detailAddress: 详细地址
//
// 返回:
//   - *UserAddressesBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesBuilder) WithDetailAddress(detailAddress string) *UserAddressesBuilder {
	b.instance.DetailAddress = detailAddress
	return b
}

// WithIsDefault 设置 is_default 字段
// 参数:
//   - isDefault: 是否默认：0-否 1-是
//
// 返回:
//   - *UserAddressesBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesBuilder) WithIsDefault(isDefault *int8) *UserAddressesBuilder {
	b.instance.IsDefault = isDefault
	return b
}

// WithIsDefaultValue 设置 is_default 字段（便捷方法，自动转换为指针）
// 参数:
//   - isDefault: 是否默认：0-否 1-是
//
// 返回:
//   - *UserAddressesBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesBuilder) WithIsDefaultValue(isDefault int8) *UserAddressesBuilder {
	b.instance.IsDefault = &isDefault
	return b
}

// WithCreatedAt 设置 created_at 字段
// 参数:
//   - createdAt:
//
// 返回:
//   - *UserAddressesBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesBuilder) WithCreatedAt(createdAt *time.Time) *UserAddressesBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

// WithCreatedAtValue 设置 created_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - createdAt:
//
// 返回:
//   - *UserAddressesBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesBuilder) WithCreatedAtValue(createdAt time.Time) *UserAddressesBuilder {
	b.instance.CreatedAt = &createdAt
	return b
}

// WithUpdatedAt 设置 updated_at 字段
// 参数:
//   - updatedAt:
//
// 返回:
//   - *UserAddressesBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesBuilder) WithUpdatedAt(updatedAt *time.Time) *UserAddressesBuilder {
	b.instance.UpdatedAt = updatedAt
	return b
}

// WithUpdatedAtValue 设置 updated_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - updatedAt:
//
// 返回:
//   - *UserAddressesBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesBuilder) WithUpdatedAtValue(updatedAt time.Time) *UserAddressesBuilder {
	b.instance.UpdatedAt = &updatedAt
	return b
}

// WithDeletedAt 设置 deleted_at 字段
// 参数:
//   - deletedAt:
//
// 返回:
//   - *UserAddressesBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesBuilder) WithDeletedAt(deletedAt *time.Time) *UserAddressesBuilder {
	b.instance.DeletedAt = deletedAt
	return b
}

// WithDeletedAtValue 设置 deleted_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - deletedAt:
//
// 返回:
//   - *UserAddressesBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesBuilder) WithDeletedAtValue(deletedAt time.Time) *UserAddressesBuilder {
	b.instance.DeletedAt = &deletedAt
	return b
}

// Build 构建并返回 UserAddresses 实例
// 返回:
//   - *UserAddresses: 构建完成的实例
func (b *UserAddressesBuilder) Build() *UserAddresses {
	return b.instance
}
