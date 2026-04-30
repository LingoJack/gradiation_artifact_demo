package controller

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/lingojack/taobao_clone/pkg/response"
	"github.com/lingojack/taobao_clone/service"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Register 用户注册
func (c *AuthController) Register(ctx echo.Context) error {
	var req service.RegisterRequest
	if err := ctx.Bind(&req); err != nil {
		return response.BadRequest(ctx, "参数解析失败")
	}

	result, err := c.authService.Register(ctx.Request().Context(), &req)
	if err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}

	return response.OK(ctx, result)
}

// Login 用户登录
func (c *AuthController) Login(ctx echo.Context) error {
	var req service.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return response.BadRequest(ctx, "参数解析失败")
	}

	result, err := c.authService.Login(ctx.Request().Context(), &req)
	if err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}

	return response.OK(ctx, result)
}

// Logout 用户登出
func (c *AuthController) Logout(ctx echo.Context) error {
	return response.OK(ctx, nil)
}

// getUserID 从 JWT context 获取用户 ID
func getUserID(ctx echo.Context) uint64 {
	userID, ok := ctx.Get("userID").(uint64)
	if !ok {
		return 0
	}
	return userID
}

func parseUint64(s string) uint64 {
	if s == "" {
		return 0
	}
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return v
}

func parseInt(s string, defaultVal int) int {
	if s == "" {
		return defaultVal
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return defaultVal
	}
	return v
}

func parseFloat(s string) float64 {
	if s == "" {
		return 0
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return v
}