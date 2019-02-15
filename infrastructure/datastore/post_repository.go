package datastore

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"sample-clean-arch/domain/model"
)

type postRepository struct {
	db *gorm.DB
}

type PostRepository interface {
	Store(post *model.Post) error
	FindAll(posts []*model.Post) ([]*model.Post, error)
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

func (postRepository *postRepository) Store(post *model.Post) error {
	return postRepository.db.Save(post).Error

}

func (postRepository *postRepository) FindAll(posts []*model.Post) ([]*model.Post, error) {

	err := postRepository.db.Find(&posts).Error
	if err != nil {
		return nil, fmt.Errorf("sql error", err)
	}

	return posts, nil
}
