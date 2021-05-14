package model

import "time"

type ImageBg struct {
	ImageId   int64     `json:"image_id" xorm:"not null pk autoincr BIGINT(20)"`
	ImageName string    `json:"image_name" xorm:"VARCHAR(200)"`
	ImagePath string    `json:"image_path" xorm:"VARCHAR(500)"`
	ImageSize int64     `json:"image_size" xorm:"BIGINT(20)"`
	AddTime   time.Time `json:"add_time" xorm:"comment('添加时间') DATETIME"`
	AddUser   int64     `json:"add_user" xorm:"comment('添加者') BIGINT(20)"`
	EditTime  time.Time `json:"edit_time" xorm:"comment('编辑时间') DATETIME"`
	EditUser  int64     `json:"edit_user" xorm:"comment('编辑者') BIGINT(20)"`
	DelFlag   int       `json:"del_flag" xorm:"comment('是否删除（1：存在；0：删除）') INT(1)"`
}
