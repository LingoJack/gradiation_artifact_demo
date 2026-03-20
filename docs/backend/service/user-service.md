# 用户服务 API 文档

## 概述

用户服务负责处理用户认证、个人信息管理、地址管理等核心功能。

## API 端点

### 1. 用户认证

#### POST /api/v1/auth/register

用户注册

**请求体：**
```json
{
  "username": "string",    // 用户名（必填，3-20字符）
  "password": "string",    // 密码（必填，6-20字符）
  "phone": "string",       // 手机号（可选）
  "email": "string"        // 邮箱（可选）
}
```

**响应：**
```json
{
  "code": 0,
  "message": "注册成功",
  "data": {
    "userId": 1,
    "username": "testuser",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

#### POST /api/v1/auth/login

用户登录

**请求体：**
```json
{
  "username": "string",  // 用户名或手机号
  "password": "string"   // 密码
}
```

**响应：**
```json
{
  "code": 0,
  "message": "登录成功",
  "data": {
    "userId": 1,
    "username": "testuser",
    "nickname": "测试用户",
    "avatar": "https://example.com/avatar.jpg",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

#### POST /api/v1/auth/logout

用户登出

**请求头：**
```
Authorization: Bearer <token>
```

**响应：**
```json
{
  "code": 0,
  "message": "登出成功"
}
```

### 2. 用户信息

#### GET /api/v1/users/profile

获取当前用户信息

**请求头：**
```
Authorization: Bearer <token>
```

**响应：**
```json
{
  "code": 0,
  "data": {
    "id": 1,
    "username": "testuser",
    "nickname": "测试用户",
    "avatar": "https://example.com/avatar.jpg",
    "phone": "138****8000",
    "email": "test@example.com",
    "gender": 1,
    "birthday": "1990-01-01",
    "createdAt": "2024-01-01T00:00:00Z"
  }
}
```

#### PUT /api/v1/users/profile

更新用户信息

**请求头：**
```
Authorization: Bearer <token>
```

**请求体：**
```json
{
  "nickname": "string",
  "avatar": "string",
  "gender": 1,
  "birthday": "1990-01-01"
}
```

**响应：**
```json
{
  "code": 0,
  "message": "更新成功",
  "data": {
    "id": 1,
    "nickname": "新昵称",
    "avatar": "https://example.com/new-avatar.jpg"
  }
}
```

### 3. 用户地址

#### GET /api/v1/users/addresses

获取用户地址列表

**请求头：**
```
Authorization: Bearer <token>
```

**响应：**
```json
{
  "code": 0,
  "data": [
    {
      "id": 1,
      "receiverName": "张三",
      "receiverPhone": "13800138000",
      "province": "北京市",
      "city": "北京市",
      "district": "朝阳区",
      "detailAddress": "某某街道某某小区1号楼",
      "isDefault": true
    }
  ]
}
```

#### POST /api/v1/users/addresses

创建新地址

**请求头：**
```
Authorization: Bearer <token>
```

**请求体：**
```json
{
  "receiverName": "string",      // 收货人姓名（必填）
  "receiverPhone": "string",     // 收货人电话（必填）
  "province": "string",          // 省（必填）
  "city": "string",              // 市（必填）
  "district": "string",          // 区（必填）
  "detailAddress": "string",     // 详细地址（必填）
  "isDefault": false             // 是否默认地址
}
```

**响应：**
```json
{
  "code": 0,
  "message": "创建成功",
  "data": {
    "id": 1,
    "receiverName": "张三",
    "receiverPhone": "13800138000",
    "province": "北京市",
    "city": "北京市",
    "district": "朝阳区",
    "detailAddress": "某某街道某某小区1号楼",
    "isDefault": true
  }
}
```

#### PUT /api/v1/users/addresses/:id

更新地址

**请求头：**
```
Authorization: Bearer <token>
```

**请求体：**
```json
{
  "receiverName": "string",
  "receiverPhone": "string",
  "province": "string",
  "city": "string",
  "district": "string",
  "detailAddress": "string",
  "isDefault": false
}
```

**响应：**
```json
{
  "code": 0,
  "message": "更新成功"
}
```

#### DELETE /api/v1/users/addresses/:id

删除地址

**请求头：**
```
Authorization: Bearer <token>
```

**响应：**
```json
{
  "code": 0,
  "message": "删除成功"
}
```

### 4. 用户收藏

#### GET /api/v1/users/favorites

获取用户收藏列表

**请求头：**
```
Authorization: Bearer <token>
```

**查询参数：**
```
page: int     // 页码（默认1）
pageSize: int // 每页数量（默认20）
```

**响应：**
```json
{
  "code": 0,
  "data": {
    "total": 10,
    "page": 1,
    "pageSize": 20,
    "items": [
      {
        "id": 1,
        "productId": 1,
        "product": {
          "id": 1,
          "name": "商品名称",
          "mainImage": "https://example.com/product.jpg",
          "price": 99.00,
          "originalPrice": 199.00
        },
        "createdAt": "2024-01-01T00:00:00Z"
      }
    ]
  }
}
```

#### POST /api/v1/users/favorites

添加收藏

**请求头：**
```
Authorization: Bearer <token>
```

**请求体：**
```json
{
  "productId": 1
}
```

**响应：**
```json
{
  "code": 0,
  "message": "收藏成功"
}
```

#### DELETE /api/v1/users/favorites/:productId

取消收藏

**请求头：**
```
Authorization: Bearer <token>
```

**响应：**
```json
{
  "code": 0,
  "message": "取消收藏成功"
}
```

### 5. 搜索历史

#### GET /api/v1/users/search-history

获取用户搜索历史

**请求头：**
```
Authorization: Bearer <token>
```

**响应：**
```json
{
  "code": 0,
  "data": [
    {
      "id": 1,
      "keyword": "手机",
      "createdAt": "2024-01-01T00:00:00Z"
    }
  ]
}
```

#### DELETE /api/v1/users/search-history

清空搜索历史

**请求头：**
```
Authorization: Bearer <token>
```

**响应：**
```json
{
  "code": 0,
  "message": "清空成功"
}
```

## 数据模型

### User
```go
type User struct {
    ID        uint64    `json:"id" gorm:"primaryKey"`
    Username  string    `json:"username" gorm:"uniqueIndex;size:50;not null"`
    Password  string    `json:"-" gorm:"size:255;not null"` // 不暴露给前端
    Nickname  string    `json:"nickname" gorm:"size:100"`
    Avatar    string    `json:"avatar" gorm:"size:500"`
    Phone     string    `json:"phone" gorm:"size:20"`
    Email     string    `json:"email" gorm:"size:100"`
    Gender    int8      `json:"gender" gorm:"default:0"` // 0-未知 1-男 2-女
    Birthday  *time.Time `json:"birthday"`
    Status    int8      `json:"status" gorm:"default:1"` // 0-禁用 1-正常
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
    DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
```

### UserAddress
```go
type UserAddress struct {
    ID            uint64    `json:"id" gorm:"primaryKey"`
    UserID        uint64    `json:"userId" gorm:"not null;index"`
    ReceiverName  string    `json:"receiverName" gorm:"size:50;not null"`
    ReceiverPhone string    `json:"receiverPhone" gorm:"size:20;not null"`
    Province      string    `json:"province" gorm:"size:50;not null"`
    City          string    `json:"city" gorm:"size:50;not null"`
    District      string    `json:"district" gorm:"size:50;not null"`
    DetailAddress string    `json:"detailAddress" gorm:"size:200;not null"`
    IsDefault     bool      `json:"isDefault" gorm:"default:false"`
    CreatedAt     time.Time `json:"createdAt"`
    UpdatedAt     time.Time `json:"updatedAt"`
    DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}
```

## 业务逻辑

### 认证逻辑
1. 注册时密码使用 bcrypt 加密存储
2. 登录成功后生成 JWT Token，有效期 7 天
3. Token 包含用户 ID 和过期时间
4. 登出时将 Token 加入黑名单（使用 Redis）

### 地址管理逻辑
1. 每个用户最多 20 个地址
2. 设置默认地址时，自动取消其他默认地址
3. 删除地址时，如果是默认地址，自动设置第一个地址为默认

### 收藏逻辑
1. 同一商品只能收藏一次（唯一索引）
2. 返回收藏列表时，包含商品的基本信息

## 依赖关系

- 依赖 MySQL 进行数据持久化
- 依赖 Redis 进行 Token 黑名单管理
- 依赖 JWT 进行身份认证
