package schema

type DataViewQueryParam struct {
	PagingParam
	BusinessId int64
}

type ChartDataParams struct {
	// 基础字段
	DataSourceType string `json:"dataSourceType"`
	ChartType      string `json:"chartType"`
	Database       int64  `json:"database"`
	FileName       string `json:"fileName"`
	Sql            string `json:"sql"`
	OrderBy        string `json:"orderBy"`

	// 图表扩展字段
	Legend   string `json:"legend"`
	Name     string `json:"name"`
	Value    string `json:"value"`
	Max      string `json:"max"`
	Stack    string `json:"stack"`
	Data     string `json:"data"`
	Province string `json:"province"`
}
