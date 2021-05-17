package test

import (
	"data-view/model/charts"
	"encoding/json"
	"os"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func StrVal(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

const (
	Name   = "name"
	Value  = "value"
	Legend = "legend"
)

func TestChartData(t *testing.T) {
	_ = os.Setenv("CONFIG_PATH", "../config.toml")

	// 测试数据源信息
	url := "root:root@tcp(127.0.0.1:3306)/data-view-template?charset=utf8&multiStatements=true"
	db, err := sqlx.Open("mysql", url)
	if err != nil {
		t.Error(err)
		return
	}

	// dataset数据类型
	rows, err := db.Queryx("select * from line_normal limit 1000 ")
	if err != nil {
		t.Error(err)
		return
	}

	n := "x"
	v := "y"
	l := "name"

	// 返回值和中间状态
	results := make([]interface{}, 0)
	legends := make([]string, 0)
	values := make([][]string, 0)

	// 数据对象列表
	items := make([]map[string]string, 0)
	for rows.Next() {
		// 临时数据对象
		temp := make(map[string]interface{})
		err := rows.MapScan(temp)
		if err != nil {
			t.Error(err)
			continue
		}

		// 封装数据对象
		item := make(map[string]string)
		for key, value := range temp {
			if key == n {
				item[Name] = StrVal(value)
			}
			if key == v {
				item[Value] = StrVal(value)
			}
			if key == l {
				item[Legend] = StrVal(value)
				legends = append(legends, StrVal(value))
			}
		}
		items = append(items, item)
	}

	legends = charts.Duplicate(legends)

	for _, legend := range legends {
	Items:
		for _, item := range items {
			if item[Legend] == legend {
				value := make([]string, 0)
				if len(values) == 0 {
					value = append(value, item[Name], item[Value])
				} else {
					for i := 0; i < len(values); i++ {
						if item[Name] == values[i][0] {
							values[i] = append(values[i], item[Value])
							continue Items
						}
					}
					value = append(value, item[Name], item[Value])
				}
				values = append(values, value)
			}
		}
	}

	legends = append([]string{Legend}, legends...)
	results = append(results, legends, values)

	t.Log(results)
}
