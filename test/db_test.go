package test

import (
	"os"
	"testing"

	"data-view/config"
	"data-view/model"

	"github.com/satori/go.uuid"
	"xorm.io/core"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

func TestFixElId(t *testing.T) {
	_ = os.Setenv("CONFIG_PATH", "../config.toml")
	config.Init()
	connectUrl := config.Instance.Mysql.Username + ":" + config.Instance.Mysql.Password +
		"@tcp(" + config.Instance.Mysql.Address + ":" + config.Instance.Mysql.Port +
		")/" + config.Instance.Mysql.Name + "?charset=utf8&parseTime=True&loc=Local"
	if engine, err := xorm.NewEngine("mysql", connectUrl); err != nil {
		t.Fatal(err)
	} else {
		engine.SetTableMapper(core.SnakeMapper{})
		engine.SetColumnMapper(core.GonicMapper{})
		engine.SetMaxIdleConns(20)
		engine.SetMaxOpenConns(100)
		engine.ShowSQL(true)
		engine.Logger().SetLevel(log.LOG_DEBUG)

		var list = make([]*model.ViewChartItem, 0)

		if err := engine.Find(&list); err != nil {
			t.Fatal(err)
		}

		for i := 0; i < len(list); i++ {
			item := list[i]
			u4 := uuid.NewV4()
			if _, err := engine.ID(item.ItemId).Table(new(model.ViewChartItem)).Update(map[string]interface{}{
				"item_el_id": u4.String(),
			}); err != nil {
				t.Fatal(err)
			}
		}

		t.Log("success")
	}
}
