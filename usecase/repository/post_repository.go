package repository

import "sample-clean-arch/domain/model"

type PostRepository interface {
	Store(post *model.Post) error
	FindAll(posts []*model.Post) ([]*model.Post, error)
	FindById(id int)(*model.Post, error)
	Update(post *model.Post) error
	Delete(post *model.Post) error
}
