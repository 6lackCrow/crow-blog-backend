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

func (r UserRepository) GetLinkDTOs(userId uint) ([]entity.LinkDTO, error) {
	var links []entity.LinkDTO
	db := config.GetDatabaseInstance()
	err := db.Model(&entity.Link{}).Where("user_id = ?", userId).Find(&links).Error
	if err != nil {
		return links, err
	}
	return links, nil
}
