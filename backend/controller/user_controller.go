package controller

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/lingojack/taobao_clone/model"
	"github.com/lingojack/taobao_clone/pkg/response"
	"github.com/lingojack/taobao_clone/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

// GetProfile 获取用户信息
func (c *UserController) GetProfile(ctx echo.Context) error {
	userID := getUserID(ctx)
	user, err := c.userService.GetProfile(userID)
	if err != nil {
		return response.NotFound(ctx, "用户不存在")
	}
	return response.OK(ctx, user)
}

// UpdateProfile 更新用户信息
func (c *UserController) UpdateProfile(ctx echo.Context) error {
	userID := getUserID(ctx)
	var req service.UpdateProfileRequest
	if err := ctx.Bind(&req); err != nil {
		return response.BadRequest(ctx, "参数解析失败")
	}

	user, err := c.userService.UpdateProfile(userID, &req)
	if err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}
	return response.OK(ctx, user)
}

// GetAddresses 获取地址列表
func (c *UserController) GetAddresses(ctx echo.Context) error {
	userID := getUserID(ctx)
	addresses, err := c.userService.GetAddresses(userID)
	if err != nil {
		return response.Fail(ctx, 400, 1006, err.Error())
	}
	return response.OK(ctx, addresses)
}

// CreateAddress 创建地址
func (c *UserController) CreateAddress(ctx echo.Context) error {
	userID := getUserID(ctx)
	var req struct {
		ReceiverName  string `json:"receiver"`
		ReceiverPhone string `json:"phone"`
		Province      string `json:"province"`
		City          string `json:"city"`
		District      string `json:"district"`
		DetailAddress string `json:"detail"`
		IsDefault     int8   `json:"isDefault"`
	}
	if err := ctx.Bind(&req); err != nil {
		return response.BadRequest(ctx, "参数解析失败")
	}

	addr := &model.UserAddresses{
		ReceiverName:  req.ReceiverName,
		ReceiverPhone: req.ReceiverPhone,
		Province:      req.Province,
		City:          req.City,
		District:      req.District,
		DetailAddress: req.DetailAddress,
		IsDefault:     &req.IsDefault,
	}

	if err := c.userService.CreateAddress(userID, addr); err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}
	return response.OK(ctx, addr)
}

// UpdateAddress 更新地址
func (c *UserController) UpdateAddress(ctx echo.Context) error {
	userID := getUserID(ctx)
	addrID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(ctx, "无效的地址ID")
	}

	var req struct {
		ReceiverName  string `json:"receiver"`
		ReceiverPhone string `json:"phone"`
		Province      string `json:"province"`
		City          string `json:"city"`
		District      string `json:"district"`
		DetailAddress string `json:"detail"`
		IsDefault     int8   `json:"isDefault"`
	}
	if err := ctx.Bind(&req); err != nil {
		return response.BadRequest(ctx, "参数解析失败")
	}

	addr := &model.UserAddresses{
		ReceiverName:  req.ReceiverName,
		ReceiverPhone: req.ReceiverPhone,
		Province:      req.Province,
		City:          req.City,
		District:      req.District,
		DetailAddress: req.DetailAddress,
		IsDefault:     &req.IsDefault,
	}

	if err := c.userService.UpdateAddress(userID, addrID, addr); err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}
	return response.OK(ctx, addr)
}

// DeleteAddress 删除地址
func (c *UserController) DeleteAddress(ctx echo.Context) error {
	userID := getUserID(ctx)
	addrID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(ctx, "无效的地址ID")
	}

	if err := c.userService.DeleteAddress(userID, addrID); err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}
	return response.OK(ctx, nil)
}

// SetDefaultAddress 设置默认地址
func (c *UserController) SetDefaultAddress(ctx echo.Context) error {
	userID := getUserID(ctx)
	addrID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(ctx, "无效的地址ID")
	}

	if err := c.userService.SetDefaultAddress(userID, addrID); err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}
	return response.OK(ctx, nil)
}