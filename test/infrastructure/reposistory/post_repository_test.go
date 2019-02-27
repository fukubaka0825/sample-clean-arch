package test

import (
	"github.com/stretchr/testify/assert"
	"sample-clean-arch/domain/model"
	"sample-clean-arch/infrastructure/datastore"
	"testing"
)

func TestPostStore(t *testing.T) {
	//ほんとはgormDBのmock
	db := datastore.NewMySqlDB()
	defer db.Close()

	//前提がmemberIdが1のmemberデータが入っている事
	a := datastore.NewPostRepository(db)
	post := model.Post{
		MemberID:1,
		Name:"aaaa",
		Content:"aaaa",
	}
	err := a.Store(&post)
	assert.NoError(t, err)
}

