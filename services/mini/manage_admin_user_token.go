package mini

import (
	"university/global"
	"university/models/mini"
)

type MiniUserTokenService struct {
}

func (m *MiniUserTokenService) ExistUserToken(token string) (err error, miniUserToken mini.MiniUserToken) {
	err = global.GVA_DB.Where("token =?", token).First(&miniUserToken).Error
	return
}

func (m *MiniUserTokenService) DeleteMiniUserToken(token string) (err error) {
	err = global.GVA_DB.Delete(&[]mini.MiniUserToken{}, "token =?", token).Error
	return err
}
