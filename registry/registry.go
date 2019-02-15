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

type interactor struct {
	conn *gorm.DB
}

type Interactor interface {
	NewAppHandler() handler.AppHandler
}

func NewInteractor(conn *gorm.DB) Interactor {
	return &interactor{conn}
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	return i.NewAppHandler()
}

func (i *interactor) NewMemberHandler() handler.MemberHandler {
	return handler.NewMemberHandler(i.NewMemberController())
}

func (i *interactor) NewMemberController() controllers.MemberController {
	return controllers.NewMemberController(i.NewMemberService())
}

func (i *interactor) NewMemberService() service.MemberService {
	return service.NewMemberService(i.NewMemberRepository(), i.NewMemberPresenter())
}

func (i *interactor) NewMemberRepository() repository.MemberRepository {
	return datastore.NewMemberRepository(i.conn)
}

func (i *interactor) NewMemberPresenter() presenter.MemberPresenter {
	return presenters.NewMemberPresenter()
}

func (i *interactor) NewPostHandler() handler.PostHandler {
	return handler.NewPostHandler(i.NewPostController())
}

func (i *interactor) NewPostController() controllers.PostController {
	return controllers.NewPostController(i.NewPostService())
}

func (i *interactor) NewPostService() service.PostService {
	return service.NewPostService(i.NewPostRepository(), i.NewPostPresenter())
}

func (i *interactor) NewPostRepository() repository.PostRepository {
	return datastore.NewPostRepository(i.conn)
}

func (i *interactor) NewPostPresenter() presenter.PostPresenter {
	return presenters.NewPostPresenter()
}
