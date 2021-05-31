package charts

import (
	"errors"

	"data-view/logger"
	"data-view/schema"

	"github.com/jmoiron/sqlx"
)

const RadarNormal = "RadarNormal"

var RadarNormalGetDataHandle = &ChartDataHandler{RunGetDataFromDB: func(db *sqlx.DB, params *schema.ChartDataParams) (map[string]interface{}, error) {
	if len(params.Name) == 0 ||
		len(params.Value) == 0 ||
		len(params.Dimension) == 0 ||
		len(params.Max) == 0 {
		return nil, errors.New("参数不完整")
	}

	var (
		indicator = make([]map[string]string, 0)
		data      = make([]map[string][]string, 0)
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

		var (
			name      string
			value     string
			max       string
			dimension string
		)
		// 封装数据对象
		for k, v := range temp {
			if k == params.Name {
				name = StrVal(v)
			}
			if k == params.Value {
				value = StrVal(v)
			}
			if k == params.Max {
				max = StrVal(v)
			}
			if k == params.Dimension {
				dimension = StrVal(v)
			}
		}

		exist := false
		for i := 0; i < len(data); i++ {
			for k := range data[i] {
				if k == dimension {
					exist = true
					data[i][k] = append(data[i][k], value)
				}
			}
		}
		if !exist {
			data = append(data, map[string][]string{
				dimension: {value},
			})
		}

		exist = false
		for i := 0; i < len(indicator); i++ {
			if indicator[i]["name"] == name {
				exist = true
				continue
			}
		}

		if !exist {
			indicator = append(indicator, map[string]string{
				"name": name,
				"max":  max,
			})
		}
	}

	fData := make([]map[string]interface{}, 0)
	for i := 0; i < len(data); i++ {
		for k, v := range data[i] {
			fData = append(fData, map[string]interface{}{
				"name":  k,
				"value": v,
			})
		}
	}

	return map[string]interface{}{
		"data":      fData,
		"indicator": indicator,
	}, nil
}, RunGetDataFromCsv: func(params *schema.ChartDataParams) (map[string]interface{}, error) {
	return nil, nil
}}
