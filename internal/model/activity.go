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
	Location       Point `gorm:"type:point;index:idx_activity_location;not null"`
	ExpectedNumber uint
	CurrentNumber  uint
	BgImage        string
	UserId         uint
	UserName       string
	UserAvatar     string
}
