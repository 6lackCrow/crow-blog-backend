package result_type

const (
	Success = 2000
	Error   = 4999 + iota
)

var resultMap = map[int]string{
	Success: "result.success",
	Error:   "result.error.base",
}

func GetKeyByCode(code int) string {
	return resultMap[code]
}
