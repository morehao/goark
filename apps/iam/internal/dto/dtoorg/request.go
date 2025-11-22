package dtoorg

import (
	"github.com/morehao/goark/apps/iam/object/objcommon"
	"github.com/morehao/goark/apps/iam/object/objorg"
)

type CompanyCreateReq struct {
	objorg.CompanyBaseInfo
}

type CompanyUpdateReq struct {
	// ID 数据自增id
	ID uint `json:"id" validate:"required" label:"数据自增id"`

	objorg.CompanyBaseInfo
}

type CompanyDetailReq struct {
	// ID 数据自增id
	ID uint `json:"id" form:"id" validate:"required" label:"数据自增id"`
}

type CompanyPageListReq struct {
	objcommon.PageQuery
}

type CompanyDeleteReq struct {
	// ID 数据自增id
	ID uint `json:"id" form:"id" validate:"required" label:"数据自增id"`
}

type DepartmentCreateReq struct {
	objorg.DepartmentBaseInfo
}

type DepartmentUpdateReq struct {
	// ID 数据自增id
	ID uint `json:"id" validate:"required" label:"数据自增id"`

	objorg.DepartmentBaseInfo
}

type DepartmentDetailReq struct {
	// ID 数据自增id
	ID uint `json:"id" form:"id" validate:"required" label:"数据自增id"`
}

type DepartmentPageListReq struct {
	objcommon.PageQuery
}

type DepartmentDeleteReq struct {
	// ID 数据自增id
	ID uint `json:"id" form:"id" validate:"required" label:"数据自增id"`
}
