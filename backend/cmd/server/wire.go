//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"

	"github.com/lingojack/taobao_clone/config"
	"github.com/lingojack/taobao_clone/controller"
	"github.com/lingojack/taobao_clone/pkg/database"
	"github.com/lingojack/taobao_clone/pkg/logger"
	"github.com/lingojack/taobao_clone/pkg/validator"
	"github.com/lingojack/taobao_clone/router"
	"github.com/lingojack/taobao_clone/service"
)

func initEcho(cfgPath string) (*echo.Echo, func(), error) {
	wire.Build(
		config.Load,
		logger.New,
		database.New,
		validator.New,
		service.NewAuthService,
		service.NewUserService,
		service.NewProductService,
		service.NewCartService,
		service.NewOrderService,
		service.NewFavoriteService,
		service.NewCouponService,
		service.NewShopService,
		controller.NewHealthController,
		controller.NewAuthController,
		controller.NewUserController,
		controller.NewProductController,
		controller.NewCartController,
		controller.NewOrderController,
		controller.NewFavoriteController,
		controller.NewCouponController,
		controller.NewShopController,
		router.NewControllers,
		router.NewEcho,
	)
	return nil, nil, nil
}
