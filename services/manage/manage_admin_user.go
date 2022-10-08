package manage

import (
	"errors"
	"strings"
	"time"
	"university/global"
	"university/models/manage"
	"university/models/manage/request"
	"university/utils"

	"strconv"

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

func (m *ManageAdminUserService) AdminLogin(params request.ManageAdminLoginParam) (manageAdminUser manage.ManageAdminUser, adminToken manage.ManageAdminUserToken, err error) {
	err = global.GVA_DB.Where("login_user_name=? AND login_password=?", params.UserName, params.PasswordMd5).First(&manageAdminUser).Error
	if manageAdminUser != (manage.ManageAdminUser{}) {
		token := getNewToken(time.Now().UnixNano()/1e6, int(manageAdminUser.AdminUserId))
		global.GVA_DB.Where("admin_user_id", manageAdminUser.AdminUserId).First(&adminToken)
		nowDate := time.Now()
		// 48小时过期
		expireTime, _ := time.ParseDuration("48h")
		expireDate := nowDate.Add(expireTime)
		// 没有token新增，有token则更新你
		if adminToken == (manage.ManageAdminUserToken{}) {
			adminToken.AdminUserId = manageAdminUser.AdminUserId
			adminToken.Token = token
			adminToken.UpdateTime = nowDate
			adminToken.ExpireTime = expireDate
			if err = global.GVA_DB.Create(&adminToken).Error; err != nil {
				return

			}
		} else {
			adminToken.Token = token
			adminToken.UpdateTime = nowDate
			adminToken.ExpireTime = expireDate
			if err = global.GVA_DB.Save(&adminToken).Error; err != nil {
				return
			}
		}

	}
	return manageAdminUser, adminToken, err
}

func getNewToken(timeInt int64, userId int) (token string) {
	var build strings.Builder
	build.WriteString(strconv.FormatInt(timeInt, 10))
	build.WriteString(strconv.Itoa(userId))
	build.WriteString(utils.GenValidateCode(6))
	return utils.MD5V([]byte(build.String()))
}
