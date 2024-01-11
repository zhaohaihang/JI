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

type LikeUpdateProc struct {
	activityDao *dao.ActivityDao
	userDao     *dao.UserDao
	likeDao     *dao.LikeDao
	rm          *mq.RabbitMQClient
	logger      *logrus.Logger
}

func NewLikeUpdateProc(
	ad *dao.ActivityDao,
	ud *dao.UserDao,
	ld *dao.LikeDao,
	rm *mq.RabbitMQClient,
	l *logrus.Logger) *LikeUpdateProc {
	return &LikeUpdateProc{
		activityDao: ad,
		userDao:     ud,
		likeDao:     ld,
		rm:          rm,
		logger:      l,
	}
}

func (lup *LikeUpdateProc) start() error {
	if err := lup.rm.ConsumerDirect("likeExchange", "likeCreateQueue", lup.likeCreateHandler); err != nil {
		return err
	}

	if err := lup.rm.ConsumerDirect("likeExchange", "likeUpdateQueue", lup.likeUpdateHandler); err != nil {
		return err
	}

	return nil
}

func (dup *LikeUpdateProc) stop() {
	dup.rm.Close()
}

func (lup *LikeUpdateProc) likeCreateHandler(delivery amqp.Delivery) error {
	var mqlike serializer.Like
	if err := json.Unmarshal(delivery.Body, &mqlike); err != nil {
		return err
	}

	like := &model.Like{
		AcitivtyId: mqlike.ActivityId,
		UserId:     mqlike.UserId,
		Liked:      mqlike.Liked,
	}
	err := lup.likeDao.CreateLike(like)
	if err != nil {
		lup.logger.Info(err)
		return err
	}
	return nil
}

func (lup *LikeUpdateProc) likeUpdateHandler(delivery amqp.Delivery) error {
	var mqlike serializer.Like
	if err := json.Unmarshal(delivery.Body, &mqlike); err != nil {
		return err
	}

	like := &model.Like{
		AcitivtyId: mqlike.ActivityId,
		UserId:     mqlike.UserId,
		Liked:      mqlike.Liked,
	}
	err := lup.likeDao.UpdateLike(like)
	if err != nil {
		lup.logger.Info(err)
		return err
	}
	return nil
}
