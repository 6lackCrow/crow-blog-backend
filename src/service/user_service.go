package service

import (
	"crow-blog-backend/src/repository"
	"crow-blog-backend/src/utils/result"
	"crow-blog-backend/src/vo"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		UserRepository: repository.NewUserRepository(),
	}
}

func (s UserService) GetMyInfo() vo.MyInfo {
	panic("test")
	user, err := s.UserRepository.GetMyInfo()
	if err != nil {
		panic(result.Failed(err.Error()))
	}
	myInfo := vo.MyInfo{
		Avatar:        user.AvatarUrl,
		Nickname:      user.Nickname,
		Slogan:        user.Slogan,
		ArticleCount:  user.ArticleCount,
		CategoryCount: user.CategoryCount,
		TagCount:      user.TagCount,
	}
	links, err := s.UserRepository.GetLinks(user.ID)
	if err != nil {
		panic(result.Failed(err.Error()))
	}
	myInfo.Links = links
	return myInfo
}
