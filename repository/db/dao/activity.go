package dao

import (
	"context"
	"ji/repository/db/model"

	"gorm.io/gorm"
)

type ActivityDao struct {
	*gorm.DB
}

func NewActivityDao(ctx context.Context) *ActivityDao {
	return &ActivityDao{NewDBClient(ctx)}
}

func (activityDao *ActivityDao)CreateActivity(activity *model.Activity) error {
	return activityDao.DB.Model(&model.Activity{}).Create(&activity).Error
}

func (activityDao *ActivityDao)GetActivityById(aId uint)(activity *model.Activity,err error){
	err  = activityDao.DB.Model(&model.Activity{}).Where("id=?", aId).First(&activity).Error
	return 
}

func (activityDao *ActivityDao) UpdateActivityById(aId uint,activity *model.Activity) (err error) {
	err = activityDao.DB.Model(&model.Activity{}).Where("id=?",aId).Updates(&activity).Error
	return 
}

func(activityDao *ActivityDao) DeleteActivityById(aId uint)error {
	return activityDao.DB.Where("id=?", aId).Delete(&model.Activity{}).Error
}

// func (activityDao *ActivityDao)ListActivityByUserId(uId uint ,pageSize ,pageNum int ) (activitys []*model.Activity, total int64, err error) {
	
// 	err = activityDao.DB.Model(&model.Activity{}).Preload("User").Where("user_id=?", uId).Count(&total).Error
// 	if err != nil {
// 		return
// 	}
	
// 	err = activityDao.DB.Model(model.Activity{}).Preload("User").Where("user_id=?", uId).
// 		Offset((pageNum - 1) * pageSize).
// 		Limit(pageSize).Find(&activitys).Error
// 	return
// }

func (activityDao *ActivityDao) SearchActivity(searchStr string, page model.BasePage) (activitys []*model.Activity, err error) {
	err = activityDao.DB.Model(&model.Activity{}).
		Where("title LIKE ? OR introduction LIKE ?", "%"+searchStr+"%", "%"+searchStr+"%").
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Find(&activitys).Error
	return
}
