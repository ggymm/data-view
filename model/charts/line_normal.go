package charts

import (
	"data-view/schema"

	"github.com/jmoiron/sqlx"
)

const LineNormal = "LineNormal"

// LineNormalGetDataHandle 折线图集合（参照Echarts）
// 基础折线图, 基础平滑折线图, 折线图堆叠
var LineNormalGetDataHandle = &ChartDataHandler{RunGetDataFromDB: func(db *sqlx.DB, params schema.ChartDataParams) (map[string]interface{}, error) {
	return DatasetGetDataHandle.GetDataFromDB(db, params)
}, RunGetDataFromCsv: func(params schema.ChartDataParams) (map[string]interface{}, error) {
	return nil, nil
}, RunGetDataFromRest: func(params schema.ChartDataParams) (map[string]interface{}, error) {
	return nil, nil
}}
