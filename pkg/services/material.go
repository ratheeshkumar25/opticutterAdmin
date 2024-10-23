package services

import (
	"context"
	"fmt"

	materialpb "github.com/ratheeshkumar25/opti_cut_admin/pkg/clients/material/pb"
	pb "github.com/ratheeshkumar25/opti_cut_admin/pkg/pb"
)

// AddMaterialService implements interfaces.AdminServiceInter.
func (a *AdminService) AddMaterialService(p *pb.AdminMaterial) (*pb.AdminResponse, error) {
	ctx := context.Background()

	newMaterial := &materialpb.Material{
		Material_ID:   p.Material_ID,
		Material_Name: p.Material_Name,
		Description:   p.Description,
		Stock:         p.Stock,
		Price:         p.Price,
	}

	materialID, err := a.MaterialClient.AddMaterial(ctx, newMaterial)
	if err != nil {
		return &pb.AdminResponse{
			Status:  pb.AdminResponse_ERROR,
			Message: "Failed to create material",
			Payload: &pb.AdminResponse_Error{Error: err.Error()},
		}, err
	}
	// Return success response with the new item ID
	return &pb.AdminResponse{
		Status:  pb.AdminResponse_OK,
		Message: "Material created successfully",
		Payload: &pb.AdminResponse_Data{
			Data: fmt.Sprintf("ItemID:%s", materialID),
		},
	}, nil
}

// EditMaterialService implements interfaces.AdminServiceInter.
func (a *AdminService) EditMaterialService(p *pb.AdminMaterial) (*pb.AdminMaterial, error) {
	ctx := context.Background()

	// Map pb.Item to materialpb.Item for sending to MaterialClient
	updatedmaterial := &materialpb.Material{
		Material_ID:   p.Material_ID,
		Material_Name: p.Material_Name,
		Description:   p.Description,
		Stock:         p.Stock,
		Price:         p.Price,
	}

	// Call the MaterialClient's EditItem method
	_, err := a.MaterialClient.EditMaterial(ctx, updatedmaterial)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// FindAllMaterialService implements interfaces.AdminServiceInter.
func (a *AdminService) FindAllMaterialService(p *pb.AdminItemNoParams) (*pb.AdminMaterialList, error) {
	ctx := context.Background()

	result, err := a.MaterialClient.FindAllMaterial(ctx, &materialpb.MaterialNoParams{})
	if err != nil {
		return nil, err
	}

	// slice to hold the  materials
	var materials []*pb.AdminMaterial

	for _, material := range result.Materials {
		pbMaterial := &pb.AdminMaterial{
			Material_ID:   uint32(material.Material_ID),
			Material_Name: material.Material_Name,
			Description:   material.Description,
			Stock:         int32(material.Stock),
			Price:         material.Price,
		}
		materials = append(materials, pbMaterial)
	}

	return &pb.AdminMaterialList{
		Materials: materials,
	}, nil
}

// FindMaterialByIDService implements interfaces.AdminServiceInter.
func (a *AdminService) FindMaterialByIDService(p *pb.AdminMaterialID) (*pb.AdminMaterial, error) {
	// Create a new context
	ctx := context.Background()

	// Call the MaterialClient to fetch the material by ID
	result, err := a.MaterialClient.FindMaterialByID(ctx, &materialpb.MaterialID{
		ID: p.ID,
	})
	if err != nil {
		return nil, err // Return error if something goes wrong
	}

	// If the material is found, map the response to the pb.Material type
	pbMaterial := &pb.AdminMaterial{
		Material_Name: result.Material_Name,
		Description:   result.Description,
		Stock:         int32(result.Stock),
		Price:         result.Price,
	}

	// Return the pb.Material object and no error
	return pbMaterial, nil
}

// RemoveMaterialService implements interfaces.AdminServiceInter.
func (a *AdminService) RemoveMaterialService(p *pb.AdminMaterialID) (*pb.AdminResponse, error) {
	ctx := context.Background()

	// Call the MaterialClient's RemoveItem method
	_, err := a.MaterialClient.RemoveMaterial(ctx, &materialpb.MaterialID{ID: p.ID})
	if err != nil {
		return &pb.AdminResponse{
			Status:  pb.AdminResponse_ERROR,
			Message: "Failed to remove material",
			Payload: &pb.AdminResponse_Error{Error: err.Error()},
		}, err
	}

	return &pb.AdminResponse{
		Status:  pb.AdminResponse_OK,
		Message: "Material removed successfully from list",
	}, nil
}
