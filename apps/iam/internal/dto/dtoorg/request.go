package dtoorg

import (
	"github.com/morehao/goark/apps/iam/object/objorg"
	"github.com/morehao/golib/gobject"
)

type CompanyCreateReq struct {
	objorg.CompanyBaseInfo
}

type CompanyUpdateReq struct {
	// ID 数据自增 ID
	ID uint `json:"id" validate:"required" label:"数据自增id"`
	objorg.CompanyBaseInfo
}

type CompanyDetailReq struct {
	// ID 数据自增 ID
	ID uint `json:"id" form:"id" validate:"required" label:"数据自增id"`
}

type CompanyPageListReq struct {
	gobject.PageQuery
}

type CompanyDeleteReq struct {
	// ID 数据自增 ID
	ID uint `json:"id" form:"id" validate:"required" label:"数据自增id"`
}

type DepartmentCreateReq struct {
	objorg.DepartmentBaseInfo
}
type DepartmentUpdateReq struct {
	// ID 数据自增 ID
	ID uint `json:"id" validate:"required" label:"数据自增id"`
	objorg.DepartmentBaseInfo
}
type DepartmentDetailReq struct {
	// ID 数据自增 ID
	ID uint `json:"id" form:"id" validate:"required" label:"数据自增id"`
}
type DepartmentPageListReq struct {
	gobject.PageQuery
}
type DepartmentDeleteReq struct {
	// ID 数据自增 ID
	ID uint `json:"id" form:"id" validate:"required" label:"数据自增id"`
}
