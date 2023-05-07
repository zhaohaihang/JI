package model

import (
	"time"

	geo "github.com/kellydunn/golang-geo"
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model

	Title        string
	Introduction string `gorm:"size:1000"`
	Status       int
	
	User 		 User   `gorm:"ForeignKey:UserID"`
	UserId       uint 

	StartTime time.Time
	EndTime   time.Time

	Latitude float64
	Lngitude float64
	Location *geo.Point `gorm:"type:point"`

	ExpectedNumber uint
	CurrentNumber  uint
}

func (activity *Activity) BeforeSave(db *gorm.DB) error {
	activity.Location = geo.NewPoint(activity.Latitude, activity.Lngitude)
	return nil
}

func (activity *Activity) AfterFind(db *gorm.DB) error {
	activity.Latitude = activity.Location.Lat()
	activity.Lngitude = activity.Location.Lng()
	return nil
}
