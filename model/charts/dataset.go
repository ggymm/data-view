package charts

import (
	"data-view/logger"
	"data-view/schema"
	"errors"
	"github.com/jmoiron/sqlx"
)

var DatasetGetDataHandle = &ChartDataHandler{RunGetDataFromDB: func(db *sqlx.DB, params *schema.ChartDataParams) (map[string]interface{}, error) {
	if len(params.Name) == 0 || len(params.Value) == 0 {
		return nil, errors.New("参数不完整")
	}

	var (
		items   = make([]map[string]string, 0)
		results = make([]interface{}, 0)
		legends []string
		values  = make([][]string, 0)
	)
	if len(params.Legend) != 0 {
		legends = make([]string, 0)
	}

	rows, err := db.Queryx(params.Sql)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			logger.Instance.Error(err)
		}
	}()

	// 获取数据
	for rows.Next() {
		// 临时数据对象
		temp := make(map[string]interface{})
		err := rows.MapScan(temp)
		if err != nil {
			logger.Instance.Error(err)
			continue
		}

		// 封装数据对象
		item := make(map[string]string)
		for key, value := range temp {
			if key == params.Name {
				item[Name] = StrVal(value)
			}
			if key == params.Value {
				item[Value] = StrVal(value)
			}
			if len(params.Legend) != 0 {
				item[Legend] = StrVal(value)
				legends = append(legends, StrVal(value))
			}
		}
		items = append(items, item)
	}

	if len(params.Legend) != 0 {
		legends = Duplicate(legends)
		for _, legend := range legends {
		LegendItems:
			for _, item := range items {
				if item[Legend] == legend {
					value := make([]string, 0)
					if len(values) == 0 {
						value = append(value, item[Name], item[Value])
					} else {
						for i := 0; i < len(values); i++ {
							if item[Name] == values[i][0] {
								values[i] = append(values[i], item[Value])
								continue LegendItems
							}
						}
						value = append(value, item[Name], item[Value])
					}
					values = append(values, value)
				}
			}
		}
	} else {
	Items:
		for _, item := range items {
			value := make([]string, 0)
			if len(values) == 0 {
				value = append(value, item[Name], item[Value])
			} else {
				for i := 0; i < len(values); i++ {
					if item[Name] == values[i][0] {
						values[i] = append(values[i], item[Value])
						continue Items
					}
				}
				value = append(value, item[Name], item[Value])
			}
			values = append(values, value)
		}
	}

	legends = append([]string{Legend}, legends...)
	results = append(results, legends, values)
	return map[string]interface{}{
		"source": results,
	}, nil
}, RunGetDataFromCsv: func(chartDataParams *schema.ChartDataParams) (map[string]interface{}, error) {
	return nil, nil
}}
