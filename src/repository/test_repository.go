package repository

import "fmt"

type TestRepository struct {
}

func NewTestRepository() *TestRepository {
	return &TestRepository{}
}

func (t TestRepository) GetTestArr() []string {
	fmt.Println("直接执行的repository")
	return []string{
		"非",
		"缓",
		"存",
	}
}

func (t TestRepository) GetTestStr() string {
	fmt.Println("直接执行的repository")
	return "非缓存"
}

func (t TestRepository) GetTestNum() int {
	return 5
}
