package model

import (
	"errors"
	"strconv"
	"strings"
	"time"

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
	BusinessId              int64           `json:"business_id" xorm:"comment('企业ID') BIGINT(20)"`
	InstanceBackgroundColor string          `json:"instance_background_color" xorm:"comment('背景颜色') VARCHAR(20)"`
	InstanceBackgroundImg   string          `json:"instance_background_img" xorm:"comment('背景图') VARCHAR(500)"`
	InstanceWidth           int64           `json:"instance_width" xorm:"comment('宽') BIGINT(20)"`
	InstanceHeight          int64           `json:"instance_height" xorm:"comment('高') BIGINT(20)"`
	InstanceTitle           string          `json:"instance_title" xorm:"comment('标题') VARCHAR(50)"`
	InstanceViewThumbnail   string          `json:"instance_view_thumbnail" xorm:"comment('缩略图') VARCHAR(500)"`
	InstanceVersion         int64           `json:"instance_version" xorm:"comment('版本号') BIGINT(20)"`
	CreateId                int64           `json:"create_id" xorm:"comment('创建者ID') BIGINT(20)"`
	CreateTime              schema.JsonTime `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateId                int64           `json:"update_id" xorm:"comment('更新者ID') BIGINT(20)"`
	UpdateTime              schema.JsonTime `json:"update_time" xorm:"comment('更新时间') DATETIME"`
	DelFlag                 int             `json:"del_flag" xorm:"comment('删除标识（-1：不存在；1：存在）') TINYINT(4)"`
}

type ViewChartItem struct {
	ItemId        int64  `json:"item_id" xorm:"not null pk autoincr BIGINT(20)"`
	InstanceId    int64  `json:"instance_id" xorm:"BIGINT(20)"`
	ItemChartData string `json:"chartData" xorm:"TEXT"`
	ItemChartType string `json:"chartType" xorm:"VARCHAR(50)"`
	ItemRotate    int    `json:"rotate" xorm:"INT(11)"`
	ItemLock      string `json:"lock" xorm:"VARCHAR(10)"`
	ItemChoose    string `json:"choose" xorm:"VARCHAR(10)"`
	ItemData      string `json:"data" xorm:"TEXT"`
	ItemHeight    int64  `json:"height" xorm:"BIGINT(20)"`
	ItemI         string `json:"i" xorm:"VARCHAR(20)"`
	ItemInterval  int64  `json:"interval" xorm:"BIGINT(20)"`
	ItemOption    string `json:"option" xorm:"TEXT"`
	ItemRefresh   string `json:"refresh" xorm:"VARCHAR(10)"`
	ItemWidth     int64  `json:"width" xorm:"BIGINT(20)"`
	ItemX         int64  `json:"x" xorm:"BIGINT(20)"`
	ItemY         int64  `json:"y" xorm:"BIGINT(20)"`
	ItemVersion   int64  `json:"version" xorm:"BIGINT(20)"`
}

type ViewInstancePro struct {
	ViewInstance `xorm:"extends"`
	ChartItems   []*ViewChartItem `json:"chart_items" xorm:"-"`
	StartIndex   int64            `json:"start_index" xorm:"-"`
}

func (ViewInstancePro) TableName() string {
	return "view_instance"
}

func (vip *ViewInstancePro) setStartIndex() {
	chartItems := vip.ChartItems
	if len(chartItems) > 0 {
		itemI := chartItems[len(chartItems)-1].ItemI
		index := strings.Replace(itemI, "chart", "", -1)
		i, _ := strconv.ParseInt(index, 10, 64)
		vip.StartIndex = i + 1
	} else {
		vip.StartIndex = 0
	}
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
	columns := []string{"instance_id", "instance_width", "instance_height", "instance_title", "instance_view_thumbnail", "update_time"}
	if err := m.Engine.Cols(columns...).Where(query).Limit(params.Size, params.Offset()).OrderBy("create_time").Find(&list); err != nil {
		return list, count, err
	}
	if count, err := m.Engine.Where(query).Count(new(ViewInstance)); err != nil {
		return list, count, err
	} else {
		return list, count, nil
	}
}

func (m *DataViewModel) Get(id, businessId int64) (*ViewInstancePro, error) {
	// 获取数据可视化实例
	var v = new(ViewInstancePro)
	v.InstanceId = id
	v.BusinessId = businessId
	v.DelFlag = IsExist
	if has, err := m.Engine.Get(v); err == nil {
		if has {
			// 获取图表列表
			chartItems := make([]*ViewChartItem, 0)
			if err := m.Engine.Where(map[string]interface{}{
				"instance_id":  id,
				"item_version": v.InstanceVersion,
			}).OrderBy("REPLACE (item_i, 'chart', '') + 0").Find(&chartItems); err != nil {
				return v, err
			}
			v.ChartItems = chartItems
			v.setStartIndex()
			return v, err
		} else {
			return nil, errors.New("数据不存在")
		}
	} else {
		return nil, err
	}
}

func (m *DataViewModel) Create(vip *ViewInstancePro) (int64, error) {
	session := m.Engine.NewSession()
	defer func() {
		if session != nil {
			_ = session.Close()
		}
	}()
	if err := session.Begin(); err != nil {
		return 0, err
	}

	vip.InstanceVersion = 1
	vip.CreateTime = schema.JsonTime(time.Now())
	vip.UpdateTime = schema.JsonTime(time.Now())
	vip.DelFlag = IsExist

	// 保存数据可视化实例
	if _, err := session.Insert(vip); err != nil {
		_ = session.Rollback()
		return 0, err
	} else {
		for i := 0; i < len(vip.ChartItems); i++ {
			chartItem := vip.ChartItems[i]
			chartItem.InstanceId = vip.InstanceId
			chartItem.ItemVersion = vip.InstanceVersion
			if _, err := session.Insert(chartItem); err != nil {
				_ = session.Rollback()
				return 0, err
			}
		}
	}

	// 事务提交
	if err := session.Commit(); err != nil {
		return 0, err
	}
	return vip.InstanceId, nil
}

func (m *DataViewModel) Update(vip *ViewInstancePro) error {
	session := m.Engine.NewSession()
	defer func() {
		if session != nil {
			_ = session.Close()
		}
	}()
	if err := session.Begin(); err != nil {
		return err
	}

	vip.InstanceVersion = vip.InstanceVersion + 1
	vip.UpdateTime = schema.JsonTime(time.Now())
	vip.DelFlag = IsExist

	// 保存数据可视化实例
	if _, err := session.ID(vip.InstanceId).Update(vip); err != nil {
		_ = session.Rollback()
		return err
	} else {
		for i := 0; i < len(vip.ChartItems); i++ {
			chartItem := vip.ChartItems[i]
			chartItem.InstanceId = vip.InstanceId
			chartItem.ItemVersion = vip.InstanceVersion
			if _, err := session.Insert(chartItem); err != nil {
				_ = session.Rollback()
				return err
			}
		}
	}

	// 事务提交
	if err := session.Commit(); err != nil {
		return err
	}
	return nil
}
