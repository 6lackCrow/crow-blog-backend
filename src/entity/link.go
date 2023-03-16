package entity

import (
	"time"
)

type Link struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Url         string `gorm:"type:varchar(2000)"`
	Icon        string `gorm:"type:varchar(200)"`
	Description string `gorm:"type:varchar(200)"`
}
