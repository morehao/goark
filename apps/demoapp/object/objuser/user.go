package objuser

type UserBaseInfo struct {
	// CompanyID 公司id
	CompanyID uint `json:"companyId" form:"companyId"`

	// DepartmentID 部门id
	DepartmentID uint `json:"departmentId" form:"departmentId"`

	// Name 姓名
	Name string `json:"name" form:"name"`
}
