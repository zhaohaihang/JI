package model

import (
	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	UserId     uint
	AcitivtyId uint
	Liked      int8
}
