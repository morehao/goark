package objuser

type UserBaseInfo struct {
	// CompanyID 所属公司ID(租户ID), 0表示平台管理员账号
	CompanyID uint `json:"companyID" form:"companyID"`
	// DeptID 主部门ID(冗余字段,方便查询,实际关联关系在iam_user_department表)
	DeptID uint `json:"deptID" form:"deptID"`
	// EmployeeNo 工号
	EmployeeNo string `json:"employeeNo" form:"employeeNo"`
	// EntryDate 入职日期
	EntryDate int64 `json:"entryDate" form:"entryDate"`
	// JobLevel 职级
	JobLevel string `json:"jobLevel" form:"jobLevel"`
	// LastLoginAt 最后登录时间
	LastLoginAt int64 `json:"lastLoginAt" form:"lastLoginAt"`
	// LastLoginIp 最后登录IP(支持IPv6)
	LastLoginIp string `json:"lastLoginIp" form:"lastLoginIp"`
	// LoginCount 登录次数
	LoginCount int32 `json:"loginCount" form:"loginCount"`
	// PersonID 自然人ID
	PersonID uint `json:"personID" form:"personID"`
	// Position 职位
	Position string `json:"position" form:"position"`
	// Status 状态: active-正常 locked-锁定 disabled-禁用
	Status string `json:"status" form:"status"`
	// UserType 用户类型: normal-普通用户 company_admin-公司管理员 platform_admin-平台管理员(可管理所有公司)
	UserType string `json:"userType" form:"userType"`
	// Username 用户名(公司用户:公司内唯一,平台管理员:全局唯一,需应用层保证)
	Username string `json:"username" form:"username"`
}
