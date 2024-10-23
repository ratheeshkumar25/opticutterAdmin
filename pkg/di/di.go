package di

import (
	"log"

	"github.com/ratheeshkumar25/opti_cut_admin/config"
	"github.com/ratheeshkumar25/opti_cut_admin/pkg/clients/material"
	"github.com/ratheeshkumar25/opti_cut_admin/pkg/clients/user"
	"github.com/ratheeshkumar25/opti_cut_admin/pkg/db"
	"github.com/ratheeshkumar25/opti_cut_admin/pkg/handlers"
	"github.com/ratheeshkumar25/opti_cut_admin/pkg/repo"
	"github.com/ratheeshkumar25/opti_cut_admin/pkg/server"
	"github.com/ratheeshkumar25/opti_cut_admin/pkg/services"
)

func Init() {
	cnfg := config.LoadConfig()

	db := db.ConnectDB(cnfg)

	userClient, err := user.ClientDial(*cnfg)
	if err != nil {
		log.Fatalf("failed to connect to user client")
	}

	materialClient, err := material.ClientDial(*cnfg)
	if err != nil {
		log.Fatalf("failed to connect to material client")
	}
	adminRepo := repo.NewAdminRepository(db)

	adminService := services.NewAdminRepository(adminRepo, userClient, materialClient)
	adminHandler := handlers.NewAdminHandler(adminService)

	err = server.NewGrpcAdminServer(cnfg.AdminPort, adminHandler)
	if err != nil {
		log.Fatalf("failed to start gRPC server %v", err.Error())
	}
}
