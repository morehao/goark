package dtouser

import (
	"github.com/morehao/goark/apps/demoapp/object/objcommon"
	"github.com/morehao/goark/apps/demoapp/object/objuser"
)

type UserCreateResp struct {
	// ID 数据自增id
	ID uint `json:"id"`
}

type UserDetailResp struct {
	// ID 数据自增id
	ID uint `json:"id" validate:"required"`

	objuser.UserBaseInfo
	objcommon.OperatorBaseInfo
}

type UserPageListItem struct {
	// ID 数据自增id
	ID uint `json:"id" validate:"required"`

	objuser.UserBaseInfo
	objcommon.OperatorBaseInfo
}

type UserPageListResp struct {
	// List 数据列表
	List []UserPageListItem `json:"list"`

	// Total 数据总条数
	Total int64 `json:"total"`
}
