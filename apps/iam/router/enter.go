package router

import "github.com/gin-gonic/gin"

func RegisterRouter(routerGroup *gin.RouterGroup) {
	tenantRouter(routerGroup)
	companyRouter(routerGroup)
	departmentRouter(routerGroup)
	userRouter(routerGroup)
	menuRouter(routerGroup)
	roleRouter(routerGroup)
}
