package cron

import (
	"ji/internal/dao"
	"ji/pkg/es"
	"ji/pkg/mail"

	"github.com/sirupsen/logrus"
)

type Tasks struct {
	logger      *logrus.Logger
	userDao     *dao.UserDao
	activityDao *dao.ActivityDao
	mailClient  *mail.MailClient
	esClient    *es.EsClient
}

func NewTasks(l *logrus.Logger, ud *dao.UserDao, ad *dao.ActivityDao, mc *mail.MailClient, ec *es.EsClient) *Tasks {
	return &Tasks{
		logger:      l,
		userDao:     ud,
		activityDao: ad,
		mailClient:  mc,
		esClient:    ec,
	}
}

func (t *Tasks) UpdateActivityStatusFromNostartToInprocess() {
	//t.mailClient.SendRemindEmails()
}

func (t *Tasks) UpdateActivityStatusFromInprocessToEnd() {

}
