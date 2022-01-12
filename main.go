package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"xiaoliuren/config"
	"xiaoliuren/lib/calendar"
	"xiaoliuren/lib/templatekit"
	"xiaoliuren/request"
	"xiaoliuren/service"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

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
}

func main() {
	router.GET("/home/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.tpl", gin.H{})
	})

	router.GET("/home/jixiong", func(c *gin.Context) {
		req := request.NewJixiong()
		if c.ShouldBind(req) != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		svc := service.NewXiaoliuren()
		yuePan, riPan, shiPan := svc.GetSanPan(req.Qike, req.Date, req.Dizhi)

		c.JSON(http.StatusOK, gin.H{
			"lunar_time": svc.GetLunarTime(req.Date, req.Dizhi),
			"solar_time": svc.GetSolarTime(req.Date, req.Dizhi),
			"sanpan": gin.H{
				"yue_pan": yuePan,
				"ri_pan":  riPan,
				"shi_pan": shiPan,
			},
		})
	})

	router.GET("/home/dianbo", func(c *gin.Context) {
		svc := service.NewXiaoliuren()
		date := time.Now()
		dizhi := calendar.NowDizhi()

		c.JSON(http.StatusOK, gin.H{
			"lunar_time": svc.GetLunarTime(date, dizhi),
			"solar_time": svc.GetSolarTime(date, dizhi),
			"liushen":    "",
			"duanci":     "",
		})
	})

	router.GET("/home/zeshi", func(c *gin.Context) {
		req := request.NewZeshi()
		if c.ShouldBind(req) != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, gin.H{})
	})

	log.Fatal(router.Run(config.HOST + ":" + config.PORT))
}
