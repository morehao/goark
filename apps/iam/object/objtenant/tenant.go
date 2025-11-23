package objtenant

type TenantBaseInfo struct {
	// Description 租户描述
	Description string `json:"description" form:"description"`
	// SortOrder 排序
	SortOrder int32 `json:"sortOrder" form:"sortOrder"`
	// Status 状态: active-正常 inactive-停用
	Status string `json:"status" form:"status"`
	// TenantCode 租户编码
	TenantCode string `json:"tenantCode" form:"tenantCode"`
	// TenantName 租户名称
	TenantName string `json:"tenantName" form:"tenantName"`
}
