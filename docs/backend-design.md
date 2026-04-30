# 后端设计文档

## 技术栈

- **语言**: Go 1.24
- **Web框架**: Echo v4
- **ORM**: GORM v2
- **数据库**: MySQL 8.0
- **认证**: JWT (golang-jwt/jwt v5)
- **依赖注入**: Wire (google/wire)
- **文档**: Swagger (swaggo/swag)

## 目录结构

```
backend/
├── cmd/
│   └── server/
│       ├── main.go        # 入口文件
│       ├── wire.go        # Wire 依赖注入定义
│       └── wire_gen.go    # Wire 生成的代码
├── config/
│   ├── config.go          # 配置结构体
│   └── config.yaml        # 配置文件
├── controller/            # HTTP 处理器（Handler层）
│   ├── auth_controller.go
│   ├── user_controller.go
│   ├── product_controller.go
│   ├── cart_controller.go
│   ├── order_controller.go
│   ├── favorite_controller.go
│   ├── coupon_controller.go
│   ├── shop_controller.go
│   └── health.go
├── middleware/            # 中间件
│   ├── auth.go            # JWT 认证
│   ├── cors.go
│   ├── logger.go
│   ├── rate_limit.go
│   ├── request_id.go
│   ├── recover.go
├── model/                 # 数据模型
│   ├── user.go
│   ├── address.go
│   ├── product.go
│   ├── product_spec.go
│   ├── category.go
│   ├── cart_item.go
│   ├── order.go
│   ├── order_item.go
│   ├── review.go
│   ├── favorite.go
│   ├── coupon.go
│   ├── user_coupon.go
│   ├── shop.go
│   ├── banner.go
├── repository/            # 数据访问层（DAO）
│   ├── user_repository.go
│   ├── address_repository.go
│   ├── product_repository.go
│   ├── cart_repository.go
│   ├── order_repository.go
│   ├── favorite_repository.go
│   ├── coupon_repository.go
│   ├── shop_repository.go
├── service/               # 业务逻辑层
│   ├── auth_service.go
│   ├── user_service.go
│   ├── product_service.go
│   ├── cart_service.go
│   ├── order_service.go
│   ├── favorite_service.go
│   ├── coupon_service.go
│   ├── shop_service.go
├── pkg/                   # 公共包
│   ├── database/
│   ├── logger/
│   ├── response/
│   ├── validator/
├── tool/                  # 工具函数
├── docs/                  # Swagger 文档
├── Dockerfile
├── go.mod
├── go.sum
└── Makefile
```

## 分层架构

```
┌─────────────────────────────────────────────────────────────┐
│                      HTTP Request                            │
└─────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────┐
│                    Middleware Layer                          │
│  (Auth, CORS, Logger, RateLimit, RequestID, Recover)         │
└─────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────┐
│                   Controller Layer                           │
│  (请求解析、响应格式化、参数验证)                              │
└─────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────┐
│                    Service Layer                             │
│  (业务逻辑、事务管理、跨 Repository 调用)                      │
└─────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────┐
│                  Repository Layer                            │
│  (数据访问、CRUD 操作、查询封装)                               │
└─────────────────────────────────────────────────────────────┘
                              ↓
┌─────────────────────────────────────────────────────────────┐
│                    Database (MySQL)                          │
└─────────────────────────────────────────────────────────────┘
```

## 模块设计

### 1. 认证模块 (Auth)

**Service: AuthService**
```go
type AuthService interface {
    Register(username, email, phone, password, nickname string) (*model.User, string, error)
    Login(username, password string) (*model.User, string, error)
    Logout(token string) error
    GenerateToken(user *model.User) (string, error)
    ValidateToken(token string) (*JWTClaims, error)
}
```

**Controller: AuthController**
- POST `/api/v1/auth/register` - 用户注册
- POST `/api/v1/auth/login` - 用户登录
- POST `/api/v1/auth/logout` - 用户登出

### 2. 用户模块 (User)

**Service: UserService**
```go
type UserService interface {
    GetProfile(userID string) (*model.User, error)
    UpdateProfile(userID string, updates map[string]interface{}) error
    GetAddresses(userID string) ([]model.Address, error)
    CreateAddress(userID string, addr *model.Address) error
    UpdateAddress(userID string, addrID string, addr *model.Address) error
    DeleteAddress(userID string, addrID string) error
    SetDefaultAddress(userID string, addrID string) error
}
```

### 3. 商品模块 (Product)

**Service: ProductService**
```go
type ProductService interface {
    List(filter ProductFilter, page, pageSize int) (*ProductListResult, error)
    GetDetail(productID string) (*model.Product, error)
    GetReviews(productID string, page, pageSize int, rating int) (*ReviewListResult, error)
    GetCategories() ([]model.Category, error)
    GetBanners() ([]model.Banner, error)
    Search(keyword string, page, pageSize int) (*ProductListResult, error)
}
```

### 4. 购物车模块 (Cart)

**Service: CartService**
```go
type CartService interface {
    GetCart(userID string) (*CartResult, error)
    AddItem(userID string, productID, specID string, quantity int) error
    UpdateQuantity(userID string, itemID string, quantity int) error
    RemoveItem(userID string, itemID string) error
    BatchRemove(userID string, itemIDs []string) error
    ToggleSelect(userID string, itemID string, selected bool) error
    SelectAll(userID string, selected bool) error
    Clear(userID string) error
}
```

### 5. 订单模块 (Order)

**Service: OrderService**
```go
type OrderService interface {
    Create(userID string, req CreateOrderRequest) (*model.Order, error)
    List(userID string, status string, page, pageSize int) (*OrderListResult, error)
    GetDetail(userID string, orderID string) (*model.Order, error)
    Cancel(userID string, orderID string, reason string) error
    Pay(userID string, orderID string, method string) error
    Confirm(userID string, orderID string) error
    Review(userID string, orderID string, reviews []ReviewRequest) error
}
```

### 6. 收藏模块 (Favorite)

**Service: FavoriteService**
```go
type FavoriteService interface {
    List(userID string, page, pageSize int) (*FavoriteListResult, error)
    Add(userID string, productID string) error
    Remove(userID string, productID string) error
    Check(userID string, productID string) (bool, error)
}
```

### 7. 优惠券模块 (Coupon)

**Service: CouponService**
```go
type CouponService interface {
    GetAvailable(userID string) ([]model.Coupon, error)
    GetUserCoupons(userID string, status string) ([]model.UserCoupon, error)
    Claim(userID string, couponID string) error
    CalculateAvailable(userID string, items []CartItemInfo, totalAmount float64) (*CouponCalcResult, error)
}
```

### 8. 店铺模块 (Shop)

**Service: ShopService**
```go
type ShopService interface {
    GetDetail(shopID string) (*model.Shop, error)
    GetProducts(shopID string, filter ProductFilter, page, pageSize int) (*ProductListResult, error)
    Follow(userID string, shopID string, follow bool) error
}
```

## 依赖注入 (Wire)

```go
// wire.go
func initEcho(cfgPath string) (*echo.Echo, func(), error) {
    wire.Build(
        config.Load,
        logger.New,
        database.New,
        validator.New,
        // Repository
        repository.NewUserRepository,
        repository.NewAddressRepository,
        repository.NewProductRepository,
        repository.NewCartRepository,
        repository.NewOrderRepository,
        repository.NewFavoriteRepository,
        repository.NewCouponRepository,
        repository.NewShopRepository,
        // Service
        service.NewAuthService,
        service.NewUserService,
        service.NewProductService,
        service.NewCartService,
        service.NewOrderService,
        service.NewFavoriteService,
        service.NewCouponService,
        service.NewShopService,
        // Controller
        controller.NewAuthController,
        controller.NewUserController,
        controller.NewProductController,
        controller.NewCartController,
        controller.NewOrderController,
        controller.NewFavoriteController,
        controller.NewCouponController,
        controller.NewShopController,
        controller.NewHealthController,
        // Router
        router.NewControllers,
        router.NewEcho,
    )
    return nil, nil, nil
}
```

## 中间件配置

| 中间件 | 作用 | 配置项 |
|--------|------|--------|
| Recover | 异常恢复 | enabled: true |
| RequestID | 请求追踪 ID | enabled: true |
| Logger | 日志记录 | enabled: true |
| CORS | 跨域处理 | allowed_origins: http://localhost:5173 |
| RateLimit | 限流 | enabled: false (生产启用) |
| Auth | JWT 认证 | enabled: true |

## 响应格式

遵循统一响应格式：

```json
{
  "code": 0,
  "message": "ok",
  "data": { ... }
}
```

错误响应：
```json
{
  "code": 1001,
  "message": "参数错误",
  "data": null
}
```

## API 路由分组

```go
// 公开路由（无需认证）
open := e.Group("/api/v1")
open.GET("/health", ctrl.Health.Check)
open.POST("/auth/register", ctrl.Auth.Register)
open.POST("/auth/login", ctrl.Auth.Login)
open.GET("/products", ctrl.Product.List)
open.GET("/products/:id", ctrl.Product.Get)
open.GET("/categories", ctrl.Product.Categories)
open.GET("/banners", ctrl.Product.Banners)
open.GET("/shops/:id", ctrl.Shop.Get)

// 认证路由（需要 JWT）
private := e.Group("/api/v1")
private.Use(middleware.Auth(cfg))
private.POST("/auth/logout", ctrl.Auth.Logout)
private.GET("/user/profile", ctrl.User.GetProfile)
private.PUT("/user/profile", ctrl.User.UpdateProfile)
private.GET("/user/addresses", ctrl.User.GetAddresses)
// ... 其他需要认证的 API
```

## 数据库连接

```yaml
database:
  driver: "mysql"
  dsn: "root:password@tcp(mysql:3306)/taobao_clone?charset=utf8mb4&parseTime=True&loc=Local"
  max_open_conns: 25
  max_idle_conns: 5
  conn_max_lifetime_minutes: 30
```

**开发环境 DSN (本地 MySQL 容器):**
```
root:password@tcp(localhost:3306)/taobao_clone?charset=utf8mb4&parseTime=True&loc=Local
```

## Go Module 配置

```go
module github.com/lingojack/taobao_clone

go 1.24

require (
    github.com/labstack/echo/v4 v4.13.3
    gorm.io/gorm v1.25.12
    gorm.io/driver/mysql v1.5.7
    github.com/golang-jwt/jwt/v5 v5.2.2
    github.com/google/wire v0.6.0
    github.com/swaggo/swag v1.16.4
    github.com/swaggo/echo-swagger v1.4.1
    github.com/rs/zerolog v1.34.0
)
```

## 开发流程

### 本地开发（编码期）

1. 启动 MySQL 容器（仅依赖服务）：
   ```bash
   podman compose up mysql -d
   ```

2. 运行后端（本地 go run）：
   ```bash
   cd backend && go run cmd/server/main.go
   ```

3. 运行前端：
   ```bash
   cd frontend && npm run dev
   ```

### 最终验收

```bash
podman compose up -d --build
```

访问：
- 前端: http://localhost:5173
- 后端 API: http://localhost:8080/api/v1
- Swagger 文档: http://localhost:8080/swagger/index.html

## 代码生成

使用 `sql-to-go-struct-and-dao` 技能从 `docs/backend/schema.sql` 生成：
- model/*.go - 数据模型结构体
- repository/*.go - Repository 接口和实现

命令：
```bash
# 在项目根目录运行
make jen  # 如果有 jen 工具配置
# 或手动运行技能
``