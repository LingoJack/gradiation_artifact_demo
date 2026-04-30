package service

import (
	"errors"
	"time"

	"github.com/lingojack/taobao_clone/model"
	"gorm.io/gorm"
)

type FavoriteService struct {
	db *gorm.DB
}

func NewFavoriteService(db *gorm.DB) *FavoriteService {
	return &FavoriteService{db: db}
}

type FavProductBrief struct {
	ID        uint64  `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	MainImage string  `json:"mainImage"`
	Sales     int     `json:"sales"`
}

type FavoriteWithProduct struct {
	model.UserFavorites
	Product *FavProductBrief `json:"product"`
}

func (s *FavoriteService) GetUserFavorites(userID uint64, page, pageSize int) ([]FavoriteWithProduct, int64, error) {
	var total int64
	s.db.Model(&model.UserFavorites{}).Where("user_id = ?", userID).Count(&total)

	var favorites []model.UserFavorites
	offset := (page - 1) * pageSize
	if err := s.db.Where("user_id = ?", userID).Order("created_at DESC").
		Offset(offset).Limit(pageSize).Find(&favorites).Error; err != nil {
		return nil, 0, err
	}

	result := make([]FavoriteWithProduct, 0, len(favorites))
	for _, fav := range favorites {
		item := FavoriteWithProduct{UserFavorites: fav}

		var product model.Products
		if err := s.db.Select("id, name, price, main_image, sales").First(&product, fav.ProductId).Error; err == nil {
			mainImage := ""
			if product.MainImage != nil {
				mainImage = *product.MainImage
			}
			sales := 0
			if product.Sales != nil {
				sales = *product.Sales
			}
			item.Product = &FavProductBrief{
				ID:        *product.Id,
				Name:      product.Name,
				Price:     product.Price,
				MainImage: mainImage,
				Sales:     sales,
			}
		}

		result = append(result, item)
	}

	return result, total, nil
}

func (s *FavoriteService) AddFavorite(userID uint64, productID uint64) error {
	// 检查商品是否存在
	var product model.Products
	if err := s.db.First(&product, productID).Error; err != nil {
		return errors.New("商品不存在")
	}

	// 检查是否已收藏
	var count int64
	s.db.Model(&model.UserFavorites{}).Where("user_id = ? AND product_id = ?", userID, productID).Count(&count)
	if count > 0 {
		return errors.New("已收藏该商品")
	}

	now := time.Now()
	favorite := &model.UserFavorites{
		UserId:    userID,
		ProductId: productID,
		CreatedAt: &now,
	}

	return s.db.Create(favorite).Error
}

func (s *FavoriteService) RemoveFavorite(userID uint64, productID uint64) error {
	result := s.db.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&model.UserFavorites{})
	if result.RowsAffected == 0 {
		return errors.New("未收藏该商品")
	}
	return nil
}

func (s *FavoriteService) CheckFavorite(userID uint64, productID uint64) (bool, error) {
	var count int64
	s.db.Model(&model.UserFavorites{}).Where("user_id = ? AND product_id = ?", userID, productID).Count(&count)
	return count > 0, nil
}
