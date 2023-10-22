package server

import (
	"fmt"
	"github.com/comoyi/sparrow/config"
	"github.com/gin-gonic/gin"
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

	go startHttpServer()

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

func startHttpServer() {
	addr := fmt.Sprintf(":%d", config.Conf.Port)

	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	initRouter(engine)

	//err := engine.Run(addr)
	//if err != nil {
	//	slog.Error("http server start sailed")
	//}

	s := &http.Server{
		Addr:    addr,
		Handler: engine.Handler(),
	}
	err := s.ListenAndServe()
	if err != nil {
		slog.Error("http server start sailed")
		return
	}
}

func initRouter(engine *gin.Engine) {
	ga := engine.Group("/ga")
	ga.GET("/a", func(context *gin.Context) {
		context.JSON(200, "/ga/a")
	})
	ga.GET("/b", func(context *gin.Context) {
		context.JSON(200, "/ga/b")
	})

	gb := engine.Group("/gb")
	gb.GET("/a", func(context *gin.Context) {
		context.JSON(200, "/gb/a")
	})
}
