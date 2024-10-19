package services

import (
	"errors"

	"github.com/ratheeshkumar25/opti_cut_admin/config"
	pb "github.com/ratheeshkumar25/opti_cut_admin/pkg/pb"
	"github.com/ratheeshkumar25/opti_cut_admin/pkg/utils"
)

// LoginService implements interfaces.AdminServiceInter.
func (a *AdminService) LoginService(p *pb.AdminLogin) (*pb.AdminResponse, error) {
	admin, err := a.Repo.FindAdminByEmail(p.Email)
	if err != nil {
		return nil, err
	}

	if admin.Password != p.Password {
		return &pb.AdminResponse{
			Status:  pb.AdminResponse_ERROR,
			Message: "Type correct Password",
			Payload: &pb.AdminResponse_Error{Error: "Incorrect Password"},
		}, errors.New("incorrect password")
	}

	jwtToken, err := utils.GenerateToken(config.LoadConfig().SECERETKEY, admin.Email)
	if err != nil {
		return &pb.AdminResponse{
			Status:  pb.AdminResponse_ERROR,
			Message: "error in generating token",
			Payload: &pb.AdminResponse_Error{Error: err.Error()},
		}, errors.New("error generating token")
	}

	return &pb.AdminResponse{
		Status:  pb.AdminResponse_OK,
		Message: "Login successful",
		Payload: &pb.AdminResponse_Data{Data: jwtToken},
	}, nil
}
