package charts

import (
	"database/sql"
	"encoding/json"
	"strings"

	"data-view/logger"
	"data-view/schema"
)

var RelationThreeGetDataHandle = &ChartDataHandler{RunGetDataFromDB: func(db *sql.DB, chartDataParams *schema.ChartDataParams) (string, error) {
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
	dataResults, err := RelationThreeFormatRows(rows, chartDataParams)
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

func RelationThreeFormatRows(rows *sql.Rows, chartDataParams *schema.ChartDataParams) (*map[string]interface{}, error) {
	// 图表所需要的字段和数据库中字段的对应关系
	nameField := chartDataParams.Name
	valueField := chartDataParams.Value
	legendField := chartDataParams.Legend
	xField := chartDataParams.X
	yField := chartDataParams.Y

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
	idList := make([]string, 0)
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
			if strings.EqualFold(columns[i], "id") {
				nameList = append(nameList, value)
				tempResultMap["id"] = value
			}
			if strings.EqualFold(columns[i], nameField) {
				nameList = append(nameList, value)
				tempResultMap["name"] = value
			}
			if strings.EqualFold(columns[i], valueField) {
				tempResultMap["value"] = value
			}
			if strings.EqualFold(columns[i], legendField) {
				tempResultMap["legend"] = value
			}
			if strings.EqualFold(columns[i], xField) {
				tempResultMap["x"] = value
			}
			if strings.EqualFold(columns[i], yField) {
				tempResultMap["y"] = value
			}
			if strings.EqualFold(columns[i], "sourceID") {
				if value != "NULL" {
					idList = append(idList, value)
				}
				tempResultMap["sourceID"] = value
			}
			if strings.EqualFold(columns[i], "targetID") {
				if value != "NULL" {
					idList = append(idList, value)
				}
				tempResultMap["targetID"] = value
			}
		}
		tempResults = append(tempResults, tempResultMap)
	}
	//规范数据
	//去重后的结果
	attributesMap := make(map[string]interface{})
	attributesMap["length"] = 0
	edgesList := make([]map[string]interface{}, 0)
	nodeList := make([]interface{}, 0)
	normalMap := make(map[string]string)
	normalMap["show"] = "true"
	nodeIdList := make([]string, 0)
	for _, tempResult := range tempResults {
		nodeMap := make(map[string]interface{})
		edgeMap := make(map[string]interface{})
		//节点连接关系处理
		if tempResult["sourceID"] != "NULL" && tempResult["targetID"] != "NULL" {
			edgeMap["size"] = 1
			edgeMap["sourceID"] = tempResult["sourceID"]
			edgeMap["targetID"] = tempResult["targetID"]
			edgeMap["attributes"] = attributesMap
			edgesList = append(edgesList, edgeMap)
		}
		//节点处理
		if !IsExitInArray(tempResult["id"], nodeIdList) {
			nodeIdList = append(nodeIdList, tempResult["id"])
			nodeMap["size"] = CountInArray(tempResult["id"], idList) * 30
			nodeMap["color"] = CreateColor()
			nodeMap["label"] = tempResult["name"]
			nodeMap["x"] = tempResult["x"]
			nodeMap["y"] = tempResult["y"]
			nodeMap["attributes"] = attributesMap
			nodeMap["id"] = tempResult["id"]
			nodeList = append(nodeList, nodeMap)
		}
	}
	//拼接最后结果
	resultMap["nodes"] = nodeList
	resultMap["edges"] = edgesList
	return &resultMap, nil
}
