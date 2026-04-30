package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/lingojack/taobao_clone/model"
	"github.com/lingojack/taobao_clone/model/query"
	"gorm.io/gorm"
)

// UsersDao 用户表的Dao实现
type UsersDao struct {
	*gorm.DB
}

func (dao *UsersDao) Database() string {
	// TODO 补全 db 名称
	return "@database_name"
}

// NewUsersDao 创建UsersDao实例
// 参数:
//   - db: GORM数据库连接实例
//
// 返回:
//   - *UsersDao: Dao实例
func NewUsersDao(db *gorm.DB) *UsersDao {
	return &UsersDao{DB: db}
}

// ==================== 事务支持方法 ====================

// WithTx 使用指定的事务对象创建新的 DAO 实例
// 参数:
//   - tx: GORM事务对象
//
// 返回:
//   - *UsersDao: 使用事务的新 DAO 实例
//
// 使用示例:
//
//	db.Transaction(func(tx *gorm.DB) error {
//	    txDao := dao.WithTx(tx)
//	    return txDao.Insert(ctx, poBean)
//	})
func (dao *UsersDao) WithTx(tx *gorm.DB) *UsersDao {
	return &UsersDao{DB: tx}
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
//	err := dao.Transaction(ctx, func(txDao *UsersDao) error {
//	    if err := txDao.Insert(ctx, poBean1); err != nil {
//	        return err
//	    }
//	    if err := txDao.Insert(ctx, poBean2); err != nil {
//	        return err
//	    }
//	    return nil
//	})
func (dao *UsersDao) Transaction(ctx context.Context, fn func(*UsersDao) error) error {
	return dao.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txDao := &UsersDao{DB: tx}
		return fn(txDao)
	})
}

// ==================== 查询条件构建 ====================

// buildUsersQueryCondition 构建查询条件
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
func (dao *UsersDao) buildUsersQueryCondition(db *gorm.DB, queryDto *query.UsersDto) (*gorm.DB, error) {
	if queryDto == nil {
		return db, nil
	}

	// 基础字段精确查询
	if queryDto.Id != nil {
		db = db.Where("id = ?", queryDto.Id)
	}
	if queryDto.Username != "" {
		db = db.Where("username = ?", queryDto.Username)
	}
	if queryDto.Password != "" {
		db = db.Where("password = ?", queryDto.Password)
	}
	if queryDto.Nickname != nil && *queryDto.Nickname != "" {
		db = db.Where("nickname = ?", *queryDto.Nickname)
	}
	if queryDto.Avatar != nil && *queryDto.Avatar != "" {
		db = db.Where("avatar = ?", *queryDto.Avatar)
	}
	if queryDto.Phone != nil && *queryDto.Phone != "" {
		db = db.Where("phone = ?", *queryDto.Phone)
	}
	if queryDto.Email != nil && *queryDto.Email != "" {
		db = db.Where("email = ?", *queryDto.Email)
	}
	if queryDto.Gender != nil {
		db = db.Where("gender = ?", queryDto.Gender)
	}
	if queryDto.Birthday != nil && !queryDto.Birthday.IsZero() {
		db = db.Where("birthday = ?", *queryDto.Birthday)
	}
	if queryDto.Status != nil {
		db = db.Where("status = ?", queryDto.Status)
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
	if queryDto.UsernameFuzzy != "" {
		db = db.Where("username LIKE ?", "%"+queryDto.UsernameFuzzy+"%")
	}
	if queryDto.PasswordFuzzy != "" {
		db = db.Where("password LIKE ?", "%"+queryDto.PasswordFuzzy+"%")
	}
	if queryDto.NicknameFuzzy != nil && *queryDto.NicknameFuzzy != "" {
		db = db.Where("nickname LIKE ?", "%"+*queryDto.NicknameFuzzy+"%")
	}
	if queryDto.AvatarFuzzy != nil && *queryDto.AvatarFuzzy != "" {
		db = db.Where("avatar LIKE ?", "%"+*queryDto.AvatarFuzzy+"%")
	}
	if queryDto.PhoneFuzzy != nil && *queryDto.PhoneFuzzy != "" {
		db = db.Where("phone LIKE ?", "%"+*queryDto.PhoneFuzzy+"%")
	}
	if queryDto.EmailFuzzy != nil && *queryDto.EmailFuzzy != "" {
		db = db.Where("email LIKE ?", "%"+*queryDto.EmailFuzzy+"%")
	}

	// 日期范围查询
	if queryDto.BirthdayStart != nil && !queryDto.BirthdayStart.IsZero() {
		db = db.Where("birthday >= ?", queryDto.BirthdayStart)
	}
	if queryDto.BirthdayEnd != nil && !queryDto.BirthdayEnd.IsZero() {
		db = db.Where("birthday < DATE_ADD(?, INTERVAL 1 DAY)", queryDto.BirthdayEnd)
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
	// 校验 UsernameList: 如果不为 nil 但长度为 0，则报错
	if queryDto != nil && queryDto.UsernameList != nil {
		if len(queryDto.UsernameList) == 0 {
			return nil, fmt.Errorf("UsernameList 不能为空列表")
		}
		db = db.Where("username IN ?", queryDto.UsernameList)
	}
	// 校验 PhoneList: 如果不为 nil 但长度为 0，则报错
	if queryDto != nil && queryDto.PhoneList != nil {
		if len(queryDto.PhoneList) == 0 {
			return nil, fmt.Errorf("PhoneList 不能为空列表")
		}
		db = db.Where("phone IN ?", queryDto.PhoneList)
	}
	// 校验 EmailList: 如果不为 nil 但长度为 0，则报错
	if queryDto != nil && queryDto.EmailList != nil {
		if len(queryDto.EmailList) == 0 {
			return nil, fmt.Errorf("EmailList 不能为空列表")
		}
		db = db.Where("email IN ?", queryDto.EmailList)
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
//   - [] *model.Users: 查询结果列表
//   - error: 错误信息
//
// 说明:
//   - IN 查询条件校验规则:
//     1. List 为 nil: 不添加该查询条件（正常情况，表示不按此字段过滤）
//     2. List 不为 nil 且长度 > 0: 添加 IN 查询条件
//     3. List 不为 nil 但长度 = 0: 返回错误，因为空列表的 IN 查询没有意义，应提前发现此问题
func (dao *UsersDao) SelectList(ctx context.Context, queryDto *query.UsersDto) ([]*model.Users, error) {
	var resultList []*model.Users
	db := dao.WithContext(ctx).Model(&model.Users{})

	// 应用查询条件
	var err error
	db, err = dao.buildUsersQueryCondition(db, queryDto)
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
func (dao *UsersDao) SelectCount(ctx context.Context, queryDto *query.UsersDto) (int64, error) {
	var count int64
	db := dao.WithContext(ctx).Model(&model.Users{})

	// 应用查询条件
	var err error
	db, err = dao.buildUsersQueryCondition(db, queryDto)
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
//   - [] *model.Users: 查询结果列表
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
func (dao *UsersDao) SelectListWithAppendConditionFunc(ctx context.Context, queryDto *query.UsersDto, appendConditionFunc func(ctx context.Context, db *gorm.DB) *gorm.DB) ([]*model.Users, error) {
	var resultList []*model.Users
	db := dao.WithContext(ctx).Model(&model.Users{})

	// 应用查询条件
	var err error
	db, err = dao.buildUsersQueryCondition(db, queryDto)
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
func (dao *UsersDao) SelectCountWithAppendConditionFunc(ctx context.Context, queryDto *query.UsersDto, appendConditionFunc func(ctx context.Context, db *gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	db := dao.WithContext(ctx).Model(&model.Users{})

	// 应用查询条件
	var err error
	db, err = dao.buildUsersQueryCondition(db, queryDto)
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
func (dao *UsersDao) Insert(ctx context.Context, poBean *model.Users) error {
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
func (dao *UsersDao) InsertBatch(ctx context.Context, poBeanList []*model.Users) error {
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
func (dao *UsersDao) InsertOrUpdateNullable(ctx context.Context, poBean *model.Users) error {
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
func (dao *UsersDao) InsertOrUpdateBatchNullable(ctx context.Context, poBeanList []*model.Users) error {
	if len(poBeanList) == 0 {
		return fmt.Errorf("批量插入或更新列表不能为空")
	}
	// 使用 GORM 的 Save 方法批量保存:
	// - 对每条记录根据主键判断是插入还是更新
	// - 更新时会覆盖所有字段（包括零值字段）
	// - 在一个事务中执行，保证原子性
	return dao.WithContext(ctx).Save(&poBeanList).Error
}

// ==================== 普通索引 idx_username 方法 ====================

// SelectByUsername 根据索引idx_username查询列表
// 参数:
//   - ctx: 上下文对象
//   - username: 用户名
//
// 返回:
//   - [] *model.Users: 查询结果列表
//   - error: 错误信息
//
// 说明:
//   - 该索引不是唯一索引，可能返回多条记录
func (dao *UsersDao) SelectByUsername(ctx context.Context, username string) ([]*model.Users, error) {
	var resultList []*model.Users
	err := dao.WithContext(ctx).Where("username = ?", username).Find(&resultList).Error
	return resultList, err
}

// SelectByUsernameList 根据索引idx_username批量查询列表
// 参数:
//   - ctx: 上下文对象
//   - usernameList: 用户名列表
//
// 返回:
//   - [] *model.Users: 查询结果列表
//   - error: 错误信息
func (dao *UsersDao) SelectByUsernameList(ctx context.Context, usernameList []string) ([]*model.Users, error) {
	if len(usernameList) == 0 {
		return []*model.Users{}, nil
	}
	var resultList []*model.Users
	err := dao.WithContext(ctx).Where("username IN ?", usernameList).Find(&resultList).Error
	return resultList, err
}

// UpdateByUsername 根据索引idx_username更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - username: 用户名
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 只更新非零值字段，零值字段会被忽略
//   - 注意: 该索引不是唯一键，可能会更新多条记录
func (dao *UsersDao) UpdateByUsername(ctx context.Context, poBean *model.Users, username string) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	return dao.WithContext(ctx).Model(&model.Users{}).Where("username = ?", username).Updates(poBean).Error
}

// UpdateByUsernameWithMap 根据索引idx_username使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - username: 用户名
//   - updatedMap: 要更新的字段Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 使用 map 可以显式指定要更新的字段，包括零值字段
//   - 只更新 map 中指定的字段，未指定的字段保持不变
//   - 注意: 该索引不是唯一键，可能会更新多条记录
func (dao *UsersDao) UpdateByUsernameWithMap(ctx context.Context, username string, updatedMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	return dao.WithContext(ctx).Model(&model.Users{}).Where("username = ?", username).Updates(updatedMap).Error
}

// UpdateByUsernameWithCondition 根据索引idx_username和额外条件更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - username: 用户名
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 只更新非零值字段，零值字段会被忽略
//   - 适用场景: 需要在索引基础上增加额外的更新条件，缩小更新范围
//   - 注意: 可能会更新多条记录
func (dao *UsersDao) UpdateByUsernameWithCondition(ctx context.Context, poBean *model.Users, username string, conditionMap map[string]interface{}) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	db := dao.WithContext(ctx).Model(&model.Users{}).Where("username = ?", username)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(poBean).Error
}

// UpdateByUsernameWithMapAndCondition 根据唯一索引idx_username和额外条件使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - username: 用户名
//   - updatedMap: 要更新的字段Map
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 使用 map 可以显式指定要更新的字段，包括零值字段
//   - 提供最灵活的更新控制方式
func (dao *UsersDao) UpdateByUsernameWithMapAndCondition(ctx context.Context, username string, updatedMap map[string]interface{}, conditionMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	db := dao.WithContext(ctx).Model(&model.Users{}).Where("username = ?", username)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(updatedMap).Error
}

// DeleteByUsername 根据索引idx_username删除
// 参数:
//   - ctx: 上下文对象
//   - username: 用户名
//
// 返回:
//   - error: 错误信息
//
// 说明:
//   - 注意: 该索引不是唯一键，可能会删除多条记录
func (dao *UsersDao) DeleteByUsername(ctx context.Context, username string) error {
	return dao.WithContext(ctx).Where("username = ?", username).Delete(&model.Users{}).Error
}

// ==================== 普通索引 idx_phone 方法 ====================

// SelectByPhone 根据索引idx_phone查询列表
// 参数:
//   - ctx: 上下文对象
//   - phone: 手机号
//
// 返回:
//   - [] *model.Users: 查询结果列表
//   - error: 错误信息
//
// 说明:
//   - 该索引不是唯一索引，可能返回多条记录
func (dao *UsersDao) SelectByPhone(ctx context.Context, phone *string) ([]*model.Users, error) {
	var resultList []*model.Users
	err := dao.WithContext(ctx).Where("phone = ?", phone).Find(&resultList).Error
	return resultList, err
}

// SelectByPhoneList 根据索引idx_phone批量查询列表
// 参数:
//   - ctx: 上下文对象
//   - phoneList: 手机号列表
//
// 返回:
//   - [] *model.Users: 查询结果列表
//   - error: 错误信息
func (dao *UsersDao) SelectByPhoneList(ctx context.Context, phoneList []*string) ([]*model.Users, error) {
	if len(phoneList) == 0 {
		return []*model.Users{}, nil
	}
	var resultList []*model.Users
	err := dao.WithContext(ctx).Where("phone IN ?", phoneList).Find(&resultList).Error
	return resultList, err
}

// UpdateByPhone 根据索引idx_phone更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - phone: 手机号
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 只更新非零值字段，零值字段会被忽略
//   - 注意: 该索引不是唯一键，可能会更新多条记录
func (dao *UsersDao) UpdateByPhone(ctx context.Context, poBean *model.Users, phone *string) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	return dao.WithContext(ctx).Model(&model.Users{}).Where("phone = ?", phone).Updates(poBean).Error
}

// UpdateByPhoneWithMap 根据索引idx_phone使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - phone: 手机号
//   - updatedMap: 要更新的字段Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 使用 map 可以显式指定要更新的字段，包括零值字段
//   - 只更新 map 中指定的字段，未指定的字段保持不变
//   - 注意: 该索引不是唯一键，可能会更新多条记录
func (dao *UsersDao) UpdateByPhoneWithMap(ctx context.Context, phone *string, updatedMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	return dao.WithContext(ctx).Model(&model.Users{}).Where("phone = ?", phone).Updates(updatedMap).Error
}

// UpdateByPhoneWithCondition 根据索引idx_phone和额外条件更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - phone: 手机号
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 只更新非零值字段，零值字段会被忽略
//   - 适用场景: 需要在索引基础上增加额外的更新条件，缩小更新范围
//   - 注意: 可能会更新多条记录
func (dao *UsersDao) UpdateByPhoneWithCondition(ctx context.Context, poBean *model.Users, phone *string, conditionMap map[string]interface{}) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	db := dao.WithContext(ctx).Model(&model.Users{}).Where("phone = ?", phone)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(poBean).Error
}

// UpdateByPhoneWithMapAndCondition 根据唯一索引idx_phone和额外条件使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - phone: 手机号
//   - updatedMap: 要更新的字段Map
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 使用 map 可以显式指定要更新的字段，包括零值字段
//   - 提供最灵活的更新控制方式
func (dao *UsersDao) UpdateByPhoneWithMapAndCondition(ctx context.Context, phone *string, updatedMap map[string]interface{}, conditionMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	db := dao.WithContext(ctx).Model(&model.Users{}).Where("phone = ?", phone)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(updatedMap).Error
}

// DeleteByPhone 根据索引idx_phone删除
// 参数:
//   - ctx: 上下文对象
//   - phone: 手机号
//
// 返回:
//   - error: 错误信息
//
// 说明:
//   - 注意: 该索引不是唯一键，可能会删除多条记录
func (dao *UsersDao) DeleteByPhone(ctx context.Context, phone *string) error {
	return dao.WithContext(ctx).Where("phone = ?", phone).Delete(&model.Users{}).Error
}

// ==================== 普通索引 idx_email 方法 ====================

// SelectByEmail 根据索引idx_email查询列表
// 参数:
//   - ctx: 上下文对象
//   - email: 邮箱
//
// 返回:
//   - [] *model.Users: 查询结果列表
//   - error: 错误信息
//
// 说明:
//   - 该索引不是唯一索引，可能返回多条记录
func (dao *UsersDao) SelectByEmail(ctx context.Context, email *string) ([]*model.Users, error) {
	var resultList []*model.Users
	err := dao.WithContext(ctx).Where("email = ?", email).Find(&resultList).Error
	return resultList, err
}

// SelectByEmailList 根据索引idx_email批量查询列表
// 参数:
//   - ctx: 上下文对象
//   - emailList: 邮箱列表
//
// 返回:
//   - [] *model.Users: 查询结果列表
//   - error: 错误信息
func (dao *UsersDao) SelectByEmailList(ctx context.Context, emailList []*string) ([]*model.Users, error) {
	if len(emailList) == 0 {
		return []*model.Users{}, nil
	}
	var resultList []*model.Users
	err := dao.WithContext(ctx).Where("email IN ?", emailList).Find(&resultList).Error
	return resultList, err
}

// UpdateByEmail 根据索引idx_email更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - email: 邮箱
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 只更新非零值字段，零值字段会被忽略
//   - 注意: 该索引不是唯一键，可能会更新多条记录
func (dao *UsersDao) UpdateByEmail(ctx context.Context, poBean *model.Users, email *string) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	return dao.WithContext(ctx).Model(&model.Users{}).Where("email = ?", email).Updates(poBean).Error
}

// UpdateByEmailWithMap 根据索引idx_email使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - email: 邮箱
//   - updatedMap: 要更新的字段Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 使用 map 可以显式指定要更新的字段，包括零值字段
//   - 只更新 map 中指定的字段，未指定的字段保持不变
//   - 注意: 该索引不是唯一键，可能会更新多条记录
func (dao *UsersDao) UpdateByEmailWithMap(ctx context.Context, email *string, updatedMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	return dao.WithContext(ctx).Model(&model.Users{}).Where("email = ?", email).Updates(updatedMap).Error
}

// UpdateByEmailWithCondition 根据索引idx_email和额外条件更新（不会用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - poBean: 包含更新数据的PO对象
//   - email: 邮箱
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 只更新非零值字段，零值字段会被忽略
//   - 适用场景: 需要在索引基础上增加额外的更新条件，缩小更新范围
//   - 注意: 可能会更新多条记录
func (dao *UsersDao) UpdateByEmailWithCondition(ctx context.Context, poBean *model.Users, email *string, conditionMap map[string]interface{}) error {
	if poBean == nil {
		return fmt.Errorf("更新对象不能为空")
	}
	db := dao.WithContext(ctx).Model(&model.Users{}).Where("email = ?", email)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(poBean).Error
}

// UpdateByEmailWithMapAndCondition 根据唯一索引idx_email和额外条件使用Map更新指定字段（可以用零值覆盖）
// 参数:
//   - ctx: 上下文对象
//   - email: 邮箱
//   - updatedMap: 要更新的字段Map
//   - conditionMap: 额外的查询条件Map
//
// 返回:
//   - error: 错误信息
//
// 行为说明:
//   - 使用 map 可以显式指定要更新的字段，包括零值字段
//   - 提供最灵活的更新控制方式
func (dao *UsersDao) UpdateByEmailWithMapAndCondition(ctx context.Context, email *string, updatedMap map[string]interface{}, conditionMap map[string]interface{}) error {
	if len(updatedMap) == 0 {
		return fmt.Errorf("更新字段不能为空")
	}
	db := dao.WithContext(ctx).Model(&model.Users{}).Where("email = ?", email)

	// 应用额外的条件
	for key, value := range conditionMap {
		db = db.Where(key+" = ?", value)
	}

	return db.Updates(updatedMap).Error
}

// DeleteByEmail 根据索引idx_email删除
// 参数:
//   - ctx: 上下文对象
//   - email: 邮箱
//
// 返回:
//   - error: 错误信息
//
// 说明:
//   - 注意: 该索引不是唯一键，可能会删除多条记录
func (dao *UsersDao) DeleteByEmail(ctx context.Context, email *string) error {
	return dao.WithContext(ctx).Where("email = ?", email).Delete(&model.Users{}).Error
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
//	var result entity.Users
//	err := dao.ExecSql(ctx, &result, "SELECT * FROM users WHERE id = ?", 1)
//	if err != nil {
//	    // 处理错误（包括记录不存在的情况）
//	    return err
//	}
//
//	// 示例2: 查询多条记录
//	var resultList [] *entity.Users
//	err := dao.ExecSql(ctx, &resultList, "SELECT * FROM users WHERE skill_id = ?", "skill123")
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
//	err := dao.ExecSql(ctx, &countList, "SELECT skill_id, COUNT(*) as count FROM users GROUP BY skill_id")
//
// 注意事项:
//   - recvPtr 必须传指针，否则无法接收查询结果
//   - 查询单条记录时，如果返回多行，只会取第一行
//   - 查询多条记录时，如果没有结果，会返回空 slice（不是 nil）
//   - 结构体字段需要通过 gorm 标签与数据库列名匹配
//   - 如果取了别名，gorm 的 column 标签需要和 SQL 取的别名一致
//   - gorm 的 column 标签默认为下划线格式
func (dao *UsersDao) ExecSql(ctx context.Context, recvPtr any, sql string, args ...any) error {
	return dao.WithContext(ctx).Raw(sql, args...).Scan(recvPtr).Error
}

// ==================== 辅助方法 ====================

// getValidOrderByFields 获取允许排序的字段白名单
// 返回:
//   - map[string] bool: 字段白名单，key为字段名，value为true表示允许排序
func (dao *UsersDao) getValidOrderByFields() map[string]bool {
	return map[string]bool{
		"id":         true,
		"username":   true,
		"password":   true,
		"nickname":   true,
		"avatar":     true,
		"phone":      true,
		"email":      true,
		"gender":     true,
		"birthday":   true,
		"status":     true,
		"created_at": true,
		"updated_at": true,
		"deleted_at": true,
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
func (dao *UsersDao) isValidOrderBy(orderBy string) bool {
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
