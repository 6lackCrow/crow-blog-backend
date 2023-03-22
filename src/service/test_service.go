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

func (t TestService) GetTestArr() []string {
	return cache.Cacheable[[]string]("arr", cache_opt.Select, 20*time.Minute, func() []string {
		return t.TestRepository.GetTestArr()
	})
}

func (t TestService) GetTestStr() string {
	return cache.Cacheable[string]("str", cache_opt.Select, 20*time.Minute, func() string {
		return t.TestRepository.GetTestStr()
	})
}

func (t TestService) GetTestNum() int {
	return cache.Cacheable[int]("num", cache_opt.Select, 20*time.Minute, func() int {
		return t.TestRepository.GetTestNum()
	})
}
