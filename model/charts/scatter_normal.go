package charts

import (
	"data-view/schema"
	"github.com/jmoiron/sqlx"
)

const ScatterNormal = "ScatterNormal"

// ScatterNormalGetDataHandle 散点图
var ScatterNormalGetDataHandle = &ChartDataHandler{RunGetDataFromDB: func(db *sqlx.DB, params schema.ChartDataParams) (map[string]interface{}, error) {
	return DatasetGetDataHandle.GetDataFromDB(db, params)
}, RunGetDataFromCsv: func(params schema.ChartDataParams) (map[string]interface{}, error) {
	return nil, nil
}}
