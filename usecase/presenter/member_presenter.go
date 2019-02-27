package presenter

import "sample-clean-arch/domain/model"

type MemberPresenter interface {
	ResponseMembers(members []*model.Member) []*model.Member
}
