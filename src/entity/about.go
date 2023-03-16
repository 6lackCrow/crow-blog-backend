package entity

import "time"

type About struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserId    string
	Text      string `gorm:"type:longtext"`
}
