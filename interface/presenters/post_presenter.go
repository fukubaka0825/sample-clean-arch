package presenters

import (
	"sample-clean-arch/domain/model"
	"sample-clean-arch/interface/presenters/view"
)

type postPresenter struct {
}

func NewPostPresenter() PostPresenter {
	return &postPresenter{}
}

type PostPresenter interface {
	ResponsePosts(us []*model.Post) []*view.PostView
}

func (postPresenter *postPresenter) ResponsePosts(us []*model.Post) []*view.PostView {
	uvs := []*view.PostView{}
	for _, u := range us {
		vs := new(view.PostView{
			*u,
			u.Member.Name,
		})
		uvs = append(uvs,vs)
	}
	return uvs
}
