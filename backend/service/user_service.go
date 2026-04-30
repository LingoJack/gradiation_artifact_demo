package service

import (
	"errors"
	"time"

	"github.com/lingojack/taobao_clone/model"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) GetProfile(userID uint64) (*model.Users, error) {
	var user model.Users
	if err := s.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

type UpdateProfileRequest struct {
	Nickname *string `json:"nickname"`
	Avatar   *string `json:"avatar"`
	Gender   *int    `json:"gender"`
	Birthday *string `json:"birthday"`
	Bio      *string `json:"bio"`
}

func (s *UserService) UpdateProfile(userID uint64, req *UpdateProfileRequest) (*model.Users, error) {
	updates := make(map[string]interface{})
	if req.Nickname != nil {
		updates["nickname"] = *req.Nickname
	}
	if req.Avatar != nil {
		updates["avatar"] = *req.Avatar
	}
	if req.Gender != nil {
		updates["gender"] = *req.Gender
	}
	if req.Birthday != nil {
		updates["birthday"] = *req.Birthday
	}
	if req.Bio != nil {
		updates["bio"] = *req.Bio
	}

	if len(updates) == 0 {
		return s.GetProfile(userID)
	}

	if err := s.db.Model(&model.Users{}).Where("id = ?", userID).Updates(updates).Error; err != nil {
		return nil, err
	}

	return s.GetProfile(userID)
}

// Address operations

func (s *UserService) GetAddresses(userID uint64) ([]model.UserAddresses, error) {
	var addresses []model.UserAddresses
	if err := s.db.Where("user_id = ?", userID).Order("is_default DESC, created_at DESC").Find(&addresses).Error; err != nil {
		return nil, err
	}
	return addresses, nil
}

func (s *UserService) CreateAddress(userID uint64, addr *model.UserAddresses) error {
	addr.UserId = userID

	// 如果设置为默认地址，先取消其他默认地址
	if addr.IsDefault != nil && *addr.IsDefault == 1 {
		s.db.Model(&model.UserAddresses{}).Where("user_id = ? AND is_default = 1", userID).Update("is_default", 0)
	}

	// 如果是第一个地址，自动设为默认
	var count int64
	s.db.Model(&model.UserAddresses{}).Where("user_id = ?", userID).Count(&count)
	if count == 0 {
		isDefault := int8(1)
		addr.IsDefault = &isDefault
	}

	now := time.Now()
	addr.CreatedAt = &now
	addr.UpdatedAt = &now

	return s.db.Create(addr).Error
}

func (s *UserService) UpdateAddress(userID uint64, addrID uint64, addr *model.UserAddresses) error {
	var existing model.UserAddresses
	if err := s.db.Where("id = ? AND user_id = ?", addrID, userID).First(&existing).Error; err != nil {
		return errors.New("地址不存在")
	}

	// 如果设置为默认地址
	if addr.IsDefault != nil && *addr.IsDefault == 1 {
		s.db.Model(&model.UserAddresses{}).Where("user_id = ? AND is_default = 1", userID).Update("is_default", 0)
	}

	updates := map[string]interface{}{
		"receiver_name":  addr.ReceiverName,
		"receiver_phone": addr.ReceiverPhone,
		"province":       addr.Province,
		"city":           addr.City,
		"district":       addr.District,
		"detail_address": addr.DetailAddress,
	}
	if addr.IsDefault != nil {
		updates["is_default"] = *addr.IsDefault
	}

	return s.db.Model(&existing).Updates(updates).Error
}

func (s *UserService) DeleteAddress(userID uint64, addrID uint64) error {
	result := s.db.Where("id = ? AND user_id = ?", addrID, userID).Delete(&model.UserAddresses{})
	if result.RowsAffected == 0 {
		return errors.New("地址不存在")
	}
	return nil
}

func (s *UserService) SetDefaultAddress(userID uint64, addrID uint64) error {
	var addr model.UserAddresses
	if err := s.db.Where("id = ? AND user_id = ?", addrID, userID).First(&addr).Error; err != nil {
		return errors.New("地址不存在")
	}

	// 取消当前默认
	s.db.Model(&model.UserAddresses{}).Where("user_id = ? AND is_default = 1", userID).Update("is_default", 0)
	// 设置新默认
	s.db.Model(&addr).Update("is_default", 1)
	return nil
}
