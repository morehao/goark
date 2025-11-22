package router

import (
	"github.com/morehao/goark/apps/iam/internal/controller/ctrorg"

	"github.com/gin-gonic/gin"
)

// orgRouter 初始化公司管理路由信息
func companyRouter(routerGroup *gin.RouterGroup) {
	companyCtr := ctrorg.NewCompanyCtr()

	routerGroup.POST("/company/create", companyCtr.Create)    // 新建公司管理
	routerGroup.POST("/company/delete", companyCtr.Delete)    // 删除公司管理
	routerGroup.POST("/company/update", companyCtr.Update)    // 更新公司管理
	routerGroup.GET("/company/detail", companyCtr.Detail)     // 根据ID获取公司管理
	routerGroup.GET("/company/pageList", companyCtr.PageList) // 获取公司管理列表
}

// orgRouter 初始化部门管理路由信息
func departmentRouter(routerGroup *gin.RouterGroup) {
	departmentCtr := ctrorg.NewDepartmentCtr()

	routerGroup.POST("/department/create", departmentCtr.Create)    // 新建部门管理
	routerGroup.POST("/department/delete", departmentCtr.Delete)    // 删除部门管理
	routerGroup.POST("/department/update", departmentCtr.Update)    // 更新部门管理
	routerGroup.GET("/department/detail", departmentCtr.Detail)     // 根据ID获取部门管理
	routerGroup.GET("/department/pageList", departmentCtr.PageList) // 获取部门管理列表
}
