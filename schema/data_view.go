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

	// 图表扩展字段
	Legend   string `form:"legend"`
	Name     string `form:"name"`
	Value    string `form:"value"`
	Max      string `form:"max"`
	Stack    string `form:"stack"`
	Data     string `form:"data"`
	Province string `form:"province"`
}
