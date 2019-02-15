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
	ResponseMembers(us []*model.Member) []*model.Member
}

func (memberPresenter *memberPresenter) ResponseMembers(us []*model.Member) []*model.Member {
	for _, u := range us {
		u.Name = u.Name + "hoge"
	}
	return us
}
