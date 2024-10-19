package services

import (
	"context"

	userpb "github.com/ratheeshkumar25/opti_cut_admin/pkg/clients/user/pb"
	pb "github.com/ratheeshkumar25/opti_cut_admin/pkg/pb"
)

// BlockUserService handle the admin to block the users using the provided information
func (a *AdminService) BlockUserService(p *pb.AdID) (*pb.AdminResponse, error) {
	ctx := context.Background()
	user := &userpb.ID{
		ID: p.ID,
	}
	_, err := a.UserClient.BlockUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb.AdminResponse{
		Status:  pb.AdminResponse_OK,
		Message: "User blocked successfully",
	}, nil
}

// UnblockUserService implements interfaces.AdminServiceInter.
func (a *AdminService) UnblockUserService(p *pb.AdID) (*pb.AdminResponse, error) {
	ctx := context.Background()
	user := &userpb.ID{
		ID: p.ID,
	}
	_, err := a.UserClient.UnblockUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb.AdminResponse{
		Status:  pb.AdminResponse_OK,
		Message: "User unblocked successfully",
	}, nil
}
