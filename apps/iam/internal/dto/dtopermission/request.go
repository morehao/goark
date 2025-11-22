package dtopermission

import (
	"github.com/morehao/goark/apps/iam/object/objcommon"
	"github.com/morehao/goark/apps/iam/object/objpermission"
)

type MenuCreateReq struct {
	objpermission.MenuBaseInfo
}

type MenuUpdateReq struct {
	ID uint `json:"id" validate:"required" label:"数据自增id"` // 数据自增id
	objpermission.MenuBaseInfo
}

type MenuDetailReq struct {
	ID uint `json:"id" form:"id" validate:"required" label:"数据自增id"` // 数据自增id
}

type MenuPageListReq struct {
	objcommon.PageQuery
}

type MenuDeleteReq struct {
	ID uint `json:"id" form:"id" validate:"required" label:"数据自增id"` // 数据自增id
}

type RoleCreateReq struct {
	objpermission.RoleBaseInfo
}
type RoleUpdateReq struct {
	ID uint `json:"id" validate:"required" label:"数据自增id"` // 数据自增id
	objpermission.RoleBaseInfo
}
type RoleDetailReq struct {
	ID uint `json:"id" form:"id" validate:"required" label:"数据自增id"` // 数据自增id
}
type RolePageListReq struct {
	objcommon.PageQuery
}
type RoleDeleteReq struct {
	ID uint `json:"id" form:"id" validate:"required" label:"数据自增id"` // 数据自增id
}
