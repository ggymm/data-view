package model

import "github.com/google/wire"

const (
	EmptyString  = ""
	DefaultOrder = "edit_time desc"
	IsExist      = 1
)

var SetModel = wire.NewSet(
	DataSourceModelSet,
	ImageModelSet,
	DataViewModelSet,
)

const (
	DataBase  = "DataBase"
	CSV       = "CSV"
	MySQL     = "MySQL"
	Oracle    = "Oracle"
	SQLServer = "SQLServer"
	DB2       = "DB2"
)
