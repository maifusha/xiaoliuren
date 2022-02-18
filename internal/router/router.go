package router

import (
	"embed"
	"io/fs"
	"net/http"
	"xiaoliuren/internal/util/logger"

	"xiaoliuren/internal/router/handler"
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
	r.GET("/", handler.HomeRedirect)

	homeGroup := r.Group("/home")
	{
		homeGroup.GET("/index", handler.HomeIndex)

		homeGroup.GET("/jixiong", handler.HomeJixiong)

		homeGroup.GET("/dianbo", handler.HomeDianbo)

		homeGroup.GET("/zeshi", handler.HomeZeshi)
	}

	return r
}
