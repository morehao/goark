package router

import (
	"github.com/morehao/goark/apps/iam/internal/controller/ctrpermission"

	"github.com/gin-gonic/gin"
)

// menuRouter 初始化菜单管理路由信息
func menuRouter(routerGroup *gin.RouterGroup) {
	menuCtr := ctrpermission.NewMenuCtr()

	routerGroup.POST("/menu/create", menuCtr.Create)
	routerGroup.POST("/menu/delete", menuCtr.Delete)
	routerGroup.POST("/menu/update", menuCtr.Update)
	routerGroup.GET("/menu/detail", menuCtr.Detail)
	routerGroup.POST("/menu/pageList", menuCtr.PageList)
}

// roleRouter 初始化角色管理路由信息
func roleRouter(routerGroup *gin.RouterGroup) {
	roleCtr := ctrpermission.NewRoleCtr()
	routerGroup.POST("/role/create", roleCtr.Create)
	routerGroup.POST("/role/delete", roleCtr.Delete)
	routerGroup.POST("/role/update", roleCtr.Update)
	routerGroup.GET("/role/detail", roleCtr.Detail)
	routerGroup.POST("/role/pageList", roleCtr.PageList)
}
