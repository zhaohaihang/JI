package cron

import (
	"fmt"
	"ji/internal/dao"
)

type Tasks struct {
	userDao     *dao.UserDao
	activityDao *dao.ActivityDao
}

func NewTasks(ud *dao.UserDao, ad *dao.ActivityDao) *Tasks {
	return &Tasks{
		userDao:     ud,
		activityDao: ad,
	}
}

// var TasksProviderSet = wire.NewSet(NewTasks)

func (t *Tasks) PPP() {
	fmt.Println("TTT")
}

func (t *Tasks) PPP2() {

}
