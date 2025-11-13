package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wujianling/moringstaradmin/internal/mods/artivitycenter/biz"
	"github.com/wujianling/moringstaradmin/internal/mods/artivitycenter/schema"
	"github.com/wujianling/moringstaradmin/pkg/util"
)

// Defining the `Artivity` api.
type Artivity struct {
	ArtivityBIZ *biz.Artivity
}

// @Tags ArtivityAPI
// @Security ApiKeyAuth
// @Summary Query artivity list
// @Param current query int true "pagination index" default(1)
// @Param pageSize query int true "pagination size" default(10)
// @Success 200 {object} util.ResponseResult{data=[]schema.Artivity}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/artivities [get]
func (a *Artivity) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.ArtivityQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.ArtivityBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags ArtivityAPI
// @Security ApiKeyAuth
// @Summary Get artivity record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.Artivity}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/artivities/{id} [get]
func (a *Artivity) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.ArtivityBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags ArtivityAPI
// @Security ApiKeyAuth
// @Summary Create artivity record
// @Param body body schema.ArtivityForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.Artivity}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/artivities [post]
func (a *Artivity) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.ArtivityForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.ArtivityBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags ArtivityAPI
// @Security ApiKeyAuth
// @Summary Update artivity record by ID
// @Param id path string true "unique id"
// @Param body body schema.ArtivityForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/artivities/{id} [put]
func (a *Artivity) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.ArtivityForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.ArtivityBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags ArtivityAPI
// @Security ApiKeyAuth
// @Summary Delete artivity record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/artivities/{id} [delete]
func (a *Artivity) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.ArtivityBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
