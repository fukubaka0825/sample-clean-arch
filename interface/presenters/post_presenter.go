package presenters

import (
	"sample-clean-arch/domain/model"
)

type postPresenter struct {
}

func NewPostPresenter() PostPresenter {
	return &postPresenter{}
}

type PostPresenter interface {
	ResponsePosts(us []*model.Post) []*model.Post
	ResponsePost(us *model.Post) *model.Post
}

func (postPresenter *postPresenter) ResponsePosts(us []*model.Post) []*model.Post {
	//特に加工がないからそのまま返す
	return us
}

func (postPresenter *postPresenter) ResponsePost(us *model.Post) *model.Post {
	//特に加工がないからそのまま返す
	return us
}

