package dao

import (
	"context"
	"fmt"
	"ji/repository/db/model"

	"gorm.io/gorm"
)

type ActivityDao struct {
	*gorm.DB
}

func NewActivityDao(ctx context.Context) *ActivityDao {
	return &ActivityDao{NewDBClient(ctx)}
}

func (activityDao *ActivityDao) CreateActivity(activity *model.Activity) error {
	return activityDao.DB.Model(&model.Activity{}).Create(&activity).Error
}

func (activityDao *ActivityDao) GetActivityById(aId uint) (activity *model.Activity, err error) {
	err = activityDao.DB.Model(&model.Activity{}).Where("id=?", aId).First(&activity).Error
	return
}

func (activityDao *ActivityDao) UpdateActivityById(aId uint, activity *model.Activity) (err error) {
	err = activityDao.DB.Model(&model.Activity{}).Where("id=?", aId).Updates(&activity).Error
	return
}

func (activityDao *ActivityDao) DeleteActivityById(aId uint) error {
	return activityDao.DB.Where("id=?", aId).Delete(&model.Activity{}).Error
}

func (activityDao *ActivityDao) ListActivityByUserId(uId uint, page model.BasePage) (activitys []*model.Activity, total int64, err error) {

	err = activityDao.DB.Model(&model.Activity{}).Where("user_id=?", uId).Count(&total).Error
	if err != nil {
		return
	}

	err = activityDao.DB.Model(model.Activity{}).Where("user_id=?", uId).
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Find(&activitys).Error
	return
}

func (activityDao *ActivityDao) ListNearActivity(lat, lng float64, radius int) (activitys []*model.Activity, total int64, err error) {
	pointSql := fmt.Sprintf("POINT(%.6f %.6f),4326", lat, lng)

	err = activityDao.DB.Model(&model.Activity{}).Where("ST_Distance_Sphere(location, ST_PointFromText(?)) < ?", pointSql, radius).
		Count(&total).Error
	if err != nil {
		return
	}

	err = activityDao.DB.Model(&model.Activity{}).Where("ST_Distance_Sphere(location, ST_PointFromText(?)) < ?", pointSql, radius).
		Order("ST_Distance_Sphere(location, ST_PointFromText(" + pointSql + ")) DESC").Find(&activitys).Error

	return
}

func (activityDao *ActivityDao) SearchActivity(searchStr string, page model.BasePage) (activitys []*model.Activity, err error) {
	err = activityDao.DB.Model(&model.Activity{}).
		Where("title LIKE ? OR introduction LIKE ?", "%"+searchStr+"%", "%"+searchStr+"%").
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Find(&activitys).Error
	return
}
