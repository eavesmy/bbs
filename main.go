package main

import (
	"fmt"
	"net/http"

	"gorobbs/bootstrap"
	"gorobbs/package/setting"
	router "gorobbs/router/v1"

	"github.com/eavesmy/glog"
)

var log = glog.New("[app]")

func main() {
	bootstrap.SetUp()

	router := router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	log.Info("start")

	if err := s.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}
