package manage

import (
	"errors"
	"university/global"
	"university/models/manage"

	"gorm.io/gorm"
)

type ManageAdminUserService struct {
}

func (m *ManageAdminUserService) CreateManageAdminUser(manageAdminUser manage.ManageAdminUser) (err error) {
	if !errors.Is(global.GVA_DB.Where("login_user_name = ?", manageAdminUser.LoginUserName).First(&manage.ManageAdminUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同用户名")
	}
	err = global.GVA_DB.Create(&manageAdminUser).Error
	return err
}
