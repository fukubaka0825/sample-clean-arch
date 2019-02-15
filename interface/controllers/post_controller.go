package controllers

import (
	"sample-clean-arch/domain/model"
	"sample-clean-arch/usecase/repository/view"
	"sample-clean-arch/usecase/service"
)

type postController struct {
	postService service.PostService
}

type PostController interface {
	CreatePost(post *model.Post) error
	GetPosts() ([]*view.PostView, error)
}

func NewPostController(us service.PostService) PostController {
	return &postController{us}
}

func (postController *postController) CreatePost(post *model.Post) error {
	// 内側の処理のための技術的な関心事を記載
	return postController.postService.Create(post)
}

func (postController *postController) GetPosts() ([]*view.PostView, error) {
	u := []*model.Post{}
	us, err := postController.postService.Get(u)
	if err != nil {
		return nil, err
	}
	return us, nil
}
