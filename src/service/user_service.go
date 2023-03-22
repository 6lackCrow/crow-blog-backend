package service

import "crow-blog-backend/src/vo"

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s UserService) GetMyInfo() vo.MyInfo {
	return vo.MyInfo{}
}
