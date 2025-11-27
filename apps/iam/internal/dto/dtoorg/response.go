package dtoorg

import (
	"github.com/morehao/goark/apps/iam/object/objorg"
	"github.com/morehao/golib/biz/gobject"
)

type CompanyCreateResp struct {
	// ID 数据自增 ID
	ID uint `json:"id"`
}

type CompanyDetailResp struct {
	// ID 数据自增 ID
	ID uint `json:"id" validate:"required"`
	objorg.CompanyBaseInfo
	gobject.OperatorBaseInfo
}

type CompanyPageListItem struct {
	// ID 数据自增 ID
	ID uint `json:"id" validate:"required"`
	objorg.CompanyBaseInfo
	gobject.OperatorBaseInfo
}

type CompanyPageListResp struct {
	// List 数据列表
	List []CompanyPageListItem `json:"list"`
	// Total 数据总条数
	Total int64 `json:"total"`
}

type DepartmentCreateResp struct {
	// ID 数据自增 ID
	ID uint `json:"id"`
}
type DepartmentDetailResp struct {
	// ID 数据自增 ID
	ID uint `json:"id" validate:"required"`
	objorg.DepartmentBaseInfo
	gobject.OperatorBaseInfo
}
type DepartmentPageListItem struct {
	// ID 数据自增 ID
	ID uint `json:"id" validate:"required"`
	objorg.DepartmentBaseInfo
	gobject.OperatorBaseInfo
}
type DepartmentPageListResp struct {
	// List 数据列表
	List []DepartmentPageListItem `json:"list"`
	// Total 数据总条数
	Total int64 `json:"total"`
}
