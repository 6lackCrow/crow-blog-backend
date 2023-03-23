package controller

import (
	"crow-blog-backend/src/service"
	"crow-blog-backend/src/utils/result"
	"github.com/kataras/iris/v12"
)

type TestController struct {
	Ctx         iris.Context
	TestService *service.TestService
}

func NewTestController() *TestController {
	return &TestController{
		TestService: service.NewTestService(),
	}
}

func (p *TestController) GetArr() *result.Result {
	return result.Success(p.Ctx.Values().GetString("language"), p.TestService.GetTestArr())
}

func (p *TestController) GetStr() *result.Result {
	return result.Success(p.Ctx.Values().GetString("language"), p.TestService.GetTestStr())
}

func (p *TestController) GetNum() *result.Result {
	return result.Success(p.Ctx.Values().GetString("language"), p.TestService.GetTestNum())
}
