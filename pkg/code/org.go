package code

import "github.com/morehao/golib/gerror"

const (
	CompanyCreateError      = 100100
	CompanyDeleteError      = 100101
	CompanyUpdateError      = 100102
	CompanyGetDetailError   = 100103
	CompanyGetPageListError = 100104
	CompanyNotExistError    = 100105
)

var companyErrorMsgMap = gerror.CodeMsgMap{
	CompanyCreateError:      "创建公司管理失败",
	CompanyDeleteError:      "删除公司管理失败",
	CompanyUpdateError:      "修改公司管理失败",
	CompanyGetDetailError:   "查看公司管理失败",
	CompanyGetPageListError: "查看公司管理列表失败",
	CompanyNotExistError:    "公司管理不存在",
}

const (
	DepartmentCreateError      = 100100
	DepartmentDeleteError      = 100101
	DepartmentUpdateError      = 100102
	DepartmentGetDetailError   = 100103
	DepartmentGetPageListError = 100104
	DepartmentNotExistError    = 100105
)

var departmentErrorMsgMap = gerror.CodeMsgMap{
	DepartmentCreateError:      "创建部门管理失败",
	DepartmentDeleteError:      "删除部门管理失败",
	DepartmentUpdateError:      "修改部门管理失败",
	DepartmentGetDetailError:   "查看部门管理失败",
	DepartmentGetPageListError: "查看部门管理列表失败",
	DepartmentNotExistError:    "部门管理不存在",
}
