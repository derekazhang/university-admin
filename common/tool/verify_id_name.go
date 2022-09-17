package tool

import (
	"fmt"
	"time"
)

type verifyIdNameRes struct {
	State int `json:"state"`
}

func VerifyIdName(idNum, name string) bool {
	res := verifyIdNameRes{}
	reqUrl := "https://dfidveri.market.alicloudapi.com/verify_id_name"
	ret, _, errs := Request.Post(reqUrl).Timeout(time.Second*5).AppendHeader("Authorization", "APPCODE e83144fa046e42fbaa597b4071fad822").AppendHeader("content-type", "application/x-www-form-urlencoded; charset=utf-8").Send(fmt.Sprintf("id_number=%s&name=%s", idNum, name)).EndStruct(&res)
	if len(errs) != 0 {
		return false
	}
	if ret.StatusCode != 200 {
		return false
	}
	if res.State != 1 {
		return false
	}
	return true
}
