package schema

type DataViewQueryParam struct {
	PagingParam
	BusinessId int64
}

type ChartDataParams struct {
	// 基础字段
	DataSourceType string `form:"dataSourceType"`
	ChartType      string `form:"chartType"`
	Database       int64  `form:"database"`
	FileName       string `form:"fileName"`
	Sql            string `form:"sql"`
	OrderBy        string `form:"orderBy"`

	// 图表扩展字段（基础字段）
	Dimension string `form:"dimension"`
	Name      string `form:"name"`
	Value     string `form:"value"`
}

type ViewChartItem struct {
	ItemChartData string `form:"chartData"`
	ItemChartType string `form:"chartType"`
	ItemChoose    string `form:"choose"`
	ItemData      string `form:"data"`
	ItemHeight    int64  `form:"height"`
	ItemI         string `form:"i"`
	ItemInterval  int64  `form:"interval"`
	ItemOption    string `form:"option"`
	ItemRefresh   string `form:"refresh"`
	ItemWidth     int64  `form:"width"`
	ItemX         int64  `form:"x"`
	ItemY         int64  `form:"y"`
	ItemVersion   int64  `form:"version"`
}

type CreateDataViewReq struct {
	InstanceBackgroundColor string          `form:"instance_background_color"`
	InstanceBackgroundImg   int64           `form:"instance_background_img"`
	InstanceWidth           int64           `form:"instance_width" validate:"required"`
	InstanceHeight          int64           `form:"instance_height" validate:"required"`
	InstanceTitle           string          `form:"instance_title"`
	InstanceViewThumbnail   string          `form:"instance_view_thumbnail" validate:"required"`
	InstanceVersion         int64           `form:"instance_version" validate:"required"`
	ChartItems              []ViewChartItem `form:"chart_items" validate:"required"`
}

type UpdateDataViewReq struct {
	InstanceId              int64           `form:"instance_id" validate:"required"`
	InstanceBackgroundColor string          `form:"instance_background_color"`
	InstanceBackgroundImg   int64           `form:"instance_background_img"`
	InstanceWidth           int64           `form:"instance_width" validate:"required"`
	InstanceHeight          int64           `form:"instance_height" validate:"required"`
	InstanceTitle           string          `form:"instance_title"`
	InstanceViewThumbnail   string          `form:"instance_view_thumbnail" validate:"required"`
	InstanceVersion         int64           `form:"instance_version" validate:"required"`
	ChartItems              []ViewChartItem `form:"chart_items" validate:"required"`
}
