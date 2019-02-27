package test


import (
	"github.com/jinzhu/gorm"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"sample-clean-arch/domain/model"
	"sample-clean-arch/infrastructure/datastore"
	"sample-clean-arch/interface/presenters"
	"sample-clean-arch/usecase/service"
	"testing"
	"github.com/stretchr/testify/assert"

)



func TestCreate(t *testing.T) {
	//dbをスタブ用意
	var db *gorm.DB
	_,_, err := sqlmock.NewWithDSN("sqlmock_db_0")
	if err != nil {
		panic("Got an unexpected error.")
	}
	db, err = gorm.Open("sqlmock", "sqlmock_db_0")
	if err != nil {
		panic("Got an unexpected error.")
	}
	defer db.Close()
	mockPostRepo := datastore.NewPostRepository(db)
	mockPresenter := presenters.NewPostPresenter()
	mokeMember := model.Member{
		ID: 1,
		Name:"Takashi",
	}
	mockPost := model.Post{
		Member:mokeMember,
		Name:   "Hello",
		Content: "Content",
	}
	se :=service.NewPostService (mockPostRepo, mockPresenter)
	seErr := se.Create(&mockPost)
	//正常系のみ
	assert.NoError(t, seErr)

}

