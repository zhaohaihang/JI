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
	// 该用户关注了谁
	KeyUserFollowOtherUser = "follow:user:%d"

	// 谁关注了该用户
	KeyUserFollowedOtherUser = "followed:user:%d"
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

func (fs *FollowService) UserFollow(uId uint, followId uint, followed int8) serializer.Response {
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

	isFollowed, _ := fs.followDao.IsFollowed(uId, followId)
	serializeFollow := serializer.BuildFollow(&model.Follow{
		UserId:   uId,
		FollowId: followId,
		Followed: followed,
	})

	message, _ := json.Marshal(serializeFollow)
	if isFollowed == -1 {
		if err := fs.mq.SendMessageDirect(message, "followExchange", "followCreateQueue"); err != nil {
			fs.logger.Info(err)
			code = e.ErrorSendMsgToMQ
		}
	} else {
		if err := fs.mq.SendMessageDirect(message, "followExchange", "followUpdateQueue"); err != nil {
			fs.logger.Info(err)
			code = e.ErrorSendMsgToMQ
		}
	}

	conn := fs.redisPool.Get()
	defer conn.Close()

	reply, _ := redis.Bool(conn.Do("exists", fmt.Sprintf(KeyUserFollowOtherUser, uId)))
	//不存在
	if reply {
		//从数据库导入
		followIds, _, _ := fs.followDao.GetFollowIdsById(uId)
		conn.Do("sadd", fmt.Sprintf(KeyUserFollowOtherUser, uId), followIds)
	}

	// 加入当前信息
	if followed == 1 {
		conn.Do("sadd", fmt.Sprintf(KeyUserFollowOtherUser, uId), followId)
	} else {
		conn.Do("sdel", fmt.Sprintf(KeyUserFollowOtherUser, uId), followId)
	}

	reply2, _ := redis.Bool(conn.Do("exists", fmt.Sprintf(KeyUserFollowedOtherUser, followId)))
	//不存在
	if reply2 {
		//从数据库导入
		followedIds, _, _ := fs.followDao.GetFollowedIdsById(uId)
		conn.Do("sadd", fmt.Sprintf(KeyUserFollowedOtherUser, followId), followedIds)
	}

	// 加入当前信息
	if followed == 1 {
		conn.Do("sadd", fmt.Sprintf(KeyUserFollowedOtherUser, followId), uId)
	} else {
		conn.Do("sdel", fmt.Sprintf(KeyUserFollowedOtherUser, followId), uId)
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
