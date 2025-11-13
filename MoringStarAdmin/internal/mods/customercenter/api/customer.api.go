package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wujianling/moringstaradmin/internal/mods/customercenter/biz"
	"github.com/wujianling/moringstaradmin/internal/mods/customercenter/schema"
	"github.com/wujianling/moringstaradmin/pkg/util"
)

// Defining the `Customer` api.
type Customer struct {
	CustomerBIZ *biz.Customer
}

// @Tags CustomerAPI
// @Security ApiKeyAuth
// @Summary Query customer list
// @Param current query int true "pagination index" default(1)
// @Param pageSize query int true "pagination size" default(10)
// @Success 200 {object} util.ResponseResult{data=[]schema.Customer}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/customers [get]
func (a *Customer) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.CustomerQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.CustomerBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags CustomerAPI
// @Security ApiKeyAuth
// @Summary Get customer record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.Customer}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/customers/{id} [get]
func (a *Customer) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.CustomerBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags CustomerAPI
// @Security ApiKeyAuth
// @Summary Create customer record
// @Param body body schema.CustomerForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.Customer}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/customers [post]
func (a *Customer) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.CustomerForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.CustomerBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags CustomerAPI
// @Security ApiKeyAuth
// @Summary Update customer record by ID
// @Param id path string true "unique id"
// @Param body body schema.CustomerForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/customers/{id} [put]
func (a *Customer) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.CustomerForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.CustomerBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags CustomerAPI
// @Security ApiKeyAuth
// @Summary Delete customer record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/customers/{id} [delete]
func (a *Customer) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.CustomerBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
