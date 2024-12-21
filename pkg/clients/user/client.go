package user

import (
	"fmt"
	"log"

	"github.com/ratheeshkumar25/opti_cut_admin/config"
	pb "github.com/ratheeshkumar25/opti_cut_admin/pkg/clients/user/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.UserServiceClient, error) {
	grpcAddr := fmt.Sprintf("user-service:%s", cfg.GrpcUserPort)
	grpc, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc user client: %s, ", cfg.GrpcUserPort)
		return nil, err
	}
	log.Printf("Succesfully connected to user client at port: %v", cfg.GrpcUserPort)
	return pb.NewUserServiceClient(grpc), nil
}
