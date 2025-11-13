package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wujianling/moringstaradmin/internal/mods/shopcenter/biz"
	"github.com/wujianling/moringstaradmin/internal/mods/shopcenter/schema"
	"github.com/wujianling/moringstaradmin/pkg/util"
)

// Defining the `Shop` api.
type Shop struct {
	ShopBIZ *biz.Shop
}

// @Tags ShopAPI
// @Security ApiKeyAuth
// @Summary Query shop list
// @Param current query int true "pagination index" default(1)
// @Param pageSize query int true "pagination size" default(10)
// @Success 200 {object} util.ResponseResult{data=[]schema.Shop}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/shops [get]
func (a *Shop) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.ShopQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.ShopBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags ShopAPI
// @Security ApiKeyAuth
// @Summary Get shop record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.Shop}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/shops/{id} [get]
func (a *Shop) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.ShopBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags ShopAPI
// @Security ApiKeyAuth
// @Summary Create shop record
// @Param body body schema.ShopForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.Shop}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/shops [post]
func (a *Shop) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.ShopForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.ShopBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags ShopAPI
// @Security ApiKeyAuth
// @Summary Update shop record by ID
// @Param id path string true "unique id"
// @Param body body schema.ShopForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/shops/{id} [put]
func (a *Shop) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.ShopForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.ShopBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags ShopAPI
// @Security ApiKeyAuth
// @Summary Delete shop record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/shops/{id} [delete]
func (a *Shop) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.ShopBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
