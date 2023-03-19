package result

import resultType "crow-blog-backend/src/consts/result_type"

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(message string, data interface{}) *Result {
	return &Result{
		Code:    resultType.Success,
		Message: message,
		Data:    data,
	}
}

func Failed(message string) *Result {
	return &Result{
		Code:    resultType.Error,
		Message: message,
	}
}
