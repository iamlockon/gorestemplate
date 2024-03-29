package main

import (
	"fmt"
	"log"
	"syscall"

	"github.com/fvbock/endless"

	"github.com/iamlockon/gorestemplate/models"
	"github.com/iamlockon/gorestemplate/pkg/logging"
	"github.com/iamlockon/gorestemplate/pkg/setting"
	"github.com/iamlockon/gorestemplate/routers"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
