package handlers

import (
	"context"

	pb "github.com/ratheeshkumar25/opti_cut_admin/pkg/pb"
)

func (a *AdminHandler) AdminLoginRequest(ctx context.Context, p *pb.AdminLogin) (*pb.AdminResponse, error) {
	response, err := a.SVC.LoginService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (a *AdminHandler) AdminViewProfile(ctx context.Context, p *pb.AdID) (*pb.AdminProfile, error) {
	response, err := a.SVC.AdminViewProfileService(p)
	if err != nil {
		return response, nil
	}
	return response, nil
}
