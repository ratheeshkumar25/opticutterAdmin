package services

import (
	"context"

	materialpb "github.com/ratheeshkumar25/opti_cut_admin/pkg/clients/material/pb"
	pb "github.com/ratheeshkumar25/opti_cut_admin/pkg/pb"
)

// FindOrderSvc implements interfaces.AdminServiceInter.
func (a *AdminService) FindOrderSvc(p *pb.AdminItemID) (*pb.AdminOrder, error) {
	ctx := context.Background()

	order, err := a.MaterialClient.FindOrder(ctx, &materialpb.ItemID{ID: p.ID})
	if err != nil {
		return nil, err
	}
	return &pb.AdminOrder{
		Order_ID:   uint32(order.Order_ID),
		User_ID:    uint32(order.User_ID),
		Item_ID:    uint32(order.Item_ID),
		Status:     order.Status,
		Amount:     order.Amount,
		Is_Custom:  order.Is_Custom,
		Payment_ID: order.Payment_ID,
	}, nil
}

// FindOrdersByUserSvc implements interfaces.AdminServiceInter.
func (a *AdminService) FindOrdersByUserSvc(p *pb.AdminItemID) (*pb.AdminOrderList, error) {
	ctx := context.Background()

	response, err := a.MaterialClient.FindOrdersByUser(ctx, &materialpb.ItemID{ID: p.ID})
	if err != nil {
		return nil, err
	}
	var orders []*pb.AdminOrder

	for _, order := range response.Orders {

		orders = append(orders, &pb.AdminOrder{
			Order_ID:   uint32(order.Order_ID),
			User_ID:    uint32(order.User_ID),
			Item_ID:    uint32(order.Item_ID),
			Status:     order.Status,
			Amount:     order.Amount,
			Is_Custom:  order.Is_Custom,
			Payment_ID: order.Payment_ID,
		})
	}

	return &pb.AdminOrderList{
		Orders: orders,
	}, nil
}

// OrderHistoryService implements interfaces.AdminServiceInter.
func (a *AdminService) FindAllOrdersSvc(p *pb.AdminItemNoParams) (*pb.AdminOrderList, error) {
	ctx := context.Background()

	response, err := a.MaterialClient.OrderHistory(ctx, &materialpb.ItemNoParams{})
	if err != nil {
		return nil, err
	}
	var orders []*pb.AdminOrder

	for _, order := range response.Orders {

		orders = append(orders, &pb.AdminOrder{
			Order_ID:   uint32(order.Order_ID),
			User_ID:    uint32(order.User_ID),
			Item_ID:    uint32(order.Item_ID),
			Status:     order.Status,
			Amount:     order.Amount,
			Is_Custom:  order.Is_Custom,
			Payment_ID: order.Payment_ID,
		})
	}

	return &pb.AdminOrderList{
		Orders: orders,
	}, nil
}
