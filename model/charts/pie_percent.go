package charts

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"data-view/logger"
	"data-view/schema"
)

var PiePercentGetDataHandle = &ChartDataHandler{RunGetDataFromDB: func(db *sql.DB, chartDataParams *schema.ChartDataParams) (string, error) {
	sqlString := chartDataParams.Sql
	rows, err := db.Query(sqlString)
	if err != nil {
		return EmptyString, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			logger.Instance.Error(err)
		}
	}(rows)
	dataResults, err := piePercentFormatRows(rows, chartDataParams)
	if err != nil {
		return EmptyString, err
	}
	resultString, err := json.Marshal(dataResults)
	if err != nil {
		return EmptyString, err
	}
	return string(resultString), nil
}, RunGetDataFromCsv: func(chartDataParams *schema.ChartDataParams) (string, error) {
	return "", nil
}}

func piePercentFormatRows(rows *sql.Rows, chartDataParams *schema.ChartDataParams) (*map[string]interface{}, error) {
	// 图表所需要的字段和数据库中字段的对应关系
	resultMap := make(map[string]interface{})
	valueField := chartDataParams.Value
	// 返回值列表
	dataResults := make([]map[string]interface{}, 0)
	columns, err := rows.Columns()
	if err != nil {
		return &resultMap, err
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
			return &resultMap, err
		}
		dataResult := make(map[string]interface{})
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			if strings.EqualFold(columns[i], valueField) {
				dataResult["value"] = value
			}
		}
		dataResults = append(dataResults, dataResult)
	}
	fmt.Println(dataResults)
	if len(dataResults) > 0 {
		resultMap = dataResults[0]
		return &resultMap, nil
	} else {
		resultMap["value"] = 0
		return &resultMap, nil
	}
}
