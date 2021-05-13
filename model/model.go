package model

import "github.com/google/wire"

const (
	DefaultOrder = "edit_time desc"
	IsExist      = 1
)

var SetModel = wire.NewSet(
	DataSourceModelSet,
	DataViewModelSet,
)
