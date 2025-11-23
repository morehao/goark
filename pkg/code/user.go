package code

import "github.com/morehao/golib/gerror"

const (
	UserCreateError      = 110400
	UserDeleteError      = 110401
	UserUpdateError      = 110402
	UserGetDetailError   = 110403
	UserGetPageListError = 110404
	UserNotExistError    = 110405
)

var userErrorMsgMap = gerror.CodeMsgMap{
	UserCreateError:      "创建用户管理失败",
	UserDeleteError:      "删除用户管理失败",
	UserUpdateError:      "修改用户管理失败",
	UserGetDetailError:   "查看用户管理失败",
	UserGetPageListError: "查看用户管理列表失败",
	UserNotExistError:    "用户管理不存在",
}
