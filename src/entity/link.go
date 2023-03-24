package entity

import (
	"time"
)

type LinkT interface {
	Link | LinkDTO
}

type Link struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserId      uint
	Url         string `gorm:"type:varchar(2000)"`
	Icon        string `gorm:"type:varchar(200)"`
	Description string `gorm:"type:varchar(200)"`
}

type LinkDTO struct {
	ID          uint   `json:"id"`
	UserId      uint   `json:"userId"`
	Url         string `json:"url"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
}
