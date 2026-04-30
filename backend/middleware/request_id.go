package middleware

import (
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"

	"github.com/lingojack/taobao_clone/config"
)

func RequestID(cfg *config.Config) echo.MiddlewareFunc {
	if !cfg.Middleware.RequestID.Enabled {
		return Passthrough()
	}
	return echomw.RequestID()
}
