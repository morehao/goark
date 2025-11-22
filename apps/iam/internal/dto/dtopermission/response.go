package dtopermission

import (
	"github.com/morehao/goark/apps/iam/object/objcommon"
	"github.com/morehao/goark/apps/iam/object/objpermission"
)

type MenuCreateResp struct {
	// ID 数据自增id
	ID uint `json:"id"`
}

type MenuDetailResp struct {
	// ID 数据自增id
	ID uint `json:"id" validate:"required"`

	objpermission.MenuBaseInfo
	objcommon.OperatorBaseInfo
}

type MenuPageListItem struct {
	// ID 数据自增id
	ID uint `json:"id" validate:"required"`

	objpermission.MenuBaseInfo
	objcommon.OperatorBaseInfo
}

type MenuPageListResp struct {
	// List 数据列表
	List []MenuPageListItem `json:"list"`

	// Total 数据总条数
	Total int64 `json:"total"`
}

type RoleCreateResp struct {
	// ID 数据自增id
	ID uint `json:"id"`
}

type RoleDetailResp struct {
	// ID 数据自增id
	ID uint `json:"id" validate:"required"`

	objpermission.RoleBaseInfo
	objcommon.OperatorBaseInfo
}

type RolePageListItem struct {
	// ID 数据自增id
	ID uint `json:"id" validate:"required"`

	objpermission.RoleBaseInfo
	objcommon.OperatorBaseInfo
}

type RolePageListResp struct {
	// List 数据列表
	List []RolePageListItem `json:"list"`

	// Total 数据总条数
	Total int64 `json:"total"`
}
