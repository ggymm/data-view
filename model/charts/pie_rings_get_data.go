package charts

import (
	"database/sql"
	"encoding/json"
	"strconv"
	"strings"

	"data-view/logger"
	"data-view/schema"
)

var PieRingsGetDataHandle = &ChartDataHandler{RunGetDataFromDB: func(db *sql.DB, chartDataParams *schema.ChartDataParams) (string, error) {
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
	dataResults, err := PieRingsFormatRows(rows, chartDataParams)
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

func PieRingsFormatRows(rows *sql.Rows, chartDataParams *schema.ChartDataParams) (*[]interface{}, error) {
	// 图表所需要的字段和数据库中字段的对应关系
	nameField := chartDataParams.Name
	valueField := chartDataParams.Value
	// 返回值列表
	dataResults := make([]interface{}, 0)
	columns, err := rows.Columns()
	if err != nil {
		return &dataResults, err
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
			return &dataResults, err
		}
		dataResult := make(map[string]interface{})
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			if strings.EqualFold(columns[i], nameField) {
				dataResult["name"] = value
			}
			if strings.EqualFold(columns[i], valueField) {
				other, _ := strconv.Atoi(value)
				if other < 1 && other > 0 {
					dataResult["value"] = other * 100
					dataResult["other"] = 100 - other*100
				} else if other < 10 && other > 1 {
					dataResult["value"] = other * 10
					dataResult["other"] = 100 - other*10
				} else if other <= 100 && other > 10 {
					dataResult["value"] = other
					dataResult["other"] = 100 - other
				} else {
					dataResult["value"] = 100
					dataResult["other"] = 0
				}
			}
		}
		dataResults = append(dataResults, dataResult)
	}
	return &dataResults, nil
}
