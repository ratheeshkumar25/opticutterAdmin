package repo

import (
	"github.com/ratheeshkumar25/opti_cut_admin/pkg/repo/interfaces"
	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepoInter {
	return &AdminRepository{
		DB: db,
	}
}
