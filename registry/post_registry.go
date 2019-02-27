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

type postInteractor struct {
	conn *gorm.DB
}

type PostInteractor interface {
	NewPostHandler() handler.PostHandler
}

func NewPostInteractor(conn *gorm.DB) PostInteractor {
	return &postInteractor{conn}
}


func (i *postInteractor) NewPostHandler() handler.PostHandler {
	return handler.NewPostHandler(i.NewPostController())
}

func (i *postInteractor) NewPostController() controllers.PostController {
	return controllers.NewPostController(i.NewPostService())
}

func (i *postInteractor) NewPostService() service.PostService {
	return service.NewPostService(i.NewPostRepository(), i.NewPostPresenter())
}

func (i *postInteractor) NewPostRepository() repository.PostRepository {
	return datastore.NewPostRepository(i.conn)
}

func (i *postInteractor) NewPostPresenter() presenter.PostPresenter {
	return presenters.NewPostPresenter()
}
