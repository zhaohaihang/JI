package dao

import (
	"fmt"
	"ji/internal/model"
	"ji/pkg/consts"
	"ji/pkg/database"

	"gorm.io/gorm"
)

type ActivityDao struct {
	DB *gorm.DB
}

func NewActivityDao(db *database.Database) *ActivityDao {
	return &ActivityDao{
		DB: db.Mysql,
	}
}

func (ad *ActivityDao) CreateActivity(activity *model.Activity) error {
	return ad.DB.Model(&model.Activity{}).Create(&activity).Error
}

func (ad *ActivityDao) ExistOrNotByActivityId(aId uint) (activity *model.Activity, exist bool, err error) {
	var count int64
	err = ad.DB.Model(&model.Activity{}).Where("id = ?", aId).Count(&count).Error
	if count == 0 {
		return activity, false, err
	}
	err = ad.DB.Model(&model.Activity{}).Where("id =?", aId).First(&activity).Error
	if err != nil {
		return activity, false, err
	}
	return activity, true, nil
}

func (ad *ActivityDao) GetActivityById(aId uint) (activity *model.Activity, err error) {
	err = ad.DB.Model(&model.Activity{}).Where("id=?", aId).First(&activity).Error
	return
}

func (ad *ActivityDao) UpdateActivityById(aId uint, activity *model.Activity) (err error) {
	err = ad.DB.Model(&model.Activity{}).Where("id=?", aId).Updates(&activity).Error
	return
}

func (ad *ActivityDao) UpdateActivityCurrentNumById(aId uint, delta int64) (err error) {
	err = ad.DB.Model(&model.Activity{}).Where("id = ?", aId).
		UpdateColumn("current_number", gorm.Expr("current_number + ?", delta)).Error
	return
}

func (ad *ActivityDao) DeleteActivityById(aId uint) error {
	return ad.DB.Where("id=?", aId).Delete(&model.Activity{}).Error
}

func (ad *ActivityDao) ListActivityByUserId(uId uint, page model.BasePage) (activitys []*model.Activity, total int64, err error) {

	err = ad.DB.Model(&model.Activity{}).Where("user_id=?", uId).Count(&total).Error
	if err != nil {
		return
	}

	err = ad.DB.Model(model.Activity{}).Where("user_id=?", uId).
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Find(&activitys).Error
	return
}

func (ad *ActivityDao) ListNearActivity(lat, lng float64, radius int) (activitys []*model.Activity, total int64, err error) {
	pointSql := fmt.Sprintf("POINT(%.6f %.6f),4326", lat, lng)

	err = ad.DB.Model(&model.Activity{}).Where("ST_Distance_Sphere(location, ST_PointFromText(?)) < ?", pointSql, radius).
		Count(&total).Error
	if err != nil {
		return
	}

	err = ad.DB.Model(&model.Activity{}).Where("ST_Distance_Sphere(location, ST_PointFromText(?)) < ?", pointSql, radius).
		Order("ST_Distance_Sphere(location, ST_PointFromText(" + pointSql + ")) DESC").Find(&activitys).Error

	return
}

func (ad *ActivityDao) SearchActivity(searchStr string, page model.BasePage) (activitys []*model.Activity, err error) {
	err = ad.DB.Model(&model.Activity{}).
		Where("title LIKE ? OR introduction LIKE ?", "%"+searchStr+"%", "%"+searchStr+"%").
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Find(&activitys).Error
	return
}

//
func (ad *ActivityDao) UpdateActivityStatusFromNostartToInprocess(aId uint) (err error) {
	err = ad.DB.Model(&model.Activity{}).Where("id = ?", aId).
		Update("status", consts.ACTIVITY_STATUS_INPROCESS).Error
	return
}

func (ad *ActivityDao) UpdateActivityStatusFromInprocessToEnd(aId uint) (err error) {
	err = ad.DB.Model(&model.Activity{}).Where("id = ?", aId).
		Update("status", consts.ACTIVITY_STATUS_ENDED).Error
	return
}

func (ad *ActivityDao) ListActivitysByIds(aIds []uint) (activitys []*model.Activity, err error) {
	err = ad.DB.Model(&model.Activity{}).Find(&activitys, aIds).Error
	return
}
