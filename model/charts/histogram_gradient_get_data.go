package charts

import (
	"database/sql"
	"encoding/json"
	"sort"
	"strconv"
	"strings"

	"data-view/logger"
	"data-view/schema"
)

var HistogramGradientGetDataHandle = &ChartDataHandler{RunGetDataFromDB: func(db *sql.DB, chartDataParams *schema.ChartDataParams) (string, error) {
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
	dataResults, err := HistogramGradientFormatRows(rows, chartDataParams)
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

func HistogramGradientFormatRows(rows *sql.Rows, chartDataParams *schema.ChartDataParams) (*map[string]interface{}, error) {
	// 图表所需要的字段和数据库中字段的对应关系
	xField := chartDataParams.X
	yField := chartDataParams.Y
	// 返回值列表
	resultMap := make(map[string]interface{})
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
	xList := make([]string, 0)
	yList := make([]int, 0)
	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
			return &resultMap, err
		}
		var value string
		//获取数据初步规范
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			if strings.EqualFold(columns[i], xField) {
				xList = append(xList, value)
			}
			if strings.EqualFold(columns[i], yField) {
				yInt, err := strconv.Atoi(value)
				if err == nil {
					yList = append(yList, yInt)
				}
			}
		}
	}
	resultMap["x"] = Duplicate(xList)
	resultMap["y"] = yList
	if len(yList) != 0 {
		sortList := make([]int, len(yList))
		copy(sortList, yList)
		sort.Ints(sortList)
		yMax := sortList[len(sortList)-2 : len(sortList)-1]
		resultMap["max"] = yMax
	} else {
		resultMap["max"] = nil
	}
	return &resultMap, nil
}
