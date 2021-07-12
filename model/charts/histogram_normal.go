package charts

import (
	"data-view/schema"

	"github.com/jmoiron/sqlx"
)

const HistogramNormal = "HistogramNormal"

// HistogramNormalGetDataHandle 柱状图集合（参照Echarts）
// 基础柱状图
var HistogramNormalGetDataHandle = &ChartDataHandler{RunGetDataFromDB: func(db *sqlx.DB, params schema.ChartDataParams) (map[string]interface{}, error) {
	return DatasetGetDataHandle.GetDataFromDB(db, params)
}, RunGetDataFromCsv: func(params schema.ChartDataParams) (map[string]interface{}, error) {
	return nil, nil
}}
