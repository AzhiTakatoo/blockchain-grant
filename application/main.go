package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/orangebottle/blockchain-grant/application/blockchain"
	"github.com/orangebottle/blockchain-grant/application/pkg/setting"
	"github.com/orangebottle/blockchain-grant/application/routers"
	"log"
	"net/http"
	"time"
)

func init() {
	setting.Setup()
}

func main() {
	timeLocal, err := time.LoadLocation("Asia/Hong_Kong")
	if err != nil {
		log.Printf("时区设置失败 %s", err)
	}
	time.Local = timeLocal
	blockchain.Init()
	//go service.Init()
	gin.SetMode(setting.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	if err := server.ListenAndServe(); err != nil {
		log.Printf("start http server failed %s", err)
	}
}
