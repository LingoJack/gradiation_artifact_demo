package service

import (
	"strconv"

	"github.com/lingojack/taobao_clone/model"
	"gorm.io/gorm"
)

type ProductService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{db: db}
}

type ProductListResult struct {
	Products []ProductWithCover `json:"products"`
	Total    int64              `json:"total"`
}

type ProductWithCover struct {
	model.Products
	CoverImage string `json:"coverImage"`
}

type ProductDetail struct {
	model.Products
	Skus    []model.ProductSkus `json:"skus"`
	Reviews []model.Reviews     `json:"reviews"`
}

func (s *ProductService) GetProducts(categoryID uint64, keyword string, page, pageSize int, sort string) (*ProductListResult, error) {
	query := s.db.Model(&model.Products{}).Where("status = 1")

	if categoryID > 0 {
		// 获取该分类及所有子分类
		var catIDs []uint64
		s.db.Model(&model.Categories{}).Where("parent_id = ?", categoryID).Pluck("id", &catIDs)
		catIDs = append(catIDs, categoryID)
		query = query.Where("category_id IN ?", catIDs)
	}

	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
	}

	var total int64
	query.Count(&total)

	switch sort {
	case "price_asc":
		query = query.Order("price ASC")
	case "price_desc":
		query = query.Order("price DESC")
	case "sales":
		query = query.Order("sales DESC")
	default:
		query = query.Order("created_at DESC")
	}

	var products []model.Products
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&products).Error; err != nil {
		return nil, err
	}

	items := make([]ProductWithCover, 0, len(products))
	for _, p := range products {
		coverImage := ""
		if p.MainImage != nil {
			coverImage = *p.MainImage
		}
		items = append(items, ProductWithCover{
			Products:   p,
			CoverImage: coverImage,
		})
	}

	return &ProductListResult{Products: items, Total: total}, nil
}

func (s *ProductService) GetProductDetail(productID uint64) (*ProductDetail, error) {
	var product model.Products
	if err := s.db.First(&product, productID).Error; err != nil {
		return nil, err
	}

	result := &ProductDetail{Products: product}

	// 加载 SKU
	var skus []model.ProductSkus
	s.db.Where("product_id = ?", productID).Find(&skus)
	result.Skus = skus

	// 加载评价（最新5条）
	var reviews []model.Reviews
	s.db.Where("product_id = ?", productID).Order("created_at DESC").Limit(5).Find(&reviews)
	result.Reviews = reviews

	return result, nil
}

func (s *ProductService) SearchSuggestions(keyword string) ([]string, error) {
	var products []model.Products
	if err := s.db.Where("name LIKE ? AND status = 1", "%"+keyword+"%").
		Select("name").Limit(10).Find(&products).Error; err != nil {
		return nil, err
	}

	names := make([]string, 0, len(products))
	for _, p := range products {
		names = append(names, p.Name)
	}
	return names, nil
}

func (s *ProductService) GetCategories() ([]model.Categories, error) {
	var categories []model.Categories
	if err := s.db.Where("status = 1").Order("sort_order ASC").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *ProductService) GetBanners() ([]model.Banners, error) {
	var banners []model.Banners
	if err := s.db.Where("status = 1").Order("sort_order ASC").Find(&banners).Error; err != nil {
		return nil, err
	}
	return banners, nil
}

// Helper: parse uint64 from string
func parseUint64(s string) uint64 {
	n, _ := strconv.ParseUint(s, 10, 64)
	return n
}
