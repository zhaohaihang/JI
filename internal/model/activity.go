package model

import (
	"ji/pkg/utils/datetime"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model

	Title          string
	Introduction   string `gorm:"size:1000"`
	Status         int
	StartTime      datetime.DateTime
	EndTime        datetime.DateTime
	Location       Point `gorm:"type:point"`
	ExpectedNumber uint
	CurrentNumber  uint
	UserId         uint
	UserName       string
	UserAvatar     string
}
