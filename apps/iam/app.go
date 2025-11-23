package iam

import (
	"github.com/gin-gonic/gin"
	_ "github.com/morehao/goark/apps/iam/docs"
	"github.com/morehao/goark/apps/iam/router"
)

const AppName = "iam"

func Routers(engine *gin.Engine) {
	routerGroup := engine.Group(AppName)
	routerGroups := &router.RouterGroups{
		AuthGroup:   routerGroup,
		NoAuthGroup: routerGroup,
	}
	router.RegisterRouter(routerGroups, AppName)
}
