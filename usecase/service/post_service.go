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
	Create(post *model.Post) error
	Get(posts []*model.Post) ([]*model.Post, error)
	GetForEditPost(id int) (*model.Post, error)
	Update(post *model.Post) error
	Delete(post *model.Post) error
}

func NewPostService(repo repository.PostRepository, pre presenter.PostPresenter) PostService {
	return &postService{repo, pre}
}

func (postService *postService) Create(post *model.Post) error {
	return postService.PostRepository.Store(post)
}

func (postService *postService) Get(posts []*model.Post) ([]*model.Post, error) {
	pos, err := postService.PostRepository.FindAll(posts)
	if err != nil {
		return nil, err
	}
	return postService.PostPresenter.ResponsePosts(pos), nil
}

func (postService *postService) GetForEditPost(id int)(*model.Post, error){
	po, err := postService.PostRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return postService.PostPresenter.ResponsePost(po),nil
}

func (postService *postService) Update(post *model.Post) error{

	return postService.PostRepository.Update(post)
}

func (postService *postService) Delete(post *model.Post) error{
	return postService.PostRepository.Delete(post)
}