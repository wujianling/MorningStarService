package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wujianling/moringstaradmin/internal/mods/ordercenter/biz"
	"github.com/wujianling/moringstaradmin/internal/mods/ordercenter/schema"
	"github.com/wujianling/moringstaradmin/pkg/util"
)

// Defining the `Orders` api.
type Orders struct {
	OrdersBIZ *biz.Orders
}

// @Tags OrdersAPI
// @Security ApiKeyAuth
// @Summary Query orders list
// @Param current query int true "pagination index" default(1)
// @Param pageSize query int true "pagination size" default(10)
// @Success 200 {object} util.ResponseResult{data=[]schema.Orders}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/orders [get]
func (a *Orders) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.OrdersQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.OrdersBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags OrdersAPI
// @Security ApiKeyAuth
// @Summary Get orders record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.Orders}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/orders/{id} [get]
func (a *Orders) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.OrdersBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags OrdersAPI
// @Security ApiKeyAuth
// @Summary Create orders record
// @Param body body schema.OrdersForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.Orders}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/orders [post]
func (a *Orders) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.OrdersForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.OrdersBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags OrdersAPI
// @Security ApiKeyAuth
// @Summary Update orders record by ID
// @Param id path string true "unique id"
// @Param body body schema.OrdersForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/orders/{id} [put]
func (a *Orders) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.OrdersForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.OrdersBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags OrdersAPI
// @Security ApiKeyAuth
// @Summary Delete orders record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/orders/{id} [delete]
func (a *Orders) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.OrdersBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
