package dtotenant

import (
	"github.com/morehao/goark/apps/iam/object/objtenant"
	"github.com/morehao/golib/gobject"
)

type TenantCreateResp struct {
	// ID 数据自增 ID
	ID uint `json:"id"`
}

type TenantDetailResp struct {
	// ID 数据自增 ID
	ID uint `json:"id" validate:"required"`
	objtenant.TenantBaseInfo
	gobject.OperatorBaseInfo
}

type TenantPageListItem struct {
	// ID 数据自增 ID
	ID uint `json:"id" validate:"required"`
	objtenant.TenantBaseInfo
	gobject.OperatorBaseInfo
}

type TenantPageListResp struct {
	// List 数据列表
	List []TenantPageListItem `json:"list"`
	// Total 数据总条数
	Total int64 `json:"total"`
}
