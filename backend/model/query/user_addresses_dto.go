package query

import (
	"encoding/json"
	"time"
)

// UserAddressesDto 用户地址表 数据传输对象
type UserAddressesDto struct {
	Id                 *uint64    `json:"id"`                   //
	UserId             uint64     `json:"user_id"`              // 用户ID
	ReceiverName       string     `json:"receiver_name"`        // 收货人姓名
	ReceiverPhone      string     `json:"receiver_phone"`       // 收货人电话
	Province           string     `json:"province"`             // 省
	City               string     `json:"city"`                 // 市
	District           string     `json:"district"`             // 区
	DetailAddress      string     `json:"detail_address"`       // 详细地址
	IsDefault          *int8      `json:"is_default"`           // 是否默认：0-否 1-是
	CreatedAt          *time.Time `json:"created_at"`           //
	UpdatedAt          *time.Time `json:"updated_at"`           //
	DeletedAt          *time.Time `json:"deleted_at"`           //
	UserIdList         []uint64   `json:"user_id_list"`         // 用户ID IN 查询
	ReceiverNameFuzzy  string     `json:"receiver_name_fuzzy"`  // 收货人姓名 模糊查询
	ReceiverPhoneFuzzy string     `json:"receiver_phone_fuzzy"` // 收货人电话 模糊查询
	ProvinceFuzzy      string     `json:"province_fuzzy"`       // 省 模糊查询
	CityFuzzy          string     `json:"city_fuzzy"`           // 市 模糊查询
	DistrictFuzzy      string     `json:"district_fuzzy"`       // 区 模糊查询
	DetailAddressFuzzy string     `json:"detail_address_fuzzy"` // 详细地址 模糊查询
	CreatedAtStart     *time.Time `json:"created_at_start"`     //  开始时间
	CreatedAtEnd       *time.Time `json:"created_at_end"`       //  结束时间
	UpdatedAtStart     *time.Time `json:"updated_at_start"`     //  开始时间
	UpdatedAtEnd       *time.Time `json:"updated_at_end"`       //  结束时间
	DeletedAtStart     *time.Time `json:"deleted_at_start"`     //  开始时间
	DeletedAtEnd       *time.Time `json:"deleted_at_end"`       //  结束时间
	OrderBy            string     `json:"order_by"`             // 排序字段
	PageOffset         int        `json:"page_offset"`          // 分页偏移量
	PageSize           int        `json:"page_size"`            // 每页数量
}

// Jsonify 将结构体序列化为 JSON 字符串（紧凑格式）
// 返回:
//   - string: JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *UserAddressesDto) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// JsonifyIndent 将结构体序列化为格式化的 JSON 字符串（带缩进）
// 返回:
//   - string: 格式化的 JSON 字符串，如果序列化失败则返回错误信息的 JSON
func (t *UserAddressesDto) JsonifyIndent() string {
	byts, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// UserAddressesDtoBuilder 用于构建 UserAddressesDto 实例的 Builder
type UserAddressesDtoBuilder struct {
	instance *UserAddressesDto
}

// NewUserAddressesDtoBuilder 创建一个新的 UserAddressesDtoBuilder 实例
// 返回:
//   - *UserAddressesDtoBuilder: Builder 实例，用于链式调用
func NewUserAddressesDtoBuilder() *UserAddressesDtoBuilder {
	return &UserAddressesDtoBuilder{
		instance: &UserAddressesDto{},
	}
}

// WithUserId 设置 user_id 字段
// 参数:
//   - userId: 用户ID
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithUserId(userId uint64) *UserAddressesDtoBuilder {
	b.instance.UserId = userId
	return b
}

// WithReceiverName 设置 receiver_name 字段
// 参数:
//   - receiverName: 收货人姓名
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithReceiverName(receiverName string) *UserAddressesDtoBuilder {
	b.instance.ReceiverName = receiverName
	return b
}

// WithReceiverPhone 设置 receiver_phone 字段
// 参数:
//   - receiverPhone: 收货人电话
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithReceiverPhone(receiverPhone string) *UserAddressesDtoBuilder {
	b.instance.ReceiverPhone = receiverPhone
	return b
}

// WithProvince 设置 province 字段
// 参数:
//   - province: 省
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithProvince(province string) *UserAddressesDtoBuilder {
	b.instance.Province = province
	return b
}

// WithCity 设置 city 字段
// 参数:
//   - city: 市
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithCity(city string) *UserAddressesDtoBuilder {
	b.instance.City = city
	return b
}

// WithDistrict 设置 district 字段
// 参数:
//   - district: 区
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithDistrict(district string) *UserAddressesDtoBuilder {
	b.instance.District = district
	return b
}

// WithDetailAddress 设置 detail_address 字段
// 参数:
//   - detailAddress: 详细地址
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithDetailAddress(detailAddress string) *UserAddressesDtoBuilder {
	b.instance.DetailAddress = detailAddress
	return b
}

// WithIsDefault 设置 is_default 字段
// 参数:
//   - isDefault: 是否默认：0-否 1-是
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithIsDefault(isDefault *int8) *UserAddressesDtoBuilder {
	b.instance.IsDefault = isDefault
	return b
}

// WithIsDefaultValue 设置 is_default 字段（便捷方法，自动转换为指针）
// 参数:
//   - isDefault: 是否默认：0-否 1-是
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithIsDefaultValue(isDefault int8) *UserAddressesDtoBuilder {
	b.instance.IsDefault = &isDefault
	return b
}

// WithCreatedAt 设置 created_at 字段
// 参数:
//   - createdAt:
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithCreatedAt(createdAt *time.Time) *UserAddressesDtoBuilder {
	b.instance.CreatedAt = createdAt
	return b
}

// WithCreatedAtValue 设置 created_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - createdAt:
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithCreatedAtValue(createdAt time.Time) *UserAddressesDtoBuilder {
	b.instance.CreatedAt = &createdAt
	return b
}

// WithUpdatedAt 设置 updated_at 字段
// 参数:
//   - updatedAt:
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithUpdatedAt(updatedAt *time.Time) *UserAddressesDtoBuilder {
	b.instance.UpdatedAt = updatedAt
	return b
}

// WithUpdatedAtValue 设置 updated_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - updatedAt:
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithUpdatedAtValue(updatedAt time.Time) *UserAddressesDtoBuilder {
	b.instance.UpdatedAt = &updatedAt
	return b
}

// WithDeletedAt 设置 deleted_at 字段
// 参数:
//   - deletedAt:
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithDeletedAt(deletedAt *time.Time) *UserAddressesDtoBuilder {
	b.instance.DeletedAt = deletedAt
	return b
}

// WithDeletedAtValue 设置 deleted_at 字段（便捷方法，自动转换为指针）
// 参数:
//   - deletedAt:
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithDeletedAtValue(deletedAt time.Time) *UserAddressesDtoBuilder {
	b.instance.DeletedAt = &deletedAt
	return b
}

// WithUserIdList 设置 user_idList 字段
// 参数:
//   - userIdList: 用户ID IN 查询
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithUserIdList(userIdList []uint64) *UserAddressesDtoBuilder {
	b.instance.UserIdList = userIdList
	return b
}

// WithReceiverNameFuzzy 设置 receiver_name_fuzzy 字段
// 参数:
//   - receiverNameFuzzy: 收货人姓名 模糊查询
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithReceiverNameFuzzy(receiverNameFuzzy string) *UserAddressesDtoBuilder {
	b.instance.ReceiverNameFuzzy = receiverNameFuzzy
	return b
}

// WithReceiverPhoneFuzzy 设置 receiver_phone_fuzzy 字段
// 参数:
//   - receiverPhoneFuzzy: 收货人电话 模糊查询
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithReceiverPhoneFuzzy(receiverPhoneFuzzy string) *UserAddressesDtoBuilder {
	b.instance.ReceiverPhoneFuzzy = receiverPhoneFuzzy
	return b
}

// WithProvinceFuzzy 设置 province_fuzzy 字段
// 参数:
//   - provinceFuzzy: 省 模糊查询
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithProvinceFuzzy(provinceFuzzy string) *UserAddressesDtoBuilder {
	b.instance.ProvinceFuzzy = provinceFuzzy
	return b
}

// WithCityFuzzy 设置 city_fuzzy 字段
// 参数:
//   - cityFuzzy: 市 模糊查询
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithCityFuzzy(cityFuzzy string) *UserAddressesDtoBuilder {
	b.instance.CityFuzzy = cityFuzzy
	return b
}

// WithDistrictFuzzy 设置 district_fuzzy 字段
// 参数:
//   - districtFuzzy: 区 模糊查询
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithDistrictFuzzy(districtFuzzy string) *UserAddressesDtoBuilder {
	b.instance.DistrictFuzzy = districtFuzzy
	return b
}

// WithDetailAddressFuzzy 设置 detail_address_fuzzy 字段
// 参数:
//   - detailAddressFuzzy: 详细地址 模糊查询
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithDetailAddressFuzzy(detailAddressFuzzy string) *UserAddressesDtoBuilder {
	b.instance.DetailAddressFuzzy = detailAddressFuzzy
	return b
}

// WithCreatedAtStart 设置 created_atStart 字段
// 参数:
//   - createdAtStart:  开始时间
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithCreatedAtStart(createdAtStart *time.Time) *UserAddressesDtoBuilder {
	b.instance.CreatedAtStart = createdAtStart
	return b
}

// WithCreatedAtEnd 设置 created_atEnd 字段
// 参数:
//   - createdAtEnd:  结束时间
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithCreatedAtEnd(createdAtEnd *time.Time) *UserAddressesDtoBuilder {
	b.instance.CreatedAtEnd = createdAtEnd
	return b
}

// WithUpdatedAtStart 设置 updated_atStart 字段
// 参数:
//   - updatedAtStart:  开始时间
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithUpdatedAtStart(updatedAtStart *time.Time) *UserAddressesDtoBuilder {
	b.instance.UpdatedAtStart = updatedAtStart
	return b
}

// WithUpdatedAtEnd 设置 updated_atEnd 字段
// 参数:
//   - updatedAtEnd:  结束时间
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithUpdatedAtEnd(updatedAtEnd *time.Time) *UserAddressesDtoBuilder {
	b.instance.UpdatedAtEnd = updatedAtEnd
	return b
}

// WithDeletedAtStart 设置 deleted_atStart 字段
// 参数:
//   - deletedAtStart:  开始时间
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithDeletedAtStart(deletedAtStart *time.Time) *UserAddressesDtoBuilder {
	b.instance.DeletedAtStart = deletedAtStart
	return b
}

// WithDeletedAtEnd 设置 deleted_atEnd 字段
// 参数:
//   - deletedAtEnd:  结束时间
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithDeletedAtEnd(deletedAtEnd *time.Time) *UserAddressesDtoBuilder {
	b.instance.DeletedAtEnd = deletedAtEnd
	return b
}

// WithOrderBy 设置 orderBy 字段
// 参数:
//   - orderBy: 排序字段
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithOrderBy(orderBy string) *UserAddressesDtoBuilder {
	b.instance.OrderBy = orderBy
	return b
}

// WithPageOffset 设置 pageOffset 字段
// 参数:
//   - pageOffset: 分页偏移量
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithPageOffset(pageOffset int) *UserAddressesDtoBuilder {
	b.instance.PageOffset = pageOffset
	return b
}

// WithPageSize 设置 pageSize 字段
// 参数:
//   - pageSize: 每页数量
//
// 返回:
//   - *UserAddressesDtoBuilder: 返回 Builder 实例，支持链式调用
func (b *UserAddressesDtoBuilder) WithPageSize(pageSize int) *UserAddressesDtoBuilder {
	b.instance.PageSize = pageSize
	return b
}

// Build 构建并返回 UserAddressesDto 实例
// 返回:
//   - *UserAddressesDto: 构建完成的实例
func (b *UserAddressesDtoBuilder) Build() *UserAddressesDto {
	return b.instance
}
