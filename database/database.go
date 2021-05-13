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
	connectUrl := config.Instance.Database.Username + ":" + config.Instance.Database.Password +
		"@tcp(" + config.Instance.Database.Address + ":" + config.Instance.Database.Port +
		")/" + config.Instance.Database.Name + "?charset=utf8&parseTime=True&loc=Local"
	if engine, err := xorm.NewEngine("mysql", connectUrl); err != nil {
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
