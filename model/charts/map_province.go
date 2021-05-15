package charts

import (
	"database/sql"
	"encoding/json"
	"strings"

	"data-view/logger"
	"data-view/schema"
)

var MapProvinceGetDataHandle = &ChartDataHandler{RunGetDataFromDB: func(db *sql.DB, chartDataParams *schema.ChartDataParams) (string, error) {
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
	dataResults, err := MapProvincelFormatRows(rows, chartDataParams)
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

func MapProvincelFormatRows(rows *sql.Rows, chartDataParams *schema.ChartDataParams) (*map[string]interface{}, error) {
	// 图表所需要的字段和数据库中字段的对应关系
	nameField := chartDataParams.Name
	valueField := chartDataParams.Value
	legendField := chartDataParams.Legend
	provinceField := chartDataParams.Province
	// 返回值列表
	resultMap := make(map[string]interface{})
	valueResultList := make([]map[string]interface{}, 0)
	tempResults := make([]map[string]string, 0)
	// 列名
	columns, err := rows.Columns()
	if err != nil {
		return &resultMap, err
	}
	// 读取到的数据库数据切片
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
			return &resultMap, err
		}
		var value string
		// 获取数据初步规范
		tempResultMap := make(map[string]string)
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			if strings.EqualFold(columns[i], nameField) {
				tempResultMap["name"] = value
			}
			if strings.EqualFold(columns[i], valueField) {
				tempResultMap["value"] = value
			}
			if strings.EqualFold(columns[i], legendField) {
				tempResultMap["legend"] = value
			}
		}
		tempResults = append(tempResults, tempResultMap)
	}
	// 规范数据
	for _, tempResult := range tempResults {
		valueMap := make(map[string]interface{})
		valueMap["name"] = tempResult["name"]
		valueMap["value"] = tempResult["value"]
		valueResultList = append(valueResultList, valueMap)
	}
	resultMap["province"] = provinceField
	resultMap["data"] = valueResultList
	return &resultMap, nil
}
