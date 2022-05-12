package main

import (
	"embed"
	"fmt"
	"net/http"
	"os"
	"path"
	"time"

	"xiaoliuren/internal/config"
	"xiaoliuren/internal/router"
	"xiaoliuren/internal/util"
	"xiaoliuren/internal/util/file"
	"xiaoliuren/internal/util/logger"
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
		gin.DefaultWriter = file.MustOpen(path.Join(os.Getenv("GOPATH"), config.Conf.Logfile.Request))
		gin.DefaultErrorWriter = file.MustOpen(path.Join(os.Getenv("GOPATH"), config.Conf.Logfile.Panic))
	}
}

func main() {
	logger.Printf("Application start: %s\n", util.NewNow())

	handler := router.New(gin.Default()).SetRenderWithEmbed(&templateFS)
	srv := &http.Server{
		Addr:        fmt.Sprintf("%s:%s", config.Conf.Server.Host, config.Conf.Server.Port),
		Handler:     handler.StaticBind(&staticFS).HandleBind(),
		IdleTimeout: time.Second * time.Duration(config.Conf.Server.IdleTimeout),
	}

	go grace.New(srv).ListenDown()

	if err := srv.ListenAndServe(); err != nil {
		logger.Printf("Server error: %s\n", err)
	}
	logger.Printf("Application down: %s\n", util.NewNow())
}
