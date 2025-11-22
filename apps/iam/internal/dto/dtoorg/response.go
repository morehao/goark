package dtoorg

import (
	"github.com/morehao/goark/apps/iam/object/objcommon"
	"github.com/morehao/goark/apps/iam/object/objorg"
)

type CompanyCreateResp struct {
	ID uint `json:"id"` // 数据自增id
}

type CompanyDetailResp struct {
	ID uint `json:"id" validate:"required"` // 数据自增id
	objorg.CompanyBaseInfo
	objcommon.OperatorBaseInfo
}

type CompanyPageListItem struct {
	ID uint `json:"id" validate:"required"` // 数据自增id
	objorg.CompanyBaseInfo
	objcommon.OperatorBaseInfo
}

type CompanyPageListResp struct {
	List  []CompanyPageListItem `json:"list"`  // 数据列表
	Total int64                 `json:"total"` // 数据总条数
}

type DepartmentCreateResp struct {
	ID uint `json:"id"` // 数据自增id
}
type DepartmentDetailResp struct {
	ID uint `json:"id" validate:"required"` // 数据自增id
	objorg.DepartmentBaseInfo
	objcommon.OperatorBaseInfo
}
type DepartmentPageListItem struct {
	ID uint `json:"id" validate:"required"` // 数据自增id
	objorg.DepartmentBaseInfo
	objcommon.OperatorBaseInfo
}
type DepartmentPageListResp struct {
	List  []DepartmentPageListItem `json:"list"`  // 数据列表
	Total int64                    `json:"total"` // 数据总条数
}
