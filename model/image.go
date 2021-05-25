package model

import (
	"data-view/schema"

	"github.com/google/wire"
	"xorm.io/xorm"
)

var ImageModelSet = wire.NewSet(wire.Struct(new(ImageModel), "*"))

type ImageModel struct {
	Engine *xorm.Engine
}

type Image struct {
	ImageId    int64           `json:"image_id" xorm:"not null pk autoincr BIGINT(20)"`
	ImageName  string          `json:"image_name" xorm:"comment('图片名称') VARCHAR(200)"`
	ImagePath  string          `json:"image_path" xorm:"comment('图片位置') VARCHAR(500)"`
	ImageSize  int64           `json:"image_size" xorm:"comment('图片大小') BIGINT(20)"`
	BusinessId int64           `json:"business_id" xorm:"comment('企业ID') BIGINT(20)"`
	CreateId   int64           `json:"create_id" xorm:"comment('创建者ID') BIGINT(20)"`
	CreateTime schema.JsonTime `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateId   int64           `json:"update_id" xorm:"comment('更新者ID') BIGINT(20)"`
	UpdateTime schema.JsonTime `json:"update_time" xorm:"comment('更新时间') DATETIME"`
	DelFlag    int             `json:"del_flag" xorm:"comment('删除标识（-1：不存在；1：存在）') TINYINT(4)"`
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
