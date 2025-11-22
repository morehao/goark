package objtenant

type TenantBaseInfo struct {
	Description string `json:"description" form:"description"` // 租户描述
	SortOrder   int32  `json:"sortOrder" form:"sortOrder"`     // 排序
	Status      string `json:"status" form:"status"`           // 状态: active-正常 inactive-停用
	TenantCode  string `json:"tenantCode" form:"tenantCode"`   // 租户编码
	TenantName  string `json:"tenantName" form:"tenantName"`   // 租户名称
}
