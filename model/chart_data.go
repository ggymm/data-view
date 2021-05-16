package model

import (
	"data-view/model/charts"
	"data-view/schema"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

var ChartDataHandlers = make(map[string]charts.ChartData)

func init() {
	//ChartDataHandlers["plotBubble"] = charts.PlotBubbleGetDataHandle
	//ChartDataHandlers["plotMap"] = charts.PlotMapGetDataHandle
	ChartDataHandlers["lineNormal"] = charts.LineNormalGetDataHandle
	//ChartDataHandlers["lineStacking"] = charts.LineStackingGetDataHandle
	//ChartDataHandlers["lineStackingArea"] = charts.LineStackingAreaGetDataHandle
	//ChartDataHandlers["histogramGradient"] = charts.HistogramGradientGetDataHandle
	//ChartDataHandlers["histogramGradientHorizontal"] = charts.HistogramGradientGetDataHandle
	//ChartDataHandlers["histogramStacking"] = charts.HistogramStackingGetDataHandle
	//ChartDataHandlers["histogramComplex"] = charts.HistogramComplexGetDataHandle
	//ChartDataHandlers["histogramTwoBar"] = charts.HistogramTwoBarGetDataHandle
	//ChartDataHandlers["mapChina"] = charts.MapChinaGetDataHandle
	//ChartDataHandlers["mapProvince"] = charts.MapProvinceGetDataHandle
	//ChartDataHandlers["pieNormal"] = charts.PieNormalGetDataHandle
	//ChartDataHandlers["pieRing"] = charts.PieNormalGetDataHandle
	//ChartDataHandlers["pieRings"] = charts.PieRingsGetDataHandle
	//ChartDataHandlers["pie2D"] = charts.PieNormalGetDataHandle
	//ChartDataHandlers["piePercent"] = charts.PiePercentGetDataHandle
	//ChartDataHandlers["radarBasic"] = charts.RadarBasicGetDataHandle
	//ChartDataHandlers["heatBasic"] = charts.HeatBasicGetDataHandle
	//ChartDataHandlers["relationOne"] = charts.RelationOneGetDataHandle
	//ChartDataHandlers["relationTwo"] = charts.RelationTwoGetDataHandle
	//ChartDataHandlers["relationThree"] = charts.RelationThreeGetDataHandle
	//ChartDataHandlers["relationFour"] = charts.RelationFourGetDataHandle
	//ChartDataHandlers["relationFive"] = charts.RelationFiveGetDataHandle
	//ChartDataHandlers["wordCloud"] = charts.PieNormalGetDataHandle
	//ChartDataHandlers["rotationList"] = charts.RotationListGetDataHandle
	//ChartDataHandlers["counter"] = charts.CounterGetDataHandle
	//ChartDataHandlers["gauge"] = charts.GaugeGetDataHandle
}

func (m *DataViewModel) GetChartData(params *schema.ChartDataParams, dataSource *DataSource) (string, error) {
	var (
		err        error
		db         *sql.DB
		dataResult string
	)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	// 此处判断图表类型
	chartType := params.ChartType
	// 此处判断图表数据源类型
	dataSourceType := params.DataSourceType
	if strings.EqualFold(dataSourceType, DataBase) {
		// 构建数据库访问对象
		dataSourceType := dataSource.DataSourceType
		if strings.EqualFold(dataSourceType, MySQL) {
			urlTemplate := "%s:%s@tcp(%s:%d)/%s?charset=utf8&multiStatements=true"
			url := fmt.Sprintf(urlTemplate,
				dataSource.DataSourceUsername,
				dataSource.DataSourcePassword,
				dataSource.DataSourceIp,
				dataSource.DataSourcePort,
				dataSource.DataSourceDatabaseName)
			if db, err = sql.Open("mysql", url); err != nil {
				return EmptyString, err
			}
		} else if strings.EqualFold(dataSourceType, Oracle) {
		} else if strings.EqualFold(dataSourceType, SQLServer) {
		} else if strings.EqualFold(dataSourceType, DB2) {
		} else {
			return EmptyString, errors.New("数据源类型错误")
		}
		if dataResult, err = ChartDataHandlers[chartType].GetDataFromDB(db, params); err != nil {
			return dataResult, err
		}
		return dataResult, nil
	} else if strings.EqualFold(dataSourceType, CSV) {
		if dataResult, err = ChartDataHandlers[chartType].GetDataFromCsv(params); err != nil {
			return dataResult, err
		}
		return dataResult, nil
	} else {
		return EmptyString, errors.New("数据源类型错误")
	}
}
