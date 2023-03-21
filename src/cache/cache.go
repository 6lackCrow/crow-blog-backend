package cache

import (
	"crow-blog-backend/src/consts/cache_opt"
)

func Cacheable(cacheKey string, cacheOpt int, fn func() interface{}) interface{} {
	var tmp interface{}
	switch cacheOpt {
	case cache_opt.Select:
		if cacheKey != "" {
			// 尝试取缓存
			tmp = []string{
				"asdas",
				"qweqwas",
			}
		} else {
			tmp = fn()
			// 写入缓存
		}
	case cache_opt.Remove:
		//一般用于修改 更新 删除方法 - 删除对应的缓存
		tmp = fn()
	default:
		tmp = fn() // 不操作缓存,直接执行相应逻辑
	}
	return tmp
}
