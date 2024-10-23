package handlers

import (
	"context"

	pb "github.com/ratheeshkumar25/opti_cut_admin/pkg/pb"
)

func (a *AdminHandler) AddMaterial(ctx context.Context, p *pb.AdminMaterial) (*pb.AdminResponse, error) {
	response, err := a.SVC.AddMaterialService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
func (a *AdminHandler) FindMaterialByID(ctx context.Context, p *pb.AdminMaterialID) (*pb.AdminMaterial, error) {
	response, err := a.SVC.FindMaterialByIDService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
func (a *AdminHandler) FindAllMaterial(ctx context.Context, p *pb.AdminItemNoParams) (*pb.AdminMaterialList, error) {
	response, err := a.SVC.FindAllMaterialService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
func (a *AdminHandler) EditMaterial(ctx context.Context, p *pb.AdminMaterial) (*pb.AdminMaterial, error) {
	response, err := a.SVC.EditMaterialService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
func (a *AdminHandler) RemoveMaterial(ctx context.Context, p *pb.AdminMaterialID) (*pb.AdminResponse, error) {
	response, err := a.SVC.RemoveMaterialService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
