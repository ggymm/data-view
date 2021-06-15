package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"data-view/config"
	"data-view/logger"
)

func main() {
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	config.Init()
	injector, cleanFunc, err := BuildInjector()
	if err != nil {
		panic(err)
	}
	router := injector.Router.NewRouter()
	go func() {
		_ = router.Run(config.Instance.App.Addr)
	}()

EXIT:
	for {
		sig := <-sc
		logger.Instance.Infof("接收到信号[%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	cleanFunc()
	logger.Instance.Infof("服务退出")
	time.Sleep(time.Second)
	os.Exit(state)
}
