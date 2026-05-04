package service

import (
	"context"
	"time"

	"github.com/lingojack/taobao_clone/model"
	"github.com/lingojack/taobao_clone/repository"
	"gorm.io/gorm"
)

const (
	maxSearchHistoryCount = 50
	maxKeywordLength      = 200
)

type SearchHistoryService struct {
	db  *gorm.DB
	dao *repository.SearchHistoriesDao
}

func NewSearchHistoryService(db *gorm.DB) *SearchHistoryService {
	return &SearchHistoryService{
		db:  db,
		dao: repository.NewSearchHistoriesDao(db),
	}
}

// SearchHistoryItem 搜索历史响应项
type SearchHistoryItem struct {
	ID        uint64    `json:"id"`
	Keyword   string    `json:"keyword"`
	CreatedAt time.Time `json:"created_at"`
}

// GetList 获取用户搜索历史（最多20条，倒序）
func (s *SearchHistoryService) GetList(ctx context.Context, userID uint64) ([]SearchHistoryItem, error) {
	list, err := s.dao.SelectList(ctx, &repository.SearchHistoriesQuery{
		UserId:    userID,
		OrderBy:   "created_at DESC",
		PageSize:  20,
		PageOffset: 0,
	})
	if err != nil {
		return nil, err
	}

	items := make([]SearchHistoryItem, 0, len(list))
	for _, record := range list {
		if record.CreatedAt != nil {
			items = append(items, SearchHistoryItem{
				ID:        *record.Id,
				Keyword:   record.Keyword,
				CreatedAt: *record.CreatedAt,
			})
		}
	}
	return items, nil
}

// Add 添加搜索记录（关键词去重更新时间）
func (s *SearchHistoryService) Add(ctx context.Context, userID uint64, keyword string) (*SearchHistoryItem, error) {
	now := time.Now()

	// 查找是否已存在相同关键词
	existingList, err := s.dao.SelectList(ctx, &repository.SearchHistoriesQuery{
		UserId:  userID,
		Keyword: keyword,
	})
	if err != nil {
		return nil, err
	}

	if len(existingList) > 0 {
		// 更新已存在记录的 created_at
		existing := existingList[0]
		existing.CreatedAt = &now
		if err := s.dao.InsertOrUpdateNullable(ctx, existing); err != nil {
			return nil, err
		}
		return &SearchHistoryItem{
			ID:        *existing.Id,
			Keyword:   existing.Keyword,
			CreatedAt: now,
		}, nil
	}

	// 插入新记录
	record := &model.SearchHistories{
		UserId:    userID,
		Keyword:   keyword,
		CreatedAt: &now,
	}
	if err := s.dao.Insert(ctx, record); err != nil {
		return nil, err
	}

	// 检查是否超过限制，删除最旧的
	s.cleanupOldRecords(ctx, userID)

	return &SearchHistoryItem{
		ID:        *record.Id,
		Keyword:   record.Keyword,
		CreatedAt: now,
	}, nil
}

// Delete 删除单条搜索历史（验证权限）
func (s *SearchHistoryService) Delete(ctx context.Context, userID uint64, id uint64) error {
	// 先查询记录
	list, err := s.dao.SelectList(ctx, &repository.SearchHistoriesQuery{
		Id: &id,
	})
	if err != nil {
		return err
	}
	if len(list) == 0 {
		return ErrNotFound
	}
	if list[0].UserId != userID {
		return ErrForbidden
	}

	return s.dao.WithContext(ctx).Delete(&model.SearchHistories{}, id).Error
}

// Clear 清空用户所有搜索历史
func (s *SearchHistoryService) Clear(ctx context.Context, userID uint64) error {
	return s.dao.DeleteByUserId(ctx, userID)
}

// cleanupOldRecords 清理超出限制的旧记录
func (s *SearchHistoryService) cleanupOldRecords(ctx context.Context, userID uint64) {
	count, err := s.dao.SelectCount(ctx, &repository.SearchHistoriesQuery{
		UserId: userID,
	})
	if err != nil {
		return
	}
	if count <= maxSearchHistoryCount {
		return
	}

	// 查询需要删除的记录（按时间正序，取超出部分）
	overflow := int(count - maxSearchHistoryCount)
	list, err := s.dao.SelectList(ctx, &repository.SearchHistoriesQuery{
		UserId:     userID,
		OrderBy:    "created_at ASC",
		PageSize:   overflow,
		PageOffset: 0,
	})
	if err != nil {
		return
	}
	for _, record := range list {
		if record.Id != nil {
			s.dao.WithContext(ctx).Delete(&model.SearchHistories{}, *record.Id)
		}
	}
}
