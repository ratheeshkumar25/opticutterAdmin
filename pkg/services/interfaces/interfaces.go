package interfaces

import (
	pb "github.com/ratheeshkumar25/opti_cut_admin/pkg/pb"
)

type AdminServiceInter interface {
	LoginService(p *pb.AdminLogin) (*pb.AdminResponse, error)
	BlockUserService(p *pb.AdID) (*pb.AdminResponse, error)
	UnblockUserService(p *pb.AdID) (*pb.AdminResponse, error)

	AddMaterialService(p *pb.AdminMaterial) (*pb.AdminResponse, error)
	FindMaterialByIDService(p *pb.AdminMaterialID) (*pb.AdminMaterial, error)
	FindAllMaterialService(p *pb.AdminItemNoParams) (*pb.AdminMaterialList, error)
	EditMaterialService(p *pb.AdminMaterial) (*pb.AdminMaterial, error)
	RemoveMaterialService(p *pb.AdminMaterialID) (*pb.AdminResponse, error)

	FindAllItemService(p *pb.AdminItemNoParams) (*pb.AdminItemList, error)

	FindAllOrdersSvc(p *pb.AdminItemNoParams) (*pb.AdminOrderList, error)
	FindOrderSvc(p *pb.AdminItemID) (*pb.AdminOrder, error)
	FindOrdersByUserSvc(p *pb.AdminItemID) (*pb.AdminOrderList, error)
}
