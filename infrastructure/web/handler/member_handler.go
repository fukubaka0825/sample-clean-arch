package handler

import (
	"context"

	"net/http"

	"github.com/labstack/echo"
	"sample-clean-arch/domain/model"
	"sample-clean-arch/interface/controllers"
)

type memberHandler struct {
	memberController controllers.MemberController
}

type MemberHandler interface {
	CreateMember(c echo.Context) error
	GetMembers(c echo.Context) error
}

func NewMemberHandler(uc controllers.MemberController) MemberHandler {
	return &memberHandler{memberController: uc}
}

func (uh *memberHandler) CreateMember(c echo.Context) error {

	// リクエスト用のEntityを作成
	req := &model.Member{}

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

	err := uh.memberController.CreateMember(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, "success")
}
func (uh *memberHandler) GetMembers(c echo.Context) error {

	req := &model.Member{}

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	u, err := uh.memberController.GetMembers()
	if err != nil {
		// システム内のエラー
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, u)
}
