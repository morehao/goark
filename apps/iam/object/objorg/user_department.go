package objorg

type UserDepartmentBaseInfo struct {
	CompanyID uint   `json:"companyId" form:"companyId"` // 公司ID(租户ID,冗余)
	DeptID    uint   `json:"deptId" form:"deptId"`       // 部门ID
	DeptType  string `json:"deptType" form:"deptType"`   // 部门类型: primary-主部门 secondary-其他部门
	UserID    uint   `json:"userId" form:"userId"`       // 用户ID
}
