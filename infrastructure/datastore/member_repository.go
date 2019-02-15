package datastore

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"sample-clean-arch/domain/model"
)

type memberRepository struct {
	db *gorm.DB
}

type MemberRepository interface {
	Store(member *model.Member) error
	FindAll(members []*model.Member) ([]*model.Member, error)
}

func NewMemberRepository(db *gorm.DB) MemberRepository {
	return &memberRepository{db}
}

func (memberRepository *memberRepository) Store(member *model.Member) error {
	return memberRepository.db.Save(member).Error

}

func (memberRepository *memberRepository) FindAll(members []*model.Member) ([]*model.Member, error) {

	err := memberRepository.db.Find(&members).Error
	if err != nil {
		return nil, fmt.Errorf("sql error", err)
	}

	return members, nil
}
