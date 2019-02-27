package registry

import (
	"github.com/jinzhu/gorm"
	"sample-clean-arch/infrastructure/web/handler"
	"sample-clean-arch/infrastructure/datastore"
	"sample-clean-arch/interface/controllers"
	"sample-clean-arch/interface/presenters"
	"sample-clean-arch/usecase/presenter"
	"sample-clean-arch/usecase/repository"
	"sample-clean-arch/usecase/service"
)

type memberInteractor struct {
	conn *gorm.DB
}

type MemberInteractor interface {
	NewMemberHandler()handler.MemberHandler
}

func NewMemberInteractor(conn *gorm.DB) MemberInteractor {
	return &memberInteractor{conn}
}


func (i *memberInteractor) NewMemberHandler() handler.MemberHandler {
	return handler.NewMemberHandler(i.NewMemberController())
}

func (i *memberInteractor) NewMemberController() controllers.MemberController {
	return controllers.NewMemberController(i.NewMemberService())
}

func (i *memberInteractor) NewMemberService() service.MemberService {
	return service.NewMemberService(i.NewMemberRepository(), i.NewMemberPresenter())
}

func (i *memberInteractor) NewMemberRepository() repository.MemberRepository {
	return datastore.NewMemberRepository(i.conn)
}

func (i *memberInteractor) NewMemberPresenter() presenter.MemberPresenter {
	return presenters.NewMemberPresenter()
}

