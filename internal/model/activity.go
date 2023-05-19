package model

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model

	Title          string
	Introduction   string `gorm:"size:1000"`
	Status         int
	StartTime      time.Time
	EndTime        time.Time
	Location       Point `gorm:"type:point"`
	ExpectedNumber uint
	CurrentNumber  uint
	UserId         uint
	UserName       string
	UserAvatar     string
}
