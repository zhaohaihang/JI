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

type DBSyncProc struct {
	activityDao *dao.ActivityDao
	userDao     *dao.UserDao
	engageDao   *dao.EngageDao
	rm          *mq.RabbitMQClient
	logger      *logrus.Logger
}

func NewDBSyncProc(
	ad *dao.ActivityDao,
	ud *dao.UserDao,
	ed *dao.EngageDao,
	rm *mq.RabbitMQClient,
	l *logrus.Logger) *DBSyncProc {
	return &DBSyncProc{
		activityDao: ad,
		userDao:     ud,
		engageDao:   ed,
		rm:          rm,
		logger:      l,
	}
}

func (dsp *DBSyncProc) start() error {
	if err := dsp.rm.ConsumerDirect("activityExChange", "activityStatusToStart", dsp.activityStatusToStartHandler); err != nil {
		return err
	}

	if err := dsp.rm.ConsumerDirect("activityExChange", "activityStatusToEnd", dsp.activityStatusToEndHandler); err != nil {
		return err
	}

	return nil
}

func (dsp *DBSyncProc) stop() {
	dsp.rm.Close()
}

func (dsp *DBSyncProc) activityStatusToStartHandler(delivery amqp.Delivery) error {

	var mqactivity serializer.Activity
	if err := json.Unmarshal(delivery.Body, &mqactivity); err != nil {
		return err
	}

	// 判断活动是否存在
	activity, exist, err := dsp.activityDao.ExistOrNotByActivityId(mqactivity.ID)
	if err != nil {
		dsp.logger.Info(err)
		return err
	}
	if !exist {
		dsp.logger.Info("activity %d is not exists")
		return nil
	}
	// 判断是否应该更新状态
	dlta := activity.StartTime - time.Now().UnixMilli()
	if !(-1*ONE_MINITE_MILL <= dlta && dlta <= ONE_MINITE_MILL) {
		dsp.logger.Info("activity %d start time has change")
		return nil
	}

	if err := dsp.activityDao.UpdateActivityStatusFromNostartToInprocess(activity.ID); err != nil {
		dsp.logger.Info(err)
		return err
	}
	
	return nil
}

func (dsp *DBSyncProc) activityStatusToEndHandler(delivery amqp.Delivery) error {
	var mqactivity serializer.Activity
	if err := json.Unmarshal(delivery.Body, &mqactivity); err != nil {
		return err
	}

	// 判断活动是否存在
	activity, exist, err := dsp.activityDao.ExistOrNotByActivityId(mqactivity.ID)
	if err != nil {
		dsp.logger.Info(err)
		return err
	}
	if !exist {
		dsp.logger.Info("activity %d is not exists")
		return nil
	}
	// 判断是否应该更新状态
	dlta := activity.EndTime - time.Now().UnixMilli()
	if !(-1*ONE_MINITE_MILL <= dlta && dlta <= ONE_MINITE_MILL) {
		dsp.logger.Info("activity %d end time has change")
		return nil
	}

	if err := dsp.activityDao.UpdateActivityStatusFromInprocessToEnd(activity.ID); err != nil {
		dsp.logger.Info(err)
		return err
	}

	return nil
}
