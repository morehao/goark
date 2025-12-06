package objtenant

type TenantConfigBaseInfo struct {
	// ConfigGroup 配置分组: general-通用/auth-认证/theme-主题等
	ConfigGroup string `json:"configGroup" form:"configGroup"`
	// ConfigKey 配置键
	ConfigKey string `json:"configKey" form:"configKey"`
	// ConfigType 配置类型: string/json/boolean/number
	ConfigType string `json:"configType" form:"configType"`
	// ConfigValue 配置值(支持JSON)
	ConfigValue string `json:"configValue" form:"configValue"`
	// Description 配置说明
	Description string `json:"description" form:"description"`
	// SortOrder 排序
	SortOrder int32 `json:"sortOrder" form:"sortOrder"`
	// TenantID 租户ID
	TenantID uint `json:"tenantID" form:"tenantID"`
}
