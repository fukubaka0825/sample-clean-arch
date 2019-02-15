package repository

import "sample-clean-arch/domain/model"

type PostRepository interface {
	Store(post *model.Post) error
	FindAll(posts []*model.Post) ([]*model.Post, error)
}
