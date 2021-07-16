package model

import (
	"data-view/chart"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"data-view/model/charts"
	"data-view/schema"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var ChartDataHandlers = make(map[string]charts.ChartData)

func init() {
	ChartDataHandlers[charts.ScatterNormal] = charts.ScatterNormalGetDataHandle

	ChartDataHandlers[charts.LineNormal] = charts.LineNormalGetDataHandle
	ChartDataHandlers[charts.LineArea] = charts.LineAreaGetDataHandle

	ChartDataHandlers[charts.HistogramNormal] = charts.HistogramNormalGetDataHandle
	ChartDataHandlers[charts.PictorialBar] = charts.PictorialBarGetDataHandle

	ChartDataHandlers[charts.MapChina] = charts.MapChinaGetDataHandle

	ChartDataHandlers[charts.PieNormal] = charts.PieNormalGetDataHandle

	ChartDataHandlers[charts.RadarNormal] = charts.RadarNormalGetDataHandle

	ChartDataHandlers[charts.CarouselList] = charts.CarouselListGetDataHandle
	ChartDataHandlers[charts.Counter] = charts.CounterGetDataHandle
	ChartDataHandlers[charts.Progress] = charts.ProgressGetDataHandle
}

func (m *DataViewModel) GetChartData(params schema.ChartDataParams) (map[string]interface{}, error) {
	var db *sqlx.DB
	defer func() {
		if db != nil {
			_ = db.Close()
		}
	}()

	// 此处判断图表类型
	chartType := params.ChartType
	// 此处判断图表数据源类型
	dataSourceType := params.DataSourceType
	if strings.EqualFold(dataSourceType, DataBase) {
		dataSource, err := m.DataSourceModel.Get(params.Database)
		if err != nil {
			return nil, err
		}
		// 构建数据库访问对象
		dataSourceType := dataSource.DataSourceType
		if strings.EqualFold(dataSourceType, MySQL) {
			url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&multiStatements=true",
				dataSource.DataSourceUsername,
				dataSource.DataSourcePassword,
				dataSource.DataSourceIp,
				dataSource.DataSourcePort,
				dataSource.DataSourceDatabaseName)
			if db, err = sqlx.Open("mysql", url); err != nil {
				return nil, err
			}
		} else if strings.EqualFold(dataSourceType, Oracle) {
		} else if strings.EqualFold(dataSourceType, SQLServer) {
		} else if strings.EqualFold(dataSourceType, DB2) {
		} else {
			return nil, errors.New("数据源类型错误")
		}
		chartData := ChartDataHandlers[chartType]
		if chartData == nil {
			return nil, errors.New("图表类型错误")
		}
		return chartData.GetDataFromDB(db, params)
	} else if strings.EqualFold(dataSourceType, File) {
		return ChartDataHandlers[chartType].GetDataFromCsv(params)
	} else if strings.EqualFold(dataSourceType, Rest) {
		jsonData := chart.JsonDataMap[params.ChartType]
		dataset := make([]interface{}, 0)
		if err := json.Unmarshal([]byte(jsonData), &dataset); err != nil {
			return nil, err
		}
		return map[string]interface{}{"source": dataset}, nil
	} else {
		return nil, errors.New("数据源类型错误")
	}
}
