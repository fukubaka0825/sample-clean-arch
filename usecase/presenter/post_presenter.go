package presenter

import (
	"sample-clean-arch/domain/model"
)

type PostPresenter interface {
	ResponsePosts(posts []*model.Post) []*model.Post
	ResponsePost(post *model.Post) *model.Post
}