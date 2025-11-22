package dtotenant

import (
	"github.com/morehao/goark/apps/iam/object/objcommon"
	"github.com/morehao/goark/apps/iam/object/objtenant"
)

type TenantCreateResp struct {
	ID uint `json:"id"` // 数据自增id
}

type TenantDetailResp struct {
	ID uint `json:"id" validate:"required"` // 数据自增id
	objtenant.TenantBaseInfo
	objcommon.OperatorBaseInfo
}

type TenantPageListItem struct {
	ID uint `json:"id" validate:"required"` // 数据自增id
	objtenant.TenantBaseInfo
	objcommon.OperatorBaseInfo
}

type TenantPageListResp struct {
	List  []TenantPageListItem `json:"list"`  // 数据列表
	Total int64                `json:"total"` // 数据总条数
}
