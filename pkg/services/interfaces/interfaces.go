package interfaces

import (
	pb "github.com/ratheeshkumar25/opti_cut_admin/pkg/pb"
)

type AdminServiceInter interface {
	LoginService(p *pb.AdminLogin) (*pb.AdminResponse, error)
	BlockUserService(p *pb.AdID) (*pb.AdminResponse, error)
	UnblockUserService(p *pb.AdID) (*pb.AdminResponse, error)
}
