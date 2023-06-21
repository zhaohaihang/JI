package backproc

import (
	"encoding/json"
	"ji/internal/dao"
	"ji/internal/serializer"
	"ji/pkg/mq"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

const (
	ONE_MINITE_MILL = 1 * 60 * 1000
)

type DBUpdateProc struct {
	activityDao *dao.ActivityDao
	userDao     *dao.UserDao
	engageDao   *dao.EngageDao
	rm          *mq.RabbitMQClient
	logger      *logrus.Logger
}

func NewDBUpdateProc(
	ad *dao.ActivityDao,
	ud *dao.UserDao,
	ed *dao.EngageDao,
	rm *mq.RabbitMQClient,
	l *logrus.Logger) *DBUpdateProc {
	return &DBUpdateProc{
		activityDao: ad,
		userDao:     ud,
		engageDao:   ed,
		rm:          rm,
		logger:      l,
	}
}

func (dup *DBUpdateProc) start() error {
	if err := dup.rm.ConsumerDirect("activityExChange", "activityStatusToStart", dup.activityStatusToStartHandler); err != nil {
		return err
	}

	if err := dup.rm.ConsumerDirect("activityExChange", "activityStatusToEnd", dup.activityStatusToEndHandler); err != nil {
		return err
	}

	return nil
}

func (dup *DBUpdateProc) stop() {
	dup.rm.Close()
}

func (dup *DBUpdateProc) activityStatusToStartHandler(delivery amqp.Delivery) error {

	var mqactivity serializer.Activity
	if err := json.Unmarshal(delivery.Body, &mqactivity); err != nil {
		return err
	}

	// 判断活动是否存在
	activity, exist, err := dup.activityDao.ExistOrNotByActivityId(mqactivity.ID)
	if err != nil {
		dup.logger.Info(err)
		return err
	}
	if !exist {
		dup.logger.Info("activity %d is not exists")
		return nil
	}
	// 判断是否应该更新状态
	dlta := activity.StartTime - time.Now().UnixMilli()
	if !(-1*ONE_MINITE_MILL <= dlta && dlta <= ONE_MINITE_MILL) {
		dup.logger.Info("activity %d start time has change")
		return nil
	}

	if err := dup.activityDao.UpdateActivityStatusFromNostartToInprocess(activity.ID); err != nil {
		dup.logger.Info(err)
		return err
	}
	
	return nil
}

func (dup *DBUpdateProc) activityStatusToEndHandler(delivery amqp.Delivery) error {
	var mqactivity serializer.Activity
	if err := json.Unmarshal(delivery.Body, &mqactivity); err != nil {
		return err
	}

	// 判断活动是否存在
	activity, exist, err := dup.activityDao.ExistOrNotByActivityId(mqactivity.ID)
	if err != nil {
		dup.logger.Info(err)
		return err
	}
	if !exist {
		dup.logger.Info("activity %d is not exists")
		return nil
	}
	// 判断是否应该更新状态
	dlta := activity.EndTime - time.Now().UnixMilli()
	if !(-1*ONE_MINITE_MILL <= dlta && dlta <= ONE_MINITE_MILL) {
		dup.logger.Info("activity %d end time has change")
		return nil
	}

	if err := dup.activityDao.UpdateActivityStatusFromInprocessToEnd(activity.ID); err != nil {
		dup.logger.Info(err)
		return err
	}

	return nil
}
