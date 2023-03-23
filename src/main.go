package main

import (
	"crow-blog-backend/src/config"
	globalLogger "crow-blog-backend/src/logger"
	"crow-blog-backend/src/route"
	panicUtil "crow-blog-backend/src/utils/painc"
	"github.com/kataras/iris/v12"
)

func main() {
	config.InitConfig()
	irisInstance := iris.New()

	if err := irisInstance.I18n.Load("./locales/*/*", "en-US", "zh-CN"); err != nil {
		globalLogger.Error("Failed load i18n")
	}
	irisInstance.I18n.SetDefault("zh-CN")
	config.SetApp(irisInstance)
	envConfig := config.GetEnvConfig()
	route.InitRoute(irisInstance)
	if err := irisInstance.Listen(":" + envConfig.Server.Port); err != nil {
		panicUtil.CustomPanic("Listen on port failed", err)
	}
}
