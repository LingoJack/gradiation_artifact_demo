# 项目关系文档

本文档描述淘宝克隆项目各模块之间的调用关系、数据流向和架构层次。

## 一、整体架构

```
┌─────────────────────────────────────────────────────────────────────────┐
│                              用户浏览器                                   │
│                         http://localhost:5173                            │
└─────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                          前端 (React + Vite)                              │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐     │
│  │   Pages     │──│   Stores    │──│    API      │──│  Types      │     │
│  │  (15 pages) │  │ (Zustand)   │  │  (8 模块)   │  │ (TS 接口)   │     │
│  └─────────────┘  └─────────────┘  └─────────────┘  └─────────────┘     │
│                         Vite Proxy → http://localhost:8080              │
└─────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                          后端 (Go + Echo)                                 │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐     │
│  │ Controllers │──│  Services   │──│   Model     │──│  GORM DB    │     │
│  │  (9 ctrl)   │  │ (9 service) │  │ (15 model)  │  │             │     │
│  └─────────────┘  └─────────────┘  └─────────────┘  └─────────────┘     │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐                       │
│  │   Router    │──│ Middleware  │──│    Wire     │                       │
│  │             │  │ (JWT/CORS)  │  │    (DI)     │                       │
│  └─────────────┘  └─────────────┘  └─────────────┘                       │
└─────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                          数据库 (MySQL 8.0)                               │
│                         14 张表，utf8mb4 编码                             │
└─────────────────────────────────────────────────────────────────────────┘
```

## 二、API 路由与前端页面映射

### 公开路由（无需 JWT 认证）

| 路由 | 方法 | Controller | 前端页面 | 说明 |
|------|------|------------|----------|------|
| `/api/v1/health` | GET | HealthController.Check | - | 健康检查 |
| `/api/v1/auth/register` | POST | AuthController.Register | Register.tsx | 用户注册 |
| `/api/v1/auth/login` | POST | AuthController.Login | Login.tsx | 用户登录 |
| `/api/v1/products` | GET | ProductController.List | Home.tsx, ProductList.tsx | 商品列表 |
| `/api/v1/products/:id` | GET | ProductController.Get | ProductDetail.tsx | 商品详情 |
| `/api/v1/products/search` | GET | ProductController.Search | Home.tsx (搜索框) | 商品搜索 |
| `/api/v1/categories` | GET | ProductController.Categories | Home.tsx, ProductList.tsx | 分类列表 |
| `/api/v1/banners` | GET | ProductController.Banners | Home.tsx | 首页轮播图 |
| `/api/v1/coupons/available` | GET | CouponController.GetAvailable | Coupons.tsx | 可领取优惠券 |
| `/api/v1/shops/:id` | GET | ShopController.Get | Shop.tsx | 店铺详情 |
| `/api/v1/shops/:id/products` | GET | ShopController.Products | Shop.tsx | 店铺商品 |

### 私有路由（需要 JWT 认证）

| 路由 | 方法 | Controller | 前端页面 | 说明 |
|------|------|------------|----------|------|
| `/api/v1/auth/logout` | POST | AuthController.Logout | - | 登出 |
| `/api/v1/user/profile` | GET | UserController.GetProfile | UserCenter.tsx | 用户信息 |
| `/api/v1/user/profile` | PUT | UserController.UpdateProfile | Settings.tsx | 更新信息 |
| `/api/v1/user/addresses` | GET | UserController.GetAddresses | Addresses.tsx, Checkout.tsx | 地址列表 |
| `/api/v1/user/addresses` | POST | UserController.CreateAddress | Addresses.tsx | 新增地址 |
| `/api/v1/user/addresses/:id` | PUT | UserController.UpdateAddress | Addresses.tsx | 更新地址 |
| `/api/v1/user/addresses/:id` | DELETE | UserController.DeleteAddress | Addresses.tsx | 删除地址 |
| `/api/v1/user/addresses/:id/default` | PUT | UserController.SetDefaultAddress | Addresses.tsx | 设为默认 |
| `/api/v1/cart` | GET | CartController.GetCart | Cart.tsx, Checkout.tsx | 购物车 |
| `/api/v1/cart` | POST | CartController.AddItem | ProductDetail.tsx | 加入购物车 |
| `/api/v1/cart/:id` | PUT | CartController.UpdateQuantity | Cart.tsx | 更新数量 |
| `/api/v1/cart/:id` | DELETE | CartController.RemoveItem | Cart.tsx | 删除商品 |
| `/api/v1/cart/selected` | PUT | CartController.UpdateSelected | Cart.tsx | 更新选中 |
| `/api/v1/cart/select-all` | PUT | CartController.SelectAll | Cart.tsx | 全选/取消 |
| `/api/v1/cart/clear` | DELETE | CartController.Clear | Cart.tsx | 清空购物车 |
| `/api/v1/orders` | POST | OrderController.Create | Checkout.tsx | 创建订单 |
| `/api/v1/orders` | GET | OrderController.List | Order.tsx | 订单列表 |
| `/api/v1/orders/:id` | GET | OrderController.Get | Order.tsx | 订单详情 |
| `/api/v1/orders/:id/cancel` | PUT | OrderController.Cancel | Order.tsx | 取消订单 |
| `/api/v1/orders/:id/pay` | PUT | OrderController.Pay | Order.tsx | 支付订单 |
| `/api/v1/orders/:id/confirm` | PUT | OrderController.Confirm | Order.tsx | 确认收货 |
| `/api/v1/orders/:id` | DELETE | OrderController.Delete | Order.tsx | 删除订单 |
| `/api/v1/favorites` | GET | FavoriteController.List | Favorites.tsx | 收藏列表 |
| `/api/v1/favorites` | POST | FavoriteController.Add | ProductDetail.tsx | 添加收藏 |
| `/api/v1/favorites/:productId` | DELETE | FavoriteController.Remove | Favorites.tsx | 取消收藏 |
| `/api/v1/favorites/:productId/check` | GET | FavoriteController.Check | ProductDetail.tsx | 检查收藏状态 |
| `/api/v1/coupons/mine` | GET | CouponController.GetUserCoupons | Coupons.tsx | 我的优惠券 |
| `/api/v1/coupons/:id/claim` | POST | CouponController.Claim | Coupons.tsx | 领取优惠券 |
| `/api/v1/coupons/:id/use` | PUT | CouponController.Use | Checkout.tsx | 使用优惠券 |
| `/api/v1/shops/:id/follow` | POST | ShopController.Follow | Shop.tsx | 关注/取消店铺 |
| `/api/v1/shops/:id/follow/check` | GET | ShopController.CheckFollow | Shop.tsx | 检查关注状态 |

## 三、数据库表与 Model 映射

| 数据库表 | Go Model 文件 | 主要字段 | 说明 |
|----------|---------------|----------|------|
| `users` | model/users.go | id, username, password, nickname, avatar, phone, email | 用户表 |
| `user_addresses` | model/user_addresses.go | id, user_id, receiver_name, receiver_phone, province, city, district, detail_address | 地址表 |
| `categories` | model/categories.go | id, name, parent_id, icon, sort_order | 分类表 |
| `products` | model/products.go | id, category_id, name, description, main_image, images, price, stock, sales | 商品表 |
| `product_skus` | model/product_skus.go | id, product_id, sku_code, spec_values, price, stock | SKU 表 |
| `cart_items` | model/cart_items.go | id, user_id, product_id, sku_id, quantity, selected | 购物车表 |
| `orders` | model/orders.go | id, order_no, user_id, total_amount, pay_amount, status, receiver_* | 订单表 |
| `order_items` | model/order_items.go | id, order_id, product_id, sku_id, product_name, price, quantity | 订单项表 |
| `user_favorites` | model/user_favorites.go | id, user_id, product_id | 收藏表 |
| `search_histories` | model/search_histories.go | id, user_id, keyword | 搜索历史表 |
| `banners` | model/banners.go | id, title, image_url, link_url, sort_order | Banner 表 |
| `shops` | model/shops.go | id, name, description, avatar, rating, sales, fans, location | 店铺表 |
| `coupons` | model/coupons.go | id, name, discount, min_spend, total, claimed, start_time, end_time | 优惠券表 |
| `user_coupons` | model/user_coupons.go | id, user_id, coupon_id, status, used_at | 用户优惠券表 |

**注意**：`reviews.go` 中定义了 `Reviews` 和 `ShopFollows` 两个 Model，但数据库 schema 中未创建对应表（预留扩展）。

## 四、后端分层架构

```
┌─────────────────────────────────────────────────────────────┐
│                     Router (路由层)                          │
│  router/router.go                                           │
│  - Register() 注册所有路由                                   │
│  - 分组：open(公开) / private(需JWT)                         │
│  - 中间件：Recover → RequestID → Logger → RateLimit → Auth  │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                   Controller (控制器层)                      │
│  controller/*.go                                            │
│  - 参数解析/绑定                                             │
│  - 调用 Service                                              │
│  - 统一响应格式：{code, message, data}                       │
│  - 错误码：0(成功), 1001-1006(业务错误)                      │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                     Service (服务层)                         │
│  service/*.go                                               │
│  - 业务逻辑处理                                              │
│  - 数据验证/转换                                             │
│  - 直接使用 gorm.DB 操作数据库                               │
│  - 不依赖 Repository 层（简化架构）                          │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                      Model (模型层)                          │
│  model/*.go                                                 │
│  - GORM 结构体定义                                           │
│  - TableName() 方法                                          │
│  - JSON 序列化标签 (snake_case)                              │
│  - 时间字段使用 *time.Time (允许 NULL)                       │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                    Database (数据库层)                       │
│  pkg/database/database.go                                   │
│  - GORM 连接池                                               │
│  - DSN 配置                                                  │
│  - utf8mb4 字符集                                            │
└─────────────────────────────────────────────────────────────┘
```

### Wire 依赖注入关系

```
wire.go 依赖链:

config.Load ─────────────────────────────────────────────────────┐
                                                                 │
logger.New ──────────────────────────────────────────────────────│
                                                                 │
database.New ────────────────────────────────────────────────────│
                                                                 │
validator.New ───────────────────────────────────────────────────│
                                                                 ▼
service.New* (9个) ──────────────────────────────────────────────► controller.New* (9个)
                                                                 │
                                                                 ▼
router.NewControllers ───────────────────────────────────────────► router.NewEcho
```

## 五、前端状态管理

### Zustand Store 与 API 映射

| Store | 文件 | 关联 API | 关联页面 |
|-------|------|----------|----------|
| useUserStore | stores/useUserStore.ts | authApi, userApi | Login, Register, UserCenter, Settings |
| useCartStore | stores/useCartStore.ts | cartApi | Cart, Checkout, ProductDetail |
| useOrderStore | stores/useOrderStore.ts | orderApi | Order, Checkout |
| useFavoriteStore | stores/useFavoriteStore.ts | favoriteApi | Favorites, ProductDetail |
| useCouponStore | stores/useCouponStore.ts | couponApi | Coupons, Checkout |

### 前端 API 层结构

```
frontend/src/api/
├── client.ts      # 基础请求封装 (get/post/put/del)
├── index.ts       # 导出所有 API 模块
├── auth.ts        # 认证 API (login/register/logout)
├── user.ts        # 用户 API (profile/addresses)
├── product.ts     # 商品 API (list/detail/search/categories/banners)
├── cart.ts        # 购物车 API (get/add/update/delete/clear)
├── order.ts       # 订单 API (create/list/detail/cancel/pay/confirm)
├── favorite.ts    # 收藏 API (list/add/remove/check)
├── coupon.ts      # 优惠券 API (available/mine/claim/use)
└── shop.ts        # 店铺 API (detail/products/follow/check)
```

## 六、前端页面与组件依赖

### 页面列表 (15 个)

| 页面 | 路由 | 主要组件 | Store | API |
|------|------|----------|-------|-----|
| Home | `/` | Header, ProductCard, Footer | useUserStore | productApi, bannerApi |
| ProductList | `/products` | ProductCard, CustomSelect | - | productApi |
| ProductDetail | `/products/:id` | - | useCartStore, useFavoriteStore | productApi, cartApi, favoriteApi |
| Cart | `/cart` | - | useCartStore | cartApi |
| Checkout | `/checkout` | - | useCartStore, useUserStore, useCouponStore | cartApi, orderApi, couponApi |
| Order | `/orders` | - | useOrderStore | orderApi |
| Login | `/login` | - | useUserStore | authApi |
| Register | `/register` | - | useUserStore | authApi |
| Profile | `/profile` | Header, Footer | useUserStore | userApi |
| UserCenter | `/user-center` | Header, Footer | useUserStore | userApi |
| Settings | `/settings` | Header, Footer | useUserStore | userApi |
| Addresses | `/addresses` | Header, Footer | useUserStore | userApi |
| Favorites | `/favorites` | Header, Footer, ProductCard | useFavoriteStore | favoriteApi |
| Coupons | `/coupons` | Header, Footer | useCouponStore | couponApi |
| Shop | `/shop/:id` | Header, Footer | - | shopApi |

### 共用组件

| 组件 | 文件 | 使用页面 |
|------|------|----------|
| Header | components/Header/Header.tsx | Home, Profile, UserCenter, Settings, Addresses, Favorites, Coupons, Shop |
| Footer | components/Footer/Footer.tsx | Home, Profile, UserCenter, Settings, Addresses, Favorites, Coupons, Shop |
| Layout | components/Layout/Layout.tsx | 所有页面（可选包裹） |
| ProductCard | components/ProductCard/ProductCard.tsx | Home, ProductList, Favorites |
| CustomSelect | components/CustomSelect/CustomSelect.tsx | ProductList |

## 七、容器化部署

### Podman Compose 服务

```yaml
services:
  mysql:       # MySQL 8.0，端口 3306，utf8mb4 编码
  backend:     # Go Echo，端口 8080，依赖 mysql
  frontend:    # React Vite + nginx，端口 5173，代理到 backend
```

### 启动命令

```bash
# 开发阶段（单独启动）
podman compose up -d mysql          # 启动数据库
cd backend && go run ./cmd/server/  # 启动后端
cd frontend && npm run dev          # 启动前端

# 生产阶段（全栈容器化）
podman compose up -d --build        # 构建并启动所有服务
podman compose ps                   # 查看状态
podman compose logs -f              # 查看日志
podman compose down                 # 停止服务
```

## 八、文件目录结构

```
gradiation_artifact_demo/
├── docs/                          # 文档目录
│   ├── requirement.md             # 需求文档
│   ├── api.md                     # API 设计文档
│   ├── frontend_design.md         # 前端设计文档
│   ├── backend-design.md          # 后端设计文档
│   ├── backend/schema.sql         # 数据库 schema（旧）
│   └── relations.md               # 本文档
│
├── docker/                        # Docker 配置
│   └── mysql-init/
│       └── 01_schema.sql          # MySQL 初始化脚本
│
├── backend/                       # Go 后端
│   ├── cmd/server/                # 入口
│   │   ├── main.go
│   │   ├── wire.go                # Wire DI 定义
│   │   └── wire_gen.go            # Wire 生成的代码
│   ├── config/
│   │   └── config.yaml            # 配置文件
│   ├── controller/                # 控制器层（9 个）
│   ├── service/                   # 服务层（9 个）
│   ├── model/                     # 模型层（15 个）
│   ├── repository/                # 仓库层（预留）
│   ├── middleware/                # 中间件
│   ├── router/
│   │   └── router.go              # 路由注册
│   ├── pkg/                       # 公共包
│   │   ├── database/
│   │   ├── logger/
│   │   ├── response/
│   │   ├── tool/
│   │   └── validator/
│   ├── go.mod
│   ├── go.sum
│   └── Dockerfile
│
├── frontend/                      # React 前端
│   ├── src/
│   │   ├── api/                   # API 层（8 个模块）
│   │   ├── components/            # 组件（5 个）
│   │   ├── pages/                 # 页面（15 个）
│   │   ├── stores/                # Zustand Store（5 个）
│   │   ├── types/                 # TypeScript 类型定义
│   │   ├── App.tsx                # 路由配置
│   │   └── main.tsx               # 入口
│   ├── vite.config.ts             # Vite 配置（proxy）
│   ├── package.json
│   └── Dockerfile
│
├── docker-compose.yml             # Podman Compose 配置
└── Makefile                       # 开发命令
```

## 九、关键技术点总结

1. **后端技术栈**：Go 1.21+ / Echo v4 / GORM v2 / Wire DI / JWT Auth
2. **前端技术栈**：React 19 / TypeScript / Vite / Zustand / Tailwind CSS
3. **数据库**：MySQL 8.0，utf8mb4 编码，14 张表，无外键约束
4. **API 格式**：统一 `{code, message, data}`，JSON 字段名 snake_case
5. **认证方式**：JWT Token，前端 localStorage 存储，请求头 `Authorization: Bearer <token>`
6. **容器化**：Podman Compose，本地开发单独启动，生产全栈容器化

---

*文档生成时间：2025-01*