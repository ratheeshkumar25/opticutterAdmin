package interfaces

import "github.com/ratheeshkumar25/opti_cut_admin/pkg/model"

type AdminRepoInter interface {
	FindAdminByEmail(email string) (*model.Admin, error)
}
