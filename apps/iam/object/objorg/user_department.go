package objorg

type UserDepartmentBaseInfo struct {
	// CompanyID 公司ID(租户ID,冗余)
	CompanyID uint `json:"companyID" form:"companyID"`
	// DeptID 部门ID
	DeptID uint `json:"deptID" form:"deptID"`
	// DeptType 部门类型: primary-主部门 secondary-其他部门
	DeptType string `json:"deptType" form:"deptType"`
	// UserID 用户ID
	UserID uint `json:"userID" form:"userID"`
}
