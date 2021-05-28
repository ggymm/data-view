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
	DataBase  = "DataBase"
	CSV       = "Csv"
	Rest      = "Rest"
	MySQL     = "MySQL"
	Oracle    = "Oracle"
	SQLServer = "SQLServer"
	DB2       = "DB2"
)
