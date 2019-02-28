package test

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"sample-clean-arch/domain/model"
	"sample-clean-arch/interface/presenters"
	"sample-clean-arch/usecase/service"
	"sample-clean-arch/mocks"
	"testing"
)



func TestCreate(t *testing.T) {
	//dbをスタブ用意
	var db *gorm.DB
	_,mock, err := sqlmock.NewWithDSN("sqlmock_db_0")
	if err != nil {
		panic("Got an unexpected error.")
	}
	db, err = gorm.Open("sqlmock", "sqlmock_db_0")
	if err != nil {
		panic("Got an unexpected error.")
	}
	defer db.Close()

	mockPost := model.Post{}
	mockPostRepo := mocks.NewMockPostRepository(db)
	mockPresenter := presenters.NewPostPresenter()

	se :=service.NewPostService (mockPostRepo, mockPresenter)
	assert.NoError(t, se.Create(&mockPost))
	assert.NoError(t, mock.ExpectationsWereMet())
}

