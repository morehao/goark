package objpermission

type RoleMenuBaseInfo struct {
	// CompanyID 公司ID(租户ID,冗余)
	CompanyID uint `json:"companyId" form:"companyId"`

	// MenuID 菜单ID
	MenuID uint `json:"menuId" form:"menuId"`

	// RoleID 角色ID
	RoleID uint `json:"roleId" form:"roleId"`
}
