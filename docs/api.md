# API 设计文档

## 概述

本文档定义了淘宝风格电商应用的所有 RESTful API 端点、请求/响应格式。

**基础信息：**
- Base URL: `/api`
- 认证方式: JWT Token (Bearer Authorization)
- 响应格式: JSON
- 时间格式: `YYYY-MM-DD HH:mm:ss`

## 统一响应格式

### 成功响应
```json
{
  "code": 0,
  "message": "success",
  "data": { ... }
}
```

### 错误响应
```json
{
  "code": 1001,
  "message": "错误描述",
  "data": null
}
```

### 错误码定义
| 错误码 | 说明 |
|--------|------|
| 0 | 成功 |
| 1001 | 参数错误 |
| 1002 | 认证失败 |
| 1003 | 权限不足 |
| 1004 | 资源不存在 |
| 1005 | 业务逻辑错误 |
| 1006 | 服务器内部错误 |

---

## 1. 认证模块 (Auth)

### 1.1 用户注册
**POST** `/auth/register`

**请求体：**
```json
{
  "username": "string",      // 用户名，4-20字符
  "email": "string",         // 邮箱
  "phone": "string",         // 手机号
  "password": "string",      // 密码，6-20字符
  "nickname": "string?"      // 昵称，可选
}
```

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "user": {
      "id": "string",
      "username": "string",
      "email": "string",
      "phone": "string",
      "avatar": "string?",
      "nickname": "string?",
      "createdAt": "string"
    },
    "token": "string"
  }
}
```

### 1.2 用户登录
**POST** `/auth/login`

**请求体：**
```json
{
  "username": "string",      // 用户名或手机号或邮箱
  "password": "string"
}
```

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "user": {
      "id": "string",
      "username": "string",
      "email": "string",
      "phone": "string",
      "avatar": "string?",
      "nickname": "string?",
      "gender": "male|female|other?",
      "birthday": "string?",
      "bio": "string?",
      "createdAt": "string"
    },
    "token": "string"
  }
}
```

### 1.3 用户登出
**POST** `/auth/logout`

**请求头：** `Authorization: Bearer <token>`

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": null
}
```

---

## 2. 用户模块 (User)

### 2.1 获取用户信息
**GET** `/user/profile`

**请求头：** `Authorization: Bearer <token>`

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": "string",
    "username": "string",
    "email": "string",
    "phone": "string",
    "avatar": "string?",
    "nickname": "string?",
    "gender": "male|female|other?",
    "birthday": "string?",
    "bio": "string?",
    "createdAt": "string"
  }
}
```

### 2.2 更新用户信息
**PUT** `/user/profile`

**请求头：** `Authorization: Bearer <token>`

**请求体：**
```json
{
  "nickname": "string?",
  "avatar": "string?",
  "gender": "male|female|other?",
  "birthday": "string?",
  "bio": "string?"
}
```

**响应：** 同 2.1

### 2.3 获取用户地址列表
**GET** `/user/addresses`

**请求头：** `Authorization: Bearer <token>`

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "id": "string",
      "userId": "string",
      "receiver": "string",
      "phone": "string",
      "province": "string",
      "city": "string",
      "district": "string",
      "detail": "string",
      "isDefault": boolean
    }
  ]
}
```

### 2.4 创建地址
**POST** `/user/addresses`

**请求头：** `Authorization: Bearer <token>`

**请求体：**
```json
{
  "receiver": "string",
  "phone": "string",
  "province": "string",
  "city": "string",
  "district": "string",
  "detail": "string",
  "isDefault": boolean
}
```

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": "string",
    "userId": "string",
    "receiver": "string",
    "phone": "string",
    "province": "string",
    "city": "string",
    "district": "string",
    "detail": "string",
    "isDefault": boolean
  }
}
```

### 2.5 更新地址
**PUT** `/user/addresses/:id`

**请求头：** `Authorization: Bearer <token>`

**请求体：** 同 2.4

**响应：** 同 2.4

### 2.6 删除地址
**DELETE** `/user/addresses/:id`

**请求头：** `Authorization: Bearer <token>`

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": null
}
```

### 2.7 设置默认地址
**PUT** `/user/addresses/:id/default`

**请求头：** `Authorization: Bearer <token>`

**响应：** 同 2.4

---

## 3. 商品模块 (Product)

### 3.1 获取商品列表
**GET** `/products`

**查询参数：**
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| categoryId | string | 否 | 分类 ID |
| keyword | string | 否 | 搜索关键词 |
| shopId | string | 否 | 店铺 ID |
| minPrice | number | 否 | 最低价格 |
| maxPrice | number | 否 | 最高价格 |
| sort | string | 否 | 排序字段：price/sales/createdAt |
| order | string | 否 | 排序方向：asc/desc |
| page | number | 否 | 页码，默认 1 |
| pageSize | number | 否 | 每页数量，默认 20 |

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "string",
        "categoryId": "string",
        "name": "string",
        "description": "string",
        "price": number,
        "originalPrice": number?,
        "stock": number,
        "sales": number,
        "mainImage": "string",
        "images": ["string"],
        "specs": [
          {
            "id": "string",
            "productId": "string",
            "name": "string",
            "value": "string",
            "stock": number,
            "price": number
          }
        ],
        "status": "active|inactive",
        "createdAt": "string",
        "shopName": "string?"
      }
    ],
    "total": number,
    "page": number,
    "pageSize": number
  }
}
```

### 3.2 获取商品详情
**GET** `/products/:id`

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": "string",
    "categoryId": "string",
    "shopId": "string",
    "name": "string",
    "description": "string",
    "price": number,
    "originalPrice": number?,
    "stock": number,
    "sales": number,
    "mainImage": "string",
    "images": ["string"],
    "specs": [
      {
        "id": "string",
        "productId": "string",
        "name": "string",
        "value": "string",
        "stock": number,
        "price": number
      }
    ],
    "status": "active|inactive",
    "createdAt": "string",
    "shopName": "string?",
    "shop": {
      "id": "string",
      "name": "string",
      "avatar": "string",
      "rating": number,
      "sales": number,
      "fans": number
    }
  }
}
```

### 3.3 获取商品评价
**GET** `/products/:id/reviews`

**查询参数：**
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| page | number | 否 | 页码，默认 1 |
| pageSize | number | 否 | 每页数量，默认 10 |
| rating | number | 否 | 评分筛选：1-5 |

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "string",
        "userId": "string",
        "userName": "string",
        "userAvatar": "string",
        "productId": "string",
        "orderId": "string",
        "rating": number,
        "content": "string",
        "images": ["string"],
        "specName": "string?",
        "createdAt": "string",
        "reply": "string?",
        "likes": number
      }
    ],
    "total": number,
    "page": number,
    "pageSize": number,
    "stats": {
      "average": number,
      "total": number,
      "distribution": {
        "5": number,
        "4": number,
        "3": number,
        "2": number,
        "1": number
      }
    }
  }
}
```

### 3.4 获取分类列表
**GET** `/categories`

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "id": "string",
      "name": "string",
      "parentId": "string?",
      "icon": "string?",
      "sortOrder": number,
      "children": [
        {
          "id": "string",
          "name": "string",
          "parentId": "string",
          "icon": "string?",
          "sortOrder": number
        }
      ]
    }
  ]
}
```

### 3.5 获取首页 Banner
**GET** `/banners`

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "id": "string",
      "image": "string",
      "link": "string",
      "sortOrder": number
    }
  ]
}
```

---

## 4. 购物车模块 (Cart)

### 4.1 获取购物车
**GET** `/cart`

**请求头：** `Authorization: Bearer <token>`

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": [
      {
        "id": "string",
        "userId": "string",
        "productId": "string",
        "specId": "string?",
        "quantity": number,
        "selected": boolean,
        "product": {
          "id": "string",
          "name": "string",
          "price": number,
          "mainImage": "string",
          "stock": number
        },
        "spec": {
          "id": "string",
          "name": "string",
          "value": "string",
          "price": number,
          "stock": number
        }
      }
    ],
    "total": number,
    "selectedCount": number,
    "selectedAmount": number
  }
}
```

### 4.2 添加商品到购物车
**POST** `/cart`

**请求头：** `Authorization: Bearer <token>`

**请求体：**
```json
{
  "productId": "string",
  "specId": "string?",
  "quantity": number
}
```

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": "string",
    "userId": "string",
    "productId": "string",
    "specId": "string?",
    "quantity": number,
    "selected": boolean
  }
}
```

### 4.3 更新购物车商品数量
**PUT** `/cart/:id`

**请求头：** `Authorization: Bearer <token>`

**请求体：**
```json
{
  "quantity": number
}
```

**响应：** 同 4.2

### 4.4 删除购物车商品
**DELETE** `/cart/:id`

**请求头：** `Authorization: Bearer <token>`

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": null
}
```

### 4.5 批量删除购物车商品
**POST** `/cart/batch-delete`

**请求头：** `Authorization: Bearer <token>`

**请求体：**
```json
{
  "ids": ["string"]
}
```

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": null
}
```

### 4.6 选择/取消选择购物车商品
**PUT** `/cart/:id/select`

**请求头：** `Authorization: Bearer <token>`

**请求体：**
```json
{
  "selected": boolean
}
```

**响应：** 同 4.2

### 4.7 全选/取消全选
**PUT** `/cart/select-all`

**请求头：** `Authorization: Bearer <token>`

**请求体：**
```json
{
  "selected": boolean
}
```

**响应：** 同 4.1

### 4.8 清空购物车
**DELETE** `/cart/clear`

**请求头：** `Authorization: Bearer <token>`

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": null
}
```

---

## 5. 订单模块 (Order)

### 5.1 创建订单
**POST** `/orders`

**请求头：** `Authorization: Bearer <token>`

**请求体：**
```json
{
  "items": [
    {
      "productId": "string",
      "specId": "string?",
      "quantity": number,
      "price": number
    }
  ],
  "addressId": "string",
  "couponId": "string?",
  "remark": "string?"
}
```

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": "string",
    "orderNo": "string",
    "userId": "string",
    "totalAmount": number,
    "discountAmount": number,
    "payAmount": number,
    "status": "pending",
    "receiverName": "string",
    "receiverPhone": "string",
    "receiverAddress": "string",
    "items": [
      {
        "id": "string",
        "orderId": "string",
        "productId": "string",
        "productName": "string",
        "productImage": "string",
        "specId": "string?",
        "specName": "string?",
        "price": number,
        "quantity": number
      }
    ],
    "createdAt": "string"
  }
}
```

### 5.2 获取订单列表
**GET** `/orders`

**请求头：** `Authorization: Bearer <token>`

**查询参数：**
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| status | string | 否 | 订单状态：pending/paid/shipped/completed/cancelled |
| page | number | 否 | 页码，默认 1 |
| pageSize | number | 否 | 每页数量，默认 10 |

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "string",
        "orderNo": "string",
        "userId": "string",
        "totalAmount": number,
        "payAmount": number,
        "status": "string",
        "receiverName": "string",
        "receiverPhone": "string",
        "receiverAddress": "string",
        "items": [
          {
            "id": "string",
            "orderId": "string",
            "productId": "string",
            "productName": "string",
            "productImage": "string",
            "specName": "string?",
            "price": number,
            "quantity": number
          }
        ],
        "createdAt": "string",
        "paidAt": "string?",
        "shippedAt": "string?",
        "completedAt": "string?",
        "cancelledAt": "string?"
      }
    ],
    "total": number,
    "page": number,
    "pageSize": number
  }
}
```

### 5.3 获取订单详情
**GET** `/orders/:id`

**请求头：** `Authorization: Bearer <token>`

**响应：** 同 5.2 的单个订单对象

### 5.4 取消订单
**PUT** `/orders/:id/cancel`

**请求头：** `Authorization: Bearer <token>`

**请求体：**
```json
{
  "reason": "string?"
}
```

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": "string",
    "status": "cancelled",
    "cancelledAt": "string"
  }
}
```

### 5.5 支付订单（模拟）
**POST** `/orders/:id/pay`

**请求头：** `Authorization: Bearer <token>`

**请求体：**
```json
{
  "paymentMethod": "string"  // 支付方式：alipay/wechat/balance
}
```

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": "string",
    "status": "paid",
    "paidAt": "string"
  }
}
```

### 5.6 确认收货
**PUT** `/orders/:id/confirm`

**请求头：** `Authorization: Bearer <token>`

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": "string",
    "status": "completed",
    "completedAt": "string"
  }
}
```

### 5.7 订单评价
**POST** `/orders/:id/review`

**请求头：** `Authorization: Bearer <token>`

**请求体：**
```json
{
  "reviews": [
    {
      "productId": "string",
      "rating": number,
      "content": "string",
      "images": ["string"]
    }
  ]
}
```

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": null
}
```

---

## 6. 收藏模块 (Favorite)

### 6.1 获取收藏列表
**GET** `/favorites`

**请求头：** `Authorization: Bearer <token>`

**查询参数：**
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| page | number | 否 | 页码，默认 1 |
| pageSize | number | 否 | 每页数量，默认 20 |

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "string",
        "userId": "string",
        "productId": "string",
        "product": {
          "id": "string",
          "name": "string",
          "price": number,
          "originalPrice": number?,
          "mainImage": "string",
          "sales": number,
          "stock": number
        },
        "createdAt": "string"
      }
    ],
    "total": number,
    "page": number,
    "pageSize": number
  }
}
```

### 6.2 添加收藏
**POST** `/favorites`

**请求头：** `Authorization: Bearer <token>`

**请求体：**
```json
{
  "productId": "string"
}
```

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": "string",
    "productId": "string",
    "createdAt": "string"
  }
}
```

### 6.3 取消收藏
**DELETE** `/favorites/:productId`

**请求头：** `Authorization: Bearer <token>`

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": null
}
```

### 6.4 检查是否已收藏
**GET** `/favorites/check/:productId`

**请求头：** `Authorization: Bearer <token>`

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "isFavorite": boolean
  }
}
```

---

## 7. 优惠券模块 (Coupon)

### 7.1 获取可领取优惠券列表
**GET** `/coupons/available`

**请求头：** `Authorization: Bearer <token>`

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "id": "string",
      "name": "string",
      "discount": number,
      "minSpend": number,
      "scope": "string",
      "scopeType": "all|category",
      "scopeValue": "string?",
      "startTime": "string",
      "endTime": "string",
      "total": number,
      "claimed": number,
      "status": "active"
    }
  ]
}
```

### 7.2 获取用户优惠券列表
**GET** `/coupons/user`

**请求头：** `Authorization: Bearer <token>`

**查询参数：**
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| status | string | 否 | 状态：unused/used/expired |

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "id": "string",
      "couponId": "string",
      "userId": "string",
      "status": "unused|used|expired",
      "claimedAt": "string",
      "usedAt": "string?",
      "coupon": {
        "id": "string",
        "name": "string",
        "discount": number,
        "minSpend": number,
        "scope": "string",
        "scopeType": "string",
        "scopeValue": "string?",
        "startTime": "string",
        "endTime": "string"
      }
    }
  ]
}
```

### 7.3 领取优惠券
**POST** `/coupons/:id/claim`

**请求头：** `Authorization: Bearer <token>`

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": "string",
    "couponId": "string",
    "userId": "string",
    "status": "unused",
    "claimedAt": "string"
  }
}
```

### 7.4 计算可用优惠券
**POST** `/coupons/check`

**请求头：** `Authorization: Bearer <token>`

**请求体：**
```json
{
  "items": [
    {
      "productId": "string",
      "categoryId": "string",
      "price": number,
      "quantity": number
    }
  ],
  "totalAmount": number
}
```

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "availableCoupons": [
      {
        "id": "string",
        "couponId": "string",
        "coupon": {
          "name": "string",
          "discount": number,
          "minSpend": number
        },
        "discountAmount": number
      }
    ],
    "bestCoupon": {
      "id": "string",
      "discountAmount": number
    }
  }
}
```

---

## 8. 店铺模块 (Shop)

### 8.1 获取店铺详情
**GET** `/shops/:id`

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": "string",
    "name": "string",
    "description": "string",
    "avatar": "string",
    "coverImage": "string",
    "rating": number,
    "sales": number,
    "fans": number,
    "createdAt": "string",
    "location": "string",
    "isFollowed": boolean
  }
}
```

### 8.2 获取店铺商品列表
**GET** `/shops/:id/products`

**查询参数：**
| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| categoryId | string | 否 | 分类 ID |
| sort | string | 否 | 排序字段 |
| order | string | 否 | 排序方向 |
| page | number | 否 | 页码 |
| pageSize | number | 否 | 每页数量 |

**响应：** 同 3.1 商品列表格式

### 8.3 关注/取消关注店铺
**POST** `/shops/:id/follow`

**请求头：** `Authorization: Bearer <token>`

**请求体：**
```json
{
  "follow": boolean
}
```

**响应：**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "isFollowed": boolean,
    "fansCount": number
  }
}
```

---

## API 汇总

| 模块 | API 数量 |
|------|----------|
| 认证模块 | 3 |
| 用户模块 | 7 |
| 商品模块 | 5 |
| 购物车模块 | 8 |
| 订单模块 | 7 |
| 收藏模块 | 4 |
| 优惠券模块 | 4 |
| 店铺模块 | 3 |
| **总计** | **39** |