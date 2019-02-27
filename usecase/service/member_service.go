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
	Create(member *model.Member) error
	Get(members []*model.Member) ([]*model.Member, error)
}

func NewMemberService(repo repository.MemberRepository, pre presenter.MemberPresenter) MemberService {
	return &memberService{repo, pre}
}

func (memberService *memberService) Create(member *model.Member) error {
	return memberService.MemberRepository.Store(member)
}

func (memberService *memberService) Get(members []*model.Member) ([]*model.Member, error) {
	mems, err := memberService.MemberRepository.FindAll(members)
	if err != nil {
		return nil, err
	}
	return memberService.MemberPresenter.ResponseMembers(mems), nil
}
