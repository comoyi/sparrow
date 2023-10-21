package server

import (
	"fmt"
	"github.com/comoyi/sparrow/config"
	"log/slog"
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
}
