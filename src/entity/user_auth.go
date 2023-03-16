package entity

import "time"

type UserAuth struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserId     uint
	Type       int8
	Identifier string `gorm:"type:varchar(200)"` // 权限标识
	Credential string `gorm:"type:varchar(200)"` // 密码凭证
}
