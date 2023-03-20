package repository

type TestRepository struct {
}

func NewTestRepository() *TestRepository {
	return &TestRepository{}
}

func (t TestRepository) GetTestArr() []string {
	return []string{
		"go ",
		"is ",
		"best! ",
	}
}
