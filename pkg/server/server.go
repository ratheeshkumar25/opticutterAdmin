package server

import (
	"fmt"
	"log"
	"net"

	"github.com/ratheeshkumar25/opti_cut_admin/pkg/handlers"
	pb "github.com/ratheeshkumar25/opti_cut_admin/pkg/pb"
	"google.golang.org/grpc"
)

func NewGrpcAdminServer(port string, handlr *handlers.AdminHandler) error {
	log.Println("connecting to gRPC server")
	addr := fmt.Sprintf(":%s", port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("error creating listener on %v", addr)
		return err
	}

	grpc := grpc.NewServer()
	pb.RegisterAdminServiceServer(grpc, handlr)

	log.Printf("listening on gRPC server %v", port)
	err = grpc.Serve(lis)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}
	return nil
}
