package model

type Post struct {
	ID         int `gorm:"primary_key"`
	Member  Member `member`
	MemberID   int `memberId`
	Name    string `"name"`
	Content string `"content"`
}
