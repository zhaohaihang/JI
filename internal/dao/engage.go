package dao

import (
	"ji/internal/model"
	"ji/pkg/database"

	"gorm.io/gorm"
)

type EngageDao struct {
	DB *gorm.DB
}

func NewEngageDao(db *database.Database) *EngageDao {
	return &EngageDao{
		DB: db.Mysql,
	}
}

func (ed *EngageDao) CreateEngage(engage *model.Engage) (err error) {
	return ed.DB.Model(&model.Engage{}).Create(&engage).Error
}

func (ed *EngageDao) ListActivityIdsByUserId(uId uint) ([]uint, int64, error) {

	var aIds []uint
	var count int64
	result := ed.DB.Model(&model.Engage{}).Where("user_id = ?", uId).Pluck("activity_id", &aIds)
	count = result.RowsAffected

	if nil != result.Error {
		return nil, 0, result.Error
	}

	return aIds, count, nil
}

func (ed *EngageDao) ListUserIdsByActivityId(aId uint) ([]uint, int64, error) {

	var uIds []uint
	var count int64
	result := ed.DB.Model(&model.Engage{}).Where("activity_id = ?", aId).Pluck("user_id", &uIds)
	count = result.RowsAffected

	if nil != result.Error {
		return nil, 0, result.Error
	}

	return uIds, count, nil
}

func (ed *EngageDao) DelEngageByIds(uId uint, aId uint) error {
	err := ed.DB.Model(&model.Engage{}).Where("user_id = ? and activity_id = ?", uId, aId).
		Delete(&model.Engage{}).Error
	return err
}
