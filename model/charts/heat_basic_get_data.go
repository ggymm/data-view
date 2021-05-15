package charts

import (
	"database/sql"
	"encoding/json"
	"strings"

	"data-view/logger"
	"data-view/schema"
)

var HeatBasicGetDataHandle = &ChartDataHandler{RunGetDataFromDB: func(db *sql.DB, chartDataParams *schema.ChartDataParams) (string, error) {
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
	dataResults, err := HeatBasicFormatRows(rows, chartDataParams)
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

func HeatBasicFormatRows(rows *sql.Rows, chartDataParams *schema.ChartDataParams) (*[]interface{}, error) {
	// 图表所需要的字段和数据库中字段的对应关系
	xField := chartDataParams.X
	yField := chartDataParams.Y
	valueField := chartDataParams.Value
	// 返回值列表
	dataResults := make([]map[string]interface{}, 0)
	dataList := make([]interface{}, 0)
	resultList := make([]interface{}, 0)
	columns, err := rows.Columns()
	if err != nil {
		return &resultList, err
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
			return &resultList, err
		}
		dataResult := make(map[string]interface{})
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			if strings.EqualFold(columns[i], xField) {
				dataResult["x"] = value
			}
			if strings.EqualFold(columns[i], yField) {
				dataResult["y"] = value
			}
			if strings.EqualFold(columns[i], valueField) {
				dataResult["value"] = value
			}
		}
		dataResults = append(dataResults, dataResult)
	}
	for _, dataResult := range dataResults {
		dataMap := make(map[string]interface{})
		coordList := make([]interface{}, 0)
		dataMap["elevation"] = dataResult["value"]
		coordList = append(coordList, dataResult["x"], dataResult["y"])
		dataMap["coord"] = coordList
		dataList = append(dataList, dataMap)
	}
	resultList = append(resultList, dataList)
	return &resultList, nil
}
