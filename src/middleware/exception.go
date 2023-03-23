package middleware

import (
	globalLogger "crow-blog-backend/src/logger"
	"crow-blog-backend/src/utils/result"
	"github.com/kataras/iris/v12"
	"reflect"
)

func ExceptionMiddleware(ctx iris.Context) {
	defer func() {
		err := recover()
		if err == nil {
			ctx.Next()
			return
		}
		r := &result.Result{}

		if reflect.TypeOf(err) == reflect.TypeOf(r) {
			r = err.(*result.Result)
		} else {
			r.Message = err.(string)
		}
		if rsErr := ctx.JSON(r); rsErr != nil {
			globalLogger.Error("拦截到异常，返回json出错")
		}
		ctx.Done()
	}()
	ctx.Next()
}
