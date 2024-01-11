package model

import (
	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	UserId    uint
	FollowId uint
	Followed  int8
}
