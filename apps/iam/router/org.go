package router

import (
	"github.com/morehao/goark/apps/iam/internal/controller/ctrorg"

	"github.com/gin-gonic/gin"
)

// orgRouter 初始化公司管理路由信息
func companyRouter(routerGroup *gin.RouterGroup) {
	companyCtr := ctrorg.NewCompanyCtr()
	companyGroup := routerGroup.Group("company")
	{
		companyGroup.POST("create", companyCtr.Create)    // 新建公司管理
		companyGroup.POST("delete", companyCtr.Delete)    // 删除公司管理
		companyGroup.POST("update", companyCtr.Update)    // 更新公司管理
		companyGroup.GET("detail", companyCtr.Detail)     // 根据ID获取公司管理
		companyGroup.GET("pageList", companyCtr.PageList) // 获取公司管理列表
	}
}

// orgRouter 初始化部门管理路由信息
func departmentRouter(routerGroup *gin.RouterGroup) {
	departmentCtr := ctrorg.NewDepartmentCtr()
	departmentGroup := routerGroup.Group("department")
	{
		departmentGroup.POST("create", departmentCtr.Create)    // 新建部门管理
		departmentGroup.POST("delete", departmentCtr.Delete)    // 删除部门管理
		departmentGroup.POST("update", departmentCtr.Update)    // 更新部门管理
		departmentGroup.GET("detail", departmentCtr.Detail)     // 根据ID获取部门管理
		departmentGroup.GET("pageList", departmentCtr.PageList) // 获取部门管理列表
	}
}
