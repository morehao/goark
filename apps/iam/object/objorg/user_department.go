package objorg

type UserDepartmentBaseInfo struct {
	// CompanyID 公司ID(租户ID,冗余)
	CompanyID uint `json:"companyId" form:"companyId"`

	// DeptID 部门ID
	DeptID uint `json:"deptId" form:"deptId"`

	// DeptType 部门类型: primary-主部门 secondary-其他部门
	DeptType string `json:"deptType" form:"deptType"`

	// UserID 用户ID
	UserID uint `json:"userId" form:"userId"`
}
