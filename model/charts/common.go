package charts

import (
	"data-view/schema"
	"encoding/json"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type ChartData interface {
	GetDataFromDB(db *sqlx.DB, chartDataParams schema.ChartDataParams) (map[string]interface{}, error)
	GetDataFromCsv(chartDataParams schema.ChartDataParams) (map[string]interface{}, error)
}

type ChartDataHandler struct {
	RunGetDataFromDB  func(db *sqlx.DB, chartDataParams schema.ChartDataParams) (map[string]interface{}, error)
	RunGetDataFromCsv func(chartDataParams schema.ChartDataParams) (map[string]interface{}, error)
}

func (handler *ChartDataHandler) GetDataFromDB(db *sqlx.DB, chartDataParams schema.ChartDataParams) (map[string]interface{}, error) {
	return handler.RunGetDataFromDB(db, chartDataParams)
}

func (handler *ChartDataHandler) GetDataFromCsv(chartDataParams schema.ChartDataParams) (map[string]interface{}, error) {
	return handler.RunGetDataFromCsv(chartDataParams)
}

// StrVal interface 转 string
func StrVal(value interface{}) string {
	var str string
	if value == nil {
		return str
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		str = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		str = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		str = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		str = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		str = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		str = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		str = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		str = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		str = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		str = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		str = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		str = strconv.FormatUint(it, 10)
	case string:
		str = value.(string)
	case []byte:
		str = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		str = string(newValue)
	}
	return str
}

// Duplicate 去除重复数据，可能需要排序
func Duplicate(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
