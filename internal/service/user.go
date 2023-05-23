package service

import (
	"context"
	"ji/internal/dao"
	"ji/internal/model"
	"ji/internal/serializer"
	"ji/pkg/e"
	"ji/pkg/storages/qiniu"
	"ji/pkg/utils/tokenutil.go"
	"strconv"
	"strings"

	"mime/multipart"
	"time"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	userDao      *dao.UserDao
	activityDao  *dao.ActivityDao
	qiniuStroage *qiniu.QiNiuStroage
}

func NewUserService(ud *dao.UserDao, ad *dao.ActivityDao, qs *qiniu.QiNiuStroage) *UserService {
	return &UserService{
		userDao:      ud,
		activityDao:  ad,
		qiniuStroage: qs,
	}
}

var UserServiceProviderSet = wire.NewSet(NewUserService)

// Login 用户登陆函数
func (service *UserService) Login(ctx context.Context, loginUserInfo serializer.LoginUserInfo) serializer.Response {
	var user *model.User
	code := e.SUCCESS
	// userDao := dao.NewUserDao(ctx)
	user, exist, _ := service.userDao.ExistOrNotByUserName(loginUserInfo.UserName)
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
		if err := service.userDao.CreateUser(user); err != nil {
			logrus.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		// 新注册用户，返回的最后一次登录时间为当前时间
		user.LastLogin = time.Now().UnixMilli()
	}

	token, err := tokenutil.GenerateToken(user.ID, loginUserInfo.UserName, 0)
	if err != nil {
		logrus.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	service.userDao.UpdateLastLoginById(user.ID, time.Now().UnixMilli())

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
	// userDao := dao.NewUserDao(ctx)
	user, err = service.userDao.GetUserById(uId)

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

	err = service.userDao.UpdateUserById(uId, user)
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
	// userDao := dao.NewUserDao(ctx)
	user, err = service.userDao.GetUserById(uId)

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

func (service *UserService) UploadUserAvatar(ctx context.Context, uId uint, file multipart.File, fileHeader *multipart.FileHeader) serializer.Response {
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
	filename := "user_avatar/" + strconv.Itoa(int(uId)) + "_" + ti + "_" + strconv.FormatInt(time.Now().Unix(), 10) + ext

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
