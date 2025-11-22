package router

import (
	"github.com/morehao/goark/apps/iam/internal/controller/ctrpermission"

	"github.com/gin-gonic/gin"
)

// permissionRouter 初始化菜单管理路由信息
func menuRouter(routerGroup *gin.RouterGroup) {
	menuCtr := ctrpermission.NewMenuCtr()

	routerGroup.POST("/menu/create", menuCtr.Create)    // 新建菜单管理
	routerGroup.POST("/menu/delete", menuCtr.Delete)    // 删除菜单管理
	routerGroup.POST("/menu/update", menuCtr.Update)    // 更新菜单管理
	routerGroup.GET("/menu/detail", menuCtr.Detail)     // 根据ID获取菜单管理
	routerGroup.GET("/menu/pageList", menuCtr.PageList) // 获取菜单管理列表
}

// permissionRouter 初始化角色管理路由信息
func roleRouter(routerGroup *gin.RouterGroup) {
	roleCtr := ctrpermission.NewRoleCtr()

	routerGroup.POST("/role/create", roleCtr.Create)    // 新建角色管理
	routerGroup.POST("/role/delete", roleCtr.Delete)    // 删除角色管理
	routerGroup.POST("/role/update", roleCtr.Update)    // 更新角色管理
	routerGroup.GET("/role/detail", roleCtr.Detail)     // 根据ID获取角色管理
	routerGroup.GET("/role/pageList", roleCtr.PageList) // 获取角色管理列表
}
