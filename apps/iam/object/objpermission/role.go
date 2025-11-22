package objpermission

type RoleBaseInfo struct {
	CompanyID   uint   `json:"companyId" form:"companyId"`     // 所属公司ID(租户ID)
	DataScope   string `json:"dataScope" form:"dataScope"`     // 数据权限范围: all-全部 dept_and_sub-本部门及以下 dept-本部门 self-仅本人 custom-自定义
	Description string `json:"description" form:"description"` // 角色描述
	RoleCode    string `json:"roleCode" form:"roleCode"`       // 角色编码
	RoleName    string `json:"roleName" form:"roleName"`       // 角色名称
	RoleType    string `json:"roleType" form:"roleType"`       // 角色类型: custom-自定义 system-系统内置
	SortOrder   int32  `json:"sortOrder" form:"sortOrder"`     // 排序
	Status      string `json:"status" form:"status"`           // 状态: active-正常 inactive-停用
}
