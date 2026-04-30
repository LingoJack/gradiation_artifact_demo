package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/lingojack/taobao_clone/model"
	"github.com/lingojack/taobao_clone/model/query"
	"gorm.io/gorm"
)

// OrdersDao 订单表的Dao实现
type OrdersDao struct {
	*gorm.DB
}

func (dao *OrdersDao) Database() string {
	// TODO 补全 db 名称
	return "@database_name"
}

// NewOrdersDao 创建OrdersDao实例
// 参数:
//   - db: GORM数据库连接实例
//
// 返回:
//   - *OrdersDao: Dao实例
func NewOrdersDao(db *gorm.DB) *OrdersDao {
	return &OrdersDao{DB: db}
}

// ==================== 事务支持方法 ====================

// WithTx 使用指定的事务对象创建新的 DAO 实例
// 参数:
//   - tx: GORM事务对象
//
// 返回:
//   - *OrdersDao: 使用事务的新 DAO 实例
//
// 使用示例:
//
//	db.Transaction(func(tx *gorm.DB) error {
//	    txDao := dao.WithTx(tx)
//	    return txDao.Insert(ctx, poBean)
//	})
func (dao *OrdersDao) WithTx(tx *gorm.DB) *OrdersDao {
	return &OrdersDao{DB: tx}
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
//	err := dao.Transaction(ctx, func(txDao *OrdersDao) error {
//	    if err := txDao.Insert(ctx, poBean1); err != nil {
//	        return err
//	    }
//	    if err := txDao.Insert(ctx, poBean2); err != nil {
//	        return err
//	    }
//	    return nil
//	})
func (dao *OrdersDao) Transaction(ctx context.Context, fn func(*OrdersDao) error) error {
	return dao.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txDao := &OrdersDao{DB: tx}
		return fn(txDao)
	})
}

// ==================== 查询条件构建 ====================

// buildOrdersQueryCondition 构建查询条件
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
func (dao *OrdersDao) buildOrdersQueryCondition(db *gorm.DB, queryDto *query.OrdersDto) (*gorm.DB, error) {
	if queryDto == nil {
		return db, nil
	}

	// 基础字段精确查询
	if queryDto.Id != nil {
		db = db.Where("id = ?", queryDto.Id)
	}
	if queryDto.OrderNo != "" {
		db = db.Where("order_no = ?", queryDto.OrderNo)
	}
	if queryDto.UserId != 0 {
		db = db.Where("user_id = ?", queryDto.UserId)
	}
	if queryDto.TotalAmount != 0 {
		db = db.Where("total_amount = ?", queryDto.TotalAmount)
	}
	if queryDto.PayAmount != 0 {
		db = db.Where("pay_amount = ?", queryDto.PayAmount)
	}
	if queryDto.Status != "" {
		db = db.Where("status = ?", queryDto.Status)
	}
	if queryDto.ReceiverName != "" {
		db = db.Where("receiver_name = ?", queryDto.ReceiverName)
	}
	if queryDto.ReceiverPhone != "" {
		db = db.Where("receiver_phone = ?", queryDto.ReceiverPhone)
	}
	if queryDto.ReceiverAddress != "" {
		db = db.Where("receiver_address = ?", queryDto.ReceiverAddress)
	}
	if queryDto.Remark != nil && *queryDto.Remark != "" {
		db = db.Where("remark = ?", *queryDto.Remark)
	}
	if queryDto.PayTime != nil && !queryDto.PayTime.IsZero() {
		db = db.Where("pay_time = ?", *queryDto.PayTime)
	}
	if queryDto.DeliveryTime != nil && !queryDto.DeliveryTime.IsZero() {
		db = db.Where("delivery_time = ?", *queryDto.DeliveryTime)
	}
	if queryDto.ReceiveTime != nil && !queryDto.ReceiveTime.IsZero() {
		db = db.Where("receive_time = ?", *queryDto.ReceiveTime)
	}
	if queryDto.CreatedAt != nil && !queryDto.CreatedAt.IsZero() {
		db = db.Where("created_at = ?", *queryDto.CreatedAt)
	}
	if queryDto.UpdatedAt != nil && !queryDto.UpdatedAt.IsZero() {
		db = db.Where("updated_at = ?", *queryDto.UpdatedAt)
	}
	if queryDto.DeletedAt != nil && !queryDto.DeletedAt.IsZero() {
		db = db.Where("deleted_at = ?", *queryDto.DeletedAt)
	}

	// 模糊查询条件
	if queryDto.OrderNoFuzzy != "" {
		db = db.Where("order_no LIKE ?", "%"+queryDto.OrderNoFuzzy+"%")
	}
	if queryDto.StatusFuzzy != "" {
		db = db.Where("status LIKE ?", "%"+queryDto.StatusFuzzy+"%")
	}
	if queryDto.ReceiverNameFuzzy != "" {
		db = db.Where("receiver_name LIKE ?", "%"+queryDto.ReceiverNameFuzzy+"%")
	}
	if queryDto.ReceiverPhoneFuzzy != "" {
		db = db.Where("receiver_phone LIKE ?", "%"+queryDto.ReceiverPhoneFuzzy+"%")
	}
	if queryDto.ReceiverAddressFuzzy != "" {
		db = db.Where("receiver_address LIKE ?", "%"+queryDto.ReceiverAddressFuzzy+"%")
	}
	if queryDto.RemarkFuzzy != nil && *queryDto.RemarkFuzzy != "" {
		db = db.Where("remark LIKE ?", "%"+*queryDto.RemarkFuzzy+"%")
	}

	// 日期范围查询
	if queryDto.PayTimeStart != nil && !queryDto.PayTimeStart.IsZero() {
		db = db.Where("pay_time >= ?", queryDto.PayTimeStart)
	}
	if queryDto.PayTimeEnd != nil && !queryDto.PayTimeEnd.IsZero() {
		db = db.Where("pay_time < DATE_ADD(?, INTERVAL 1 DAY)", queryDto.PayTimeEnd)
	}
	if queryDto.DeliveryTimeStart != nil && !queryDto.DeliveryTimeStart.IsZero() {
		db = db.Where("delivery_time >= ?", queryDto.DeliveryTimeStart)
	}
	if queryDto.DeliveryTimeEnd != nil && !queryDto.DeliveryTimeEnd.IsZero() {
		db = db.Where("delivery_time < DATE_ADD(?, INTERVAL 1 DAY)", queryDto.DeliveryTimeEnd)
	}
	if queryDto.ReceiveTimeStart != nil && !queryDto.ReceiveTimeStart.IsZero() {
		db = db.Where("receive_time >= ?", queryDto.ReceiveTimeStart)
	}
	if queryDto.ReceiveTimeEnd != nil && !queryDto.ReceiveTimeEnd.IsZero() {
		db = db.Where("receive_time < DATE_ADD(?, INTERVAL 1 DAY)", queryDto.ReceiveTimeEnd)
	}
	if queryDto.CreatedAtStart != nil && !queryDto.CreatedAtStart.IsZero() {
		db = db.Where("created_at >= ?", queryDto.CreatedAtStart)
	}
	if queryDto.CreatedAtEnd != nil && !queryDto.CreatedAtEnd.IsZero() {
		db = db.Where("created_at < DATE_ADD(?, INTERVAL 1 DAY)", queryDto.CreatedAtEnd)
	}
	if queryDto.UpdatedAtStart != nil && !queryDto.UpdatedAtStart.IsZero() {
		db = db.Where("updated_at >= ?", queryDto.UpdatedAtStart)
	}
	if queryDto.UpdatedAtEnd != nil && !queryDto.UpdatedAtEnd.IsZero() {
		db = db.Where("updated_at < DATE_ADD(?, INTERVAL 1 DAY)", queryDto.UpdatedAtEnd)
	}
	if queryDto.DeletedAtStart != nil && !queryDto.DeletedAtStart.IsZero() {
		db = db.Where("deleted_at >= ?", queryDto.DeletedAtStart)
	}
	if queryDto.DeletedAtEnd != nil && !queryDto.DeletedAtEnd.IsZero() {
		db = db.Where("deleted_at < DATE_ADD(?, INTERVAL 1 DAY)", queryDto.DeletedAtEnd)
	}

	// IN 查询条件
	// 校验 OrderNoList: 如果不为 nil 但长度为 0，则报错
	if queryDto != nil && queryDto.OrderNoList != nil {
		if len(queryDto.OrderNoList) == 0 {
			return nil, fmt.Errorf("OrderNoList 不能为空列表")
		}
		db = db.Where("order_no IN ?", queryDto.OrderNoList)
	}
	// 校验 UserIdList: 如果不为 nil 但长度为 0，则报错
	if queryDto != nil && queryDto.UserIdList != nil {
		if len(queryDto.UserIdList) == 0 {
			return nil, fmt.Errorf("UserIdList 不能为空列表")
		}
		db = db.Where("user_id IN ?", queryDto.UserIdList)
	}
	// 校验 StatusList: 如果不为 nil 但长度为 0，则报错
	if queryDto != nil && queryDto.StatusList != nil {
		if len(queryDto.StatusList) == 0 {
			return nil, fmt.Errorf("StatusList 不能为空列表")
		}
		db = db.Where("status IN ?", queryDto.StatusList)
	}
	// 校验 CreatedAtList: 如果不为 nil 但长度为 0，则报错
	if queryDto != nil && queryDto.CreatedAtList != nil {
		if len(queryDto.CreatedAtList) == 0 {
			return nil, fmt.Errorf("CreatedAtList 不能为空列表")
		}
		db = db.Where("created_at IN ?", queryDto.CreatedAtList)
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
//   - [] *model.Orders: 查询结果列表
//   - error: 错误信息
//
// 说明:
//   - IN 查询条件校验规则:
//     1. List 为 nil: 不添加该查询条件（正常情况，表示不按此字段过滤）
//     2. List 不为 nil 且长度 > 0: 添加 IN 查询条件
//     3. List 不为 nil 但长度 = 0: 返回错误，因为空列表的 IN 查询没有意义，应提前发现此问题
func (dao *OrdersDao) SelectList(ctx context.Context, queryDto *query.OrdersDto) ([]*model.Orders, error) {
	var resultList []*model.Orders
	db := dao.WithContext(ctx).Model(&model.Orders{})

	// 应用查询条件
	var err error
	db, err = dao.buildOrdersQueryCondition(db, queryDto)
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
func (dao *OrdersDao) SelectCount(ctx context.Context, queryDto *query.OrdersDto) (int64, error) {
	var count int64
	db := dao.WithContext(ctx).Model(&model.Orders{})

	// 应用查询条件
	var err error
	db, err = dao.buildOrdersQueryCondition(db, queryDto)
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
//   - [] *model.Orders: 查询结果列表
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
func (dao *OrdersDao) SelectListWithAppendConditionFunc(ctx context.Context, queryDto *query.OrdersDto, appendConditionFunc func(ctx context.Context, db *gorm.DB) *gorm.DB) ([]*model.Orders, error) {
	var resultList []*model.Orders
	db := dao.WithContext(ctx).Model(&model.Orders{})

	// 应用查询条件
	var err error
	db, err = dao.buildOrdersQueryCondition(db, queryDto)
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
func (dao *OrdersDao) SelectCountWithAppendConditionFunc(ctx context.Context, queryDto *query.OrdersDto, appendConditionFunc func(ctx context.Context, db *gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	db := dao.WithContext(ctx).Model(&model.Orders{})

	// 应用查询条件
	var err error
	db, err = dao.buildOrdersQueryCondition(db, queryDto)
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
func (dao *OrdersDao) Insert(ctx context.Context, poBean *model.Orders) error {
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
func (dao *OrdersDao) InsertBatch(ctx context.Context, poBeanList []*model.Orders) error {
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
func (dao *OrdersDao) InsertOrUpdateNullable(ctx context.Context, poBean *model.Orders) error {
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
func (dao *OrdersDao) InsertOrUpdateBatchNullable(ctx context.Context, poBeanList []*model.Orders) error {
	if len(poBeanList) == 0 {
		return fmt.Errorf("批量插入或更新列表不能为空")
	}
	// 使用 GORM 的 Save 方法批量保存:
	// - 对每条记录根据主键判断是插入还是更新
	// - 更新时会覆盖所有字段（包括零值字段）
	// - 在一个事务中执行，保证原子性
	return dao.WithContext(ctx).Save(&poBeanList).Error
}

// ==================== 普通索引 idx_user_id 方法 ====================

// SelectByUserId 根据索引idx_user_id查询列表
// 参数:
//   - ctx: 上下文对象
//   - userId: 用户ID
//
// 返回:
//   - [] *model.Orders: 查询结果列表
//   - error: 错误信息
//
// 说明:
//   - 该索引不是唯一索引，可能返回多条记录
func (dao *OrdersDao) SelectByUserId(ctx context.Context, userId uint64) ([]*model.Orders, error) {
	var resultList []*model.Orders
	err := dao.WithContext(ctx).Where("user_id = ?", userId).Find(&resultList).Error
	return resultList, err
}

// SelectByUserIdList 根据索引idx_user_id批量查询列表
// 参数:
//   - ctx: 上下文对象
//   - userIdList: 用户ID列表
//
// 返回:
//   - [] *model.Orders: 查询结果列表
//   - error: 错误信息
func (dao *OrdersDao) SelectByUserIdList(ctx context.Context, userIdList []uint64) ([]*model.Orders, error) {
	if len(userIdList) == 0 {
		return []*model.Orders{}, nil
	}
	var resultList []*model.Orders
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
func (dao *OrdersDao) UpdateByUserId(ctx context.Context, poBean *model.Orders, userId uint64) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	return dao.WithContext(ctx).Model(&model.Orders{}).Where("user_id = ?", userId).Updates(poBean).Error
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
func (dao *OrdersDao) UpdateByUserIdWithMap(ctx context.Context, userId uint64, updatedMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	return dao.WithContext(ctx).Model(&model.Orders{}).Where("user_id = ?", userId).Updates(updatedMap).Error
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
func (dao *OrdersDao) UpdateByUserIdWithCondition(ctx context.Context, poBean *model.Orders, userId uint64, conditionMap map[string]interface{}) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	db := dao.WithContext(ctx).Model(&model.Orders{}).Where("user_id = ?", userId)

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
func (dao *OrdersDao) UpdateByUserIdWithMapAndCondition(ctx context.Context, userId uint64, updatedMap map[string]interface{}, conditionMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	db := dao.WithContext(ctx).Model(&model.Orders{}).Where("user_id = ?", userId)

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
func (dao *OrdersDao) DeleteByUserId(ctx context.Context, userId uint64) error {
	return dao.WithContext(ctx).Where("user_id = ?", userId).Delete(&model.Orders{}).Error
}

// ==================== 普通索引 idx_order_no 方法 ====================

// SelectByOrderNo 根据索引idx_order_no查询列表
// 参数:
//   - ctx: 上下文对象
//   - orderNo: 订单号
//
// 返回:
//   - [] *model.Orders: 查询结果列表
//   - error: 错误信息
//
// 说明:
//   - 该索引不是唯一索引，可能返回多条记录
func (dao *OrdersDao) SelectByOrderNo(ctx context.Context, orderNo string) ([]*model.Orders, error) {
	var resultList []*model.Orders
	err := dao.WithContext(ctx).Where("order_no = ?", orderNo).Find(&resultList).Error
	return resultList, err
}

// SelectByOrderNoList 根据索引idx_order_no批量查询列表
// 参数:
//   - ctx: 上下文对象
//   - orderNoList: 订单号列表
//
// 返回:
//   - [] *model.Orders: 查询结果列表
//   - error: 错误信息
func (dao *OrdersDao) SelectByOrderNoList(ctx context.Context, orderNoList []string) ([]*model.Orders, error) {
	if len(orderNoList) == 0 {
		return []*model.Orders{}, nil
	}
	var resultList []*model.Orders
	err := dao.WithContext(ctx).Where("order_no IN ?", orderNoList).Find(&resultList).Error
	return resultList, err
}

// UpdateByOrderNo 根据索引idx_order_no更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - orderNo: 订单号
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 只更新非零值字段，零值字段会被忽略
//   - 注意: 该索引不是唯一键，可能会更新多条记录
func (dao *OrdersDao) UpdateByOrderNo(ctx context.Context, poBean *model.Orders, orderNo string) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	return dao.WithContext(ctx).Model(&model.Orders{}).Where("order_no = ?", orderNo).Updates(poBean).Error
}

// UpdateByOrderNoWithMap 根据索引idx_order_no使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - orderNo: 订单号
//   - updatedMap: 要更新的字段Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 使用 map 可以显式指定要更新的字段，包括零值字段
//   - 只更新 map 中指定的字段，未指定的字段保持不变
//   - 注意: 该索引不是唯一键，可能会更新多条记录
func (dao *OrdersDao) UpdateByOrderNoWithMap(ctx context.Context, orderNo string, updatedMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	return dao.WithContext(ctx).Model(&model.Orders{}).Where("order_no = ?", orderNo).Updates(updatedMap).Error
}

// UpdateByOrderNoWithCondition 根据索引idx_order_no和额外条件更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - orderNo: 订单号
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 只更新非零值字段，零值字段会被忽略
//   - 适用场景: 需要在索引基础上增加额外的更新条件，缩小更新范围
//   - 注意: 可能会更新多条记录
func (dao *OrdersDao) UpdateByOrderNoWithCondition(ctx context.Context, poBean *model.Orders, orderNo string, conditionMap map[string]interface{}) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	db := dao.WithContext(ctx).Model(&model.Orders{}).Where("order_no = ?", orderNo)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(poBean).Error
}

// UpdateByOrderNoWithMapAndCondition 根据唯一索引idx_order_no和额外条件使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - orderNo: 订单号
//   - updatedMap: 要更新的字段Map
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 使用 map 可以显式指定要更新的字段，包括零值字段
//   - 提供最灵活的更新控制方式
func (dao *OrdersDao) UpdateByOrderNoWithMapAndCondition(ctx context.Context, orderNo string, updatedMap map[string]interface{}, conditionMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	db := dao.WithContext(ctx).Model(&model.Orders{}).Where("order_no = ?", orderNo)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(updatedMap).Error
}

// DeleteByOrderNo 根据索引idx_order_no删除
// 参数:
//   - ctx: 上下文对象
//   - orderNo: 订单号
//
// 返回:
//   - error: 错误信息
//
// 说明:
//   - 注意: 该索引不是唯一键，可能会删除多条记录
func (dao *OrdersDao) DeleteByOrderNo(ctx context.Context, orderNo string) error {
	return dao.WithContext(ctx).Where("order_no = ?", orderNo).Delete(&model.Orders{}).Error
}

// ==================== 普通索引 idx_status 方法 ====================

// SelectByStatus 根据索引idx_status查询列表
// 参数:
//   - ctx: 上下文对象
//   - status: 订单状态
//
// 返回:
//   - [] *model.Orders: 查询结果列表
//   - error: 错误信息
//
// 说明:
//   - 该索引不是唯一索引，可能返回多条记录
func (dao *OrdersDao) SelectByStatus(ctx context.Context, status string) ([]*model.Orders, error) {
	var resultList []*model.Orders
	err := dao.WithContext(ctx).Where("status = ?", status).Find(&resultList).Error
	return resultList, err
}

// SelectByStatusList 根据索引idx_status批量查询列表
// 参数:
//   - ctx: 上下文对象
//   - statusList: 订单状态列表
//
// 返回:
//   - [] *model.Orders: 查询结果列表
//   - error: 错误信息
func (dao *OrdersDao) SelectByStatusList(ctx context.Context, statusList []string) ([]*model.Orders, error) {
	if len(statusList) == 0 {
		return []*model.Orders{}, nil
	}
	var resultList []*model.Orders
	err := dao.WithContext(ctx).Where("status IN ?", statusList).Find(&resultList).Error
	return resultList, err
}

// UpdateByStatus 根据索引idx_status更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - status: 订单状态
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 只更新非零值字段，零值字段会被忽略
//   - 注意: 该索引不是唯一键，可能会更新多条记录
func (dao *OrdersDao) UpdateByStatus(ctx context.Context, poBean *model.Orders, status string) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	return dao.WithContext(ctx).Model(&model.Orders{}).Where("status = ?", status).Updates(poBean).Error
}

// UpdateByStatusWithMap 根据索引idx_status使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - status: 订单状态
//   - updatedMap: 要更新的字段Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 使用 map 可以显式指定要更新的字段，包括零值字段
//   - 只更新 map 中指定的字段，未指定的字段保持不变
//   - 注意: 该索引不是唯一键，可能会更新多条记录
func (dao *OrdersDao) UpdateByStatusWithMap(ctx context.Context, status string, updatedMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	return dao.WithContext(ctx).Model(&model.Orders{}).Where("status = ?", status).Updates(updatedMap).Error
}

// UpdateByStatusWithCondition 根据索引idx_status和额外条件更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - status: 订单状态
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 只更新非零值字段，零值字段会被忽略
//   - 适用场景: 需要在索引基础上增加额外的更新条件，缩小更新范围
//   - 注意: 可能会更新多条记录
func (dao *OrdersDao) UpdateByStatusWithCondition(ctx context.Context, poBean *model.Orders, status string, conditionMap map[string]interface{}) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	db := dao.WithContext(ctx).Model(&model.Orders{}).Where("status = ?", status)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(poBean).Error
}

// UpdateByStatusWithMapAndCondition 根据唯一索引idx_status和额外条件使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - status: 订单状态
//   - updatedMap: 要更新的字段Map
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 使用 map 可以显式指定要更新的字段，包括零值字段
//   - 提供最灵活的更新控制方式
func (dao *OrdersDao) UpdateByStatusWithMapAndCondition(ctx context.Context, status string, updatedMap map[string]interface{}, conditionMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	db := dao.WithContext(ctx).Model(&model.Orders{}).Where("status = ?", status)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(updatedMap).Error
}

// DeleteByStatus 根据索引idx_status删除
// 参数:
//   - ctx: 上下文对象
//   - status: 订单状态
//
// 返回:
//   - error: 错误信息
//
// 说明:
//   - 注意: 该索引不是唯一键，可能会删除多条记录
func (dao *OrdersDao) DeleteByStatus(ctx context.Context, status string) error {
	return dao.WithContext(ctx).Where("status = ?", status).Delete(&model.Orders{}).Error
}

// ==================== 普通索引 idx_created_at 方法 ====================

// SelectByCreatedAt 根据索引idx_created_at查询列表
// 参数:
//   - ctx: 上下文对象
//   - createdAt:
//
// 返回:
//   - [] *model.Orders: 查询结果列表
//   - error: 错误信息
//
// 说明:
//   - 该索引不是唯一索引，可能返回多条记录
func (dao *OrdersDao) SelectByCreatedAt(ctx context.Context, createdAt *time.Time) ([]*model.Orders, error) {
	var resultList []*model.Orders
	err := dao.WithContext(ctx).Where("created_at = ?", createdAt).Find(&resultList).Error
	return resultList, err
}

// SelectByCreatedAtList 根据索引idx_created_at批量查询列表
// 参数:
//   - ctx: 上下文对象
//   - createdAtList: 列表
//
// 返回:
//   - [] *model.Orders: 查询结果列表
//   - error: 错误信息
func (dao *OrdersDao) SelectByCreatedAtList(ctx context.Context, createdAtList []*time.Time) ([]*model.Orders, error) {
	if len(createdAtList) == 0 {
		return []*model.Orders{}, nil
	}
	var resultList []*model.Orders
	err := dao.WithContext(ctx).Where("created_at IN ?", createdAtList).Find(&resultList).Error
	return resultList, err
}

// UpdateByCreatedAt 根据索引idx_created_at更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - createdAt:
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 只更新非零值字段，零值字段会被忽略
//   - 注意: 该索引不是唯一键，可能会更新多条记录
func (dao *OrdersDao) UpdateByCreatedAt(ctx context.Context, poBean *model.Orders, createdAt *time.Time) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	return dao.WithContext(ctx).Model(&model.Orders{}).Where("created_at = ?", createdAt).Updates(poBean).Error
}

// UpdateByCreatedAtWithMap 根据索引idx_created_at使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - createdAt:
//   - updatedMap: 要更新的字段Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 使用 map 可以显式指定要更新的字段，包括零值字段
//   - 只更新 map 中指定的字段，未指定的字段保持不变
//   - 注意: 该索引不是唯一键，可能会更新多条记录
func (dao *OrdersDao) UpdateByCreatedAtWithMap(ctx context.Context, createdAt *time.Time, updatedMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	return dao.WithContext(ctx).Model(&model.Orders{}).Where("created_at = ?", createdAt).Updates(updatedMap).Error
}

// UpdateByCreatedAtWithCondition 根据索引idx_created_at和额外条件更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - createdAt:
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 只更新非零值字段，零值字段会被忽略
//   - 适用场景: 需要在索引基础上增加额外的更新条件，缩小更新范围
//   - 注意: 可能会更新多条记录
func (dao *OrdersDao) UpdateByCreatedAtWithCondition(ctx context.Context, poBean *model.Orders, createdAt *time.Time, conditionMap map[string]interface{}) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	db := dao.WithContext(ctx).Model(&model.Orders{}).Where("created_at = ?", createdAt)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(poBean).Error
}

// UpdateByCreatedAtWithMapAndCondition 根据唯一索引idx_created_at和额外条件使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - createdAt:
//   - updatedMap: 要更新的字段Map
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 使用 map 可以显式指定要更新的字段，包括零值字段
//   - 提供最灵活的更新控制方式
func (dao *OrdersDao) UpdateByCreatedAtWithMapAndCondition(ctx context.Context, createdAt *time.Time, updatedMap map[string]interface{}, conditionMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	db := dao.WithContext(ctx).Model(&model.Orders{}).Where("created_at = ?", createdAt)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(updatedMap).Error
}

// DeleteByCreatedAt 根据索引idx_created_at删除
// 参数:
//   - ctx: 上下文对象
//   - createdAt:
//
// 返回:
//   - error: 错误信息
//
// 说明:
//   - 注意: 该索引不是唯一键，可能会删除多条记录
func (dao *OrdersDao) DeleteByCreatedAt(ctx context.Context, createdAt *time.Time) error {
	return dao.WithContext(ctx).Where("created_at = ?", createdAt).Delete(&model.Orders{}).Error
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
//	var result entity.Orders
//	err := dao.ExecSql(ctx, &result, "SELECT * FROM orders WHERE id = ?", 1)
//	if err != nil {
//	    // 处理错误（包括记录不存在的情况）
//	    return err
//	}
//
//	// 示例2: 查询多条记录
//	var resultList [] *entity.Orders
//	err := dao.ExecSql(ctx, &resultList, "SELECT * FROM orders WHERE skill_id = ?", "skill123")
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
//	err := dao.ExecSql(ctx, &countList, "SELECT skill_id, COUNT(*) as count FROM orders GROUP BY skill_id")
//
// 注意事项:
//   - recvPtr 必须传指针，否则无法接收查询结果
//   - 查询单条记录时，如果返回多行，只会取第一行
//   - 查询多条记录时，如果没有结果，会返回空 slice（不是 nil）
//   - 结构体字段需要通过 gorm 标签与数据库列名匹配
//   - 如果取了别名，gorm 的 column 标签需要和 SQL 取的别名一致
//   - gorm 的 column 标签默认为下划线格式
func (dao *OrdersDao) ExecSql(ctx context.Context, recvPtr any, sql string, args ...any) error {
	return dao.WithContext(ctx).Raw(sql, args...).Scan(recvPtr).Error
}

// ==================== 辅助方法 ====================

// getValidOrderByFields 获取允许排序的字段白名单
// 返回:
//   - map[string] bool: 字段白名单，key为字段名，value为true表示允许排序
func (dao *OrdersDao) getValidOrderByFields() map[string]bool {
	return map[string]bool{
		"id":               true,
		"order_no":         true,
		"user_id":          true,
		"total_amount":     true,
		"pay_amount":       true,
		"status":           true,
		"receiver_name":    true,
		"receiver_phone":   true,
		"receiver_address": true,
		"remark":           true,
		"pay_time":         true,
		"delivery_time":    true,
		"receive_time":     true,
		"created_at":       true,
		"updated_at":       true,
		"deleted_at":       true,
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
func (dao *OrdersDao) isValidOrderBy(orderBy string) bool {
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
