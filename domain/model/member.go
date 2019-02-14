package model

type Member struct {
	ID       int `gorm:"primary_key"`
	Name  string `"name"`
	Posts []Post
	//Email string
}

