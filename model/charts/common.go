package charts

import (
	"database/sql"
	"math/rand"
	"strconv"
	"strings"

	"data-view/schema"
)

const EmptyString = ""

type ChartDataHandler struct {
	RunGetDataFromDB  func(db *sql.DB, chartDataParams *schema.ChartDataParams) (string, error)
	RunGetDataFromCsv func(chartDataParams *schema.ChartDataParams) (string, error)
}

func (handler *ChartDataHandler) GetDataFromDB(db *sql.DB, chartDataParams *schema.ChartDataParams) (string, error) {
	return handler.RunGetDataFromDB(db, chartDataParams)
}

func (handler *ChartDataHandler) GetDataFromCsv(chartDataParams *schema.ChartDataParams) (string, error) {
	return handler.RunGetDataFromCsv(chartDataParams)
}

type ChartData interface {
	GetDataFromDB(db *sql.DB, chartDataParams *schema.ChartDataParams) (string, error)
	GetDataFromCsv(chartDataParams *schema.ChartDataParams) (string, error)
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

// CreateColor 生成随机颜色
func CreateColor() string {
	red := strings.ToUpper(strconv.FormatInt(int64(rand.Intn(256)), 16))
	if len(red) == 1 {
		red = "0" + red
	}
	blue := strings.ToUpper(strconv.FormatInt(int64(rand.Intn(256)), 16))
	if len(blue) == 1 {
		blue = "0" + blue
	}
	green := strings.ToUpper(strconv.FormatInt(int64(rand.Intn(256)), 16))
	if len(green) == 1 {
		green = "0" + green
	}
	return "#" + red + green + blue
}

// CountInArray 判断元素在数组中出现的次数
func CountInArray(a string, b []string) int {
	var result = 0
	for i := 0; i < len(b); i++ {
		if a == b[i] {
			result++
		}
	}
	return result
}

// GetIndexInArray 判断元素在数组中的位置
func GetIndexInArray(a string, b []string) int {
	var result = 0
	for i := 0; i < len(b); i++ {
		if a == b[i] {
			result = i
			break
		}
	}
	return result
}

// IsExitInArray 判断元素在数组中是否存在
func IsExitInArray(a string, b []string) bool {
	var result = false
	for i := 0; i < len(b); i++ {
		if a == b[i] {
			result = true
			break
		}
	}
	return result
}
