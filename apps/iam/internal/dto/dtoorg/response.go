package dtoorg

import (
	"github.com/morehao/goark/apps/iam/object/objcommon"
	"github.com/morehao/goark/apps/iam/object/objorg"
)

type CompanyCreateResp struct {
	// ID 数据自增id
	ID uint `json:"id"`
}

type CompanyDetailResp struct {
	// ID 数据自增id
	ID uint `json:"id" validate:"required"`

	objorg.CompanyBaseInfo
	objcommon.OperatorBaseInfo
}

type CompanyPageListItem struct {
	// ID 数据自增id
	ID uint `json:"id" validate:"required"`

	objorg.CompanyBaseInfo
	objcommon.OperatorBaseInfo
}

type CompanyPageListResp struct {
	// List 数据列表
	List []CompanyPageListItem `json:"list"`

	// Total 数据总条数
	Total int64 `json:"total"`
}

type DepartmentCreateResp struct {
	// ID 数据自增id
	ID uint `json:"id"`
}

type DepartmentDetailResp struct {
	// ID 数据自增id
	ID uint `json:"id" validate:"required"`

	objorg.DepartmentBaseInfo
	objcommon.OperatorBaseInfo
}

type DepartmentPageListItem struct {
	// ID 数据自增id
	ID uint `json:"id" validate:"required"`

	objorg.DepartmentBaseInfo
	objcommon.OperatorBaseInfo
}

type DepartmentPageListResp struct {
	// List 数据列表
	List []DepartmentPageListItem `json:"list"`

	// Total 数据总条数
	Total int64 `json:"total"`
}
