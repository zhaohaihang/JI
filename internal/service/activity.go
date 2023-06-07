package service

import (
	"encoding/json"
	"ji/internal/dao"
	"ji/internal/model"
	"ji/internal/serializer"
	"ji/pkg/consts"
	"ji/pkg/e"
	"ji/pkg/mq"
	"ji/pkg/storages/qiniu"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

type ActivityService struct {
	logger       *logrus.Logger
	userDao      *dao.UserDao
	activityDao  *dao.ActivityDao
	redisPool    *redis.Pool
	qiniuStroage *qiniu.QiNiuStroage
	mq           *mq.RabbitMQClient
}

func NewActivityService(
	l *logrus.Logger,
	ud *dao.UserDao,
	ad *dao.ActivityDao,
	rp *redis.Pool,
	qs *qiniu.QiNiuStroage,
	mq *mq.RabbitMQClient) *ActivityService {
	return &ActivityService{
		logger:       l,
		userDao:      ud,
		activityDao:  ad,
		redisPool:    rp,
		qiniuStroage: qs,
		mq:           mq,
	}
}

var ActivityServiceProviderSet = wire.NewSet(NewActivityService)

func (as *ActivityService) CreateActivity(uId uint, activityInfo serializer.CreateActivityInfo) serializer.Response {

	code := e.SUCCESS

	user, err := as.userDao.GetUserById(uId)
	if err != nil {
		as.logger.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	activity := &model.Activity{
		Title:          activityInfo.Title,
		Introduction:   activityInfo.Introduction,
		Status:         consts.ACTIVITY_STATUS_NOSTART,
		StartTime:      activityInfo.StartTime,
		EndTime:        activityInfo.EndTime,
		Location:       model.Point(activityInfo.Location),
		ExpectedNumber: activityInfo.ExpectedNumber,
		CurrentNumber:  0,
		UserId:         user.ID,
		UserName:       user.UserName,
		UserAvatar:     user.Avatar,
	}

	if err := as.activityDao.CreateActivity(activity); err != nil {
		as.logger.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// TODO send to mq
	a := serializer.BuildActivity(activity)
	b,_:=json.Marshal(a)
	as.mq.SendMessageDirect(b,"activityExChange", "activityCreateQueue")

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildActivity(activity),
		Msg:    e.GetMsg(code),
	}
}

func (as *ActivityService) GetActivityById(aId uint) serializer.Response {
	code := e.SUCCESS
	activity, err := as.activityDao.GetActivityById(aId)
	if err != nil {
		as.logger.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildActivity(activity),
		Msg:    e.GetMsg(code),
	}
}

func (as *ActivityService) ListActivityByUserId(uId uint, basePage serializer.BasePage) serializer.Response {
	code := e.SUCCESS
	activitys, total, err := as.activityDao.ListActivityByUserId(uId, model.BasePage(basePage))
	if err != nil {
		as.logger.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.BuildListResponse(serializer.BuildActivitys(activitys), uint(total))
}

func (as *ActivityService) ListNearActivity(nearInfo serializer.NearInfo) serializer.Response {
	code := e.SUCCESS
	activitys, total, err := as.activityDao.ListNearActivity(nearInfo.Lat, nearInfo.Lng, nearInfo.Rad)
	if err != nil {
		as.logger.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.BuildListResponse(serializer.BuildActivitys(activitys), uint(total))
}

func (as *ActivityService) UploadActivityCover(uId uint, file multipart.File, fileHeader *multipart.FileHeader) serializer.Response {
	code := e.SUCCESS
	var err error

	//重命名文件的名称
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	ti := tm.Format("2006010203040501")
	//提取文件后缀类型
	var ext string
	if pos := strings.LastIndexByte(fileHeader.Filename, '.'); pos != -1 {
		ext = fileHeader.Filename[pos:]
		if ext == "." {
			ext = ""
		}
	}
	filename := "activity_bg/" + strconv.Itoa(int(uId)) + "_" + ti + "_" + strconv.FormatInt(time.Now().Unix(), 10) + ext

	path, err := as.qiniuStroage.UploadToQiNiu(filename, file, fileHeader.Size)

	if err != nil {
		as.logger.Info(err)
		code = e.ErrorUploadFile
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  path,
		}
	}

	return serializer.Response{
		Status: code,
		Data:   path,
		Msg:    e.GetMsg(code),
	}
}
