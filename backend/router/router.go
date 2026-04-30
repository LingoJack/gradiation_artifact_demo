package router

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/lingojack/taobao_clone/config"
	"github.com/lingojack/taobao_clone/controller"
	mw "github.com/lingojack/taobao_clone/middleware"
	"github.com/lingojack/taobao_clone/pkg/validator"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Controllers struct {
	Health   *controller.HealthController
	Auth     *controller.AuthController
	User     *controller.UserController
	Product  *controller.ProductController
	Cart     *controller.CartController
	Order    *controller.OrderController
	Favorite *controller.FavoriteController
	Coupon   *controller.CouponController
	Shop     *controller.ShopController
}

func NewControllers(
	health *controller.HealthController,
	auth *controller.AuthController,
	user *controller.UserController,
	product *controller.ProductController,
	cart *controller.CartController,
	order *controller.OrderController,
	favorite *controller.FavoriteController,
	coupon *controller.CouponController,
	shop *controller.ShopController,
) *Controllers {
	return &Controllers{
		Health:   health,
		Auth:     auth,
		User:     user,
		Product:  product,
		Cart:     cart,
		Order:    order,
		Favorite: favorite,
		Coupon:   coupon,
		Shop:     shop,
	}
}

// NewEcho 创建 Echo 实例，注册中间件和路由，返回 (echo实例, cleanup函数, error)
func NewEcho(
	cfg *config.Config,
	log zerolog.Logger,
	db *gorm.DB,
	v *validator.CustomValidator,
	ctrl *Controllers,
) (*echo.Echo, func(), error) {
	e := echo.New()
	e.Validator = v
	e.HideBanner = true

	// 全局中间件
	e.Use(mw.Recover(cfg))
	e.Use(mw.RequestID(cfg))
	e.Use(mw.Logger())
	e.Use(mw.RateLimit(cfg))

	// 注册路由
	Register(e, cfg, ctrl)

	// cleanup 函数
	cleanup := func() {
		log.Info().Msg("running cleanup")
	}

	return e, cleanup, nil
}

func Register(e *echo.Echo, cfg *config.Config, ctrl *Controllers) {
	// 公开路由 — 无需鉴权
	open := e.Group(cfg.API.Prefix)
	open.Use(mw.CORS(cfg))
	open.GET("/health", ctrl.Health.Check)

	// Auth（公开）
	open.POST("/auth/register", ctrl.Auth.Register)
	open.POST("/auth/login", ctrl.Auth.Login)

	// Product（公开）
	open.GET("/products", ctrl.Product.List)
	open.GET("/products/:id", ctrl.Product.Get)
	open.GET("/products/search", ctrl.Product.Search)
	open.GET("/categories", ctrl.Product.Categories)
	open.GET("/banners", ctrl.Product.Banners)

	// Coupon（公开 — 可领取的优惠券）
	open.GET("/coupons/available", ctrl.Coupon.GetAvailable)

	// Shop（公开）
	open.GET("/shops/:id", ctrl.Shop.Get)
	open.GET("/shops/:id/products", ctrl.Shop.Products)

	// Swagger API 文档
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// 需要鉴权的路由
	private := e.Group(cfg.API.Prefix)
	private.Use(mw.CORS(cfg))
	private.Use(mw.Auth(cfg))

	// Auth（私有）
	private.POST("/auth/logout", ctrl.Auth.Logout)

	// User
	private.GET("/user/profile", ctrl.User.GetProfile)
	private.PUT("/user/profile", ctrl.User.UpdateProfile)
	private.GET("/user/addresses", ctrl.User.GetAddresses)
	private.POST("/user/addresses", ctrl.User.CreateAddress)
	private.PUT("/user/addresses/:id", ctrl.User.UpdateAddress)
	private.DELETE("/user/addresses/:id", ctrl.User.DeleteAddress)
	private.PUT("/user/addresses/:id/default", ctrl.User.SetDefaultAddress)

	// Cart
	private.GET("/cart", ctrl.Cart.GetCart)
	private.POST("/cart", ctrl.Cart.AddItem)
	private.PUT("/cart/:id", ctrl.Cart.UpdateQuantity)
	private.DELETE("/cart/:id", ctrl.Cart.RemoveItem)
	private.PUT("/cart/selected", ctrl.Cart.UpdateSelected)
	private.PUT("/cart/select-all", ctrl.Cart.SelectAll)
	private.DELETE("/cart/clear", ctrl.Cart.Clear)

	// Order
	private.POST("/orders", ctrl.Order.Create)
	private.GET("/orders", ctrl.Order.List)
	private.GET("/orders/:id", ctrl.Order.Get)
	private.PUT("/orders/:id/cancel", ctrl.Order.Cancel)
	private.PUT("/orders/:id/pay", ctrl.Order.Pay)
	private.PUT("/orders/:id/confirm", ctrl.Order.Confirm)
	private.DELETE("/orders/:id", ctrl.Order.Delete)

	// Favorite
	private.GET("/favorites", ctrl.Favorite.List)
	private.POST("/favorites", ctrl.Favorite.Add)
	private.DELETE("/favorites/:productId", ctrl.Favorite.Remove)
	private.GET("/favorites/:productId/check", ctrl.Favorite.Check)

	// Coupon（私有）
	private.GET("/coupons/mine", ctrl.Coupon.GetUserCoupons)
	private.POST("/coupons/:id/claim", ctrl.Coupon.Claim)
	private.PUT("/coupons/:id/use", ctrl.Coupon.Use)

	// Shop（私有 — 关注相关）
	private.POST("/shops/:id/follow", ctrl.Shop.Follow)
	private.GET("/shops/:id/follow/check", ctrl.Shop.CheckFollow)
}
