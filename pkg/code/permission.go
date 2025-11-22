package code

import "github.com/morehao/golib/gerror"

const (
	MenuCreateError      = 100100
	MenuDeleteError      = 100101
	MenuUpdateError      = 100102
	MenuGetDetailError   = 100103
	MenuGetPageListError = 100104
	MenuNotExistError    = 100105
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
	RoleCreateError      = 100100
	RoleDeleteError      = 100101
	RoleUpdateError      = 100102
	RoleGetDetailError   = 100103
	RoleGetPageListError = 100104
	RoleNotExistError    = 100105
)

var roleErrorMsgMap = gerror.CodeMsgMap{
	RoleCreateError:      "创建角色管理失败",
	RoleDeleteError:      "删除角色管理失败",
	RoleUpdateError:      "修改角色管理失败",
	RoleGetDetailError:   "查看角色管理失败",
	RoleGetPageListError: "查看角色管理列表失败",
	RoleNotExistError:    "角色管理不存在",
}
