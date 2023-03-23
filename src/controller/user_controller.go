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

func (u UserController) GetMyInfo() *result.Result {
	return result.Success("", u.UserService.GetMyInfo())
}
