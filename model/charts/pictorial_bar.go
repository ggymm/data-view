package charts

import (
	"data-view/schema"

	"github.com/jmoiron/sqlx"
)

const PictorialBar = "PictorialBar"

// PictorialBarGetDataHandle
// 象形柱图
var PictorialBarGetDataHandle = &ChartDataHandler{RunGetDataFromDB: func(db *sqlx.DB, params *schema.ChartDataParams) (map[string]interface{}, error) {
	return DatasetGetDataHandle.GetDataFromDB(db, params)
}, RunGetDataFromCsv: func(params *schema.ChartDataParams) (map[string]interface{}, error) {
	return nil, nil
}}
