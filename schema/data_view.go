package schema

type DataViewQueryParam struct {
	PagingParam
	BusinessId int64
}

type ChartDataParams struct {
	DataSourceType string `json:"dataSourceType"`
	ChartType      string `json:"chartType"`
	Database       int64  `json:"database"`
	FileName       string `json:"fileName"`
	Sql            string `json:"sql"`
	Legend         string `json:"legend"`
	X              string `json:"x"`
	Y              string `json:"y"`
	Name           string `json:"name"`
	Value          string `json:"value"`
	Max            string `json:"max"`
	Stack          string `json:"stack"`
	Data           string `json:"data"`
	Province       string `json:"province"`
}
