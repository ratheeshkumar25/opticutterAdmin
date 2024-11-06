package services

import (
	"errors"

	"github.com/ratheeshkumar25/opti_cut_admin/config"
	"github.com/ratheeshkumar25/opti_cut_admin/pkg/model"
	pb "github.com/ratheeshkumar25/opti_cut_admin/pkg/pb"
	"github.com/ratheeshkumar25/opti_cut_admin/pkg/utils"
	"gorm.io/gorm"
)

// // LoginService implements interfaces.AdminServiceInter.
// func (a *AdminService) LoginService(p *pb.AdminLogin) (*pb.AdminResponse, error) {
// 	// hashedPass, err := utils.HashPassword(p.Password)
// 	// if err != nil {
// 	// 	return &pb.AdminResponse{
// 	// 		Status:  pb.AdminResponse_ERROR,
// 	// 		Message: "error in hasing password",
// 	// 		Payload: &pb.AdminResponse_Error{Error: err.Error()},
// 	// 	}, errors.New("unable to hashpassword")
// 	// }
// 	// admin := &model.Admin{
// 	// 	Email:    p.Email,
// 	// 	Password: hashedPass,
// 	// }

// 	admin, err := a.Repo.FindAdminByEmail(p.Email)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if admin.Password != p.Password {
// 		return &pb.AdminResponse{
// 			Status:  pb.AdminResponse_ERROR,
// 			Message: "Type correct Password",
// 			Payload: &pb.AdminResponse_Error{Error: "Incorrect Password"},
// 		}, errors.New("incorrect password")
// 	}

// 	jwtToken, err := utils.GenerateToken(config.LoadConfig().SECERETKEY, admin.Email)
// 	if err != nil {
// 		return &pb.AdminResponse{
// 			Status:  pb.AdminResponse_ERROR,
// 			Message: "error in generating token",
// 			Payload: &pb.AdminResponse_Error{Error: err.Error()},
// 		}, errors.New("error generating token")
// 	}

// 	return &pb.AdminResponse{
// 		Status:  pb.AdminResponse_OK,
// 		Message: "Login successful",
// 		Payload: &pb.AdminResponse_Data{Data: jwtToken},
// 	}, nil
// }

// LoginService implements interfaces.AdminServiceInter.
func (a *AdminService) LoginService(p *pb.AdminLogin) (*pb.AdminResponse, error) {
	// Attem to find the admin by email
	admin, err := a.Repo.FindAdminByEmail(p.Email)
	if err != nil {
		// Check if the error is due to a missing record
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Admin not found, so create a new admin with a hashed password
			hashedPass, err := utils.HashPassword(p.Password)
			if err != nil {
				return &pb.AdminResponse{
					Status:  pb.AdminResponse_ERROR,
					Message: "error hashing password",
					Payload: &pb.AdminResponse_Error{Error: err.Error()},
				}, errors.New("unable to hash password")
			}
			newAdmin := &model.Admin{
				Email:    p.Email,
				Password: hashedPass,
			}
			_, err = a.Repo.CreateAdmin(newAdmin)
			if err != nil {
				return &pb.AdminResponse{
					Status:  pb.AdminResponse_ERROR,
					Message: "error creating admin",
					Payload: &pb.AdminResponse_Error{Error: err.Error()},
				}, errors.New("unable to create admin")
			}
			return &pb.AdminResponse{
				Status:  pb.AdminResponse_OK,
				Message: "Admin created successfully",
				Payload: nil,
			}, nil
		}
		// Other errors, return directly
		return nil, err
	}

	// Verify the password if the admin exists
	if err := utils.CheckPassword(admin.Password, p.Password); err != nil {
		return &pb.AdminResponse{
			Status:  pb.AdminResponse_ERROR,
			Message: "Incorrect password",
			Payload: &pb.AdminResponse_Error{Error: "Incorrect Password"},
		}, errors.New("incorrect password")
	}

	// Generate a JWT token after successful password verification
	jwtToken, err := utils.GenerateToken(config.LoadConfig().SECERETKEY, admin.Email)
	if err != nil {
		return &pb.AdminResponse{
			Status:  pb.AdminResponse_ERROR,
			Message: "error generating token",
			Payload: &pb.AdminResponse_Error{Error: err.Error()},
		}, errors.New("error generating token")
	}

	return &pb.AdminResponse{
		Status:  pb.AdminResponse_OK,
		Message: "Login successful",
		Payload: &pb.AdminResponse_Data{Data: jwtToken},
	}, nil
}
