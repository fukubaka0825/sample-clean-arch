package handler

import (
"context"

"net/http"

"github.com/labstack/echo"
"sample-clean-arch/domain/model"
"sample-clean-arch/interface/controllers"
)

type postHandler struct {
	postController controllers.PostController
}

type PostHandler interface {
	CreatePost(c echo.Context) error
	GetPosts(c echo.Context) error
}

func NewPostHandler(uc controllers.PostController) PostHandler {
	return &postHandler{postController: uc}
}

func (uh *postHandler) CreatePost(c echo.Context) error {

	// リクエスト用のEntityを作成
	req := &model.Post{}

	// bind
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	// validate
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := uh.postController.CreatePost(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, "success")
}
func (uh *postHandler) GetPosts(c echo.Context) error {

	req := &model.Post{}

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	u, err := uh.postController.GetPosts()
	if err != nil {
		// システム内のエラー
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.Render(http.StatusOK, "posts_all", u)
}
