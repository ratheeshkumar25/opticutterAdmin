package services

import (
	"context"

	materialpb "github.com/ratheeshkumar25/opti_cut_admin/pkg/clients/material/pb"
	pb "github.com/ratheeshkumar25/opti_cut_admin/pkg/pb"
)

// FindAllItemService implements interfaces.AdminServiceInter.
func (a *AdminService) FindAllItemService(p *pb.AdminItemNoParams) (*pb.AdminItemList, error) {
	ctx := context.Background()

	// Call the MaterialClient's FindAllItem method
	result, err := a.MaterialClient.FindAllItem(ctx, &materialpb.ItemNoParams{})
	if err != nil {
		return nil, err
	}

	// Convert materialpb.ItemList to pb.ItemList
	var items []*pb.AdminItem
	for _, item := range result.Items {
		pbItem := &pb.AdminItem{
			Item_ID:         item.Item_ID,
			Item_Name:       item.Item_Name,
			Length:          item.Length,
			Width:           item.Width,
			Fixed_Size_ID:   item.Fixed_Size_ID,
			Is_Custom:       item.Is_Custom,
			Estimated_Price: item.Estimated_Price,
		}
		items = append(items, pbItem)
	}

	return &pb.AdminItemList{
		Items: items,
	}, nil
}
