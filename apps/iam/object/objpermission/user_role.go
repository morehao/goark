package objpermission

type UserRoleBaseInfo struct {
	// CompanyID 公司ID(租户ID,冗余)
	CompanyID uint `json:"companyId" form:"companyId"`

	// RoleID 角色ID
	RoleID uint `json:"roleId" form:"roleId"`

	// UserID 用户ID
	UserID uint `json:"userId" form:"userId"`
}
