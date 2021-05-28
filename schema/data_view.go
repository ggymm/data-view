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
	ItemChartData string `json:"chartData"`
	ItemChartType string `json:"chartType"`
	ItemLock      string `json:"lock"`
	ItemChoose    string `json:"choose"`
	ItemData      string `json:"data"`
	ItemHeight    int64  `json:"height"`
	ItemI         string `json:"i"`
	ItemInterval  int64  `json:"interval"`
	ItemOption    string `json:"option"`
	ItemRefresh   string `json:"refresh"`
	ItemWidth     int64  `json:"width"`
	ItemX         int64  `json:"x"`
	ItemY         int64  `json:"y"`
	ItemVersion   int64  `json:"version"`
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
