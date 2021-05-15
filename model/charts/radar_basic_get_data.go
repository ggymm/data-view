package charts

import (
	"database/sql"
	"encoding/json"
	"strings"

	"data-view/logger"
	"data-view/schema"
)

var RadarBasicGetDataHandle = &ChartDataHandler{RunGetDataFromDB: func(db *sql.DB, chartDataParams *schema.ChartDataParams) (string, error) {
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
	dataResults, err := RadarBasicFormatRows(rows, chartDataParams)
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

func RadarBasicFormatRows(rows *sql.Rows, chartDataParams *schema.ChartDataParams) (*map[string]interface{}, error) {
	// 图表所需要的字段和数据库中字段的对应关系
	nameField := chartDataParams.Name
	valueField := chartDataParams.Value
	maxField := chartDataParams.Max
	legendField := chartDataParams.Legend

	// 返回值列表
	resultMap := make(map[string]interface{})
	tempResults := make([]map[string]string, 0)
	//列名
	columns, err := rows.Columns()
	if err != nil {
		return &resultMap, err
	}
	//读取到的数据库数据切片
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	legendList := make([]string, 0)
	nameList := make([]string, 0)
	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
			return &resultMap, err
		}
		var value string
		//获取数据初步规范
		tempResultMap := make(map[string]string)
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			if strings.EqualFold(columns[i], nameField) {
				nameList = append(nameList, value)
				tempResultMap["name"] = value
			}
			if strings.EqualFold(columns[i], maxField) {
				tempResultMap["max"] = value
			}
			if strings.EqualFold(columns[i], valueField) {
				tempResultMap["value"] = value
			}
			if strings.EqualFold(columns[i], legendField) {
				legendList = append(legendList, value)
				tempResultMap["legend"] = value
			}
		}
		tempResults = append(tempResults, tempResultMap)
	}
	//规范数据
	indicatorList := make([]map[string]string, 0)
	dataList := make([]map[string]interface{}, 0)
	var ret = Duplicate(legendList)
	for _, legend := range ret {
		dataMap := make(map[string]interface{})
		valueList := make([]interface{}, 0)
		for _, tempResult := range tempResults {
			if strings.EqualFold(legend, tempResult["legend"]) {
				valueList = append(valueList, tempResult["value"])
			}
		}
		dataMap["name"] = legend
		dataMap["value"] = valueList
		dataList = append(dataList, dataMap)
	}
	for _, name := range Duplicate(nameList) {
		indicatorMap := make(map[string]string)
		for _, tempResult := range tempResults {
			if strings.EqualFold(name, tempResult["name"]) {
				max, ok := indicatorMap["max"]
				if ok { //max取最大值
					if max > tempResult["max"] {
						indicatorMap["max"] = max
					}
				} else {
					indicatorMap["max"] = tempResult["max"]
				}
			}
		}
		indicatorMap["name"] = name
		indicatorList = append(indicatorList, indicatorMap)
	}
	resultMap["legend"] = Duplicate(legendList)
	resultMap["indicator"] = indicatorList
	resultMap["data"] = dataList
	return &resultMap, nil
}
