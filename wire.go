// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"data-view/database"
	"data-view/handler"
	"data-view/model"

	"github.com/google/wire"
)

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		database.InitXormDB,
		model.SetModel,
		handler.SetHandler,
		RouterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
