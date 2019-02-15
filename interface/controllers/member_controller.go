package controllers

import (
	"sample-clean-arch/domain/model"
	"sample-clean-arch/usecase/service"
)

type memberController struct {
	memberService service.MemberService
}

type MemberController interface {
	CreateMember(member *model.Member) error
	GetMembers() ([]*model.Member, error)
}

func NewMemberController(us service.MemberService) MemberController {
	return &memberController{us}
}

func (memberController *memberController) CreateMember(member *model.Member) error {
	// 内側の処理のための技術的な関心事を記載
	return memberController.memberService.Create(member)
}

func (memberController *memberController) GetMembers() ([]*model.Member, error) {
	u := []*model.Member{}
	us, err := memberController.memberService.Get(u)

	if err != nil {
		return nil, err
	}
	return us, nil
}
