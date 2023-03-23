package middleware

import (
	globalLogger "crow-blog-backend/src/logger"
	"crow-blog-backend/src/utils/result"
	"encoding/json"
	"github.com/kataras/iris/v12"
	"reflect"
	"time"
)

func ExceptionMiddleware(ctx iris.Context) {

	defer func() {
		rec := recover()
		if rec == nil {
			ctx.Next()
			return
		}
		r := &result.Result{}
		start, _ := ctx.Values().GetInt64("startTimestamp")
		requestDateTime := time.UnixMilli(start).Format("2006-01-02 15:04:05")
		method := ctx.Method()
		fullUri := ctx.FullRequestURI()
		body, _ := ctx.GetBody()
		bodyStr := string(body)
		if bodyStr == "" {
			bodyStr = "null"
		}
		reqIp := getIpByContext(ctx)
		if reflect.TypeOf(rec) == reflect.TypeOf(r) {
			r = rec.(*result.Result)
		} else {
			r = result.Failed(rec.(string))
		}
		if rsErr := ctx.JSON(r); rsErr != nil {
			globalLogger.Error("拦截到异常，返回json出错")
		}
		end := time.Now().UnixMilli()
		responseDateTime := time.UnixMilli(end).Format("2006-01-02 15:04:05")
		bytes, _ := json.Marshal(r)
		outLog(r.Code, reqIp, fullUri, method, bodyStr, string(bytes), requestDateTime, responseDateTime, end-start)
		ctx.Done()
	}()
	ctx.Next()
}
