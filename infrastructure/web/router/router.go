package router

import (
	"github.com/labstack/echo"
	"sample-clean-arch/infrastructure/web/handler"
)

func NewRouter(e *echo.Echo, handler handler.AppHandler) {
	e.POST("/members", handler.CreateMember)
	e.GET("/members", handler.GetMembers)
	e.POST("/posts", handler.CreatePost)
	e.GET("/posts", handler.GetPosts)
}
