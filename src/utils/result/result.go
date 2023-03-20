package result

import (
	"crow-blog-backend/src/config"
	resultType "crow-blog-backend/src/consts/result_type"
	globalLogger "crow-blog-backend/src/logger"
	"github.com/kataras/iris/v12"
	"strings"
	"time"
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessWithMessage(message string, data interface{}) *Result {
	return &Result{
		Code:    resultType.Success,
		Message: message,
		Data:    data,
	}
}

func Success(language string, data interface{}) *Result {
	i18nTr := config.GetApp().I18n
	return &Result{
		Code:    resultType.Success,
		Message: i18nTr.Tr(language, resultType.GetKeyByCode(2000)),
		Data:    data,
	}
}

func Failed(message string) *Result {
	return &Result{
		Code:    resultType.Error,
		Message: message,
	}
}

func FailedWithConst(language string, constCode int, args ...string) *Result {
	i18nTr := config.GetApp().I18n
	return &Result{
		Code:    constCode,
		Message: i18nTr.Tr(language, resultType.GetKeyByCode(constCode), args),
	}
}

func Cacheable(language, cacheKey string, fn func() *Result) *Result {
	tmpResult := &Result{}
	if cacheKey != "" {
		// 尝试取缓存
		tmpResult.Code = resultType.Success
		tmpResult.Message = config.GetApp().I18n.Tr(language, resultType.GetKeyByCode(resultType.Success))
		tmpResult.Data = "缓存"
	} else {
		tmpResult = fn()
		// 写入缓存
	}
	return tmpResult
}

func WriteLogResult(ctx iris.Context, fn func() *Result) *Result {
	start := time.Now().UnixMilli()
	requestDateTime := time.UnixMilli(start).Format("2006-01-02 15:04:05")

	fullUri := ctx.FullRequestURI()
	body, err := ctx.GetBody()
	if err != nil {
		globalLogger.Error("获取请求体失败")
	}

	fnResult := fn()
	end := time.Now().UnixMilli()
	responseDateTime := time.UnixMilli(end).Format("2006-01-02 15:04:05")

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

	bodyStr := string(body)
	if bodyStr == "" {
		bodyStr = "null"
	}

	template := "reqIp: %s, reqUrl: %s, reqBody: %s, reqDateTime: %s, respDateTime: %s, costTime: %d"

	if fnResult.Code == resultType.Success {
		// 记录成功日志
		globalLogger.Infof(template, reqIp, fullUri, bodyStr, requestDateTime, responseDateTime, end-start)
	} else if fnResult.Code == resultType.Error {
		// 记录系统级别的错误
		globalLogger.Errorf(template, reqIp, fullUri, bodyStr, requestDateTime, responseDateTime, end-start)
	} else {
		// 记录业务级别的错误
		globalLogger.Warnf(template, reqIp, fullUri, bodyStr, requestDateTime, responseDateTime, end-start)
	}
	return fnResult
}

func WriteLogAndCacheableResult(ctx iris.Context, cacheKey string, fn func() *Result) *Result {
	return WriteLogResult(ctx, func() *Result {
		return Cacheable(ctx.Values().GetString("language"), cacheKey, fn)
	})
}
