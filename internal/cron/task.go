package cron

import (
	"ji/internal/dao"
	"time"

	"github.com/sirupsen/logrus"
)

type Tasks struct {
	logger      *logrus.Logger
	userDao     *dao.UserDao
	activityDao *dao.ActivityDao
}

func NewTasks(l *logrus.Logger, ud *dao.UserDao, ad *dao.ActivityDao) *Tasks {
	return &Tasks{
		logger:      l,
		userDao:     ud,
		activityDao: ad,
	}
}

func (t *Tasks) UpdateActivityStatusFromNostartToInprocess() {
	time := time.Now().UnixMilli()
	activitys, err := t.activityDao.UpdateActivityStatusFromNostartToInprocess(time)
	if err != nil {
		t.logger.Info(err)
	}

	for _, activity := range activitys {
		//SendMsgToParticipant()
		t.logger.Info(activity.ID)
	}
	return
}

func (t *Tasks) UpdateActivityStatusFromInprocessToEnd() {
	time := time.Now().UnixMilli()
	err := t.activityDao.UpdateActivityStatusFromInprocessToEnd(time)
	if err != nil {
		t.logger.Info(err)
	}
	return

}
