package model

import "time"

// Coupon 优惠券模型
type Coupon struct {
	ID        uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string     `gorm:"type:varchar(100);not null" json:"name"`
	Discount  float64    `gorm:"type:decimal(10,2);not null" json:"discount"`
	MinSpend  float64    `gorm:"column:min_spend;type:decimal(10,2);default:0" json:"min_spend"`
	Total     int        `gorm:"default:0" json:"total"`
	Claimed   int        `gorm:"default:0" json:"claimed"`
	StartTime *time.Time `gorm:"column:start_time;type:timestamp;not null" json:"start_time"`
	EndTime   *time.Time `gorm:"column:end_time;type:timestamp;not null" json:"end_time"`
	Status    int8       `gorm:"type:tinyint;default:1" json:"status"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:timestamp;default:NULL" json:"deleted_at"`
}

func (Coupon) TableName() string {
	return "coupons"
}
