package cron

import (
	"fmt"
	"ji/internal/dao"

	"github.com/sirupsen/logrus"
)

type Tasks struct {
	logger       *logrus.Logger
	userDao     *dao.UserDao
	activityDao *dao.ActivityDao
}

func NewTasks(l *logrus.Logger,ud *dao.UserDao, ad *dao.ActivityDao) *Tasks {
	return &Tasks{
		logger: l,
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
