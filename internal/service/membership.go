package service

import "github.com/formulationapp/formulationapp/internal/repository"

type MembershipService interface {
}

type membershipService struct {
	membershipRepository repository.MembershipRepository
}

func newMembershipService(membershipRepository repository.MembershipRepository) MembershipService {
	return &membershipService{membershipRepository}
}
