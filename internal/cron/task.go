package cron

import (
	"ji/internal/dao"
	"ji/internal/model"
	"net/textproto"
	"time"

	"github.com/jordan-wright/email"
	"github.com/sirupsen/logrus"
)

type Tasks struct {
	logger      *logrus.Logger
	userDao     *dao.UserDao
	activityDao *dao.ActivityDao
	mailPool    *email.Pool
}

func NewTasks(l *logrus.Logger, ud *dao.UserDao, ad *dao.ActivityDao, mp *email.Pool) *Tasks {
	return &Tasks{
		logger:      l,
		userDao:     ud,
		activityDao: ad,
		mailPool:    mp,
	}
}

func (t *Tasks) UpdateActivityStatusFromNostartToInprocess() {
	time := time.Now().UnixMilli()
	activitys, err := t.activityDao.UpdateActivityStatusFromNostartToInprocess(time)
	if err != nil {
		t.logger.Info(err)
	}

	for _, activity := range activitys {
		//TODO SendMsgToParticipant()
		// e := &email.Email{
		// 	To:      []string{"1335569551@qq.com"},
		// 	From:    "JI office<1932859223@qq.com>",
		// 	Subject: "Awesome Subject",
		// 	HTML:    []byte("<h1>Fancy HTML is supported, too!</h1>"),
		// 	Headers: textproto.MIMEHeader{},
		// }
		// // getusers
		// // e := newRemindEmail(activity,users)
		// t.mailPool.Send(e, 10)

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

func (t *Tasks) newRemindEmail(acticity *model.Activity, users []*model.User) *email.Email {

	tos := make([]string, 0)
	for _, user := range users {
		tos = append(tos, user.Email)
	}

	e := &email.Email{
		To:      tos,
		From:    "JI office<1932859223@qq.com>",
		Subject: acticity.Title + " is coming",
		HTML:    []byte("<h1>" + acticity.Title + "</h1>"),
		Headers: textproto.MIMEHeader{},
	}

	return e
}
