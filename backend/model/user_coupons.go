package model

import "time"

// UserCoupon 用户优惠券模型
type UserCoupon struct {
	ID       uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID   uint64    `gorm:"column:user_id;not null;index" json:"user_id"`
	CouponID uint64    `gorm:"column:coupon_id;not null;index" json:"coupon_id"`
	Status   int8      `gorm:"type:tinyint;default:0" json:"status"` // 0-未使用 1-已使用 2-已过期
	UsedAt   time.Time `gorm:"column:used_at" json:"used_at"`
	OrderID  uint64    `gorm:"column:order_id" json:"order_id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (UserCoupon) TableName() string {
	return "user_coupons"
}
