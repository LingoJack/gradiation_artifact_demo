package service

import (
	"time"

	"github.com/lingojack/taobao_clone/model"
	"gorm.io/gorm"
)

type ShopService struct {
	db *gorm.DB
}

func NewShopService(db *gorm.DB) *ShopService {
	return &ShopService{db: db}
}

type ShopWithProducts struct {
	model.Shops
	Products []ShopProduct `json:"products"`
}

type ShopProduct struct {
	model.Products
	CoverImage string `json:"coverImage"`
}

func (s *ShopService) GetDetail(shopID uint64) (*model.Shops, error) {
	var shop model.Shops
	if err := s.db.First(&shop, shopID).Error; err != nil {
		return nil, err
	}
	return &shop, nil
}

func (s *ShopService) GetShopProducts(shopID uint64, page, pageSize int) ([]ShopProduct, int64, error) {
	var total int64
	s.db.Model(&model.Products{}).Where("shop_id = ? AND status = 1", shopID).Count(&total)

	var products []model.Products
	offset := (page - 1) * pageSize
	if err := s.db.Where("shop_id = ? AND status = 1", shopID).
		Order("sales DESC").Offset(offset).Limit(pageSize).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	items := make([]ShopProduct, 0, len(products))
	for _, p := range products {
		coverImage := ""
		if p.MainImage != nil {
			coverImage = *p.MainImage
		}
		items = append(items, ShopProduct{
			Products:   p,
			CoverImage: coverImage,
		})
	}

	return items, total, nil
}

func (s *ShopService) ToggleFollow(userID uint64, shopID uint64, follow bool) error {
	var shop model.Shops
	if err := s.db.First(&shop, shopID).Error; err != nil {
		return err
	}

	if follow {
		// 关注（使用 FirstOrCreate 防止重复）
		now := time.Now()
		f := &model.ShopFollows{
			UserId:    userID,
			ShopId:    shopID,
			CreatedAt: &now,
		}
		s.db.Where("user_id = ? AND shop_id = ?", userID, shopID).FirstOrCreate(f)
		s.db.Model(&shop).UpdateColumn("fans", gorm.Expr("fans + 1"))
	} else {
		// 取消关注
		s.db.Where("user_id = ? AND shop_id = ?", userID, shopID).Delete(&model.ShopFollows{})
		s.db.Model(&shop).UpdateColumn("fans", gorm.Expr("GREATEST(fans - 1, 0)"))
	}

	return nil
}

func (s *ShopService) CheckFollow(userID uint64, shopID uint64) (bool, error) {
	var count int64
	s.db.Model(&model.ShopFollows{}).Where("user_id = ? AND shop_id = ?", userID, shopID).Count(&count)
	return count > 0, nil
}