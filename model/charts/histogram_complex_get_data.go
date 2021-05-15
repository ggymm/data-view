package charts

import (
	"database/sql"
	"encoding/json"
	"strconv"
	"strings"

	"data-view/logger"
	"data-view/schema"
)

var HistogramComplexGetDataHandle = &ChartDataHandler{RunGetDataFromDB: func(db *sql.DB, chartDataParams *schema.ChartDataParams) (string, error) {
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
	dataResults, err := HistogramComplexFormatRows(rows, chartDataParams)
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

func HistogramComplexFormatRows(rows *sql.Rows, chartDataParams *schema.ChartDataParams) (*map[string]interface{}, error) {
	// 图表所需要的字段和数据库中字段的对应关系
	xField := chartDataParams.X
	yField := chartDataParams.Y
	legendField := chartDataParams.Legend

	// 返回值列表
	resultMap := make(map[string]interface{})
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
	legendList := make([]string, 0)
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
			if strings.EqualFold(columns[i], xField) {
				tempResultMap[xField] = value
			}
			if strings.EqualFold(columns[i], legendField) {
				tempResultMap[legendField] = value
				legendList = append(legendList, value)
			}
			if strings.EqualFold(columns[i], yField) {
				tempResultMap[yField] = value
			}
		}
		tempResults = append(tempResults, tempResultMap)
	}
	// 规范数据
	valueList := make([]interface{}, 0)
	for _, legend := range Duplicate(legendList) {
		valueMap := make(map[string]interface{})
		xResultList := make([]string, 0)
		yResultList := make([]int, 0)
		for _, tempResult := range tempResults {
			if strings.EqualFold(legend, tempResult[legendField]) {
				yInt, err := strconv.Atoi(tempResult[yField])
				if err == nil {
					yResultList = append(yResultList, yInt)
					xResultList = append(xResultList, tempResult[xField])
				}
			}
		}
		valueMap["name"] = legend
		valueMap["x"] = xResultList
		valueMap["y"] = yResultList
		valueList = append(valueList, valueMap)
	}
	resultMap["legend"] = Duplicate(legendList)
	resultMap["value"] = valueList
	return &resultMap, nil
}
