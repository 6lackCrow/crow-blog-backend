package route

import (
	"crow-blog-backend/src/controller"
	"crow-blog-backend/src/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func InitRoute(app *iris.Application) {
	app.UseGlobal(middleware.CorsMiddleware,
		middleware.I18nMiddleware,
		middleware.ExceptionMiddleware,
		middleware.WriteLogMiddleware)
	baseUrl := "/api/v1"

	mvc.New(app.Party(baseUrl + "/test")).Handle(controller.NewTestController())

	mvc.New(app.Party(baseUrl + "/user")).Handle(controller.NewUserController())
}
