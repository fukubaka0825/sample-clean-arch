package mocks

import (
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

func NewMockPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

func (postRepository *postRepository) Store(post *model.Post) error {
	return nil
}

func (postRepository *postRepository) FindAll(posts []*model.Post) ([]*model.Post, error) {
	//postRepository.db.Table("posts").Select("members.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)

	return posts, nil
}

func (postRepository *postRepository) FindById(id int) (*model.Post, error) {
	//postRepository.db.Table("posts").Select("members.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)
	post := model.Post{}

	return &post, nil
}

func (postRepository *postRepository) Update(post *model.Post) error {

	return nil
}

func (postRepository *postRepository) Delete(post *model.Post) error {

	return  nil
}

