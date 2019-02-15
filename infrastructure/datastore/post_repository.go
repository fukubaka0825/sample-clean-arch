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
	FindAll(postds []*model.Post) ([]*model.Post, error)
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

func (postRepository *postRepository) Store(post *model.Post) error {
	return postRepository.db.Save(post).Error

}

func (postRepository *postRepository) FindAll(posts []*model.Post) ([]*model.Post, error) {
	//postRepository.db.Table("posts").Select("members.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)
	err := postRepository.db.Preload("members").Find(&posts).Scan(&posts).Error
	if err != nil {
		return nil, fmt.Errorf("sql error", err)
	}

	return posts, nil
}
