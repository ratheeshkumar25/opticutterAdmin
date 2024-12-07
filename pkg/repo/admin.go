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

// CreateAdmin implements interfaces.AdminRepoInter.
func (a *AdminRepository) CreateAdmin(admin *model.Admin) (uint, error) {
	if err := a.DB.Create(&admin).Error; err != nil {
		return 0, err
	}
	return admin.ID, nil
}

func (a *AdminRepository) FindAdminByID(Admin_ID uint) (*model.Admin, error) {
	var admin model.Admin

	if err := a.DB.First(&admin, Admin_ID).Error; err != nil {
		return nil, err
	}
	return &admin, nil

}
