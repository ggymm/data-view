package model

import (
	"data-view/schema"
	"time"

	"github.com/google/wire"
	"xorm.io/xorm"
)

var ImageModelSet = wire.NewSet(wire.Struct(new(ImageModel), "*"))

type ImageModel struct {
	Engine *xorm.Engine
}

type Image struct {
	ImageId    int64     `json:"image_id" xorm:"not null pk autoincr BIGINT(20)"`
	ImageName  string    `json:"image_name" xorm:"comment('图片名称') VARCHAR(200)"`
	ImagePath  string    `json:"image_path" xorm:"comment('图片位置') VARCHAR(500)"`
	ImageSize  int64     `json:"image_size" xorm:"comment('图片大小') BIGINT(20)"`
	BusinessId int64     `json:"business_id" xorm:"comment('企业ID') BIGINT(20)"`
	AddTime    time.Time `json:"add_time" xorm:"comment('添加时间') DATETIME"`
	AddUser    int64     `json:"add_user" xorm:"comment('添加者') BIGINT(20)"`
	EditTime   time.Time `json:"edit_time" xorm:"comment('编辑时间') DATETIME"`
	EditUser   int64     `json:"edit_user" xorm:"comment('编辑者') BIGINT(20)"`
	DelFlag    int       `json:"del_flag" xorm:"comment('是否删除（1：存在；0：删除）') INT(1)"`
}

func (m *ImageModel) GetPage(params schema.ImageQueryParam) ([]*Image, int64, error) {
	var (
		list        = make([]*Image, 0)
		count int64 = 0
		query       = make(map[string]interface{})
	)
	// 构造查询参数
	query["business_id"] = params.BusinessId
	query["del_flag"] = IsExist
	if err := m.Engine.Where(query).Limit(params.Size, params.Offset()).OrderBy(DefaultOrder).Find(&list); err != nil {
		return list, count, err
	}
	if count, err := m.Engine.Where(query).Count(new(Image)); err != nil {
		return list, count, err
	} else {
		return list, count, nil
	}
}

func (m *ImageModel) GetList(businessId int64) ([]*Image, error) {
	var (
		list  = make([]*Image, 0)
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
