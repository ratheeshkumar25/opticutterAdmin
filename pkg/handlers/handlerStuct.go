package handlers

import (
	pb "github.com/ratheeshkumar25/opti_cut_admin/pkg/pb"
	inter "github.com/ratheeshkumar25/opti_cut_admin/pkg/services/interfaces"
)

type AdminHandler struct {
	SVC inter.AdminServiceInter
	pb.AdminServiceServer
}

func NewAdminHandler(svc inter.AdminServiceInter) *AdminHandler {
	return &AdminHandler{
		SVC: svc,
	}
}
