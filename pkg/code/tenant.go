package code

import "github.com/morehao/golib/gerror"

const (
	TenantCreateError      = 110100
	TenantDeleteError      = 110101
	TenantUpdateError      = 110102
	TenantGetDetailError   = 110103
	TenantGetPageListError = 110104
	TenantNotExistError    = 110105
)

var tenantErrorMsgMap = gerror.CodeMsgMap{
	TenantCreateError:      "创建租户管理失败",
	TenantDeleteError:      "删除租户管理失败",
	TenantUpdateError:      "修改租户管理失败",
	TenantGetDetailError:   "查看租户管理失败",
	TenantGetPageListError: "查看租户管理列表失败",
	TenantNotExistError:    "租户管理不存在",
}
