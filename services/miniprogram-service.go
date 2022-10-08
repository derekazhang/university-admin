package services

import (
	"log"
	"university/global"

	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram"
)

var MiniProgramApp *miniProgram.MiniProgram

const TIMEZONE = "asia/shanghai"
const DATETIME_FORMAT = "20060102"

func NewMiniMiniProgramService() (*miniProgram.MiniProgram, error) {
	log.Printf("miniprogram app_id: %s", global.CONFIG.MiniProgram.AppID)
	var cache kernel.CacheInterface
	if global.CONFIG.MiniProgram.RedisAddr != "" {
		cache = kernel.NewRedisClient(&kernel.RedisOptions{
			Addr: global.CONFIG.MiniProgram.RedisAddr,
		})
	}
	app, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:        global.CONFIG.MiniProgram.AppID,
		Secret:       global.CONFIG.MiniProgram.Secret,
		ResponseType: response.TYPE_MAP,
		Log: miniProgram.Log{
			Level: "debug",
			File:  "./wechat.log",
		},
		Cache:     cache,
		HttpDebug: true,
		Debug:     false,
	})

	return app, err
}
