package controller

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/lingojack/taobao_clone/pkg/response"
	"github.com/lingojack/taobao_clone/service"
)

type ProductController struct {
	productService *service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

// List 商品列表
func (c *ProductController) List(ctx echo.Context) error {
	categoryID := parseUint64(ctx.QueryParam("categoryId"))
	keyword := ctx.QueryParam("keyword")
	page := parseInt(ctx.QueryParam("page"), 1)
	pageSize := parseInt(ctx.QueryParam("pageSize"), 20)
	sort := ctx.QueryParam("sort")

	result, err := c.productService.GetProducts(categoryID, keyword, page, pageSize, sort)
	if err != nil {
		return response.Fail(ctx, 400, 1006, err.Error())
	}
	return response.OK(ctx, result)
}

// Get 商品详情
func (c *ProductController) Get(ctx echo.Context) error {
	productID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return response.BadRequest(ctx, "无效的商品ID")
	}

	result, err := c.productService.GetProductDetail(productID)
	if err != nil {
		return response.NotFound(ctx, "商品不存在")
	}
	return response.OK(ctx, result)
}

// Search 商品搜索建议
func (c *ProductController) Search(ctx echo.Context) error {
	keyword := ctx.QueryParam("keyword")
	if keyword == "" {
		return response.BadRequest(ctx, "搜索关键词不能为空")
	}

	result, err := c.productService.SearchSuggestions(keyword)
	if err != nil {
		return response.Fail(ctx, 400, 1006, err.Error())
	}
	return response.OK(ctx, result)
}

// Categories 分类列表
func (c *ProductController) Categories(ctx echo.Context) error {
	categories, err := c.productService.GetCategories()
	if err != nil {
		return response.Fail(ctx, 400, 1006, err.Error())
	}
	return response.OK(ctx, categories)
}

// Banners Banner列表
func (c *ProductController) Banners(ctx echo.Context) error {
	banners, err := c.productService.GetBanners()
	if err != nil {
		return response.Fail(ctx, 400, 1006, err.Error())
	}
	return response.OK(ctx, banners)
}