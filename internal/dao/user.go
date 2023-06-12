package dao

import (
	"ji/internal/model"
	"ji/pkg/database"

	"gorm.io/gorm"
)

type UserDao struct {
	DB *gorm.DB
}

func NewUserDao(db *database.Database) *UserDao {
	return &UserDao{
		DB: db.Mysql,
	}
}

// GetUserById 根据 id 获取用户
func (ud *UserDao) GetUserById(uId uint) (user *model.User, err error) {
	//err = ud.DB.Model(&model.User{}).Select("id,ST_AsTEXT(location)").Where("id=?", uId).
	err = ud.DB.Model(&model.User{}).Where("id=?", uId).
		First(&user).Error
	return
}

// UpdateUserById 根据 id 更新用户信息
func (ud *UserDao) UpdateUserById(uId uint, user *model.User) (err error) {
	return ud.DB.Model(&model.User{}).Where("id=?", uId).
		Updates(&user).Error
}

// ExistOrNotByUserName 根据username判断是否存在该名字
func (ud *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = ud.DB.Model(&model.User{}).Where("user_name=?", userName).Count(&count).Error
	if count == 0 {
		return user, false, err
	}
	err = ud.DB.Model(&model.User{}).Where("user_name=?", userName).First(&user).Error
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}

// CreateUser 创建用户
func (ud *UserDao) CreateUser(user *model.User) error {
	return ud.DB.Model(&model.User{}).Create(&user).Error
}

// UpdateLastLoginById 根据ID更新最后登录时间
func (ud *UserDao) UpdateLastLoginById(uId uint, loginTime int64) (err error) {
	return ud.DB.Model(&model.User{}).Select("last_login").Where("id=?", uId).Updates(&model.User{LastLogin: loginTime}).Error
}

func (ud *UserDao) UpdateUserAvatarById(uId uint, path string) (err error) {
	return ud.DB.Model(&model.User{}).Select("avatar").Where("id=?", uId).Updates(&model.User{Avatar: path}).Error
}

// func GetNearestUsers(lat, lng float64) ([]User, error) {
//     var users []User
//     if err := db.Where(`ST_Distance_Sphere(location, ST_GeomFromText(?)) < 1609.34`, fmt.Sprintf("POINT(%.6f %.6f)", lng, lat)).Order("ST_Distance_Sphere(location, ST_GeomFromText(?)) DESC", fmt.Sprintf("POINT(%.6f %.6f)", lng, lat)).Find(&users).Error; err != nil {
//         return nil, err
//     }
//     return users, nil
// }

func (ud *UserDao) ListUsersByIds(uIds[]uint)(users []*model.User, err error){
	err = ud.DB.Model(&model.User{}).Find(&users, uIds).Error
	return
}