package router

import (
	"embed"
	"io/fs"
	"net/http"

	homehandler "xiaoliuren/internal/router/handler/home"
	"xiaoliuren/internal/util/logger"
	"xiaoliuren/pkg/templatekit"

	"github.com/gin-gonic/gin"
)

type router struct {
	*gin.Engine
}

func New(engine *gin.Engine) *router {
	return &router{Engine: engine}
}

func (r *router) SetRenderWithEmbed(templateFS *embed.FS) *router {
	r.HTMLRender = templatekit.New(templateFS).MultiRender()

	return r
}

func (r *router) StaticBind(staticFS *embed.FS) *router {
	subStatic, err := fs.Sub(staticFS, "static")
	if err != nil {
		logger.Fatalln(err)
	}

	r.StaticFS("/static", http.FS(subStatic))

	return r
}

func (r *router) HandleBind() *router {
	r.GET("/", homehandler.Redirect)

	homeGroup := r.Group("/home")
	{
		homeGroup.GET("/index", homehandler.Index)

		homeGroup.GET("/jixiong", homehandler.Jixiong)

		homeGroup.GET("/dianbo", homehandler.Dianbo)

		homeGroup.GET("/zeshi", homehandler.Zeshi)
	}

	return r
}
