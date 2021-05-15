package model

import (
	"data-view/model/charts"
	"data-view/schema"
)

var ChartDataHandlers = make(map[string]charts.ChartData)

const (
	PlotBubble                  = "plotBubble"
	PlotMap                     = "plotMap"
	LineNormal                  = "lineNormal"
	LineStacking                = "lineStacking"
	LineStackingArea            = "lineStackingArea"
	HistogramGradient           = "histogramGradient"
	HistogramGradientHorizontal = "histogramGradientHorizontal"
	HistogramStacking           = "histogramStacking"
	HistogramComplex            = "histogramComplex"
	HistogramTwoBar             = "histogramTwoBar"
	MapChina                    = "mapChina"
	MapProvince                 = "mapProvince"
	PieNormal                   = "pieNormal"
	PieRing                     = "pieRing"
	PieRings                    = "pieRings"
	Pie2D                       = "pie2D"
	PiePercent                  = "piePercent"
	RadarBasic                  = "radarBasic"
	HeatBasic                   = "heatBasic"
	RelationOne                 = "relationOne"
	RelationTwo                 = "relationTwo"
	RelationThree               = "relationThree"
	RelationFour                = "relationFour"
	RelationFive                = "relationFive"
	WordCloud                   = "wordCloud"
	RotationList                = "rotationList"
	Counter                     = "counter"
	Gauge                       = "gauge"
)

func init() {
	ChartDataHandlers[PlotBubble] = charts.PlotBubbleGetDataHandle
	ChartDataHandlers[PlotMap] = charts.PlotMapGetDataHandle
	ChartDataHandlers[LineNormal] = charts.LineNormalGetDataHandle
	ChartDataHandlers[LineStacking] = charts.LineStackingGetDataHandle
	ChartDataHandlers[LineStackingArea] = charts.LineStackingAreaGetDataHandle
	ChartDataHandlers[HistogramGradient] = charts.HistogramGradientGetDataHandle
	ChartDataHandlers[HistogramGradientHorizontal] = charts.HistogramGradientGetDataHandle
	ChartDataHandlers[HistogramStacking] = charts.HistogramStackingGetDataHandle
	ChartDataHandlers[HistogramComplex] = charts.HistogramComplexGetDataHandle
	ChartDataHandlers[HistogramTwoBar] = charts.HistogramTwoBarGetDataHandle
	ChartDataHandlers[MapChina] = charts.MapChinaGetDataHandle
	ChartDataHandlers[MapProvince] = charts.MapProvinceGetDataHandle
	ChartDataHandlers[PieNormal] = charts.PieNormalGetDataHandle
	ChartDataHandlers[PieRing] = charts.PieNormalGetDataHandle
	ChartDataHandlers[PieRings] = charts.PieRingsGetDataHandle
	ChartDataHandlers[Pie2D] = charts.PieNormalGetDataHandle
	ChartDataHandlers[PiePercent] = charts.PiePercentGetDataHandle
	ChartDataHandlers[RadarBasic] = charts.RadarBasicGetDataHandle
	ChartDataHandlers[HeatBasic] = charts.HeatBasicGetDataHandle
	ChartDataHandlers[RelationOne] = charts.RelationOneGetDataHandle
	ChartDataHandlers[RelationTwo] = charts.RelationTwoGetDataHandle
	ChartDataHandlers[RelationThree] = charts.RelationThreeGetDataHandle
	ChartDataHandlers[RelationFour] = charts.RelationFourGetDataHandle
	ChartDataHandlers[RelationFive] = charts.RelationFiveGetDataHandle
	ChartDataHandlers[WordCloud] = charts.PieNormalGetDataHandle
	ChartDataHandlers[RotationList] = charts.RotationListGetDataHandle
	ChartDataHandlers[Counter] = charts.CounterGetDataHandle
	ChartDataHandlers[Gauge] = charts.GaugeGetDataHandle
}

func (m *DataViewModel) GetChartData(chartDataParams schema.ChartDataParams) (string, error) {
	return "", nil
}
