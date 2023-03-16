package entity

import (
	"time"
)

type User struct {
	ID            uint `gorm:"primarykey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Nickname      string `gorm:"type:varchar(50)"`
	AvatarUrl     string `gorm:"type:varchar(2000)"`
	Slogan        string `gorm:"type:varchar(200)"`
	ArticleCount  int64
	CategoryCount int64
	TagCount      int64
}
