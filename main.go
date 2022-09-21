package main

import (
	"university/core"
	"university/global"
	"university/initialize"
)

func main() {

	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()
	global.GVA_DB = initialize.Gorm()

	core.RunWindowsServer()
}
