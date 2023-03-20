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

func (p *TestController) Get() *result.Result {
	resultFn := func() *result.Result {
		return result.Success(p.Ctx.Values().GetString("language"), p.TestService.GetTestArr())
	}
	return result.WriteLogAndCacheableResult(p.Ctx, "", resultFn)
}
