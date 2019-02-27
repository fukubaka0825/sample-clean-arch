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
	ResponsePosts(posts []*model.Post) []*model.Post
	ResponsePost(posts *model.Post) *model.Post
}

func (postPresenter *postPresenter) ResponsePosts(posts []*model.Post) []*model.Post {
	for _,v := range posts{
		v.Member.Name = "Mr." + v.Name
		v.Content += "!!!!!!!!!"
	}
	return posts
}

func (postPresenter *postPresenter) ResponsePost(post *model.Post) *model.Post {
	post.Name = "Mr." + post.Name
	post.Content += "!!!!!!!!!"
	return post
}

