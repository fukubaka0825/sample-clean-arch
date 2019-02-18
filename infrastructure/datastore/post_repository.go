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
	FindById(id int) (*model.Post, error)
	Update(post *model.Post) error
	Delete(post *model.Post) error
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

func (postRepository *postRepository) Store(post *model.Post) error {
	return postRepository.db.Save(post).Error

}

func (postRepository *postRepository) FindAll(posts []*model.Post) ([]*model.Post, error) {
	//postRepository.db.Table("posts").Select("members.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)
	err := postRepository.db.Preload("Member").Find(&posts).Error
	if err != nil {
		return nil, fmt.Errorf("sql error", err)
	}

	return posts, nil
}

func (postRepository *postRepository) FindById(id int) (*model.Post, error) {
	//postRepository.db.Table("posts").Select("members.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)
	post := model.Post{}
	err := postRepository.db.Find(&post,id).Error
	if err != nil {
		return nil, fmt.Errorf("sql error", err)
	}

	return &post, nil
}

func (postRepository *postRepository) Update(post *model.Post) error {
	//postRepository.db.Table("posts").Select("members.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)
	err := postRepository.db.Save(&post).Error
	if err != nil {
		return fmt.Errorf("sql error", err)
	}

	return nil
}

func (postRepository *postRepository) Delete(post *model.Post) error {
	//postRepository.db.Table("posts").Select("members.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)
	err := postRepository.db.Delete(&post).Error
	if err != nil {
		return fmt.Errorf("sql error", err)
	}

	return  nil
}

