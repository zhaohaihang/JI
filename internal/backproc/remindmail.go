package backproc

import (
	"encoding/json"
	"ji/internal/dao"
	"ji/internal/serializer"
	"ji/pkg/mail"
	"ji/pkg/mq"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

const (
	ONE_HONOR_MILL = 1 * 60 * 60 * 1000
)

type RemindMailProc struct {
	activityDao    *dao.ActivityDao
	userDao        *dao.UserDao
	engageDao      *dao.EngageDao
	rabbitmqClient *mq.RabbitMQClient
	mailClient     *mail.MailClient
	logger         *logrus.Logger
}

func NewRemindMailProc(
	ad *dao.ActivityDao,
	ud *dao.UserDao,
	ed *dao.EngageDao,
	rc *mq.RabbitMQClient,
	mc *mail.MailClient,
	l *logrus.Logger,
) *RemindMailProc {
	return &RemindMailProc{
		activityDao:    ad,
		userDao:        ud,
		engageDao:      ed,
		rabbitmqClient: rc,
		mailClient:     mc,
		logger:         l,
	}
}

func (rmp *RemindMailProc) start() error {
	if err := rmp.rabbitmqClient.ConsumerDelay("activityExchange", "activityStartRemind", rmp.activityStartRemindHandler); err != nil {
		return err
	}

	if err := rmp.rabbitmqClient.ConsumerDirect("activityExChange", "activityTimeOrLocationChangeRemind", rmp.activityTimeOrLocationChangeRemindHandler); err != nil {
		return err
	}
	return nil
}

func (rmp *RemindMailProc) stop() {
	rmp.rabbitmqClient.Close()
}

// 提前一小时通知
func (rmp *RemindMailProc) activityStartRemindHandler(delivery amqp.Delivery) error {
	var mqActivity serializer.Activity
	if err := json.Unmarshal(delivery.Body, &mqActivity); err != nil {
		rmp.logger.Info(err)
		return err
	}

	// 判断活动是否存在
	activity, exist, err := rmp.activityDao.ExistOrNotByActivityId(mqActivity.ID)
	if err != nil {
		rmp.logger.Info(err)
		return err
	}
	if !exist {
		rmp.logger.Info("activity %d is not exists")
		return nil
	}
	// 判断时间是否是提前时间一小时
	dlta := activity.StartTime - time.Now().UnixMilli()
	if !(ONE_HONOR_MILL-30*1000 <= dlta && dlta <= ONE_HONOR_MILL+30*1000) {
		rmp.logger.Info("activity %d start time has change")
		return nil
	}

	// 获取该活动的所有参加人员id
	uIds, _, err := rmp.engageDao.ListUserIdsByActivityId(activity.ID)
	if err != nil {
		rmp.logger.Info(err)
		return err
	}
	rmp.logger.Debugf("particspat user ids: %v",uIds)


	//获取参加人员email
	emails, err := rmp.userDao.ListUsersEmailsByIds(uIds)
	if err != nil {
		rmp.logger.Info(err)
		return err
	}
	rmp.logger.Debugf("particspat user emails: %v",emails)

	// 发送消息至每一个人
	if err := rmp.mailClient.SendRemindEmails(emails, "Activity Start", activity.Title+"The event will start in one hour"); err != nil {
		rmp.logger.Info(err)
		return err
	}
	return nil
}

// 时间或地点发生变化 ，应立即通知参加人
// activityTimeOrLocationChange
func (rmp *RemindMailProc) activityTimeOrLocationChangeRemindHandler(delivery amqp.Delivery) error {
	var mqActivity serializer.Activity
	if err := json.Unmarshal(delivery.Body, &mqActivity); err != nil {
		rmp.logger.Info(err)
		return err
	}

	// 获取该活动的所有参加人员id
	uIds, _, err := rmp.engageDao.ListUserIdsByActivityId(mqActivity.ID)
	if err != nil {
		rmp.logger.Info(err)
		return err
	}
	rmp.logger.Debugf("particspat user ids: %v",uIds)

	//获取参加人员email
	emails, err := rmp.userDao.ListUsersEmailsByIds(uIds)
	if err != nil {
		rmp.logger.Info(err)
		return err
	}
	rmp.logger.Debugf("particspat user emails: %v",emails)

	// 发送消息至每一个人
	if err := rmp.mailClient.SendRemindEmails(
		emails,
		"Activity Tiem or Location Change",
		mqActivity.Title+"The activity's start time or location change"); err != nil {
			rmp.logger.Info(err)
			return err
	}
	return nil
}
