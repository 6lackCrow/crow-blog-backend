package controller

import (
	"crow-blog-backend/src/service"
	"crow-blog-backend/src/utils/result"
	"github.com/kataras/iris/v12"
)

type UserController struct {
	Ctx         iris.Context
	UserService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		UserService: service.NewUserService(),
	}
}

func GetMyInfo() *result.Result {
	return new(result.Result)
}
