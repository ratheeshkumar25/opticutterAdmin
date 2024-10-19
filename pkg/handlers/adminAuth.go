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
