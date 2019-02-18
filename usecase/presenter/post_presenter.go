package presenter

import (
	"sample-clean-arch/domain/model"
)

type PostPresenter interface {
	ResponsePosts(us []*model.Post) []*model.Post
	ResponsePost(us *model.Post) *model.Post
}