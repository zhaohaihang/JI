package dao

import (
	"errors"
	"ji/internal/model"
	"ji/pkg/database"

	"gorm.io/gorm"
)

type LikeDao struct {
	DB *gorm.DB
}

func NewLikeDao(db *database.Database) *LikeDao {
	return &LikeDao{
		DB: db.Mysql,
	}
}

func (ld *LikeDao) CreateLike(like *model.Like) error {
	return ld.DB.Model(&model.Like{}).Create(&like).Error
}

func (ld *LikeDao) UpdateLike(like *model.Like) error {
	return ld.DB.Model(&model.Like{}).Where("user_id=? and activity_id=?", like.UserId, like.AcitivtyId).
		Updates(&like).Error
}

func (ld *LikeDao) IsLikedByUser(uId uint, aId uint) (int8, error) {
	var isLiked int8
	result := ld.DB.Model(&model.Activity{}).Select("liked").Where("user_id= ? and actvity_id= ?", uId, aId).First(&isLiked)
	c := result.RowsAffected
	if c == 0 {
		return -1, errors.New("current user haven not liked current actvity")
	}
	if result.Error != nil {
		return -1, result.Error
	}
	return isLiked, nil
}
