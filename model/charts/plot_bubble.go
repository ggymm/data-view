package charts

import (
	"database/sql"
	"encoding/json"
	"strings"

	"data-view/logger"
	"data-view/schema"
)

var PlotBubbleGetDataHandle = &ChartDataHandler{RunGetDataFromDB: func(db *sql.DB, chartDataParams *schema.ChartDataParams) (string, error) {
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
	dataResults, err := PlotBubbleFormatRows(rows, chartDataParams)
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

func PlotBubbleFormatRows(rows *sql.Rows, chartDataParams *schema.ChartDataParams) (*map[string]interface{}, error) {
	// 普通饼图==环形饼图==2D饼图==词云
	// 图表所需要的字段和数据库中字段的对应关系
	xField := chartDataParams.X
	yField := chartDataParams.Y
	valueField := chartDataParams.Value
	legendField := chartDataParams.Legend
	// 返回值列表
	resultMap := make(map[string]interface{})
	tempResults := make([]map[string]string, 0)
	// 列名
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
			if strings.EqualFold(columns[i], xField) {
				tempResultMap["x"] = value
			}
			if strings.EqualFold(columns[i], yField) {
				tempResultMap["y"] = value
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
	list := make([]interface{}, 0)
	for _, legend := range Duplicate(legendList) {
		yMap := make(map[string]interface{})
		valueResultList := make([]interface{}, 0)
		for _, tempResult := range tempResults {
			if strings.EqualFold(legend, tempResult["legend"]) {
				valueList := make([]interface{}, 0)
				valueList = append(valueList, tempResult["x"])
				valueList = append(valueList, tempResult["y"])
				valueList = append(valueList, tempResult["value"])
				valueList = append(valueList, tempResult["legend"])
				valueResultList = append(valueResultList, valueList)
			}
		}
		yMap["name"] = legend
		yMap["value"] = valueResultList
		list = append(list, yMap)
	}
	resultMap["legend"] = Duplicate(legendList)
	resultMap["data"] = list
	return &resultMap, nil
}
