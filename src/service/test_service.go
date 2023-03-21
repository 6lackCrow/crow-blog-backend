package service

import (
	"crow-blog-backend/src/cache"
	"crow-blog-backend/src/consts/cache_opt"
	"crow-blog-backend/src/repository"
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
	return cache.Cacheable("asdada", cache_opt.Select, func() interface{} {
		return t.TestRepository.GetTestArr()
	}).([]string)
}
