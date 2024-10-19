package repo

import "github.com/ratheeshkumar25/opti_cut_admin/pkg/model"

// FindAdminByEmail implements interfaces.AdminRepoInter.
func (a *AdminRepository) FindAdminByEmail(email string) (*model.Admin, error) {
	var admin model.Admin
	if err := a.DB.Model(&model.Admin{}).Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
