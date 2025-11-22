package objorg

type DepartmentBaseInfo struct {
	CompanyID uint   `json:"companyId" form:"companyId"` // 所属公司ID(租户ID)
	DeptCode  string `json:"deptCode" form:"deptCode"`   // 部门编码
	DeptLevel int32  `json:"deptLevel" form:"deptLevel"` // 部门层级
	DeptName  string `json:"deptName" form:"deptName"`   // 部门名称
	DeptPath  string `json:"deptPath" form:"deptPath"`   // 部门路径: /1/2/3/
	LeaderID  uint   `json:"leaderId" form:"leaderId"`   // 部门负责人ID
	ParentID  uint   `json:"parentId" form:"parentId"`   // 父部门ID,0表示根部门
	SortOrder int32  `json:"sortOrder" form:"sortOrder"` // 排序
	Status    string `json:"status" form:"status"`       // 状态: active-正常 inactive-停用
}
