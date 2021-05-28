package charts

import (
	"data-view/schema"

	"github.com/jmoiron/sqlx"
)

const Progress = "Progress"

var ProgressGetDataHandle = &ChartDataHandler{RunGetDataFromDB: func(db *sqlx.DB, params *schema.ChartDataParams) (map[string]interface{}, error) {
	return CounterGetDataHandle.GetDataFromDB(db, params)
}, RunGetDataFromCsv: func(params *schema.ChartDataParams) (map[string]interface{}, error) {
	return nil, nil
}}
