package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/lingojack/taobao_clone/model"
	"github.com/lingojack/taobao_clone/model/query"
	"gorm.io/gorm"
)

// UserFavoritesDao 用户收藏表的Dao实现
type UserFavoritesDao struct {
	*gorm.DB
}

func (dao *UserFavoritesDao) Database() string {
	// TODO 补全 db 名称
	return "@database_name"
}

// NewUserFavoritesDao 创建UserFavoritesDao实例
// 参数:
//   - db: GORM数据库连接实例
//
// 返回:
//   - *UserFavoritesDao: Dao实例
func NewUserFavoritesDao(db *gorm.DB) *UserFavoritesDao {
	return &UserFavoritesDao{DB: db}
}

// ==================== 事务支持方法 ====================

// WithTx 使用指定的事务对象创建新的 DAO 实例
// 参数:
//   - tx: GORM事务对象
//
// 返回:
//   - *UserFavoritesDao: 使用事务的新 DAO 实例
//
// 使用示例:
//
//	db.Transaction(func(tx *gorm.DB) error {
//	    txDao := dao.WithTx(tx)
//	    return txDao.Insert(ctx, poBean)
//	})
func (dao *UserFavoritesDao) WithTx(tx *gorm.DB) *UserFavoritesDao {
	return &UserFavoritesDao{DB: tx}
}

// Transaction 在事务中执行操作
// 参数:
//   - ctx: 上下文对象
//   - fn: 事务处理函数，接收使用事务的 DAO 实例
//
// 返回:
//   - error: 错误信息
//
// 说明:
//   - 自动管理事务的开始、提交和回滚
//   - 如果 fn 返回 error，事务会自动回滚
//   - 如果 fn 执行成功，事务会自动提交
//
// 使用示例:
//
//	err := dao.Transaction(ctx, func(txDao *UserFavoritesDao) error {
//	    if err := txDao.Insert(ctx, poBean1); err != nil {
//	        return err
//	    }
//	    if err := txDao.Insert(ctx, poBean2); err != nil {
//	        return err
//	    }
//	    return nil
//	})
func (dao *UserFavoritesDao) Transaction(ctx context.Context, fn func(*UserFavoritesDao) error) error {
	return dao.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txDao := &UserFavoritesDao{DB: tx}
		return fn(txDao)
	})
}

// ==================== 查询条件构建 ====================

// buildUserFavoritesQueryCondition 构建查询条件
// 参数:
//   - db: GORM数据库连接实例
//   - queryDto: 查询条件Dto对象
//
// 返回:
//   - *gorm.DB: 应用了查询条件的数据库连接
//   - error: 错误信息
//
// 说明:
//   - 支持精确匹配、模糊查询、IN查询、范围查询等多种查询方式
//   - 零值字段会被忽略，不会作为查询条件
//   - IN 查询条件校验规则:
//     1. List 为 nil: 不添加该查询条件（正常情况，表示不按此字段过滤）
//     2. List 不为 nil 且长度 > 0: 添加 IN 查询条件
//     3. List 不为 nil 但长度 = 0: 返回错误，因为空列表的 IN 查询没有意义，应提前发现此问题
func (dao *UserFavoritesDao) buildUserFavoritesQueryCondition(db *gorm.DB, queryDto *query.UserFavoritesDto) (*gorm.DB, error) {
	if queryDto == nil {
		return db, nil
	}

	// 基础字段精确查询
	if queryDto.Id != nil {
		db = db.Where("id = ?", queryDto.Id)
	}
	if queryDto.UserId != 0 {
		db = db.Where("user_id = ?", queryDto.UserId)
	}
	if queryDto.ProductId != 0 {
		db = db.Where("product_id = ?", queryDto.ProductId)
	}
	if queryDto.CreatedAt != nil && !queryDto.CreatedAt.IsZero() {
		db = db.Where("created_at = ?", *queryDto.CreatedAt)
	}

	// 模糊查询条件

	// 日期范围查询
	if queryDto.CreatedAtStart != nil && !queryDto.CreatedAtStart.IsZero() {
		db = db.Where("created_at >= ?", queryDto.CreatedAtStart)
	}
	if queryDto.CreatedAtEnd != nil && !queryDto.CreatedAtEnd.IsZero() {
		db = db.Where("created_at < DATE_ADD(?, INTERVAL 1 DAY)", queryDto.CreatedAtEnd)
	}

	// IN 查询条件
	// 校验 UserIdList: 如果不为 nil 但长度为 0，则报错
	if queryDto != nil && queryDto.UserIdList != nil {
		if len(queryDto.UserIdList) == 0 {
			return nil, fmt.Errorf("UserIdList 不能为空列表")
		}
		db = db.Where("user_id IN ?", queryDto.UserIdList)
	}
	// 校验 ProductIdList: 如果不为 nil 但长度为 0，则报错
	if queryDto != nil && queryDto.ProductIdList != nil {
		if len(queryDto.ProductIdList) == 0 {
			return nil, fmt.Errorf("ProductIdList 不能为空列表")
		}
		db = db.Where("product_id IN ?", queryDto.ProductIdList)
	}

	return db, nil
}

// ==================== 基础查询方法 ====================

// SelectList 查询列表
// 参数:
//   - ctx: 上下文对象
//   - queryDto: 查询条件Dto对象，支持分页、排序、多条件查询
//
// 返回:
//   - [] *model.UserFavorites: 查询结果列表
//   - error: 错误信息
//
// 说明:
//   - IN 查询条件校验规则:
//     1. List 为 nil: 不添加该查询条件（正常情况，表示不按此字段过滤）
//     2. List 不为 nil 且长度 > 0: 添加 IN 查询条件
//     3. List 不为 nil 但长度 = 0: 返回错误，因为空列表的 IN 查询没有意义，应提前发现此问题
func (dao *UserFavoritesDao) SelectList(ctx context.Context, queryDto *query.UserFavoritesDto) ([]*model.UserFavorites, error) {
	var resultList []*model.UserFavorites
	db := dao.WithContext(ctx).Model(&model.UserFavorites{})

	// 应用查询条件
	var err error
	db, err = dao.buildUserFavoritesQueryCondition(db, queryDto)
	if err != nil {
		return nil, err
	}

	// 排序
	if queryDto != nil && queryDto.OrderBy != "" {
		if dao.isValidOrderBy(queryDto.OrderBy) {
			db = db.Order(queryDto.OrderBy)
		}
	}

	// 分页
	if queryDto != nil && queryDto.PageSize > 0 {
		db = db.Offset(queryDto.PageOffset * queryDto.PageSize).Limit(queryDto.PageSize)
	}

	err = db.Find(&resultList).Error
	return resultList, err
}

// SelectCount 查询数量
// 参数:
//   - ctx: 上下文对象
//   - queryDto: 查询条件Dto对象
//
// 返回:
//   - int64: 符合条件的记录数量
//   - error: 错误信息
//
// 说明:
//   - IN 查询条件校验规则:
//     1. List 为 nil: 不添加该查询条件（正常情况，表示不按此字段过滤）
//     2. List 不为 nil 且长度 > 0: 添加 IN 查询条件
//     3. List 不为 nil 但长度 = 0: 返回错误，因为空列表的 IN 查询没有意义，应提前发现此问题
func (dao *UserFavoritesDao) SelectCount(ctx context.Context, queryDto *query.UserFavoritesDto) (int64, error) {
	var count int64
	db := dao.WithContext(ctx).Model(&model.UserFavorites{})

	// 应用查询条件
	var err error
	db, err = dao.buildUserFavoritesQueryCondition(db, queryDto)
	if err != nil {
		return 0, err
	}

	err = db.Count(&count).Error
	return count, err
}

// SelectListWithAppendConditionFunc 查询列表（支持自定义条件函数）
// 参数:
//   - ctx: 上下文对象
//   - queryDto: 查询条件Dto对象，支持分页、排序、多条件查询
//   - appendConditionFunc: 自定义条件函数，用于添加额外的查询条件
//   - 如果为 nil，则不添加额外条件
//   - 函数签名: func(ctx context.Context, db *gorm.DB) *gorm.DB
//
// 返回:
//   - [] *model.UserFavorites: 查询结果列表
//   - error: 错误信息
//
// 使用示例:
//
//	// 示例1: 添加复杂的自定义条件
//	resultList, err := dao.SelectListWithAppendConditionFunc(ctx, queryDto, func(ctx context.Context, db *gorm.DB) *gorm.DB {
//	    return db.Where("status IN (?, ?)", "active", "pending").
//	              Where("created_at > ?", time.Now().AddDate(0, -1, 0))
//	})
//
//	// 示例2: 不添加额外条件
//	resultList, err := dao.SelectListWithAppendConditionFunc(ctx, queryDto, nil)
//
// 说明:
//   - 自定义条件函数在基础查询条件、排序和分页之后执行
//   - 适用于需要动态添加复杂查询条件的场景
//   - IN 查询条件校验规则同 SelectList 方法
func (dao *UserFavoritesDao) SelectListWithAppendConditionFunc(ctx context.Context, queryDto *query.UserFavoritesDto, appendConditionFunc func(ctx context.Context, db *gorm.DB) *gorm.DB) ([]*model.UserFavorites, error) {
	var resultList []*model.UserFavorites
	db := dao.WithContext(ctx).Model(&model.UserFavorites{})

	// 应用查询条件
	var err error
	db, err = dao.buildUserFavoritesQueryCondition(db, queryDto)
	if err != nil {
		return nil, err
	}

	// 排序
	if queryDto != nil && queryDto.OrderBy != "" {
		if dao.isValidOrderBy(queryDto.OrderBy) {
			db = db.Order(queryDto.OrderBy)
		}
	}

	// 分页
	if queryDto != nil && queryDto.PageSize > 0 {
		db = db.Offset(queryDto.PageOffset * queryDto.PageSize).Limit(queryDto.PageSize)
	}

	// 应用自定义条件函数
	if appendConditionFunc != nil {
		db = appendConditionFunc(ctx, db)
	}

	err = db.Find(&resultList).Error
	return resultList, err
}

// SelectCountWithAppendConditionFunc 查询数量（支持自定义条件函数）
// 参数:
//   - ctx: 上下文对象
//   - queryDto: 查询条件Dto对象
//   - appendConditionFunc: 自定义条件函数，用于添加额外的查询条件
//   - 如果为 nil，则不添加额外条件
//   - 函数签名: func(ctx context.Context, db *gorm.DB) *gorm.DB
//
// 返回:
//   - int64: 符合条件的记录数量
//   - error: 错误信息
//
// 使用示例:
//
//	// 示例1: 添加自定义条件统计
//	count, err := dao.SelectCountWithAppendConditionFunc(ctx, queryDto, func(ctx context.Context, db *gorm.DB) *gorm.DB {
//	    return db.Where("status = ?", "active")
//	})
//
//	// 示例2: 不添加额外条件
//	count, err := dao.SelectCountWithAppendConditionFunc(ctx, queryDto, nil)
//
// 说明:
//   - 自定义条件函数在基础查询条件之后执行
//   - 适用于需要动态添加复杂查询条件的统计场景
//   - IN 查询条件校验规则同 SelectCount 方法
func (dao *UserFavoritesDao) SelectCountWithAppendConditionFunc(ctx context.Context, queryDto *query.UserFavoritesDto, appendConditionFunc func(ctx context.Context, db *gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	db := dao.WithContext(ctx).Model(&model.UserFavorites{})

	// 应用查询条件
	var err error
	db, err = dao.buildUserFavoritesQueryCondition(db, queryDto)
	if err != nil {
		return 0, err
	}

	// 应用自定义条件函数
	if appendConditionFunc != nil {
		db = appendConditionFunc(ctx, db)
	}

	err = db.Count(&count).Error
	return count, err
}

// ==================== 基础插入方法 ====================

// Insert 单行插入
// 参数:
//   - ctx: 上下文对象
//   - poBean: 要插入的PO对象
//
// 返回:
//   - error: 错误信息
//
// 说明:
//   - 插入所有字段，包括零值字段
//   - 自增主键会在插入后自动填充到poBean中
func (dao *UserFavoritesDao) Insert(ctx context.Context, poBean *model.UserFavorites) error {
	if poBean == nil {
		return fmt.Errorf("插入对象不能为空")
	}
	return dao.WithContext(ctx).Create(poBean).Error
}

// InsertBatch 批量插入
// 参数:
//   - ctx: 上下文对象
//   - poBeanList: 要插入的PO对象列表
//
// 返回:
//   - error: 错误信息
//
// 说明:
//   - 批量插入所有记录，在一个事务中执行
//   - 自增主键会在插入后自动填充到各个poBean中
func (dao *UserFavoritesDao) InsertBatch(ctx context.Context, poBeanList []*model.UserFavorites) error {
	if len(poBeanList) == 0 {
		return fmt.Errorf("批量插入列表不能为空")
	}
	return dao.WithContext(ctx).Create(&poBeanList).Error
}

// InsertOrUpdateNullable 插入或更新（会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 要插入或更新的PO对象
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//  1. 如果记录不存在（根据主键判断），则执行插入操作
//  2. 如果记录已存在，则执行全字段更新操作
//  3. **重要**: 更新时会用传入对象的所有字段值覆盖数据库中的值，包括零值（nil、""、0、false等）
//     例如: 如果 poBean.Content = nil，会将数据库中的 content 字段更新为 NULL
//     例如: 如果 poBean.ArtifactName = ""，会将数据库中的 artifactName 字段更新为空字符串
//  4. 这种行为适用于需要"完整替换"记录的场景
//  5. 如果不希望零值覆盖数据库中的非零值，应使用 UpdateByXxx 等方法（内部使用 Updates）
func (dao *UserFavoritesDao) InsertOrUpdateNullable(ctx context.Context, poBean *model.UserFavorites) error {
	if poBean == nil {
		return fmt.Errorf("插入或更新对象不能为空")
	}
	// 使用 GORM 的 Save 方法:
	// - 根据主键判断记录是否存在
	// - 存在则更新所有字段（包括零值字段）
	// - 不存在则插入新记录
	return dao.WithContext(ctx).Save(poBean).Error
}

// InsertOrUpdateBatchNullable 批量插入或更新（会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBeanList: 要插入或更新的PO对象列表
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//  1. 对列表中的每条记录，根据主键判断是插入还是更新
//  2. 如果记录不存在，则执行插入操作
//  3. 如果记录已存在，则执行全字段更新操作
//  4. **重要**: 更新时会用传入对象的所有字段值覆盖数据库中的值，包括零值（nil、""、0、false等）
//     这意味着如果某个字段在传入对象中为零值，会将数据库中对应字段更新为零值
//  5. 批量操作在一个事务中执行，要么全部成功，要么全部失败
//  6. 适用场景: 需要完整替换多条记录的场景
//  7. 性能提示: 批量操作比逐条调用 InsertOrUpdateNullable 效率更高
//  8. 如果不希望零值覆盖，建议逐条调用 UpdateByXxx 等方法
func (dao *UserFavoritesDao) InsertOrUpdateBatchNullable(ctx context.Context, poBeanList []*model.UserFavorites) error {
	if len(poBeanList) == 0 {
		return fmt.Errorf("批量插入或更新列表不能为空")
	}
	// 使用 GORM 的 Save 方法批量保存:
	// - 对每条记录根据主键判断是插入还是更新
	// - 更新时会覆盖所有字段（包括零值字段）
	// - 在一个事务中执行，保证原子性
	return dao.WithContext(ctx).Save(&poBeanList).Error
}

// ==================== 唯一索引 uk_user_product 方法 ====================

// SelectByUserIdAndProductId 根据唯一索引uk_user_product查询单条记录
// 参数:
//   - ctx: 上下文对象
//   - userId: 用户ID
//   - productId: 商品ID
//
// 返回:
//   - *model.UserFavorites: 查询结果，如果不存在返回nil
//   - error: 错误信息
func (dao *UserFavoritesDao) SelectByUserIdAndProductId(ctx context.Context, userId uint64, productId uint64) (*model.UserFavorites, error) {
	var resultBean model.UserFavorites
	err := dao.WithContext(ctx).Where("user_id = ? AND product_id = ?", userId, productId).First(&resultBean).Error
	if err != nil {
		return nil, err
	}
	return &resultBean, nil
}

// UpdateByUserIdAndProductId 根据唯一索引uk_user_product更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - userId: 用户ID
//   - productId: 商品ID
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 只更新非零值字段，零值字段会被忽略
func (dao *UserFavoritesDao) UpdateByUserIdAndProductId(ctx context.Context, poBean *model.UserFavorites, userId uint64, productId uint64) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	return dao.WithContext(ctx).Model(&model.UserFavorites{}).Where("user_id = ? AND product_id = ?", userId, productId).Updates(poBean).Error
}

// UpdateByUserIdAndProductIdWithMap 根据唯一索引uk_user_product使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - userId: 用户ID
//   - productId: 商品ID
//   - updatedMap: 要更新的字段Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 使用 map 可以显式指定要更新的字段，包括零值字段
//   - 只更新 map 中指定的字段，未指定的字段保持不变
func (dao *UserFavoritesDao) UpdateByUserIdAndProductIdWithMap(ctx context.Context, userId uint64, productId uint64, updatedMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	return dao.WithContext(ctx).Model(&model.UserFavorites{}).Where("user_id = ? AND product_id = ?", userId, productId).Updates(updatedMap).Error
}

// UpdateByUserIdAndProductIdWithCondition 根据唯一索引uk_user_product和额外条件更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - userId: 用户ID
//   - productId: 商品ID
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 只更新非零值字段，零值字段会被忽略
//   - 适用场景: 需要在唯一键基础上增加额外的更新条件
func (dao *UserFavoritesDao) UpdateByUserIdAndProductIdWithCondition(ctx context.Context, poBean *model.UserFavorites, userId uint64, productId uint64, conditionMap map[string]interface{}) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	db := dao.WithContext(ctx).Model(&model.UserFavorites{}).Where("user_id = ? AND product_id = ?", userId, productId)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(poBean).Error
}

// UpdateByUserIdAndProductIdWithMapAndCondition 根据唯一索引uk_user_product和额外条件使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - userId: 用户ID
//   - productId: 商品ID
//   - updatedMap: 要更新的字段Map
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 使用 map 可以显式指定要更新的字段，包括零值字段
//   - 提供最灵活的更新控制方式
func (dao *UserFavoritesDao) UpdateByUserIdAndProductIdWithMapAndCondition(ctx context.Context, userId uint64, productId uint64, updatedMap map[string]interface{}, conditionMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	db := dao.WithContext(ctx).Model(&model.UserFavorites{}).Where("user_id = ? AND product_id = ?", userId, productId)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(updatedMap).Error
}

// DeleteByUserIdAndProductId 根据唯一索引uk_user_product删除
// 参数:
//   - ctx: 上下文对象
//   - userId: 用户ID
//   - productId: 商品ID
//
// 返回:
//   - error: 错误信息
func (dao *UserFavoritesDao) DeleteByUserIdAndProductId(ctx context.Context, userId uint64, productId uint64) error {
	return dao.WithContext(ctx).Where("user_id = ? AND product_id = ?", userId, productId).Delete(&model.UserFavorites{}).Error
}

// ==================== 普通索引 idx_user_id 方法 ====================

// SelectByUserId 根据索引idx_user_id查询列表
// 参数:
//   - ctx: 上下文对象
//   - userId: 用户ID
//
// 返回:
//   - [] *model.UserFavorites: 查询结果列表
//   - error: 错误信息
//
// 说明:
//   - 该索引不是唯一索引，可能返回多条记录
func (dao *UserFavoritesDao) SelectByUserId(ctx context.Context, userId uint64) ([]*model.UserFavorites, error) {
	var resultList []*model.UserFavorites
	err := dao.WithContext(ctx).Where("user_id = ?", userId).Find(&resultList).Error
	return resultList, err
}

// SelectByUserIdList 根据索引idx_user_id批量查询列表
// 参数:
//   - ctx: 上下文对象
//   - userIdList: 用户ID列表
//
// 返回:
//   - [] *model.UserFavorites: 查询结果列表
//   - error: 错误信息
func (dao *UserFavoritesDao) SelectByUserIdList(ctx context.Context, userIdList []uint64) ([]*model.UserFavorites, error) {
	if len(userIdList) == 0 {
		return []*model.UserFavorites{}, nil
	}
	var resultList []*model.UserFavorites
	err := dao.WithContext(ctx).Where("user_id IN ?", userIdList).Find(&resultList).Error
	return resultList, err
}

// UpdateByUserId 根据索引idx_user_id更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - userId: 用户ID
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 只更新非零值字段，零值字段会被忽略
//   - 注意: 该索引不是唯一键，可能会更新多条记录
func (dao *UserFavoritesDao) UpdateByUserId(ctx context.Context, poBean *model.UserFavorites, userId uint64) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	return dao.WithContext(ctx).Model(&model.UserFavorites{}).Where("user_id = ?", userId).Updates(poBean).Error
}

// UpdateByUserIdWithMap 根据索引idx_user_id使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - userId: 用户ID
//   - updatedMap: 要更新的字段Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 使用 map 可以显式指定要更新的字段，包括零值字段
//   - 只更新 map 中指定的字段，未指定的字段保持不变
//   - 注意: 该索引不是唯一键，可能会更新多条记录
func (dao *UserFavoritesDao) UpdateByUserIdWithMap(ctx context.Context, userId uint64, updatedMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	return dao.WithContext(ctx).Model(&model.UserFavorites{}).Where("user_id = ?", userId).Updates(updatedMap).Error
}

// UpdateByUserIdWithCondition 根据索引idx_user_id和额外条件更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - userId: 用户ID
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 只更新非零值字段，零值字段会被忽略
//   - 适用场景: 需要在索引基础上增加额外的更新条件，缩小更新范围
//   - 注意: 可能会更新多条记录
func (dao *UserFavoritesDao) UpdateByUserIdWithCondition(ctx context.Context, poBean *model.UserFavorites, userId uint64, conditionMap map[string]interface{}) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	db := dao.WithContext(ctx).Model(&model.UserFavorites{}).Where("user_id = ?", userId)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(poBean).Error
}

// UpdateByUserIdWithMapAndCondition 根据唯一索引idx_user_id和额外条件使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - userId: 用户ID
//   - updatedMap: 要更新的字段Map
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 使用 map 可以显式指定要更新的字段，包括零值字段
//   - 提供最灵活的更新控制方式
func (dao *UserFavoritesDao) UpdateByUserIdWithMapAndCondition(ctx context.Context, userId uint64, updatedMap map[string]interface{}, conditionMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	db := dao.WithContext(ctx).Model(&model.UserFavorites{}).Where("user_id = ?", userId)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(updatedMap).Error
}

// DeleteByUserId 根据索引idx_user_id删除
// 参数:
//   - ctx: 上下文对象
//   - userId: 用户ID
//
// 返回:
//   - error: 错误信息
//
// 说明:
//   - 注意: 该索引不是唯一键，可能会删除多条记录
func (dao *UserFavoritesDao) DeleteByUserId(ctx context.Context, userId uint64) error {
	return dao.WithContext(ctx).Where("user_id = ?", userId).Delete(&model.UserFavorites{}).Error
}

// ==================== 普通索引 idx_product_id 方法 ====================

// SelectByProductId 根据索引idx_product_id查询列表
// 参数:
//   - ctx: 上下文对象
//   - productId: 商品ID
//
// 返回:
//   - [] *model.UserFavorites: 查询结果列表
//   - error: 错误信息
//
// 说明:
//   - 该索引不是唯一索引，可能返回多条记录
func (dao *UserFavoritesDao) SelectByProductId(ctx context.Context, productId uint64) ([]*model.UserFavorites, error) {
	var resultList []*model.UserFavorites
	err := dao.WithContext(ctx).Where("product_id = ?", productId).Find(&resultList).Error
	return resultList, err
}

// SelectByProductIdList 根据索引idx_product_id批量查询列表
// 参数:
//   - ctx: 上下文对象
//   - productIdList: 商品ID列表
//
// 返回:
//   - [] *model.UserFavorites: 查询结果列表
//   - error: 错误信息
func (dao *UserFavoritesDao) SelectByProductIdList(ctx context.Context, productIdList []uint64) ([]*model.UserFavorites, error) {
	if len(productIdList) == 0 {
		return []*model.UserFavorites{}, nil
	}
	var resultList []*model.UserFavorites
	err := dao.WithContext(ctx).Where("product_id IN ?", productIdList).Find(&resultList).Error
	return resultList, err
}

// UpdateByProductId 根据索引idx_product_id更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - productId: 商品ID
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 只更新非零值字段，零值字段会被忽略
//   - 注意: 该索引不是唯一键，可能会更新多条记录
func (dao *UserFavoritesDao) UpdateByProductId(ctx context.Context, poBean *model.UserFavorites, productId uint64) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	return dao.WithContext(ctx).Model(&model.UserFavorites{}).Where("product_id = ?", productId).Updates(poBean).Error
}

// UpdateByProductIdWithMap 根据索引idx_product_id使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - productId: 商品ID
//   - updatedMap: 要更新的字段Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 使用 map 可以显式指定要更新的字段，包括零值字段
//   - 只更新 map 中指定的字段，未指定的字段保持不变
//   - 注意: 该索引不是唯一键，可能会更新多条记录
func (dao *UserFavoritesDao) UpdateByProductIdWithMap(ctx context.Context, productId uint64, updatedMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	return dao.WithContext(ctx).Model(&model.UserFavorites{}).Where("product_id = ?", productId).Updates(updatedMap).Error
}

// UpdateByProductIdWithCondition 根据索引idx_product_id和额外条件更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - productId: 商品ID
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 只更新非零值字段，零值字段会被忽略
//   - 适用场景: 需要在索引基础上增加额外的更新条件，缩小更新范围
//   - 注意: 可能会更新多条记录
func (dao *UserFavoritesDao) UpdateByProductIdWithCondition(ctx context.Context, poBean *model.UserFavorites, productId uint64, conditionMap map[string]interface{}) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	db := dao.WithContext(ctx).Model(&model.UserFavorites{}).Where("product_id = ?", productId)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(poBean).Error
}

// UpdateByProductIdWithMapAndCondition 根据唯一索引idx_product_id和额外条件使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - productId: 商品ID
//   - updatedMap: 要更新的字段Map
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 使用 map 可以显式指定要更新的字段，包括零值字段
//   - 提供最灵活的更新控制方式
func (dao *UserFavoritesDao) UpdateByProductIdWithMapAndCondition(ctx context.Context, productId uint64, updatedMap map[string]interface{}, conditionMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	db := dao.WithContext(ctx).Model(&model.UserFavorites{}).Where("product_id = ?", productId)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(updatedMap).Error
}

// DeleteByProductId 根据索引idx_product_id删除
// 参数:
//   - ctx: 上下文对象
//   - productId: 商品ID
//
// 返回:
//   - error: 错误信息
//
// 说明:
//   - 注意: 该索引不是唯一键，可能会删除多条记录
func (dao *UserFavoritesDao) DeleteByProductId(ctx context.Context, productId uint64) error {
	return dao.WithContext(ctx).Where("product_id = ?", productId).Delete(&model.UserFavorites{}).Error
}

// ==================== 原生SQL执行方法 ====================

// ExecSql 执行原生SQL查询
// 参数:
//   - ctx: 上下文对象
//   - recvPtr: 接收查询结果的指针（必须是指针类型）
//   - 查询单条记录时传入结构体指针，如 &entity.TNode{}
//   - 查询多条记录时传入 slice 指针，如 &[] *entity.TNode{}
//   - sql: SQL语句，支持占位符 ?
//   - args: SQL参数，按顺序对应 SQL 中的占位符
//
// 返回:
//   - error: 错误信息，查询失败或记录不存在时返回错误
//
// 使用示例:
//
//	// 示例1: 查询单条记录
//	var result entity.UserFavorites
//	err := dao.ExecSql(ctx, &result, "SELECT * FROM user_favorites WHERE id = ?", 1)
//	if err != nil {
//	    // 处理错误（包括记录不存在的情况）
//	    return err
//	}
//
//	// 示例2: 查询多条记录
//	var resultList [] *entity.UserFavorites
//	err := dao.ExecSql(ctx, &resultList, "SELECT * FROM user_favorites WHERE skill_id = ?", "skill123")
//	if err != nil {
//	    return err
//	}
//
//	// 示例3: 查询聚合结果
//	type CountResult struct {
//	    SkillId string `gorm:"column:skill_id"`
//	    Count   int64  `gorm:"column:count"`
//	}
//	var countList [] *CountResult
//	err := dao.ExecSql(ctx, &countList, "SELECT skill_id, COUNT(*) as count FROM user_favorites GROUP BY skill_id")
//
// 注意事项:
//   - recvPtr 必须传指针，否则无法接收查询结果
//   - 查询单条记录时，如果返回多行，只会取第一行
//   - 查询多条记录时，如果没有结果，会返回空 slice（不是 nil）
//   - 结构体字段需要通过 gorm 标签与数据库列名匹配
//   - 如果取了别名，gorm 的 column 标签需要和 SQL 取的别名一致
//   - gorm 的 column 标签默认为下划线格式
func (dao *UserFavoritesDao) ExecSql(ctx context.Context, recvPtr any, sql string, args ...any) error {
	return dao.WithContext(ctx).Raw(sql, args...).Scan(recvPtr).Error
}

// ==================== 辅助方法 ====================

// getValidOrderByFields 获取允许排序的字段白名单
// 返回:
//   - map[string] bool: 字段白名单，key为字段名，value为true表示允许排序
func (dao *UserFavoritesDao) getValidOrderByFields() map[string]bool {
	return map[string]bool{
		"id":         true,
		"user_id":    true,
		"product_id": true,
		"created_at": true,
	}
}

// isValidOrderBy 验证排序字符串是否安全（基于字段白名单）
// 支持格式:
//   - 单字段: id DESC
//   - 多字段: id DESC, createTime ASC
//
// 参数:
//   - orderBy: 排序字符串
//
// 返回:
//   - true: 排序字符串合法且所有字段都在白名单中
//   - false: 排序字符串不合法或包含非白名单字段
func (dao *UserFavoritesDao) isValidOrderBy(orderBy string) bool {
	if orderBy == "" {
		return false
	}

	// 获取字段白名单
	validFields := dao.getValidOrderByFields()

	// 按逗号分割多个排序字段
	orderParts := strings.Split(orderBy, ",")

	for _, part := range orderParts {
		part = strings.TrimSpace(part)
		if part == "" {
			return false
		}

		// 按空格分割字段名和排序方向
		tokens := strings.Fields(part)
		if len(tokens) == 0 || len(tokens) > 2 {
			// 格式错误: 必须是 "字段名" 或 "字段名 方向"
			return false
		}

		// 验证字段名是否在白名单中
		fieldName := tokens[0]
		if !validFields[fieldName] {
			// 字段不在白名单中
			return false
		}

		// 如果指定了排序方向，验证是否为 ASC 或 DESC
		if len(tokens) == 2 {
			direction := strings.ToUpper(tokens[1])
			if direction != "ASC" && direction != "DESC" {
				// 排序方向无效
				return false
			}
		}
	}

	return true
}
