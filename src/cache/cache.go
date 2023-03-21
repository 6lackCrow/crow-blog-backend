package cache

import (
	"crow-blog-backend/src/cache/redis_cache"
	"crow-blog-backend/src/config"
	"crow-blog-backend/src/consts/cache_opt"
	globalLogger "crow-blog-backend/src/logger"
	"github.com/redis/go-redis/v9"
	"time"
)

var cacheLockChan = make(chan bool)

func getLock(ch chan bool) bool {
	return <-ch
}

func setLock(ch chan bool, lock bool) {
	ch <- lock
}

// Cacheable 需要解决的问题: 1.缓存击穿 2.缓存穿透 3.缓存雪崩
func Cacheable[T any](cacheKey string, cacheOpt int, expireTime time.Duration, fn func() T) T {
	if !config.GetEnvConfig().Cache.Use {
		return fn()
	}
	var tmp T
	switch cacheOpt {
	case cache_opt.Select:
		err := redis_cache.GetScan(cacheKey, &tmp)

		switch {
		case err == redis.Nil:
			// key不存在 添加缓存 需要处理缓存击穿、缓存雪崩问题
			// 缓存雪崩
			//定义：同一时间大量的缓存失效导致数据库压力升高。
			//解决：随机过期时间保证同一时间不会有大量的key失效。
			//缓存击穿
			//定义：热点key失效的瞬间大量读请求落到数据库上。
			//解决： 加锁让少量的请求来重构缓存，例如分布式锁或者读写锁、本地锁等等，粒度根据实际情况选择。 热点key不过期，只有写请求会重构缓存。 做缓存时间的续期。
			lockStr := "----lock"
			nx, nxErr := redis_cache.SetNX(cacheKey+lockStr, lockStr, 1*time.Minute)
			if nxErr != nil {
				globalLogger.Errorf("设置锁%s失败: %s", cacheKey+lockStr, nxErr.Error())
				return fn()
			}
			if nx {
				// 加锁成功 处理缓存
				go setLock(cacheLockChan, true)
				tmp = fn()
				if setErr := redis_cache.Set(cacheKey, tmp, expireTime); setErr != nil {
					globalLogger.Errorf("缓存插入失败: %s", setErr.Error())
				}
				// 删除锁
				if rmLockErr := redis_cache.Remove(cacheKey + lockStr); rmLockErr != nil {
					globalLogger.Errorf("删除锁失败: %s", rmLockErr.Error())
				}
				go setLock(cacheLockChan, false)
				return tmp
			} else {
				// 等待锁释放
				for {
					if !getLock(cacheLockChan) {
						_ = redis_cache.GetScan(cacheKey, &tmp)
						break
					}
				}

			}

		case err != nil:
			globalLogger.Errorf("获取缓存出错: %s", err.Error())
			return fn()
			//case err == nil && str == "":
			//	//返回空缓存 防止缓存穿透
			//	//缓存穿透
			//	//定义：访问一个数据库中不存在的值，由于每次在缓存中都查询不到会穿透到数据库中进行查询。
			//	//解决：缓存空值（最简单），或者引入布隆过滤器。接口基本的参数校验，比如过滤负值等等。
			//	return tmp
		}
	case cache_opt.Remove:
		//一般用于修改 更新 删除方法 - 删除对应的缓存
		tmp = fn()
	default:
		tmp = fn() // 不操作缓存,直接执行相应逻辑
	}
	return tmp
}
