package router

import (
	"github.com/labstack/echo"
	"sample-clean-arch/registry"
)

func NewRouter(e *echo.Echo, r registry.Interactor) {
	//h := r.NewMemberHandler()
	g := r.NewPostHandler()

	//e.POST("/member", h.CreateMember)
	//e.GET("/members", h.GetMembers)
	e.POST("/posts", g.CreatePost)
	e.GET("/", g.GetPosts)
	e.POST("/post/edit",g.GetPost)
	e.POST("/post/update",g.UpdatePost)
	e.POST("/post/delete",g.DeletePost)
}
