package router

import (
	"github.com/gin-gonic/gin"
	"github.com/morehao/goark/apps/iam/config"
	"github.com/morehao/golib/grouter/ginrouter"
)

type RouterGroups struct {
	AuthGroup	*gin.RouterGroup
	NoAuthGroup	*gin.RouterGroup
}

func RegisterRouter(groups *RouterGroups, appName string) {
	if config.Conf.Server.Env == "dev" {
		ginrouter.RegisterSwagger(groups.AuthGroup, appName)
	}
	v1AuthGroup := groups.AuthGroup.Group("/v1")
	// v1NoAuth := groups.NoAuthGroup.Group("/v1")
	tenantRouter(v1AuthGroup)
	companyRouter(v1AuthGroup)
	departmentRouter(v1AuthGroup)
	userRouter(v1AuthGroup)
	menuRouter(v1AuthGroup)
	roleRouter(v1AuthGroup)
}
