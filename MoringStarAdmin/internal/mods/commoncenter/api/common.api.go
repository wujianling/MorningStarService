package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wujianling/moringstaradmin/internal/mods/commoncenter/biz"
	"github.com/wujianling/moringstaradmin/internal/mods/commoncenter/schema"
	"github.com/wujianling/moringstaradmin/pkg/util"
)

// Defining the `Common` api.
type Common struct {
	CommonBIZ *biz.Common
}

// @Tags CommonAPI
// @Security ApiKeyAuth
// @Summary Query common list
// @Param current query int true "pagination index" default(1)
// @Param pageSize query int true "pagination size" default(10)
// @Success 200 {object} util.ResponseResult{data=[]schema.Common}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/commons [get]
func (a *Common) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.CommonQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.CommonBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags CommonAPI
// @Security ApiKeyAuth
// @Summary Get common record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.Common}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/commons/{id} [get]
func (a *Common) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.CommonBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags CommonAPI
// @Security ApiKeyAuth
// @Summary Create common record
// @Param body body schema.CommonForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.Common}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/commons [post]
func (a *Common) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.CommonForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.CommonBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags CommonAPI
// @Security ApiKeyAuth
// @Summary Update common record by ID
// @Param id path string true "unique id"
// @Param body body schema.CommonForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/commons/{id} [put]
func (a *Common) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.CommonForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.CommonBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags CommonAPI
// @Security ApiKeyAuth
// @Summary Delete common record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/commons/{id} [delete]
func (a *Common) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.CommonBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
