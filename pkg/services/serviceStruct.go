package services

import (
	materialpb "github.com/ratheeshkumar25/opti_cut_admin/pkg/clients/material/pb"
	userpb "github.com/ratheeshkumar25/opti_cut_admin/pkg/clients/user/pb"
	inter "github.com/ratheeshkumar25/opti_cut_admin/pkg/repo/interfaces"
	"github.com/ratheeshkumar25/opti_cut_admin/pkg/services/interfaces"
)

type AdminService struct {
	Repo           inter.AdminRepoInter
	UserClient     userpb.UserServiceClient
	MaterialClient materialpb.MaterialServiceClient
}

func NewAdminRepository(repo inter.AdminRepoInter, userClient userpb.UserServiceClient, materialClient materialpb.MaterialServiceClient) interfaces.AdminServiceInter {
	return &AdminService{
		Repo:           repo,
		UserClient:     userClient,
		MaterialClient: materialClient,
	}
}
