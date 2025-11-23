package code

import "github.com/morehao/golib/gerror"

const (
	CompanyCreateError      = 110200
	CompanyDeleteError      = 110201
	CompanyUpdateError      = 110202
	CompanyGetDetailError   = 110203
	CompanyGetPageListError = 110204
	CompanyNotExistError    = 110205
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
	DepartmentCreateError      = 110300
	DepartmentDeleteError      = 110301
	DepartmentUpdateError      = 110302
	DepartmentGetDetailError   = 110303
	DepartmentGetPageListError = 110304
	DepartmentNotExistError    = 110305
)

var departmentErrorMsgMap = gerror.CodeMsgMap{
	DepartmentCreateError:      "创建部门管理失败",
	DepartmentDeleteError:      "删除部门管理失败",
	DepartmentUpdateError:      "修改部门管理失败",
	DepartmentGetDetailError:   "查看部门管理失败",
	DepartmentGetPageListError: "查看部门管理列表失败",
	DepartmentNotExistError:    "部门管理不存在",
}
