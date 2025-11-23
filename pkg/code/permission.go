package code

import "github.com/morehao/golib/gerror"

const (
	MenuCreateError      = 110500
	MenuDeleteError      = 110501
	MenuUpdateError      = 110502
	MenuGetDetailError   = 110503
	MenuGetPageListError = 110504
	MenuNotExistError    = 110505
)

var menuErrorMsgMap = gerror.CodeMsgMap{
	MenuCreateError:      "创建菜单管理失败",
	MenuDeleteError:      "删除菜单管理失败",
	MenuUpdateError:      "修改菜单管理失败",
	MenuGetDetailError:   "查看菜单管理失败",
	MenuGetPageListError: "查看菜单管理列表失败",
	MenuNotExistError:    "菜单管理不存在",
}

const (
	RoleCreateError      = 110600
	RoleDeleteError      = 110601
	RoleUpdateError      = 110602
	RoleGetDetailError   = 110603
	RoleGetPageListError = 110604
	RoleNotExistError    = 110605
)

var roleErrorMsgMap = gerror.CodeMsgMap{
	RoleCreateError:      "创建角色管理失败",
	RoleDeleteError:      "删除角色管理失败",
	RoleUpdateError:      "修改角色管理失败",
	RoleGetDetailError:   "查看角色管理失败",
	RoleGetPageListError: "查看角色管理列表失败",
	RoleNotExistError:    "角色管理不存在",
}
