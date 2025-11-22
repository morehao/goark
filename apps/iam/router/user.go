package router

import (
	"github.com/morehao/goark/apps/iam/internal/controller/ctruser"

	"github.com/gin-gonic/gin"
)

// userRouter 初始化用户管理路由信息
func userRouter(routerGroup *gin.RouterGroup) {
	userCtr := ctruser.NewUserCtr()

	routerGroup.POST("/user/create", userCtr.Create)    // 新建用户管理
	routerGroup.POST("/user/delete", userCtr.Delete)    // 删除用户管理
	routerGroup.POST("/user/update", userCtr.Update)    // 更新用户管理
	routerGroup.GET("/user/detail", userCtr.Detail)     // 根据ID获取用户管理
	routerGroup.GET("/user/pageList", userCtr.PageList) // 获取用户管理列表
}
