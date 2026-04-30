package service

import (
	"errors"
	"time"

	"github.com/lingojack/taobao_clone/model"
	"gorm.io/gorm"
)

type CouponService struct {
	db *gorm.DB
}

func NewCouponService(db *gorm.DB) *CouponService {
	return &CouponService{db: db}
}

type CouponWithDetail struct {
	model.Coupons
	Status    string  `json:"status"`
	ClaimedAt *string `json:"claimedAt,omitempty"`
	UsedAt    *string `json:"usedAt,omitempty"`
}

type UserCouponWithDetail struct {
	model.UserCoupons
	Coupon *model.Coupons `json:"coupon"`
}

func (s *CouponService) GetAvailableCoupons() ([]model.Coupons, error) {
	var coupons []model.Coupons
	now := time.Now()
	if err := s.db.Where("status = 1 AND start_time <= ? AND end_time >= ? AND claimed < total", now, now).
		Order("discount DESC").Find(&coupons).Error; err != nil {
		return nil, err
	}
	return coupons, nil
}

func (s *CouponService) GetUserCoupons(userID uint64) ([]UserCouponWithDetail, error) {
	var userCoupons []model.UserCoupons
	if err := s.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&userCoupons).Error; err != nil {
		return nil, err
	}

	result := make([]UserCouponWithDetail, 0, len(userCoupons))
	for _, uc := range userCoupons {
		item := UserCouponWithDetail{UserCoupons: uc}

		var coupon model.Coupons
		if err := s.db.First(&coupon, uc.CouponId).Error; err == nil {
			item.Coupon = &coupon
		}

		result = append(result, item)
	}

	return result, nil
}

func (s *CouponService) ClaimCoupon(userID uint64, couponID uint64) (*model.UserCoupons, error) {
	// 检查优惠券是否存在
	var coupon model.Coupons
	if err := s.db.First(&coupon, couponID).Error; err != nil {
		return nil, errors.New("优惠券不存在")
	}

	// 检查是否已领取
	var count int64
	s.db.Model(&model.UserCoupons{}).Where("user_id = ? AND coupon_id = ?", userID, couponID).Count(&count)
	if count > 0 {
		return nil, errors.New("已领取该优惠券")
	}

	// 检查库存
	if coupon.Claimed >= coupon.Total {
		return nil, errors.New("优惠券已被领完")
	}

	// 创建用户优惠券记录
	now := time.Now()
	userCoupon := &model.UserCoupons{
		UserId:    userID,
		CouponId:  couponID,
		Status:    "unused",
		ClaimedAt: &now,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	if err := s.db.Create(userCoupon).Error; err != nil {
		return nil, err
	}

	// 更新已领取数量
	s.db.Model(&coupon).UpdateColumn("claimed", gorm.Expr("claimed + 1"))

	return userCoupon, nil
}

func (s *CouponService) UseCoupon(userID uint64, userCouponID uint64) error {
	var uc model.UserCoupons
	if err := s.db.Where("id = ? AND user_id = ?", userCouponID, userID).First(&uc).Error; err != nil {
		return errors.New("优惠券不存在")
	}

	if uc.Status != "unused" {
		return errors.New("优惠券状态不可用")
	}

	now := time.Now()
	return s.db.Model(&uc).Updates(map[string]interface{}{
		"status":  "used",
		"used_at": now,
	}).Error
}
