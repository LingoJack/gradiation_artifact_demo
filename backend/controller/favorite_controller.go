package controller

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/lingojack/taobao_clone/pkg/response"
	"github.com/lingojack/taobao_clone/service"
)

type FavoriteController struct {
	favoriteService *service.FavoriteService
}

func NewFavoriteController(favoriteService *service.FavoriteService) *FavoriteController {
	return &FavoriteController{favoriteService: favoriteService}
}

// List 收藏列表
func (c *FavoriteController) List(ctx echo.Context) error {
	userID := getUserID(ctx)
	page := parseInt(ctx.QueryParam("page"), 1)
	pageSize := parseInt(ctx.QueryParam("pageSize"), 20)

	result, total, err := c.favoriteService.GetUserFavorites(userID, page, pageSize)
	if err != nil {
		return response.Fail(ctx, 400, 1006, err.Error())
	}
	return response.OK(ctx, map[string]interface{}{"favorites": result, "total": total})
}

// Add 添加收藏
func (c *FavoriteController) Add(ctx echo.Context) error {
	userID := getUserID(ctx)
	var req struct {
		ProductID uint64 `json:"productId"`
	}
	if err := ctx.Bind(&req); err != nil {
		return response.BadRequest(ctx, "参数解析失败")
	}

	if err := c.favoriteService.AddFavorite(userID, req.ProductID); err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}
	return response.OK(ctx, nil)
}

// Remove 取消收藏
func (c *FavoriteController) Remove(ctx echo.Context) error {
	userID := getUserID(ctx)
	productID, err := strconv.ParseUint(ctx.Param("productId"), 10, 64)
	if err != nil {
		return response.BadRequest(ctx, "无效的商品ID")
	}

	if err := c.favoriteService.RemoveFavorite(userID, productID); err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}
	return response.OK(ctx, nil)
}

// Check 检查是否已收藏
func (c *FavoriteController) Check(ctx echo.Context) error {
	userID := getUserID(ctx)
	productID, err := strconv.ParseUint(ctx.Param("productId"), 10, 64)
	if err != nil {
		return response.BadRequest(ctx, "无效的商品ID")
	}

	isFavorite, err := c.favoriteService.CheckFavorite(userID, productID)
	if err != nil {
		return response.Fail(ctx, 400, 1006, err.Error())
	}
	return response.OK(ctx, map[string]bool{"isFavorite": isFavorite})
}