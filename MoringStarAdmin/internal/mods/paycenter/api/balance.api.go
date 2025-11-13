package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wujianling/moringstaradmin/internal/mods/paycenter/biz"
	"github.com/wujianling/moringstaradmin/internal/mods/paycenter/schema"
	"github.com/wujianling/moringstaradmin/pkg/util"
)

// Defining the `Balance` api.
type Balance struct {
	BalanceBIZ *biz.Balance
}

// @Tags BalanceAPI
// @Security ApiKeyAuth
// @Summary Query balance list
// @Param current query int true "pagination index" default(1)
// @Param pageSize query int true "pagination size" default(10)
// @Success 200 {object} util.ResponseResult{data=[]schema.Balance}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/balances [get]
func (a *Balance) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.BalanceQueryParam
	if err := util.ParseQuery(c, &params); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.BalanceBIZ.Query(ctx, params)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResPage(c, result.Data, result.PageResult)
}

// @Tags BalanceAPI
// @Security ApiKeyAuth
// @Summary Get balance record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult{data=schema.Balance}
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/balances/{id} [get]
func (a *Balance) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.BalanceBIZ.Get(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, item)
}

// @Tags BalanceAPI
// @Security ApiKeyAuth
// @Summary Create balance record
// @Param body body schema.BalanceForm true "Request body"
// @Success 200 {object} util.ResponseResult{data=schema.Balance}
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/balances [post]
func (a *Balance) Create(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.BalanceForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.BalanceBIZ.Create(ctx, item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResSuccess(c, result)
}

// @Tags BalanceAPI
// @Security ApiKeyAuth
// @Summary Update balance record by ID
// @Param id path string true "unique id"
// @Param body body schema.BalanceForm true "Request body"
// @Success 200 {object} util.ResponseResult
// @Failure 400 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/balances/{id} [put]
func (a *Balance) Update(c *gin.Context) {
	ctx := c.Request.Context()
	item := new(schema.BalanceForm)
	if err := util.ParseJSON(c, item); err != nil {
		util.ResError(c, err)
		return
	} else if err := item.Validate(); err != nil {
		util.ResError(c, err)
		return
	}

	err := a.BalanceBIZ.Update(ctx, c.Param("id"), item)
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}

// @Tags BalanceAPI
// @Security ApiKeyAuth
// @Summary Delete balance record by ID
// @Param id path string true "unique id"
// @Success 200 {object} util.ResponseResult
// @Failure 401 {object} util.ResponseResult
// @Failure 500 {object} util.ResponseResult
// @Router /api/v1/balances/{id} [delete]
func (a *Balance) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.BalanceBIZ.Delete(ctx, c.Param("id"))
	if err != nil {
		util.ResError(c, err)
		return
	}
	util.ResOK(c)
}
