package dtopermission

import (
	"github.com/morehao/goark/apps/iam/object/objcommon"
	"github.com/morehao/goark/apps/iam/object/objpermission"
)

type MenuCreateResp struct {
	ID uint `json:"id"` // 数据自增id
}

type MenuDetailResp struct {
	ID uint `json:"id" validate:"required"` // 数据自增id
	objpermission.MenuBaseInfo
	objcommon.OperatorBaseInfo
}

type MenuPageListItem struct {
	ID uint `json:"id" validate:"required"` // 数据自增id
	objpermission.MenuBaseInfo
	objcommon.OperatorBaseInfo
}

type MenuPageListResp struct {
	List  []MenuPageListItem `json:"list"`  // 数据列表
	Total int64              `json:"total"` // 数据总条数
}

type RoleCreateResp struct {
	ID uint `json:"id"` // 数据自增id
}
type RoleDetailResp struct {
	ID uint `json:"id" validate:"required"` // 数据自增id
	objpermission.RoleBaseInfo
	objcommon.OperatorBaseInfo
}
type RolePageListItem struct {
	ID uint `json:"id" validate:"required"` // 数据自增id
	objpermission.RoleBaseInfo
	objcommon.OperatorBaseInfo
}
type RolePageListResp struct {
	List  []RolePageListItem `json:"list"`  // 数据列表
	Total int64              `json:"total"` // 数据总条数
}
