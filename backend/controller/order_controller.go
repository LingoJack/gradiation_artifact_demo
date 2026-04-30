package controller

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/lingojack/taobao_clone/pkg/response"
	"github.com/lingojack/taobao_clone/service"
)

type OrderController struct {
	orderService *service.OrderService
}

func NewOrderController(orderService *service.OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

// Create 创建订单
func (c *OrderController) Create(ctx echo.Context) error {
	userID := getUserID(ctx)
	var req service.CreateOrderRequest
	if err := ctx.Bind(&req); err != nil {
		return response.BadRequest(ctx, "参数解析失败")
	}

	result, err := c.orderService.CreateOrder(userID, &req)
	if err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}
	return response.Created(ctx, result)
}

// List 订单列表
func (c *OrderController) List(ctx echo.Context) error {
	userID := getUserID(ctx)
	status := ctx.QueryParam("status")
	page := parseInt(ctx.QueryParam("page"), 1)
	pageSize := parseInt(ctx.QueryParam("pageSize"), 10)

	result, total, err := c.orderService.GetOrders(userID, status, page, pageSize)
	if err != nil {
		return response.Fail(ctx, 400, 1006, err.Error())
	}
	return response.OK(ctx, map[string]interface{}{"orders": result, "total": total})
}

// Get 订单详情
func (c *OrderController) Get(ctx echo.Context) error {
	userID := getUserID(ctx)
	orderID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(ctx, "无效的订单ID")
	}

	result, err := c.orderService.GetOrderDetail(userID, orderID)
	if err != nil {
		return response.NotFound(ctx, err.Error())
	}
	return response.OK(ctx, result)
}

// Cancel 取消订单
func (c *OrderController) Cancel(ctx echo.Context) error {
	userID := getUserID(ctx)
	orderID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(ctx, "无效的订单ID")
	}

	if err := c.orderService.CancelOrder(userID, orderID); err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}
	return response.OK(ctx, nil)
}

// Pay 支付订单
func (c *OrderController) Pay(ctx echo.Context) error {
	userID := getUserID(ctx)
	orderID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(ctx, "无效的订单ID")
	}

	if err := c.orderService.PayOrder(userID, orderID); err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}
	return response.OK(ctx, nil)
}

// Confirm 确认收货
func (c *OrderController) Confirm(ctx echo.Context) error {
	userID := getUserID(ctx)
	orderID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(ctx, "无效的订单ID")
	}

	if err := c.orderService.ConfirmReceive(userID, orderID); err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}
	return response.OK(ctx, nil)
}

// Delete 删除订单
func (c *OrderController) Delete(ctx echo.Context) error {
	userID := getUserID(ctx)
	orderID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(ctx, "无效的订单ID")
	}

	if err := c.orderService.DeleteOrder(userID, orderID); err != nil {
		return response.Fail(ctx, 400, 1005, err.Error())
	}
	return response.OK(ctx, nil)
}