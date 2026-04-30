package controller

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/lingojack/taobao_clone/pkg/response"
	"github.com/lingojack/taobao_clone/service"
)

type CouponController struct {
	couponService *service.CouponService
}

func NewCouponController(couponService *service.CouponService) *CouponController {
	return &CouponController{couponService: couponService}
}

// GetAvailable 可领取优惠券
func (c *CouponController) GetAvailable(ctx echo.Context) error {
	coupons, err := c.couponService.GetAvailableCoupons()
	if err != nil {
		return response.Fail(ctx, 400, 1006, err.Error())
	}
	return response.OK(ctx, coupons)
}

// GetUserCoupons 用户优惠券列表
func (c *CouponController) GetUserCoupons(ctx echo.Context) error {
	userID := getUserID(ctx)

	coupons, err := c.couponService.GetUserCoupons(userID)
	if err != nil {
		return response.Fail(ctx, 400, 1006, err.Error())
	}
	return response.OK(ctx, coupons)
}

// Claim 领取优惠券
func (c *CouponController) Claim(ctx echo.Context) error {
	userID := getUserID(ctx)
	couponID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(ctx, "无效的优惠券ID")
	}

	result, err := c.couponService.ClaimCoupon(userID, couponID)
	if err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}
	return response.OK(ctx, result)
}

// Use 使用优惠券
func (c *CouponController) Use(ctx echo.Context) error {
	userID := getUserID(ctx)
	userCouponID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(ctx, "无效的优惠券ID")
	}

	if err := c.couponService.UseCoupon(userID, userCouponID); err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}
	return response.OK(ctx, nil)
}