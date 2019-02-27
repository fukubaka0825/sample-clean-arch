package presenters

import (
	"sample-clean-arch/domain/model"
)

type memberPresenter struct {

}

func NewMemberPresenter() MemberPresenter {
	return &memberPresenter{}
}

type MemberPresenter interface {
	ResponseMembers(members []*model.Member) []*model.Member
}

func (memberPresenter *memberPresenter) ResponseMembers(members []*model.Member) []*model.Member {
	for _, u := range members {
		u.Name = u.Name + "hoge"
	}
	return members
}
