package model

import (
	"encoding/json"
	"time"
)

// Reviews 商品评价
type Reviews struct {
	Id        *uint64    `gorm:"column:id;type:bigint(20) UNSIGNED;primaryKey;autoIncrement;" json:"id"`
	UserId    uint64     `gorm:"column:user_id;type:bigint(20) UNSIGNED;not null;index" json:"user_id"`
	ProductId uint64     `gorm:"column:product_id;type:bigint(20) UNSIGNED;not null;index" json:"product_id"`
	OrderId   uint64     `gorm:"column:order_id;type:bigint(20) UNSIGNED;not null;index" json:"order_id"`
	Rating    int8       `gorm:"column:rating;type:tinyint;not null" json:"rating"`
	Content   string     `gorm:"column:content;type:text" json:"content"`
	Images    string     `gorm:"column:images;type:text" json:"images"`
	SpecName  *string    `gorm:"column:spec_name;type:varchar(200);" json:"spec_name"`
	Reply     *string    `gorm:"column:reply;type:text;" json:"reply"`
	Likes     int        `gorm:"column:likes;type:int(11);default:0" json:"likes"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:timestamp;" json:"deleted_at"`
}

func (Reviews) TableName() string {
	return "reviews"
}

func (t *Reviews) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// Coupons 优惠券模板
type Coupons struct {
	Id         *uint64    `gorm:"column:id;type:bigint(20) UNSIGNED;primaryKey;autoIncrement;" json:"id"`
	Name       string     `gorm:"column:name;type:varchar(100);not null" json:"name"`
	Discount   float64    `gorm:"column:discount;type:decimal(10,2);not null" json:"discount"`
	MinSpend   float64    `gorm:"column:min_spend;type:decimal(10,2);default:0" json:"min_spend"`
	ScopeType  string     `gorm:"column:scope_type;type:varchar(50);default:all" json:"scope_type"`
	ScopeValue *string    `gorm:"column:scope_value;type:varchar(100);" json:"scope_value"`
	StartTime  *time.Time `gorm:"column:start_time;type:timestamp;" json:"start_time"`
	EndTime    *time.Time `gorm:"column:end_time;type:timestamp;" json:"end_time"`
	Total      int        `gorm:"column:total;type:int(11);default:0" json:"total"`
	Claimed    int        `gorm:"column:claimed;type:int(11);default:0" json:"claimed"`
	Status     int8       `gorm:"column:status;type:tinyint;default:1" json:"status"`
	CreatedAt  *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Coupons) TableName() string {
	return "coupons"
}

func (t *Coupons) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// UserCoupons 用户领取的优惠券
type UserCoupons struct {
	Id        *uint64    `gorm:"column:id;type:bigint(20) UNSIGNED;primaryKey;autoIncrement;" json:"id"`
	UserId    uint64     `gorm:"column:user_id;type:bigint(20) UNSIGNED;not null;index" json:"user_id"`
	CouponId  uint64     `gorm:"column:coupon_id;type:bigint(20) UNSIGNED;not null;index" json:"coupon_id"`
	Status    string     `gorm:"column:status;type:varchar(20);default:unused" json:"status"`
	ClaimedAt *time.Time `gorm:"column:claimed_at;type:timestamp;" json:"claimed_at"`
	UsedAt    *time.Time `gorm:"column:used_at;type:timestamp;" json:"used_at"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (UserCoupons) TableName() string {
	return "user_coupons"
}

func (t *UserCoupons) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// Shops 店铺
type Shops struct {
	Id          *uint64    `gorm:"column:id;type:bigint(20) UNSIGNED;primaryKey;autoIncrement;" json:"id"`
	Name        string     `gorm:"column:name;type:varchar(100);not null" json:"name"`
	Description *string    `gorm:"column:description;type:text;" json:"description"`
	Avatar      *string    `gorm:"column:avatar;type:varchar(500);" json:"avatar"`
	CoverImage  *string    `gorm:"column:cover_image;type:varchar(500);" json:"cover_image"`
	Rating      float64    `gorm:"column:rating;type:decimal(3,2);default:5.00" json:"rating"`
	Sales       int        `gorm:"column:sales;type:int(11);default:0" json:"sales"`
	Fans        int        `gorm:"column:fans;type:int(11);default:0" json:"fans"`
	Location    *string    `gorm:"column:location;type:varchar(200);" json:"location"`
	CreatedAt   *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Shops) TableName() string {
	return "shops"
}

func (t *Shops) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}

// ShopFollows 店铺关注
type ShopFollows struct {
	Id        *uint64    `gorm:"column:id;type:bigint(20) UNSIGNED;primaryKey;autoIncrement;" json:"id"`
	UserId    uint64     `gorm:"column:user_id;type:bigint(20) UNSIGNED;not null;index" json:"user_id"`
	ShopId    uint64     `gorm:"column:shop_id;type:bigint(20) UNSIGNED;not null;index" json:"shop_id"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (ShopFollows) TableName() string {
	return "shop_follows"
}

func (t *ShopFollows) Jsonify() string {
	byts, err := json.Marshal(t)
	if err != nil {
		return `{"error": "` + err.Error() + `"}`
	}
	return string(byts)
}
