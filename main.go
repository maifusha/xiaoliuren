package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"xiaoliuren/internal/config"
	"xiaoliuren/internal/router"
	"xiaoliuren/internal/util"
	"xiaoliuren/pkg/grace"

	"github.com/gin-gonic/gin"
)

//go:embed template
var templateFS embed.FS

//go:embed static
var staticFS embed.FS

func init() {
	gin.SetMode(config.Conf.Mode)

	if gin.Mode() == gin.ReleaseMode {
		gin.DefaultWriter, _ = os.Create(path.Join(os.Getenv("GOPATH"), config.Conf.Log.Request))
		gin.DefaultErrorWriter, _ = os.Create(path.Join(os.Getenv("GOPATH"), config.Conf.Log.Panic))
	}
}

func main() {
	log.Println("Application start:" + util.NewNow().String())

	handler := router.New(gin.Default()).SetRenderWithEmbed(&templateFS)
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Conf.Server.Host, config.Conf.Server.Port),
		Handler: handler.StaticBind(&staticFS).HandleBind(),
	}

	go grace.New(srv).Down()

	if err := srv.ListenAndServe(); err != nil {
		log.Println("Server exited：" + err.Error())
	}
	log.Println("Application down：" + util.NewNow().String())
}
