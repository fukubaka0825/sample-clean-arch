package service

import (
	"sample-clean-arch/domain/model"
	"sample-clean-arch/usecase/repository/view"
	"sample-clean-arch/usecase/presenter"
	"sample-clean-arch/usecase/repository"
)

type postService struct {
	PostRepository repository.PostRepository
	PostPresenter  presenter.PostPresenter
}

type PostService interface {
	Create(u *model.Post) error
	Get(u []*model.Post) ([]*view.PostView, error)
}

func NewPostService(repo repository.PostRepository, pre presenter.PostPresenter) PostService {
	return &postService{repo, pre}
}

func (postService *postService) Create(u *model.Post) error {
	return postService.PostRepository.Store(u)
}

func (postService *postService) Get(u []*model.Post) ([]*view.PostView, error) {
	us, err := postService.PostRepository.FindAll(u)
	if err != nil {
		return nil, err
	}
	return postService.PostPresenter.ResponsePosts(us), nil
}
