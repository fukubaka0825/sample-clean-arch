package registry

import "github.com/jinzhu/gorm"

type interactor struct {
	memberInteractor
	postInteractor
}

type Interactor interface {
	MemberInteractor
	PostInteractor
}
func NewInteractor(conn *gorm.DB) Interactor {

	return &interactor{
		memberInteractor{conn},
			postInteractor{conn},
	}
}
