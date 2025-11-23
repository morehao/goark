package objpermission

type UserRoleBaseInfo struct {
	// CompanyID 公司ID(租户ID,冗余)
	CompanyID uint `json:"companyID" form:"companyID"`
	// RoleID 角色ID
	RoleID uint `json:"roleID" form:"roleID"`
	// UserID 用户ID
	UserID uint `json:"userID" form:"userID"`
}
