package admin_user_service

import (
	"apis/common"
	"apis/common/jwt"
	"apis/common/redis"
	"apis/src/models/admin_user"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Logout(uid int) (err error) {
	err = redis.GetRedis().Del(redis.GetAdminLoginKey(uid)).Err()
	return
}

func Login(account, pwd string) (ret admin_user.AdminUser, token string, err error) {
	err = common.GetGorm().Where("account = ?", account).First(&ret).Error
	if err != nil {
		return
	}
	if ret.ID == 0 {
		err = fmt.Errorf("账号或密码错误")
		return
	}
	// 密码验证
	err = bcrypt.CompareHashAndPassword([]byte(ret.Pwd), []byte(pwd)) //验证（对比）
	if err != nil {
		return
	}
	token, err = jwt.CreateToken(ret.ID, 604800, "good luck")
	if err != nil {
		return
	}
	err = redis.GetRedis().Set(redis.GetAdminLoginKey(ret.ID), token, time.Second*600000).Err()
	if err != nil {
		return
	}
	ret.Pwd = ""
	return
}
