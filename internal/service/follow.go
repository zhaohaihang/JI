package service

import (
	"encoding/json"
	"ji/internal/dao"
	"ji/internal/model"
	"ji/internal/serializer"
	"ji/pkg/e"
	"ji/pkg/mq"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	// 某一个用户关注的其他用户
	KeyUserFollowOtherUser = "follow:user:%d"
)

type FollowService struct {
	logger    *logrus.Logger
	followDao *dao.FollowDao
	userDao   *dao.UserDao
	mq        *mq.RabbitMQClient
	redisPool *redis.Pool
}

func NewFollowService(
	l *logrus.Logger,
	fd *dao.FollowDao,
	rp *redis.Pool,
	mq *mq.RabbitMQClient) *FollowService {
	return &FollowService{
		logger:    l,
		followDao: fd,
		mq:        mq,
		redisPool: rp,
	}
}

func (fs *FollowService) UserFollow(uId uint, followId uint) serializer.Response {
	code := e.SUCCESS
	_, err := fs.userDao.GetUserById(uId)
	if err != nil && err != gorm.ErrRecordNotFound {
		fs.logger.Info(err)
		code = e.ErrorGetUserInfo
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	_, err = fs.userDao.GetUserById(followId)
	if err != nil && err != gorm.ErrRecordNotFound {
		fs.logger.Info(err)
		code = e.ErrorGetUserInfo
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	isFollowed, err := fs.followDao.IsFollowed(uId, followId)
	serializeFollow := serializer.BuildFollow(&model.Follow{
		UserId:   uId,
		FollowId: followId,
		Followed: 1,
	})

	message, _ := json.Marshal(serializeFollow)
	if isFollowed == -1 {
		if err := fs.mq.SendMessageDirect(message, "likeExchange", "likeCreateQueue"); err != nil {
			fs.logger.Info(err)
			code = e.ErrorSendMsgToMQ
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
