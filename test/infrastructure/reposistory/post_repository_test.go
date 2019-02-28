package test
//
//import (
//	"github.com/jinzhu/gorm"
//	"github.com/stretchr/testify/assert"
//	"gopkg.in/DATA-DOG/go-sqlmock.v1"
//	"sample-clean-arch/domain/model"
//	"sample-clean-arch/infrastructure/datastore"
//	"testing"
//)
//
//func TestPostStore(t *testing.T) {
//	//dbスタブ用意
//	var db *gorm.DB
//	_,_, err := sqlmock.NewWithDSN("sqlmock_db_0")
//	if err != nil {
//		panic("Got an unexpected error.")
//	}
//	db, err = gorm.Open("sqlmock", "sqlmock_db_0")
//if err != nil {
//panic("Got an unexpected error.")
//}
//defer db.Close()
////前提がmemberIdが1のmemberデータが入っている事
//a := datastore.NewPostRepository(db)
//post := model.Post{
//MemberID:1,
//Name:"aaaa",
//Content:"aaaa",
//}
//aErr := a.Store(&post)
////正常系のみ
//assert.NoError(t, aErr)
//}
//
