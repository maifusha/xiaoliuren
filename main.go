package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"

	"xiaoliuren/config"
	"xiaoliuren/lib/templatekit"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var router *gin.Engine

var db *gorm.DB

//go:embed template
var templateFS embed.FS

//go:embed static
var staticFS embed.FS

func init() {
	if gin.Mode() == gin.ReleaseMode {
		//必须在router初始化前配置才有效
		gin.DefaultWriter, _ = os.Create(os.Getenv("GOPATH") + config.RUNTIME_LOG)
		gin.DefaultErrorWriter, _ = os.Create(os.Getenv("GOPATH") + config.ERROR_LOG)
	}

	//从嵌入式文件系统加载模板和静态资源
	router = gin.Default()
	router.HTMLRender = templatekit.New(&templateFS).MultiRender()
	subStatic, _ := fs.Sub(staticFS, "static")
	router.StaticFS("/static", http.FS(subStatic))

	ds := sqlite.Open(config.DBPATH)
	db, _ = gorm.Open(ds, &gorm.Config{})
}

func main() {
	router.GET("/home/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.tpl", gin.H{})
	})

	log.Fatal(router.Run(config.HOST + ":" + config.PORT))
}
