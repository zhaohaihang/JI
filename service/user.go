package service

import (
	"context"
	"ji/pkg/e"
	"ji/pkg/utils"
	"ji/repository/db/dao"
	"ji/repository/db/model"
	"ji/serializer"
	"mime/multipart"
	"time"

	"github.com/sirupsen/logrus"
)

type UserService struct {
}

// Login 用户登陆函数
func (service *UserService) Login(ctx context.Context, loginUserInfo serializer.LoginUserInfo) serializer.Response {
	var user *model.User
	code := e.SUCCESS
	userDao := dao.NewUserDao(ctx)
	user, exist, _ := userDao.ExistOrNotByUserName(loginUserInfo.UserName)
	if exist { // 如果存在，则校验密码
		if !user.CheckPassword(loginUserInfo.Password) {
			code = e.ErrorPasswordNotCompare
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
		} else {
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
		// 新注册用户，返回的最后一次登录时间为当前时间
		user.LastLogin = time.Now()
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

	userDao.UpdateLastLoginById(user.ID, time.Now())

	return serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    e.GetMsg(code),
	}
}

// Update 用户修改信息
func (service UserService) UpdateUserById(ctx context.Context, uId uint, updateUserInfo serializer.UpdateUserInfo) serializer.Response {
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
	// 更新字段
	user.Biography = updateUserInfo.Biography
	user.Address = updateUserInfo.Address
	user.Email = updateUserInfo.Address
	user.Phone = updateUserInfo.Address
	user.Location = model.Point{
		Lat: updateUserInfo.Location.Lat,
		Lng: updateUserInfo.Location.Lng,
	}

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

func (service UserService) GetUserById(ctx context.Context, uId uint) serializer.Response {
	var err error
	var user *model.User
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

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}

func (service *UserService) UploadUserAvatar(ctx context.Context, uId uint, file multipart.File, fileSize int64) serializer.Response {
	code := e.SUCCESS
	var user *model.User
	var err error

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
	var path string
	
	path, err = utils.UploadAvatarToLocalStatic(file, uId, user.UserName)
	
	if err != nil {
		code = e.ErrorUploadFile
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  path,
		}
	}

	user.Avatar = path
	err = userDao.UpdateUserAvatarById(uId, path)
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