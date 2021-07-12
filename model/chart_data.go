package model

import (
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

func (m *DataViewModel) GetChartData(params schema.ChartDataParams, dataSource *DataSource) (map[string]interface{}, error) {
	var (
		err        error
		db         *sqlx.DB
		dataResult map[string]interface{}
	)
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
		if dataResult, err = chartData.GetDataFromDB(db, params); err != nil {
			return dataResult, err
		}
		return dataResult, nil
	} else if strings.EqualFold(dataSourceType, File) {
		if dataResult, err = ChartDataHandlers[chartType].GetDataFromCsv(params); err != nil {
			return dataResult, err
		}
		return dataResult, nil
	} else if strings.EqualFold(dataSourceType, Rest) {
		if dataResult, err = ChartDataHandlers[chartType].GetDataFromRest(params); err != nil {
			return dataResult, err
		}
		return dataResult, nil
	} else {
		return nil, errors.New("数据源类型错误")
	}
}
