package model

import (
	"errors"

	"data-view/schema"

	"github.com/google/wire"
	"xorm.io/xorm"
)

var DataSourceModelSet = wire.NewSet(wire.Struct(new(DataSourceModel), "*"))

type DataSourceModel struct {
	Engine *xorm.Engine
}

type DataSource struct {
	DataSourceId           int64           `json:"data_source_id" xorm:"not null pk autoincr BIGINT(20)"`
	BusinessId             int64           `json:"business_id" xorm:"comment('企业ID') BIGINT(20)"`
	DataSourceName         string          `json:"data_source_name" xorm:"comment('数据源名称') VARCHAR(50)"`
	DataSourceType         string          `json:"data_source_type" xorm:"comment('数据源类型') VARCHAR(50)"`
	DataSourceDatabaseName string          `json:"data_source_database_name" xorm:"comment('数据源的数据库名称') VARCHAR(100)"`
	DataSourceIp           string          `json:"data_source_ip" xorm:"comment('数据源的IP地址') VARCHAR(50)"`
	DataSourcePort         int             `json:"data_source_port" xorm:"comment('数据源的端口号') INT(5)"`
	DataSourceUsername     string          `json:"data_source_username" xorm:"comment('数据源的账户名称') VARCHAR(50)"`
	DataSourcePassword     string          `json:"data_source_password" xorm:"comment('数据源的账户密码') VARCHAR(50)"`
	AddTime                schema.JsonTime `json:"add_time" xorm:"comment('添加时间') DATETIME"`
	AddUser                int64           `json:"add_user" xorm:"comment('添加者') BIGINT(20)"`
	EditTime               schema.JsonTime `json:"edit_time" xorm:"comment('编辑时间') DATETIME"`
	EditUser               int64           `json:"edit_user" xorm:"comment('编辑者') BIGINT(20)"`
	DelFlag                int             `json:"del_flag" xorm:"comment('是否删除（1：存在；0：删除）') INT(1)"`
}

func (m *DataSourceModel) GetPage(params schema.DataSourceQueryParam) ([]*DataSource, int64, error) {
	var (
		list        = make([]*DataSource, 0)
		count int64 = 0
		query       = make(map[string]interface{})
	)
	// 构造查询参数
	query["business_id"] = params.BusinessId
	query["del_flag"] = IsExist
	// 获取列表
	columns := []string{"data_source_id", "data_source_name", "data_source_type", "data_source_database_name", "data_source_ip", "data_source_username", "edit_time"}
	if err := m.Engine.Cols(columns...).Where(query).Limit(params.Size, params.Offset()).OrderBy(DefaultOrder).Find(&list); err != nil {
		return list, count, err
	}
	if count, err := m.Engine.Where(query).Count(new(DataSource)); err != nil {
		return list, count, err
	} else {
		return list, count, nil
	}
}

func (m *DataSourceModel) GetList(businessId int64) ([]*DataSource, error) {
	var (
		list  = make([]*DataSource, 0)
		query = make(map[string]interface{})
	)
	query["business_id"] = businessId
	query["del_flag"] = IsExist
	// 获取查询列表
	if err := m.Engine.Where(query).OrderBy(DefaultOrder).Find(&list); err != nil {
		return list, err
	}
	return list, nil
}

func (m *DataSourceModel) Get(id int64) (*DataSource, error) {
	var d = new(DataSource)
	d.DataSourceId = id
	d.DelFlag = IsExist
	if has, err := m.Engine.Get(d); err == nil {
		if has {
			return d, err
		} else {
			return nil, errors.New("数据源不存在")
		}
	} else {
		return nil, err
	}
}
