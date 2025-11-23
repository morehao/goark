package dtopermission

import (
	"github.com/morehao/goark/apps/iam/object/objpermission"
	"github.com/morehao/golib/gobject"
)

type MenuCreateResp struct {
	// ID 数据自增 ID
	ID uint `json:"id"`
}

type MenuDetailResp struct {
	// ID 数据自增 ID
	ID uint `json:"id" validate:"required"`
	objpermission.MenuBaseInfo
	gobject.OperatorBaseInfo
}

type MenuPageListItem struct {
	// ID 数据自增 ID
	ID uint `json:"id" validate:"required"`
	objpermission.MenuBaseInfo
	gobject.OperatorBaseInfo
}

type MenuPageListResp struct {
	// List 数据列表
	List []MenuPageListItem `json:"list"`
	// Total 数据总条数
	Total int64 `json:"total"`
}

type RoleCreateResp struct {
	// ID 数据自增 ID
	ID uint `json:"id"`
}
type RoleDetailResp struct {
	// ID 数据自增 ID
	ID uint `json:"id" validate:"required"`
	objpermission.RoleBaseInfo
	gobject.OperatorBaseInfo
}
type RolePageListItem struct {
	// ID 数据自增 ID
	ID uint `json:"id" validate:"required"`
	objpermission.RoleBaseInfo
	gobject.OperatorBaseInfo
}
type RolePageListResp struct {
	// List 数据列表
	List []RolePageListItem `json:"list"`
	// Total 数据总条数
	Total int64 `json:"total"`
}
