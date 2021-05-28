package charts

import (
	"errors"

	"data-view/logger"
	"data-view/schema"

	"github.com/jmoiron/sqlx"
)

const (
	Source    = "source"
	Name      = "name"
	Value     = "value"
	Dimension = "dimension"
)

var DatasetGetDataHandle = &ChartDataHandler{RunGetDataFromDB: func(db *sqlx.DB, params *schema.ChartDataParams) (map[string]interface{}, error) {
	if len(params.Name) == 0 || len(params.Value) == 0 {
		return nil, errors.New("参数不完整")
	}

	var (
		hasDimension = len(params.Dimension) != 0 && params.Dimension != "无"
		items        = make([]map[string]string, 0)
		dataset      = make([][]string, 0)
		dimensions   = make([]string, 0)
		source       = make([][]string, 0)
	)

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
			if hasDimension {
				if key == params.Dimension {
					item[Dimension] = StrVal(value)
					dimensions = append(dimensions, StrVal(value))
				}
			}
		}
		items = append(items, item)
	}

	// 构造dataset的内容
	if hasDimension {
		dimensions = Duplicate(dimensions)
		for _, dimension := range dimensions {
		DimensionItems:
			for _, item := range items {
				if item[Dimension] == dimension {
					value := make([]string, 0)
					if len(source) == 0 {
						value = append(value, item[Name], item[Value])
					} else {
						for i := 0; i < len(source); i++ {
							if item[Name] == source[i][0] {
								source[i] = append(source[i], item[Value])
								continue DimensionItems
							}
						}
						value = append(value, item[Name], item[Value])
					}
					source = append(source, value)
				}
			}
		}
		dimensions = append([]string{Dimension}, dimensions...)
	} else {
	Items:
		for _, item := range items {
			value := make([]string, 0)
			if len(source) == 0 {
				value = append(value, item[Name], item[Value])
			} else {
				for i := 0; i < len(source); i++ {
					if item[Name] == source[i][0] {
						source[i] = append(source[i], item[Value])
						continue Items
					}
				}
				value = append(value, item[Name], item[Value])
			}
			source = append(source, value)
		}
		dimensions = append([]string{Dimension})
	}

	// 构造dataset
	if params.ChartType == PieNormal ||
		params.ChartType == MapChina {
		// 饼图特殊处理
		dimensions = append(dimensions, "")
	}
	dataset = append(dataset, dimensions)
	dataset = append(dataset, source...)
	return map[string]interface{}{Source: dataset}, nil
}, RunGetDataFromCsv: func(chartDataParams *schema.ChartDataParams) (map[string]interface{}, error) {
	return nil, nil
}}
