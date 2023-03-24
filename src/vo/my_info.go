package vo

import "crow-blog-backend/src/entity"

type MyInfo struct {
	AvatarUrl     string           `json:"avatarUrl"`
	Nickname      string           `json:"nickname"`
	Slogan        string           `json:"slogan"`
	ArticleCount  int              `json:"articleCount"`
	CategoryCount int              `json:"categoryCount"`
	TagCount      int              `json:"tagCount"`
	Links         []entity.LinkDTO `json:"links"`
}
