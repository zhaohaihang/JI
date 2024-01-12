package backproc

import (
	"encoding/json"
	"ji/internal/dao"
	"ji/internal/model"
	"ji/internal/serializer"
	"ji/pkg/mq"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type FollowUpdateProc struct {
	activityDao *dao.ActivityDao
	userDao     *dao.UserDao
	followDao   *dao.FollowDao
	rm          *mq.RabbitMQClient
	logger      *logrus.Logger
}

func NewFollowUpdateProc(
	ad *dao.ActivityDao,
	ud *dao.UserDao,
	fd *dao.FollowDao,
	rm *mq.RabbitMQClient,
	l *logrus.Logger) *FollowUpdateProc {
	return &FollowUpdateProc{
		activityDao: ad,
		userDao:     ud,
		followDao:   fd,
		rm:          rm,
		logger:      l,
	}
}

func (fup *FollowUpdateProc) start() error {
	if err := fup.rm.ConsumerDirect("followExchange", "followCreateQueue", fup.followCreateHandler); err != nil {
		return err
	}

	if err := fup.rm.ConsumerDirect("followExchange", "followUpdateQueue", fup.followUpdateHandler); err != nil {
		return err
	}

	return nil
}

func (fup *FollowUpdateProc) stop() {
	fup.rm.Close()
}

func (fup *FollowUpdateProc) followCreateHandler(delivery amqp.Delivery) error {
	var mqfollow serializer.Follow
	if err := json.Unmarshal(delivery.Body, &mqfollow); err != nil {
		return err
	}

	follow := &model.Follow{
		FollowId: mqfollow.FollowId,
		UserId:     mqfollow.UserId,
		Followed:      mqfollow.Followed,
	}
	err := fup.followDao.CreateFollow(follow)
	if err != nil {
		fup.logger.Info(err)
		return err
	}
	return nil
}

func (fup *FollowUpdateProc) followUpdateHandler(delivery amqp.Delivery) error {
	var mqfollow serializer.Follow
	if err := json.Unmarshal(delivery.Body, &mqfollow); err != nil {
		return err
	}

	follow := &model.Follow{
		FollowId: mqfollow.FollowId,
		UserId:     mqfollow.UserId,
		Followed:      mqfollow.Followed,
	}
	err := fup.followDao.UpdateFollow(follow)
	if err != nil {
		fup.logger.Info(err)
		return err
	}
	return nil
}
