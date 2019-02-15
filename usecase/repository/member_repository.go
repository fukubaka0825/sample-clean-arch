package repository

import "sample-clean-arch/domain/model"

type MemberRepository interface {
	Store(post *model.Member) error
	FindAll(posts []*model.Member) ([]*model.Member, error)
}
