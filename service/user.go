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
	
}

// Login 用户登陆函数
func (service *UserService) Login(ctx context.Context, loginUserInfo LoginUserInfo) serializer.Response {
	var user *model.User
	code := e.SUCCESS
	userDao := dao.NewUserDao(ctx)
	user, exist, _ := userDao.ExistOrNotByUserName(loginUserInfo.UserName)
	if exist { // 如果存在，则校验密码
		if !user.CheckPassword(loginUserInfo.Password) {
			code = e.ErrorNotCompare
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	} else { // 如果不存在则新建
		user = &model.User{
			UserName: loginUserInfo.UserName,
			Status:   model.Active,
		}
		user.Avatar = "avatar.jpg"
		if loginUserInfo.Type == 0 {
			user.Phone = loginUserInfo.UserName
		}else {
			user.Email = loginUserInfo.UserName
		}
		// 加密密码
		if err := user.SetPassword(loginUserInfo.Password); err != nil {
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

	token, err := utils.GenerateToken(user.ID, loginUserInfo.UserName, 0)
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

// Update 用户修改信息
func (service UserService) UpdateUserById(ctx context.Context, uId uint) serializer.Response {
	var user *model.User
	var err error
	code := e.SUCCESS
	// 找到用户
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uId)

	if err != nil {
		logrus.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// user.Biography = service.Biography
	// user.Address = service.Address

	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		logrus.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}
