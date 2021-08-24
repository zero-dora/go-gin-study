package main

import (
	"fmt"
	"github.com/zero-dora/go-gin-example/models"
	"github.com/zero-dora/go-gin-example/pkg/gredis"
	"github.com/zero-dora/go-gin-example/pkg/logging"
	"github.com/zero-dora/go-gin-example/pkg/setting"
	"github.com/zero-dora/go-gin-example/routers"
	"log"
	"net/http"
)

func main() {
	setting.SetUp()
	models.SetUp()
	logging.Setup()
	gredis.SetUp()

	router :=routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
