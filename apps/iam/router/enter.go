package router

import "github.com/gin-gonic/gin"

// RouterGroups 路由组配置
type RouterGroups struct {
	// AuthGroup 需要鉴权的路由组
	AuthGroup *gin.RouterGroup
	// NoAuthGroup 不需要鉴权的路由组
	NoAuthGroup *gin.RouterGroup
}

func RegisterRouter(groups *RouterGroups) {
	v1Auth := groups.AuthGroup.Group("/v1")
	// v1NoAuth := groups.NoAuthGroup.Group("/v1")
	tenantRouter(v1Auth)
	companyRouter(v1Auth)
	departmentRouter(v1Auth)
	userRouter(v1Auth)
	menuRouter(v1Auth)
	roleRouter(v1Auth)
}
