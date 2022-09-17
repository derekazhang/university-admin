package redis

import (
	"apis/common"
	"fmt"
	"sync"

	"github.com/go-redis/redis"
)

var (
	redisOnce   sync.Once
	redisClient *redis.Client
)

//type Redis struct {
//	Conf *Config
//}

func GetRedis() *redis.Client {
	redisOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", common.GetConfig().Redis.Host, common.GetConfig().Redis.Port),
			Password: common.GetConfig().Redis.Password,
			DB:       common.GetConfig().Redis.Db,
		})
		err := redisClient.Ping().Err()
		if err != nil {
			panic(err)
		}
	})
	return redisClient
}

func GetWxSessionKey(openid string) string {
	return fmt.Sprintf("session_%s", openid)
}

func GetTtSessionKey(openid string) string {
	return fmt.Sprintf("tt_session_#{openid}")
}

func GetTodayPatternKey(de string) string {
	return fmt.Sprintf("today_%s", de)
}

func GetWxAccessToken() string {
	return fmt.Sprintf("accesstoken")
}

func GetModelConfKey(id string) string {
	return fmt.Sprintf("mode_%s", id)
}

func GetWxPayKey(orderNum string) string {
	return fmt.Sprintf("pay_%s", orderNum)
}

func GetAdminLoginKey(uid int) string {
	return fmt.Sprintf("admin_login_%d", uid)
}

func GetCosSignKey() string {
	return fmt.Sprintf("cos:sign")
}

//func (s *Redis) Publish(name string, data interface{}) (err error) {
//	body, e := json.Marshal(data)
//	if e != nil {
//		err = e
//		return
//	}
//	client := s.Get()
//	err = client.LPush(name, body).Err()
//	return
//}
//
//func (s *Redis) Consume(name string, cb func(bs []byte)) (err error) {
//	client := s.Get()
//	for {
//		data, err := client.BRPop(time.Second*3, name).Result()
//		if err != nil {
//			continue
//		}
//		cb([]byte(data[1]))
//	}
//}
