package router

import (
	"github.com/morehao/goark/apps/iam/internal/controller/ctrtenant"

	"github.com/gin-gonic/gin"
)

// tenantRouter 初始化租户管理路由信息
func tenantRouter(routerGroup *gin.RouterGroup) {
	tenantCtr := ctrtenant.NewTenantCtr()
	routerGroup.POST("/tenant/create", tenantCtr.Create)    // 新建租户管理
	routerGroup.POST("/tenant/delete", tenantCtr.Delete)    // 删除租户管理
	routerGroup.POST("/tenant/update", tenantCtr.Update)    // 更新租户管理
	routerGroup.GET("/tenant/detail", tenantCtr.Detail)     // 根据ID获取租户管理
	routerGroup.GET("/tenant/pageList", tenantCtr.PageList) // 获取租户管理列表
}
