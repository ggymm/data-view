package database

import (
	"data-view/config"
	"data-view/logger"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/core"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

// InitXormDB 初始化数据源
func InitXormDB() (*xorm.Engine, func(), error) {
	db, cleanFunc, err := NewGormDB()
	if err != nil {
		return nil, cleanFunc, err
	}
	return db, cleanFunc, nil
}

func NewGormDB() (*xorm.Engine, func(), error) {
	var connectUrl string
	switch config.Instance.Database.Type {
	default:
	case "mysql":
		connectUrl = config.Instance.Mysql.Username + ":" + config.Instance.Mysql.Password +
			"@tcp(" + config.Instance.Mysql.Address + ":" + config.Instance.Mysql.Port +
			")/" + config.Instance.Mysql.Name + "?charset=utf8&parseTime=True&loc=Local"
	case "sqlite3":
		connectUrl = config.Instance.Sqlite.Path
	}
	if engine, err := xorm.NewEngine(config.Instance.Database.Type, connectUrl); err != nil {
		return nil, nil, err
	} else {
		engine.SetTableMapper(core.SnakeMapper{})
		engine.SetColumnMapper(core.GonicMapper{})
		engine.SetMaxIdleConns(20)
		engine.SetMaxOpenConns(100)
		engine.ShowSQL(true)
		engine.Logger().SetLevel(log.LOG_DEBUG)
		cleanFunc := func() {
			err := engine.Close()
			if err != nil {
				logger.Instance.Errorf("xorm db close error: %s", err.Error())
			}
		}
		return engine, cleanFunc, nil
	}
}
