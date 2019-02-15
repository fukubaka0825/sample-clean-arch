package presenter

import (
	"sample-clean-arch/domain/model"
	"sample-clean-arch/interface/presenters/view"
)

type PostPresenter interface {
	ResponsePosts(us []*model.Post) []*view.PostView
}