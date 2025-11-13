package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wujianling/moringstaradmin/internal/mods/productcenter/biz"
	"github.com/wujianling/moringstaradmin/internal/mods/productcenter/schema"
	"github.com/wujianling/moringstaradmin/pkg/util"
)

// Defining the `Product` api.
type Product struct {
	ProductBIZ *biz.Product
}

// @Tags ProductAPI
// @Security ApiKeyAuth
// @Summary Query product list
// @Param current query int true "pagination index" default(1)
// @Param pageSize query int true "pagination size" default(10)
// @Success 200 {object} util.ResponseResult{data=[]schema.Product}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/products [get]
func (a *Product) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.ProductQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.ProductBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags ProductAPI
// @Security ApiKeyAuth
// @Summary Get product record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.Product}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/products/{id} [get]
func (a *Product) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.ProductBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags ProductAPI
// @Security ApiKeyAuth
// @Summary Create product record
// @Param body body schema.ProductForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.Product}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/products [post]
func (a *Product) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.ProductForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.ProductBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags ProductAPI
// @Security ApiKeyAuth
// @Summary Update product record by ID
// @Param id path string true "unique id"
// @Param body body schema.ProductForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/products/{id} [put]
func (a *Product) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.ProductForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.ProductBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags ProductAPI
// @Security ApiKeyAuth
// @Summary Delete product record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/products/{id} [delete]
func (a *Product) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.ProductBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
