package router

import (
	"github.com/gin-gonic/gin"
	"github.com/morehao/goark/apps/iam/config"
	"github.com/morehao/golib/grouter/ginrouter"
)

type RouterGroups struct {
	AuthGroup   *gin.RouterGroup
	NoAuthGroup *gin.RouterGroup
}

func RegisterRouter(groups *RouterGroups, appName string) {
	if config.Conf.Server.Env == "dev" {
		ginrouter.RegisterSwagger(groups.AuthGroup, appName)
	}
	v1Auth := groups.AuthGroup.Group("/v1")
	// v1NoAuth := groups.NoAuthGroup.Group("/v1")
	tenantRouter(v1Auth)
	companyRouter(v1Auth)
	departmentRouter(v1Auth)
	userRouter(v1Auth)
	menuRouter(v1Auth)
	roleRouter(v1Auth)
}
