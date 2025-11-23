package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morehao/goark/apps/iam/config"
	_ "github.com/morehao/goark/apps/iam/docs"
	"github.com/morehao/goark/apps/iam/router"
	"github.com/morehao/golib/glog"
	"github.com/morehao/golib/gmiddleware/ginmiddleware"
	"github.com/morehao/golib/grouter/ginrouter"
)

func main() {
	if err := serverInit(); err != nil {
		panic(fmt.Sprintf("server init failed, error: %v", err))
	}
	if config.Conf.Server.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	defer glog.Close()

	engine := gin.New()
	engine.Use(gin.Recovery())
	routerGroup := engine.Group("iam")
	routerGroup.Use(ginmiddleware.AccessLog())
	if config.Conf.Server.Env == "dev" {
		ginrouter.RegisterSwagger(routerGroup, config.Conf.Server.Name)
	}
	routerGroups := &router.RouterGroups{
		AuthGroup:   routerGroup,
		NoAuthGroup: routerGroup,
	}
	router.RegisterRouter(routerGroups)

	if err := engine.Run(fmt.Sprintf(":%s", config.Conf.Server.Port)); err != nil {
		glog.Errorf(context.Background(), "%s run fail, port:%s", config.Conf.Server.Name, config.Conf.Server.Port)
		panic(err)
	} else {
		glog.Infof(context.Background(), "%s run success, port:%s", config.Conf.Server.Name, config.Conf.Server.Port)
	}
}
