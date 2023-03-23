package middleware

import (
	resultType "crow-blog-backend/src/consts/result_type"
	globalLogger "crow-blog-backend/src/logger"
	"crow-blog-backend/src/utils/result"
	"encoding/json"
	"github.com/kataras/iris/v12"
	"strings"
	"time"
)

func WriteLogMiddleware(ctx iris.Context) {
	start := time.Now().UnixMilli()
	ctx.Values().Set("startTimestamp", start)
	requestDateTime := time.UnixMilli(start).Format("2006-01-02 15:04:05")
	method := ctx.Method()
	fullUri := ctx.FullRequestURI()
	body, err := ctx.GetBody()
	bodyStr := string(body)
	if bodyStr == "" {
		bodyStr = "null"
	}
	if err != nil {
		globalLogger.Error("获取请求体失败")
	}
	ctx.Record()
	ctx.Next()
	end := time.Now().UnixMilli()
	responseDateTime := time.UnixMilli(end).Format("2006-01-02 15:04:05")
	resStr := string(ctx.Recorder().Body())

	reqIp := getIpByContext(ctx)
	resp := &result.Result{}
	_ = json.Unmarshal([]byte(resStr), resp)
	outLog(resp.Code, reqIp, fullUri, method, bodyStr, resStr, requestDateTime, responseDateTime, end-start)

}

func getIpByContext(ctx iris.Context) string {
	xForwardedFor := ctx.GetHeader("X-Forwarded-For")

	reqIp := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if reqIp == "" {
		reqIp = strings.TrimSpace(ctx.GetHeader("X-Real-Ip"))
	}
	if reqIp == "" {
		reqIp = ctx.RemoteAddr()
	}
	if reqIp == "" {
		reqIp = "unknown"
	}
	return reqIp
}

func outLog(code int, reqIp, fullUri, method, bodyStr, resStr, requestDateTime, responseDateTime string, costTime int64) {
	template := "reqIp: %s, reqUrl: %s, method: %s, reqBody: %s, respBody: %s, reqDateTime: %s, respDateTime: %s, costTime: %d"
	if code == resultType.Success {
		// 记录成功日志
		globalLogger.Infof(template, reqIp, fullUri, method, bodyStr, resStr, requestDateTime, responseDateTime, costTime)
	} else if code == resultType.Error {
		// 记录系统级别的错误
		globalLogger.Errorf(template, reqIp, fullUri, method, bodyStr, resStr, requestDateTime, responseDateTime, costTime)
	} else {
		// 记录业务级别的错误
		globalLogger.Warnf(template, reqIp, fullUri, method, bodyStr, resStr, requestDateTime, responseDateTime, costTime)
	}
}
