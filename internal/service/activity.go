package service

import (
	"context"
	"ji/internal/dao"
	"ji/internal/model"
	"ji/internal/serializer"
	"ji/pkg/e"
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
	userDao      *dao.UserDao
	activityDao  *dao.ActivityDao
	redisPool    *redis.Pool
	qiniuStroage *qiniu.QiNiuStroage
}

func NewActivityService(ud *dao.UserDao, ad *dao.ActivityDao, rp *redis.Pool, qs *qiniu.QiNiuStroage) *ActivityService {
	return &ActivityService{
		userDao:      ud,
		activityDao:  ad,
		redisPool:    rp,
		qiniuStroage: qs,
	}
}

var ActivityServiceProviderSet = wire.NewSet(NewActivityService)

func (service *ActivityService) CreateActivity(ctx context.Context, uId uint, activityInfo serializer.CreateActivityInfo) serializer.Response {

	code := e.SUCCESS
	// activityDao := dao.NewActivityDao(ctx)
	// userDao := dao.NewUserDao(ctx)

	user, err := service.userDao.GetUserById(uId)
	if err != nil {
		logrus.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	activity := &model.Activity{
		Title:          activityInfo.Title,
		Introduction:   activityInfo.Introduction,
		Status:         activityInfo.Status,
		StartTime:      activityInfo.StartTime,
		EndTime:        activityInfo.EndTime,
		Location:       model.Point(activityInfo.Location),
		ExpectedNumber: activityInfo.ExpectedNumber,
		CurrentNumber:  0,
		UserId:         user.ID,
		UserName:       user.UserName,
		UserAvatar:     user.Avatar,
	}

	if err := service.activityDao.CreateActivity(activity); err != nil {
		logrus.Info(err)
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

func (service *ActivityService) GetActivityById(ctx context.Context, aId uint) serializer.Response {
	code := e.SUCCESS
	activity, err := service.activityDao.GetActivityById(aId)
	if err != nil {
		logrus.Info(err)
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

func (service *ActivityService) ListActivityByUserId(ctx context.Context, uId uint, basePage serializer.BasePage) serializer.Response {
	code := e.SUCCESS
	// activityDao := dao.NewActivityDao(ctx)
	activitys, total, err := service.activityDao.ListActivityByUserId(uId, model.BasePage(basePage))
	if err != nil {
		logrus.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.BuildListResponse(serializer.BuildActivitys(activitys), uint(total))
}

func (service *ActivityService) ListNearActivity(ctx context.Context, nearInfo serializer.NearInfo) serializer.Response {
	code := e.SUCCESS
	// activityDao := dao.NewActivityDao(ctx)
	activitys, total, err := service.activityDao.ListNearActivity(nearInfo.Lat, nearInfo.Lng, nearInfo.Rad)
	if err != nil {
		logrus.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.BuildListResponse(serializer.BuildActivitys(activitys), uint(total))
}

func (service *ActivityService) UploadActivityCover(ctx context.Context, uId uint, file multipart.File, fileHeader *multipart.FileHeader) serializer.Response {
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

	path, err := service.qiniuStroage.UploadToQiNiu(filename, file, fileHeader.Size)

	if err != nil {
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
