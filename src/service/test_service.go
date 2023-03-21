package service

import (
	"crow-blog-backend/src/cache"
	"crow-blog-backend/src/consts/cache_opt"
	"crow-blog-backend/src/repository"
	"time"
)

type TestService struct {
	TestRepository *repository.TestRepository
}

func NewTestService() *TestService {
	return &TestService{
		TestRepository: repository.NewTestRepository(),
	}
}

func (t TestService) GetTestArr(key string) []string {
	return cache.Cacheable[[]string](key, cache_opt.Select, 20*time.Minute, func() []string {
		return t.TestRepository.GetTestArr()
	})
}

func (t TestService) GetTestStr(key string) string {
	return cache.Cacheable[string](key, cache_opt.Select, 20*time.Minute, func() string {
		return t.TestRepository.GetTestStr()
	})
}

func (t TestService) GetTestNum(key string) int {
	return cache.Cacheable[int](key, cache_opt.Select, 20*time.Minute, func() int {
		return t.TestRepository.GetTestNum()
	})
}
