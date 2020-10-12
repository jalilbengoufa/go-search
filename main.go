package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jalilbengoufa/go-search/models"
	"github.com/jalilbengoufa/go-search/pkg/redis"
	"github.com/jalilbengoufa/go-search/pkg/setting"
	"github.com/jalilbengoufa/go-search/routers"
)

func init() {
	setting.Setup()
	models.Setup()
	redis.Setup()
}

func main() {

	gin.SetMode(setting.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := &http.Server{
		Addr:    endPoint,
		Handler: routersInit,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

}
