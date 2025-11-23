package router

import (
	"github.com/morehao/goark/apps/iam/internal/controller/ctrorg"

	"github.com/gin-gonic/gin"
)

// companyRouter 初始化公司管理路由信息
func companyRouter(routerGroup *gin.RouterGroup) {
	companyCtr := ctrorg.NewCompanyCtr()

	routerGroup.POST("/company/create", companyCtr.Create)
	routerGroup.POST("/company/delete", companyCtr.Delete)
	routerGroup.POST("/company/update", companyCtr.Update)
	routerGroup.GET("/company/detail", companyCtr.Detail)
	routerGroup.POST("/company/pageList", companyCtr.PageList)
}

// departmentRouter 初始化部门管理路由信息
func departmentRouter(routerGroup *gin.RouterGroup) {
	departmentCtr := ctrorg.NewDepartmentCtr()
	routerGroup.POST("/department/create", departmentCtr.Create)
	routerGroup.POST("/department/delete", departmentCtr.Delete)
	routerGroup.POST("/department/update", departmentCtr.Update)
	routerGroup.GET("/department/detail", departmentCtr.Detail)
	routerGroup.POST("/department/pageList", departmentCtr.PageList)
}
