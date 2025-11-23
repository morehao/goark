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
	gormLogConfig := config.Conf.Log["gorm"]
	if err := storages.InitMultiMysql(config.Conf.MysqlConfigs, &gormLogConfig); err != nil {
		return fmt.Errorf("init mysql failed: " + err.Error())
	}
	redisLogConfig := config.Conf.Log["redis"]
	if err := storages.InitMultiRedis(config.Conf.RedisConfigs, &redisLogConfig); err != nil {
		return fmt.Errorf("init redis failed: " + err.Error())
	}
	return nil
}
