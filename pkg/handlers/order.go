package handlers

import (
	"context"

	pb "github.com/ratheeshkumar25/opti_cut_admin/pkg/pb"
)

func (a *AdminHandler) OrderHistory(ctx context.Context, p *pb.AdminItemNoParams) (*pb.AdminOrderList, error) {
	response, err := a.SVC.FindAllOrdersSvc(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
func (a *AdminHandler) FindOrder(ctx context.Context, p *pb.AdminItemID) (*pb.AdminOrder, error) {
	response, err := a.SVC.FindOrderSvc(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
func (a *AdminHandler) FindOrdersByUser(ctx context.Context, p *pb.AdminItemID) (*pb.AdminOrderList, error) {
	response, err := a.SVC.FindOrdersByUserSvc(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
