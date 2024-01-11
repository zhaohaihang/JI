package service

import (
	"encoding/json"
	"fmt"
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
	// 某一个用户喜欢的活动列表
	KeyUserLikeActivity = "like:activity:user:%d"

	// 喜欢某一个活动的用户列表
	keyActivityLikedUser = "liked:user:activity:%d"
)

type LikeService struct {
	logger      *logrus.Logger
	likeDao     *dao.LikeDao
	userDao     *dao.UserDao
	activityDao *dao.ActivityDao
	mq          *mq.RabbitMQClient
	redisPool   *redis.Pool
}

func NewLikeService(l *logrus.Logger,
	ld *dao.LikeDao,
	ud *dao.UserDao,
	ad *dao.ActivityDao,
	rp *redis.Pool,
	mq *mq.RabbitMQClient) *LikeService {
	return &LikeService{
		logger:      l,
		likeDao:     ld,
		userDao:     ud,
		activityDao: ad,
		mq:          mq,
		redisPool:   rp,
	}
}

func (ls *LikeService) UserLikeActivity(uId uint, aId uint, liked int8) serializer.Response {
	code := e.SUCCESS
	_, err := ls.userDao.GetUserById(uId)
	if err != nil && err != gorm.ErrRecordNotFound {
		ls.logger.Info(err)
		code = e.ErrorGetUserInfo
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	_, err = ls.activityDao.GetActivityById(aId)
	if err != nil && err != gorm.ErrRecordNotFound {
		ls.logger.Info(err)
		code = e.ErrorGetUserInfo
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	conn := ls.redisPool.Get()
	defer conn.Close()
	switch liked {
	case 1:
		// 点赞
		conn.Do("sadd", fmt.Sprintf(KeyUserLikeActivity, uId), aId)
		conn.Do("sadd", fmt.Sprintf(keyActivityLikedUser, aId), uId)
	case 2:
		// 取消点赞
		conn.Do("sadd", fmt.Sprintf(KeyUserLikeActivity, uId), aId)
		conn.Do("sadd", fmt.Sprintf(keyActivityLikedUser, aId), uId)
	default:
		ls.logger.Info("liked error")
	}

	isLiked, _ := ls.likeDao.IsLikedByUser(uId, aId)
	serializeLike := serializer.BuildLike(&model.Like{
		UserId:     uId,
		AcitivtyId: uId,
		Liked:      liked,
	})
	message, _ := json.Marshal(serializeLike)
	if isLiked == -1 {
		if err := ls.mq.SendMessageDirect(message, "likeExchange", "likeCreateQueue"); err != nil {
			ls.logger.Info(err)
			code = e.ErrorSendMsgToMQ
		}
	} else {
		if err := ls.mq.SendMessageDirect(message, "likeExchange", "likeUpdateQueue"); err != nil {
			ls.logger.Info(err)
			code = e.ErrorSendMsgToMQ
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
