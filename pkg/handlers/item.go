package handlers

import (
	"context"

	pb "github.com/ratheeshkumar25/opti_cut_admin/pkg/pb"
)

func (a *AdminHandler) FindAllItem(ctx context.Context, p *pb.AdminItemNoParams) (*pb.AdminItemList, error) {
	response, err := a.SVC.FindAllItemService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
