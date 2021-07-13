package schema

type DataViewQueryParam struct {
	PagingParam
	BusinessId int64
}

type ChartDataParams struct {
	DataSourceType string `json:"dataSourceType"` // 数据源类型
	ChartType      string `json:"chartType"`      // 图表类型
	Database       int64  `json:"database"`       // 数据库
	Sql            string `json:"sql"`            // 查询SQL
	Dimension      string `json:"dimension"`      // 分类字段
	Name           string `json:"name"`           //
	Value          string `json:"value"`          //
	Max            string `json:"max"`            // 雷达图
	FileName       string `json:"fileName"`       // 文件名称
}

type ViewChartItem struct {
	ItemElId         string `json:"elId"`
	ItemIndex        int    `json:"index"`
	ItemChartName    string `json:"chartName"`
	ItemChartType    string `json:"chartType"`
	ItemChartVersion int64  `json:"chartVersion"`
	ItemChartData    string `json:"chartData"`
	ItemData         string `json:"data"`
	ItemRotate       int    `json:"rotate"`
	ItemShow         string `json:"show"`
	ItemLock         string `json:"lock"`
	ItemChoose       string `json:"choose"`
	ItemInterval     int64  `json:"interval"`
	ItemOption       string `json:"option"`
	ItemRefresh      string `json:"refresh"`
	ItemWidth        int64  `json:"width"`
	ItemHeight       int64  `json:"height"`
	ItemX            int64  `json:"x"`
	ItemY            int64  `json:"y"`
	ItemVersion      int64  `json:"version"`
}

type CreateDataViewReq struct {
	InstanceBackgroundColor string           `json:"instance_background_color"`
	InstanceBackgroundImg   string           `json:"instance_background_img"`
	InstanceWidth           int64            `json:"instance_width" validate:"required"`
	InstanceHeight          int64            `json:"instance_height" validate:"required"`
	InstanceTitle           string           `json:"instance_title"`
	InstanceViewThumbnail   string           `json:"instance_view_thumbnail" validate:"required"`
	InstanceVersion         int64            `json:"instance_version" validate:"required"`
	ChartItems              []*ViewChartItem `json:"chart_items" validate:"required"`
}

type UpdateDataViewReq struct {
	InstanceId              int64            `json:"instance_id" validate:"required"`
	InstanceBackgroundColor string           `json:"instance_background_color"`
	InstanceBackgroundImg   string           `json:"instance_background_img"`
	InstanceWidth           int64            `json:"instance_width" validate:"required"`
	InstanceHeight          int64            `json:"instance_height" validate:"required"`
	InstanceTitle           string           `json:"instance_title"`
	InstanceViewThumbnail   string           `json:"instance_view_thumbnail" validate:"required"`
	InstanceVersion         int64            `json:"instance_version" validate:"required"`
	ChartItems              []*ViewChartItem `json:"chart_items" validate:"required"`
}
