package objpermission

type RoleMenuBaseInfo struct {
	CompanyID uint `json:"companyId" form:"companyId"` // 公司ID(租户ID,冗余)
	MenuID    uint `json:"menuId" form:"menuId"`       // 菜单ID
	RoleID    uint `json:"roleId" form:"roleId"`       // 角色ID
}
