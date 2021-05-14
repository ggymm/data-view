package model

import (
	"data-view/config"
	"data-view/schema"

	"github.com/google/wire"
	"xorm.io/xorm"
)

var DataViewModelSet = wire.NewSet(wire.Struct(new(DataViewModel), "*"))

type DataViewModel struct {
	Engine *xorm.Engine
}

type ViewInstance struct {
	InstanceId              int64           `json:"instance_id" xorm:"not null pk autoincr comment('实例ID') BIGINT(20)"`
	InstanceBackgroundColor string          `json:"instance_background_color" xorm:"comment('背景颜色') VARCHAR(20)"`
	InstanceBackgroundImg   string          `json:"instance_background_img" xorm:"comment('背景图') VARCHAR(50)"`
	InstanceWidth           int64           `json:"instance_width" xorm:"comment('宽') BIGINT(20)"`
	InstanceHeight          int64           `json:"instance_height" xorm:"comment('高') BIGINT(20)"`
	InstanceTitle           string          `json:"instance_title" xorm:"comment('标题') VARCHAR(50)"`
	InstanceViewThumbnail   string          `json:"instance_view_thumbnail" xorm:"comment('缩略图') VARCHAR(50)"`
	InstanceVersion         int64           `json:"instance_version" xorm:"comment('版本号') BIGINT(20)"`
	AddUser                 int64           `json:"add_user" xorm:"comment('添加用户') BIGINT(20)"`
	AddTime                 schema.JsonTime `json:"add_time" xorm:"comment('添加时间') DATETIME"`
	EditUser                int64           `json:"edit_user" xorm:"comment('更新用户') BIGINT(20)"`
	EditTime                schema.JsonTime `json:"edit_time" xorm:"comment('更新时间') DATETIME"`
	DelFlag                 int             `json:"del_flag" xorm:"comment('删除标识') INT(1)"`
}

type ChartItem struct {
	ItemId        int64  `json:"item_id" xorm:"not null pk autoincr BIGINT(20)"`
	InstanceId    int64  `json:"instance_id" xorm:"BIGINT(20)"`
	ItemChartData string `json:"item_chart_data" xorm:"TEXT"`
	ItemChartType string `json:"item_chart_type" xorm:"VARCHAR(50)"`
	ItemChoose    string `json:"item_choose" xorm:"VARCHAR(10)"`
	ItemData      string `json:"item_data" xorm:"TEXT"`
	ItemHeight    int64  `json:"item_height" xorm:"BIGINT(20)"`
	ItemI         string `json:"item_i" xorm:"VARCHAR(20)"`
	ItemInterval  int64  `json:"item_interval" xorm:"BIGINT(20)"`
	ItemOption    string `json:"item_option" xorm:"TEXT"`
	ItemRefresh   string `json:"item_refresh" xorm:"VARCHAR(10)"`
	ItemWidth     int64  `json:"item_width" xorm:"BIGINT(20)"`
	ItemX         int64  `json:"item_x" xorm:"BIGINT(20)"`
	ItemY         int64  `json:"item_y" xorm:"BIGINT(20)"`
	ItemVersion   int64  `json:"item_version" xorm:"BIGINT(20)"`
}

func (m *DataViewModel) GetPage(params schema.DataViewQueryParam) ([]*ViewInstance, int64, error) {
	var (
		list        = make([]*ViewInstance, 0)
		count int64 = 0
		query       = make(map[string]interface{})
	)
	// 构造查询参数
	query["business_id"] = params.BusinessId
	query["del_flag"] = IsExist
	// 获取列表
	columns := []string{"instance_id", "instance_width", "instance_height", "instance_title", "instance_view_thumbnail", "edit_time"}
	if err := m.Engine.Cols(columns...).Where(query).Limit(params.Size, params.Offset()).OrderBy(DefaultOrder).Find(&list); err != nil {
		return list, count, err
	}
	// 处理列表
	for i := 0; i < len(list); i++ {
		list[i].InstanceBackgroundImg = config.Instance.Storage.Url + list[i].InstanceBackgroundImg
		list[i].InstanceViewThumbnail = config.Instance.Storage.Url + list[i].InstanceViewThumbnail
	}
	if count, err := m.Engine.Where(query).Count(new(ViewInstance)); err != nil {
		return list, count, err
	} else {
		return list, count, nil
	}
}
