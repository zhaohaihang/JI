package service

import (
	"ji/internal/dao"
	"ji/internal/model"
	"ji/internal/serializer"
	"ji/pkg/e"

	"github.com/sirupsen/logrus"
)

type EngageService struct {
	logger      *logrus.Logger
	activityDao *dao.ActivityDao
	engageDao   *dao.EngageDao
}

func NewEngageService(
	l *logrus.Logger,
	ad *dao.ActivityDao,
	ed *dao.EngageDao,
) *EngageService {
	return &EngageService{
		logger:      l,
		activityDao: ad,
		engageDao:   ed,
	}
}

func (es *EngageService) EngageActivity(uId uint, aId uint) serializer.Response {
	var engage *model.Engage
	code := e.SUCCESS

	engage = &model.Engage{
		UserId:     uId,
		ActivityId: aId,
	}

	if err := es.engageDao.CreateEngage(engage); err != nil {
		es.logger.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if err := es.activityDao.UpdateActivityCurrentNumById(aId, 1); err != nil {
		es.logger.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (es *EngageService) DelEngageActivity(uId uint, aId uint) serializer.Response {

	code := e.SUCCESS

	if err := es.engageDao.DelEngageByIds(uId, aId); err != nil {
		es.logger.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if err := es.activityDao.UpdateActivityCurrentNumById(aId, -1); err != nil {
		es.logger.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
