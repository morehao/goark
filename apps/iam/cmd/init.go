package main

import (
	"fmt"

	"github.com/morehao/goark/apps/iam/config"
	"github.com/morehao/goark/pkg/storages"
	"github.com/morehao/golib/glog"
)

func serverInit() error {
	if err := preInit(); err != nil {
		return err
	}
	if err := resourceInit(); err != nil {
		return err
	}
	return nil
}

func preInit() error {
	config.InitConf()
	defaultLogCfg := config.Conf.Log["default"]
	if err := glog.InitLogger(&defaultLogCfg); err != nil {
		return fmt.Errorf("init logger failed: " + err.Error())
	}
	return nil
}

func resourceInit() error {
	if err := storages.InitMultiMysql(config.Conf.MysqlConfigs); err != nil {
		return fmt.Errorf("init mysql failed: " + err.Error())
	}
	if err := storages.InitMultiRedis(config.Conf.RedisConfigs); err != nil {
		return fmt.Errorf("init redis failed: " + err.Error())
	}
	return nil
}
