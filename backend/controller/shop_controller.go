package controller

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/lingojack/taobao_clone/pkg/response"
	"github.com/lingojack/taobao_clone/service"
)

type ShopController struct {
	shopService *service.ShopService
}

func NewShopController(shopService *service.ShopService) *ShopController {
	return &ShopController{shopService: shopService}
}

// Get 店铺详情
func (c *ShopController) Get(ctx echo.Context) error {
	shopID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(ctx, "无效的店铺ID")
	}

	result, err := c.shopService.GetDetail(shopID)
	if err != nil {
		return response.NotFound(ctx, "店铺不存在")
	}
	return response.OK(ctx, result)
}

// Products 店铺商品
func (c *ShopController) Products(ctx echo.Context) error {
	shopID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(ctx, "无效的店铺ID")
	}

	page := parseInt(ctx.QueryParam("page"), 1)
	pageSize := parseInt(ctx.QueryParam("pageSize"), 20)

	result, total, err := c.shopService.GetShopProducts(shopID, page, pageSize)
	if err != nil {
		return response.Fail(ctx, 400, 1006, err.Error())
	}
	return response.OK(ctx, map[string]interface{}{"products": result, "total": total})
}

// Follow 关注/取消关注店铺
func (c *ShopController) Follow(ctx echo.Context) error {
	userID := getUserID(ctx)
	shopID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(ctx, "无效的店铺ID")
	}

	var req struct {
		Follow bool `json:"follow"`
	}
	ctx.Bind(&req)

	if err := c.shopService.ToggleFollow(userID, shopID, req.Follow); err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}
	return response.OK(ctx, nil)
}

// CheckFollow 检查是否关注
func (c *ShopController) CheckFollow(ctx echo.Context) error {
	userID := getUserID(ctx)
	shopID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(ctx, "无效的店铺ID")
	}

	isFollow, err := c.shopService.CheckFollow(userID, shopID)
	if err != nil {
		return response.Fail(ctx, 400, 1006, err.Error())
	}
	return response.OK(ctx, map[string]bool{"isFollow": isFollow})
}