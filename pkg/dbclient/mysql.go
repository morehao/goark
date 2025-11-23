package dbclient

import (
	"fmt"

	"github.com/morehao/golib/database/dbmysql"
	"github.com/morehao/golib/glog"
	"gorm.io/gorm"
)

var (
	DBDemo *gorm.DB
	DBIam  *gorm.DB
)

const (
	DBNameDemo = "demo"
	DBNameIam  = "ark_iam"
)

func InitMultiMysql(configs []dbmysql.MysqlConfig, logConfig *glog.LogConfig) error {
	if len(configs) == 0 {
		return fmt.Errorf("mysql config is empty")
	}

	var opts []dbmysql.Option
	if logConfig != nil {
		opts = append(opts, dbmysql.WithLogConfig(logConfig))
	}
	for _, cfg := range configs {
		client, err := dbmysql.InitMysql(&cfg, opts...)
		if err != nil {
			return fmt.Errorf("init mysql failed: " + err.Error())
		}
		switch cfg.Database {
		case DBNameDemo:
			DBDemo = client
		case DBNameIam:
			DBIam = client
		default:
			return fmt.Errorf("unknown database: " + cfg.Database)
		}
	}
	return nil
}
