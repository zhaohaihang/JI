package dao

import (
	"errors"
	"ji/internal/model"
	"ji/pkg/database"

	"gorm.io/gorm"
)

type FollowDao struct {
	db *gorm.DB
}

func NewFollowedDao(db *database.Database) *FollowDao {
	return &FollowDao{
		db: db.Mysql,
	}
}

func (fd *FollowDao)CreateFollow(follow *model.Follow) error {
	return fd.db.Model(&model.Follow{}).Create(&follow).Error
}

func (fd *FollowDao) UpdateFollow(follow *model.Follow) error {
	return fd.db.Model(&model.Like{}).Where("user_id=? and followed_id=?", follow.UserId, follow.FollowId).
		Updates(&follow).Error
}

func (fd *FollowDao) IsFollowed(uId uint ,followdId uint) (int8,error) {
	var isFollowed int8
	result := fd.db.Model(&model.Follow{}).Select("followed").Where("user_id= ? and follow_id= ?", uId, followdId).First(&isFollowed)
	c := result.RowsAffected
	if c == 0 {
		return -1, errors.New("current user haven not liked current actvity")
	}
	if result.Error != nil {
		return -1, result.Error
	}
	return isFollowed, nil
}
