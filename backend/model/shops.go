package model

import "time"

// Shop 店铺模型
type Shop struct {
	ID          uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string     `gorm:"type:varchar(100);not null" json:"name"`
	Description string     `gorm:"type:varchar(500)" json:"description"`
	Avatar      string     `gorm:"type:varchar(500)" json:"avatar"`
	CoverImage  string     `gorm:"column:cover_image;type:varchar(500)" json:"cover_image"`
	Rating      float64    `gorm:"type:decimal(2,1);default:5.0" json:"rating"`
	Sales       int        `gorm:"default:0" json:"sales"`
	Fans        int        `gorm:"default:0" json:"fans"`
	Location    string     `gorm:"type:varchar(100)" json:"location"`
	Status      int8       `gorm:"type:tinyint;default:1" json:"status"`
	CreatedAt   *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at;type:timestamp;default:NULL" json:"deleted_at"`
}

func (Shop) TableName() string {
	return "shops"
}
