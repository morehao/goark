package router

import (
	"github.com/morehao/goark/apps/iam/internal/controller/ctrtenant"

	"github.com/gin-gonic/gin"
)

// tenantRouter 初始化租户管理路由信息
func tenantRouter(routerGroup *gin.RouterGroup) {
	tenantCtr := ctrtenant.NewTenantCtr()

	routerGroup.POST("/tenant/create", tenantCtr.Create)
	routerGroup.POST("/tenant/delete", tenantCtr.Delete)
	routerGroup.POST("/tenant/update", tenantCtr.Update)
	routerGroup.GET("/tenant/detail", tenantCtr.Detail)
	routerGroup.POST("/tenant/pageList", tenantCtr.PageList)
}
