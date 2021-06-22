package model

import "github.com/google/wire"

var SetModel = wire.NewSet(
	DataSourceModelSet,
	ImageModelSet,
	DataViewModelSet,
)

const (
	DefaultOrder = "update_time desc"
	IsExist      = 1
)

const (
	Static    = "Static"
	DataBase  = "DataBase"
	File      = "File"
	Rest      = "Rest"
	MySQL     = "MySQL"
	Oracle    = "Oracle"
	SQLServer = "SQLServer"
	DB2       = "DB2"
)
