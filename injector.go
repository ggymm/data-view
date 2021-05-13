package main

import (
	"github.com/google/wire"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Router *Router
}
