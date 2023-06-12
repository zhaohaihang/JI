package model

import "gorm.io/gorm"

type Engage struct {
	gorm.Model

	UserId     uint
	ActivityId uint
}
