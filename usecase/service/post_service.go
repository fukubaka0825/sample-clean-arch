package service

import (
	"sample-clean-arch/domain/model"
	"sample-clean-arch/usecase/presenter"
	"sample-clean-arch/usecase/repository"
)

type postService struct {
	PostRepository repository.PostRepository
	PostPresenter  presenter.PostPresenter
}

type PostService interface {
	Create(u *model.Post) error
	Get(u []*model.Post) ([]*model.Post, error)
	GetForEditPost(id int) (*model.Post, error)
	Update(u *model.Post) error
	Delete(u *model.Post) error
}

func NewPostService(repo repository.PostRepository, pre presenter.PostPresenter) PostService {
	return &postService{repo, pre}
}

func (postService *postService) Create(u *model.Post) error {
	return postService.PostRepository.Store(u)
}

func (postService *postService) Get(u []*model.Post) ([]*model.Post, error) {
	us, err := postService.PostRepository.FindAll(u)
	if err != nil {
		return nil, err
	}
	return postService.PostPresenter.ResponsePosts(us), nil
}

func (postService *postService) GetForEditPost(id int)(*model.Post, error){
	us, err := postService.PostRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return postService.PostPresenter.ResponsePost(us),nil
}

func (postService *postService) Update(u *model.Post) error{

	return postService.PostRepository.Update(u)
}

func (postService *postService) Delete(u *model.Post) error{
	return postService.PostRepository.Delete(u)
}