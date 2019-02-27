package test


import (
	"sample-clean-arch/domain/model"
	"sample-clean-arch/infrastructure/datastore"
	"sample-clean-arch/interface/presenters"
	"sample-clean-arch/usecase/service"
	"testing"
	"github.com/stretchr/testify/assert"

)



func TestCreate(t *testing.T) {
	//ほんとはgormDBのmock
	db := datastore.NewMySqlDB()
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
	err := se.Create(&mockPost)
	//正常系のみ
	assert.NoError(t, err)

}

