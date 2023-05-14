package service

import (
	"context"
	"ji/pkg/e"
	"ji/repository/db/dao"
	"ji/repository/db/model"
	"ji/serializer"

	"github.com/sirupsen/logrus"
)

type ActivityService struct {
}

func (service *ActivityService) CreateActivity(ctx context.Context, uId uint, activityInfo serializer.CreateActivityInfo) serializer.Response {

	code := e.SUCCESS
	activityDao := dao.NewActivityDao(ctx)
	userDao := dao.NewUserDao(ctx)

	user, err := userDao.GetUserById(uId)
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

	if err := activityDao.CreateActivity(activity); err != nil {
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
