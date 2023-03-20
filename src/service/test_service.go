package service

import "crow-blog-backend/src/repository"

type TestService struct {
	TestRepository *repository.TestRepository
}

func NewTestService() *TestService {
	return &TestService{
		TestRepository: repository.NewTestRepository(),
	}
}

func (t TestService) GetTestArr() []string {
	return t.TestRepository.GetTestArr()
}
