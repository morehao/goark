package objuser

type UserRoleBaseInfo struct {
	CompanyID uint `json:"companyId" form:"companyId"` // 公司ID(租户ID,冗余)
	RoleID    uint `json:"roleId" form:"roleId"`       // 角色ID
	UserID    uint `json:"userId" form:"userId"`       // 用户ID
}
