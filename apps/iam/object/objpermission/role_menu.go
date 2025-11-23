package objpermission

type RoleMenuBaseInfo struct {
	// CompanyID 公司ID(租户ID,冗余)
	CompanyID uint `json:"companyID" form:"companyID"`
	// MenuID 菜单ID
	MenuID uint `json:"menuID" form:"menuID"`
	// RoleID 角色ID
	RoleID uint `json:"roleID" form:"roleID"`
}
