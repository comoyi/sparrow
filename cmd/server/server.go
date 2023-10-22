package server

import (
	"fmt"
	"github.com/comoyi/sparrow/config"
	"log/slog"
	"net/http"
	_ "net/http/pprof"
)

func Start() {
	start()
}

func start() {
	err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	if config.Conf.Env == "dev" {
		slog.Info(fmt.Sprintf("%+v", config.Conf))
	}

	go startPprof()

	blockChan := make(chan struct{})
	<-blockChan
}

func startPprof() {
	addr := fmt.Sprintf("%s:%v", config.Conf.PprofHost, config.Conf.PprofPort)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		slog.Warn(fmt.Sprintf("err: %v", err))
	}
}
