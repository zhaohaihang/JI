package service

import (
	"context"
	"ji/pkg/e"
	"ji/pkg/utils"
	"ji/repository/db/dao"
	"ji/repository/db/model"
	"ji/serializer"

	"github.com/sirupsen/logrus"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
}

func (service UserService) Register(ctx context.Context) serializer.Response {
	var user *model.User
	code := e.SUCCESS

	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user = &model.User{
		UserName: service.UserName,
		Status:   model.Active,
	}
	// 加密密码
	if err = user.SetPassword(service.Password); err != nil {
		logrus.Info(err)
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 创建用户
	err = userDao.CreateUser(user)
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
		Msg:    e.GetMsg(code),
	}
}

// Login 用户登陆函数
func (service *UserService) Login(ctx context.Context) serializer.Response {
	var user *model.User
	code := e.SUCCESS
	userDao := dao.NewUserDao(ctx)
	user, exist, _ := userDao.ExistOrNotByUserName(service.UserName)
	if exist { // 如果存在，则校验密码
		if !user.CheckPassword(service.Password) {
			code = e.ErrorNotCompare
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	} else { // 如果不存在则新建
		user = &model.User{
			UserName: service.UserName,
			Status:   model.Active,
		}
		user.GenerateRandomNickName()
		user.Avatar = "avatar.jpg"
		// 加密密码
		if err := user.SetPassword(service.Password); err != nil {
			logrus.Info(err)
			code = e.SUCCESS
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		// 创建用户
		if err := userDao.CreateUser(user); err != nil {
			logrus.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	}

	token, err := utils.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		logrus.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    e.GetMsg(code),
	}
}
