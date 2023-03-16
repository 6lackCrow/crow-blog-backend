package main

import (
	"crow-blog-backend/src/config"
	panicUtil "crow-blog-backend/src/utils"
	"github.com/kataras/iris/v12"
)

func main() {
	config.InitConfig()
	irisInstance := iris.New()
	envConfig := config.GetEnvConfig()
	if err := irisInstance.Listen(":" + envConfig.Server.Port); err != nil {
		panicUtil.CustomPanic("Listen on port failed", err)
	}
}
