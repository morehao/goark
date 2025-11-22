package objpermission

type MenuBaseInfo struct {
	CacheType     string `json:"cacheType" form:"cacheType"`         // 缓存类型: enabled-启用 disabled-禁用
	CompanyID     uint   `json:"companyId" form:"companyId"`         // 所属公司ID(租户ID)
	ComponentPath string `json:"componentPath" form:"componentPath"` // 组件路径
	Icon          string `json:"icon" form:"icon"`                   // 菜单图标
	LinkType      string `json:"linkType" form:"linkType"`           // 链接类型: internal-内部链接 external-外部链接
	MenuCode      string `json:"menuCode" form:"menuCode"`           // 菜单编码
	MenuName      string `json:"menuName" form:"menuName"`           // 菜单名称
	MenuType      string `json:"menuType" form:"menuType"`           // 菜单类型: directory-目录 menu-菜单 button-按钮
	ParentID      uint   `json:"parentId" form:"parentId"`           // 父菜单ID
	Permission    string `json:"permission" form:"permission"`       // 权限标识: sys:user:add
	RoutePath     string `json:"routePath" form:"routePath"`         // 路由地址
	SortOrder     int32  `json:"sortOrder" form:"sortOrder"`         // 排序
	Status        string `json:"status" form:"status"`               // 状态: active-正常 inactive-停用
	Visibility    string `json:"visibility" form:"visibility"`       // 可见性: visible-可见 hidden-隐藏
}
