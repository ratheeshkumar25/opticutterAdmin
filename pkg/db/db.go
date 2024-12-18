package db

import (
	"fmt"
	"log"

	"github.com/ratheeshkumar25/opti_cut_admin/config"
	"github.com/ratheeshkumar25/opti_cut_admin/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config *config.Config) *gorm.DB {
	host := config.Host
	user := config.User
	password := config.Password
	dbname := config.Database
	port := config.Port
	sslmode := config.Sslmode

	// Print each configuration for debugging
	fmt.Printf("Connecting to DB: host=%s, user=%s, password=%s, dbname=%s, port=%s, sslmode=%s\n", host, user, password, dbname, port, sslmode)

	// Construct the DSN correctly
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)
	fmt.Println("DSN:", dsn) // Debugging line

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("connection to the database failed:", err)
	}

	err = DB.AutoMigrate(
		&model.Admin{Email: config.AdminEmail, Password: config.AdminPassword},
	)

	if err != nil {
		fmt.Printf("error while migrating %v", err.Error())
		return nil
	}
	return DB
}
