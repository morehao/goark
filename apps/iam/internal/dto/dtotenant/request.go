package dtotenant

import (
	"github.com/morehao/goark/apps/iam/object/objcommon"
	"github.com/morehao/goark/apps/iam/object/objtenant"
)

type TenantCreateReq struct {
	objtenant.TenantBaseInfo
}

type TenantUpdateReq struct {
	ID uint `json:"id" validate:"required" label:"数据自增id"` // 数据自增id
	objtenant.TenantBaseInfo
}

type TenantDetailReq struct {
	ID uint `json:"id" form:"id" validate:"required" label:"数据自增id"` // 数据自增id
}

type TenantPageListReq struct {
	objcommon.PageQuery
}

type TenantDeleteReq struct {
	ID uint `json:"id" form:"id" validate:"required" label:"数据自增id"` // 数据自增id
}
