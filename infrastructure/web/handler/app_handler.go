package handler

import "sample-clean-arch/interface/controllers"

type appHandler struct {
	memberController controllers.MemberController
	postController controllers.PostController
}

type AppHandler interface {
	MemberHandler
	PostHandler
}

//func NewAppHandler(cm controllers.MemberController,cp controllers.PostController) AppHandler {
//	return &appHandler{memberController: cm,cp}
//}
