package result

import (
	"crow-blog-backend/src/config"
	resultType "crow-blog-backend/src/consts/result_type"
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
