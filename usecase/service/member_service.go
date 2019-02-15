package service

import (
	"sample-clean-arch/domain/model"
	"sample-clean-arch/usecase/presenter"
	"sample-clean-arch/usecase/repository"
)

type memberService struct {
	MemberRepository repository.MemberRepository
	MemberPresenter  presenter.MemberPresenter
}

type MemberService interface {
	Create(u *model.Member) error
	Get(u []*model.Member) ([]*model.Member, error)
}

func NewMemberService(repo repository.MemberRepository, pre presenter.MemberPresenter) MemberService {
	return &memberService{repo, pre}
}

func (memberService *memberService) Create(u *model.Member) error {
	return memberService.MemberRepository.Store(u)
}

func (memberService *memberService) Get(u []*model.Member) ([]*model.Member, error) {
	us, err := memberService.MemberRepository.FindAll(u)
	if err != nil {
		return nil, err
	}
	return memberService.MemberPresenter.ResponseMembers(us), nil
}
