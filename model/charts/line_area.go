package charts

import (
	"data-view/schema"

	"github.com/jmoiron/sqlx"
)

const LineArea = "LineArea"

// LineAreaGetDataHandle
// 基础面积图, 堆叠面积图
var LineAreaGetDataHandle = &ChartDataHandler{RunGetDataFromDB: func(db *sqlx.DB, params *schema.ChartDataParams) (map[string]interface{}, error) {
	return DatasetGetDataHandle.GetDataFromDB(db, params)
}, RunGetDataFromCsv: func(params *schema.ChartDataParams) (map[string]interface{}, error) {
	return nil, nil
}}
