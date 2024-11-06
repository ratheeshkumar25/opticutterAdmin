package handlers

import (
	"context"

	pb "github.com/ratheeshkumar25/opti_cut_admin/pkg/pb"
)

// AdminBlockUser helps to find the user and block by Admin service.
func (a *AdminHandler) AdminBlockUser(ctx context.Context, p *pb.AdID) (*pb.AdminResponse, error) {
	response, err := a.SVC.BlockUserService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// AdminBlockUser helps to find the user and block by Admin service.
func (a *AdminHandler) AdminUnblockUser(ctx context.Context, p *pb.AdID) (*pb.AdminResponse, error) {
	response, err := a.SVC.UnblockUserService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
