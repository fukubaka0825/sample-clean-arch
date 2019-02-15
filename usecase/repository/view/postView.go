package view

import "sample-clean-arch/domain/model"

type PostView struct{
	model.Post
	MemberName string
}