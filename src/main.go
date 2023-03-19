package main

import (
	"crow-blog-backend/src/config"
	globalLogger "crow-blog-backend/src/logger"
	panicUtil "crow-blog-backend/src/utils/painc"
	"fmt"
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

	fmt.Println(config.GetApp().I18n.Tr("en-US", "result.error.base"))

	envConfig := config.GetEnvConfig()
	if err := irisInstance.Listen(":" + envConfig.Server.Port); err != nil {
		panicUtil.CustomPanic("Listen on port failed", err)
	}
}
