package model

import (
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model

	Title          string
	Introduction   string `gorm:"size:1000"`
	Status         int
	StartTime      int64
	EndTime        int64
	Location       Point `gorm:"type:point"`
	ExpectedNumber uint
	CurrentNumber  uint
	UserId         uint
	UserName       string
	UserAvatar     string
}
