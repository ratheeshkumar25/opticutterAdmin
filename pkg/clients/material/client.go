package material

import (
	"fmt"
	"log"

	"github.com/ratheeshkumar25/opti_cut_admin/config"
	pb "github.com/ratheeshkumar25/opti_cut_admin/pkg/clients/material/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.MaterialServiceClient, error) {
	grpcAddr := fmt.Sprintf("localhost:%s", cfg.MaterialPort)
	grpc, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc user client: %s, ", cfg.MaterialPort)
		return nil, err
	}
	log.Printf("Succesfully connected to user client at port: %v", cfg.MaterialPort)
	return pb.NewMaterialServiceClient(grpc), nil
}
