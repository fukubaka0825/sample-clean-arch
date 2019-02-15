package presenter

import "sample-clean-arch/domain/model"

type MemberPresenter interface {
	ResponseMembers(us []*model.Member) []*model.Member
}
