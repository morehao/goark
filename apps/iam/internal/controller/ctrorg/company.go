package ctrorg

import (
	"github.com/gin-gonic/gin"
	"github.com/morehao/goark/apps/iam/internal/dto/dtoorg"
	"github.com/morehao/goark/apps/iam/internal/service/svcorg"
	"github.com/morehao/golib/gcontext/gincontext"
)

type CompanyCtr interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	Detail(ctx *gin.Context)
	PageList(ctx *gin.Context)
}

type companyCtr struct {
	companySvc svcorg.CompanySvc
}

var _ CompanyCtr = (*companyCtr)(nil)

func NewCompanyCtr() CompanyCtr {
	return &companyCtr{
		companySvc: svcorg.NewCompanySvc(),
	}
}

// Create 创建公司管理
// @Tags 公司管理
// @Summary 创建公司管理
// @accept application/json
// @Produce application/json
// @Param req body dtoorg.CompanyCreateReq true "创建公司管理"
// @Success 200 {object} gincontext.DtoRender{data=dtoorg.CompanyCreateResp} "{"code": 0,"data": "ok","msg": "success"}"
// @Router /iam/v1/company/create [post]
func (ctr *companyCtr) Create(ctx *gin.Context) {
	var req dtoorg.CompanyCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		gincontext.Fail(ctx, err)
		return
	}
	res, err := ctr.companySvc.Create(ctx, &req)
	if err != nil {
		gincontext.Fail(ctx, err)
		return
	} else {
		gincontext.Success(ctx, res)
	}
}

// Delete 删除公司管理
// @Tags 公司管理
// @Summary 删除公司管理
// @accept application/json
// @Produce application/json
// @Param req body dtoorg.CompanyDeleteReq true "删除公司管理"
// @Success 200 {object} gincontext.DtoRender{data=string} "{"code": 0,"data": "ok","msg": "删除成功"}"
// @Router /iam/v1/company/delete [post]
func (ctr *companyCtr) Delete(ctx *gin.Context) {
	var req dtoorg.CompanyDeleteReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		gincontext.Fail(ctx, err)
		return
	}

	if err := ctr.companySvc.Delete(ctx, &req); err != nil {
		gincontext.Fail(ctx, err)
		return
	} else {
		gincontext.Success(ctx, "删除成功")
	}
}

// Update 修改公司管理
// @Tags 公司管理
// @Summary 修改公司管理
// @accept application/json
// @Produce application/json
// @Param req body dtoorg.CompanyUpdateReq true "修改公司管理"
// @Success 200 {object} gincontext.DtoRender{data=string} "{"code": 0,"data": "ok","msg": "修改成功"}"
// @Router /iam/v1/company/update [post]
func (ctr *companyCtr) Update(ctx *gin.Context) {
	var req dtoorg.CompanyUpdateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		gincontext.Fail(ctx, err)
		return
	}
	if err := ctr.companySvc.Update(ctx, &req); err != nil {
		gincontext.Fail(ctx, err)
		return
	} else {
		gincontext.Success(ctx, "修改成功")
	}
}

// Detail 公司管理详情
// @Tags 公司管理
// @Summary 公司管理详情
// @accept application/json
// @Produce application/json
// @Param req query dtoorg.CompanyDetailReq true "公司管理详情"
// @Success 200 {object} gincontext.DtoRender{data=dtoorg.CompanyDetailResp} "{"code": 0,"data": "ok","msg": "success"}"
// @Router /iam/v1/company/detail [get]
func (ctr *companyCtr) Detail(ctx *gin.Context) {
	var req dtoorg.CompanyDetailReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		gincontext.Fail(ctx, err)
		return
	}
	res, err := ctr.companySvc.Detail(ctx, &req)
	if err != nil {
		gincontext.Fail(ctx, err)
		return
	} else {
		gincontext.Success(ctx, res)
	}
}

// PageList 公司管理列表
// @Tags 公司管理
// @Summary 公司管理列表分页
// @accept application/json
// @Produce application/json
// @Param req query dtoorg.CompanyPageListReq true "公司管理列表"
// @Success 200 {object} gincontext.DtoRender{data=dtoorg.CompanyPageListResp} "{"code": 0,"data": "ok","msg": "success"}"
// @Router /iam/v1/company/pageList [get]
func (ctr *companyCtr) PageList(ctx *gin.Context) {
	var req dtoorg.CompanyPageListReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		gincontext.Fail(ctx, err)
		return
	}
	res, err := ctr.companySvc.PageList(ctx, &req)
	if err != nil {
		gincontext.Fail(ctx, err)
		return
	} else {
		gincontext.Success(ctx, res)
	}
}
