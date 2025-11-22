package router

import (
	"github.com/morehao/goark/apps/iam/internal/controller/ctrpermission"

	"github.com/gin-gonic/gin"
)

// permissionRouter 初始化菜单管理路由信息
func menuRouter(routerGroup *gin.RouterGroup) {
	menuCtr := ctrpermission.NewMenuCtr()
	menuGroup := routerGroup.Group("menu")
	{
		menuGroup.POST("create", menuCtr.Create)    // 新建菜单管理
		menuGroup.POST("delete", menuCtr.Delete)    // 删除菜单管理
		menuGroup.POST("update", menuCtr.Update)    // 更新菜单管理
		menuGroup.GET("detail", menuCtr.Detail)     // 根据ID获取菜单管理
		menuGroup.GET("pageList", menuCtr.PageList) // 获取菜单管理列表
	}
}

// permissionRouter 初始化角色管理路由信息
func roleRouter(routerGroup *gin.RouterGroup) {
	roleCtr := ctrpermission.NewRoleCtr()
	roleGroup := routerGroup.Group("role")
	{
		roleGroup.POST("create", roleCtr.Create)    // 新建角色管理
		roleGroup.POST("delete", roleCtr.Delete)    // 删除角色管理
		roleGroup.POST("update", roleCtr.Update)    // 更新角色管理
		roleGroup.GET("detail", roleCtr.Detail)     // 根据ID获取角色管理
		roleGroup.GET("pageList", roleCtr.PageList) // 获取角色管理列表
	}
}
