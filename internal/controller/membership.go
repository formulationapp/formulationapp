package controller

import "github.com/formulationapp/formulationapp/internal/service"

type MembershipController interface {
}

type membershipController struct {
	membershipService service.MembershipService
}

func newMembershipController(membershipService service.MembershipService) MembershipController {
	return &membershipController{membershipService}
}
