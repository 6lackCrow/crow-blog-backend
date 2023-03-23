package repository

import (
	"crow-blog-backend/src/config"
	"crow-blog-backend/src/entity"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r UserRepository) GetAdminUser() (*entity.User, error) {
	db := config.GetDatabaseInstance()
	user := &entity.User{}
	err := db.First(user).Error
	return user, err
}

func (r UserRepository) GetLinks(userId uint) ([]entity.Link, error) {
	var links []entity.Link
	db := config.GetDatabaseInstance()
	err := db.Where("user_id = ?", userId).Find(&links).Error
	if err != nil {
		return links, err
	}
	return links, nil
}
