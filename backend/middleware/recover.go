package middleware

import (
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"

	"github.com/lingojack/taobao_clone/config"
)

func Recover(cfg *config.Config) echo.MiddlewareFunc {
	if !cfg.Middleware.Recover.Enabled {
		return Passthrough()
	}
	return echomw.Recover()
}
