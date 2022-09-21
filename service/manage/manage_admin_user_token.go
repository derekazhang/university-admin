package manage

import (
	"university/global"
	"university/models/manage"
)

type ManageAdminUserTokenService struct {
}

func (m *ManageAdminUserTokenService) ExistAdminToken(token string) (err error, manageAdminUserToken manage.ManageAdminUserToken) {
	err = global.GVA_DB.Where("token =?", token).First(&manageAdminUserToken).Error
	return
}

func (m *ManageAdminUserTokenService) DeleteManageAdminUserToken(token string) (err error) {
	err = global.GVA_DB.Delete(&[]manage.ManageAdminUserToken{}, "token =?", token).Error
	return err
}
