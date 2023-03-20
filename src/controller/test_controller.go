package controller

import (
	"crow-blog-backend/src/service"
	"crow-blog-backend/src/utils/result"
	"fmt"
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
	fmt.Println("进入响应")

	return result.Success("", p.TestService.GetTestArr())
}
