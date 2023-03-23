package service

import (
	"crow-blog-backend/src/cache"
	"crow-blog-backend/src/consts/cache_opt"
	"crow-blog-backend/src/repository"
	"crow-blog-backend/src/vo"
	"time"
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
	return cache.Cacheable[vo.MyInfo]("myInfo-cacheKey", cache_opt.Select, 5*time.Minute, func() vo.MyInfo {
		user, err := s.UserRepository.GetAdminUser()
		if err != nil {
			cache.UnLock("myInfo-cacheKey")
			panic("查询数据库失败")
		}
		info := vo.MyInfo{
			Avatar:        user.AvatarUrl,
			Nickname:      user.Nickname,
			Slogan:        user.Slogan,
			ArticleCount:  user.ArticleCount,
			CategoryCount: user.CategoryCount,
			TagCount:      user.TagCount,
		}
		links, err := s.UserRepository.GetLinks(user.ID)
		if err != nil {
			cache.UnLock("myInfo-cacheKey")
			panic("查询数据库失败")
		}
		info.Links = links
		return info
	})
}
