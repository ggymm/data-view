package charts

import (
	"data-view/logger"
	"data-view/schema"
	"errors"
	"github.com/jmoiron/sqlx"
)

var CounterGetDataHandle = &ChartDataHandler{RunGetDataFromDB: func(db *sqlx.DB, params *schema.ChartDataParams) (map[string]interface{}, error) {
	if len(params.Value) == 0 {
		return nil, errors.New("参数不完整")
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
	values := make([]string, 0)
	for rows.Next() {
		// 临时数据对象
		temp := make(map[string]interface{})
		err := rows.MapScan(temp)
		if err != nil {
			logger.Instance.Error(err)
			continue
		}
		// 封装数据对象
		for key, value := range temp {
			if key == params.Value {
				values = append(values, StrVal(value))
			}
		}
	}

	if len(values) > 0 {
		return map[string]interface{}{
			"value": values[0],
		}, nil
	} else {
		return map[string]interface{}{
			"value": 0,
		}, nil
	}
}, RunGetDataFromCsv: func(params *schema.ChartDataParams) (map[string]interface{}, error) {
	return nil, nil
}}
