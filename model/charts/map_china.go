package charts

import (
	"data-view/schema"

	"github.com/jmoiron/sqlx"
)

const MapChina = "MapChina"

// MapChinaGetDataHandle
// 中国地图
var MapChinaGetDataHandle = &ChartDataHandler{RunGetDataFromDB: func(db *sqlx.DB, params schema.ChartDataParams) (map[string]interface{}, error) {
	return DatasetGetDataHandle.GetDataFromDB(db, params)
}, RunGetDataFromCsv: func(params schema.ChartDataParams) (map[string]interface{}, error) {
	return nil, nil
}}
