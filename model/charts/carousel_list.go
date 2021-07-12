package charts

import (
	"data-view/logger"
	"data-view/schema"

	"github.com/jmoiron/sqlx"
)

const (
	Header = "header"
	Body   = "body"
)

const CarouselList = "CarouselList"

// CarouselListGetDataHandle 轮播列表
var CarouselListGetDataHandle = &ChartDataHandler{RunGetDataFromDB: func(db *sqlx.DB, params schema.ChartDataParams) (map[string]interface{}, error) {
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

	var (
		header = make([]string, 0)
		body   = make([][]interface{}, 0)
	)

	header, _ = rows.Columns()
	// 获取数据
	for rows.Next() {
		// 数据对象
		item, err := rows.SliceScan()
		if err != nil {
			logger.Instance.Error(err)
			continue
		}
		for i := 0; i < len(item); i++ {
			item[i] = StrVal(item[i])
		}
		body = append(body, item)
	}
	return map[string]interface{}{Header: header, Body: body}, nil
}, RunGetDataFromCsv: func(params schema.ChartDataParams) (map[string]interface{}, error) {
	return nil, nil
}}
