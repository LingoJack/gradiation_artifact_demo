package controller

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/lingojack/taobao_clone/pkg/response"
	"github.com/lingojack/taobao_clone/service"
)

type CartController struct {
	cartService *service.CartService
}

func NewCartController(cartService *service.CartService) *CartController {
	return &CartController{cartService: cartService}
}

// GetCart 获取购物车
func (c *CartController) GetCart(ctx echo.Context) error {
	userID := getUserID(ctx)
	result, err := c.cartService.GetCart(userID)
	if err != nil {
		return response.Fail(ctx, 400, 1006, err.Error())
	}
	return response.OK(ctx, result)
}

// AddItem 添加商品到购物车
func (c *CartController) AddItem(ctx echo.Context) error {
	userID := getUserID(ctx)
	var req service.AddCartRequest
	if err := ctx.Bind(&req); err != nil {
		return response.BadRequest(ctx, "参数解析失败")
	}

	result, err := c.cartService.AddItem(userID, &req)
	if err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}
	return response.OK(ctx, result)
}

// UpdateQuantity 更新数量
func (c *CartController) UpdateQuantity(ctx echo.Context) error {
	userID := getUserID(ctx)
	itemID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(ctx, "无效的商品ID")
	}

	var req struct {
		Quantity int `json:"quantity"`
	}
	if err := ctx.Bind(&req); err != nil {
		return response.BadRequest(ctx, "参数解析失败")
	}

	result, err := c.cartService.UpdateQuantity(userID, itemID, req.Quantity)
	if err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}
	return response.OK(ctx, result)
}

// RemoveItem 删除购物车商品
func (c *CartController) RemoveItem(ctx echo.Context) error {
	userID := getUserID(ctx)
	itemID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(ctx, "无效的商品ID")
	}

	if err := c.cartService.RemoveItem(userID, itemID); err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}
	return response.OK(ctx, nil)
}

// UpdateSelected 更新选中状态
func (c *CartController) UpdateSelected(ctx echo.Context) error {
	userID := getUserID(ctx)
	var req struct {
		ItemIDs  []uint64 `json:"itemIds"`
		Selected bool     `json:"selected"`
	}
	if err := ctx.Bind(&req); err != nil {
		return response.BadRequest(ctx, "参数解析失败")
	}

	if err := c.cartService.UpdateSelected(userID, req.ItemIDs, req.Selected); err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}
	return response.OK(ctx, nil)
}

// SelectAll 全选/取消全选
func (c *CartController) SelectAll(ctx echo.Context) error {
	userID := getUserID(ctx)
	var req struct {
		Selected bool `json:"selected"`
	}
	if err := ctx.Bind(&req); err != nil {
		return response.BadRequest(ctx, "参数解析失败")
	}

	if err := c.cartService.SelectAll(userID, req.Selected); err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}
	return response.OK(ctx, nil)
}

// Clear 清空购物车
func (c *CartController) Clear(ctx echo.Context) error {
	userID := getUserID(ctx)
	if err := c.cartService.ClearCart(userID); err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}
	return response.OK(ctx, nil)
}