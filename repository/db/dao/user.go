package dao

import (
	"context"
	"ji/repository/db/model"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// GetUserById 根据 id 获取用户
func (userDao *UserDao) GetUserById(uId uint) (user *model.User, err error) {
	//err = userDao.DB.Model(&model.User{}).Select("id,ST_AsTEXT(location)").Where("id=?", uId).
	err = userDao.DB.Model(&model.User{}).Where("id=?", uId).
		First(&user).Error
	return
}

// UpdateUserById 根据 id 更新用户信息
func (userDao *UserDao) UpdateUserById(uId uint, user *model.User) (err error) {
	return userDao.DB.Model(&model.User{}).Where("id=?", uId).
		Updates(&user).Error
}

// ExistOrNotByUserName 根据username判断是否存在该名字
func (userDao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = userDao.DB.Model(&model.User{}).Where("user_name=?", userName).Count(&count).Error
	if count == 0 {
		return user, false, err
	}
	err = userDao.DB.Model(&model.User{}).Where("user_name=?", userName).First(&user).Error
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}

// CreateUser 创建用户
func (userDao *UserDao) CreateUser(user *model.User) error {
	return userDao.DB.Model(&model.User{}).Create(&user).Error
}

// UpdateLastLoginById 根据ID更新最后一次登录时间
func (userDao *UserDao) UpdateLastLoginById(uId uint, user *model.User) (err error) {
	return userDao.DB.Model(&model.User{}).Where("id=?", uId).Updates(user).Error
}

// func GetNearestUsers(lat, lng float64) ([]User, error) {
//     var users []User
//     if err := db.Where(`ST_Distance_Sphere(location, ST_GeomFromText(?)) < 1609.34`, fmt.Sprintf("POINT(%.6f %.6f)", lng, lat)).Order("ST_Distance_Sphere(location, ST_GeomFromText(?)) DESC", fmt.Sprintf("POINT(%.6f %.6f)", lng, lat)).Find(&users).Error; err != nil {
//         return nil, err
//     }
//     return users, nil
// }