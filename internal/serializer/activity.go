package serializer

import (
	"ji/internal/model"
)

type Activity struct {
	ID             uint   `json:"id"`
	Title          string `json:"title"`
	Introduction   string `json:"introduction"`
	Status         int    `json:"status"`
	StartTime      int64
	EndTime        int64
	Location       Point  `json:"location"`
	ExpectedNumber uint   `json:"expected_number"`
	CurrentNumber  uint   `json:"current_Number"`
	UserId         uint   `json:"user_id"`
	UserName       string `json:"user_name"`
	UserAvatar     string `json:"user_avatar"`
}

// BuildUser 序列化用户
func BuildActivity(activity *model.Activity) *Activity {
	a := &Activity{
		ID:             activity.ID,
		Title:          activity.Title,
		Introduction:   activity.Introduction,
		Status:         activity.Status,
		StartTime:      activity.StartTime,
		EndTime:        activity.EndTime,
		Location:       Point(activity.Location),
		ExpectedNumber: activity.ExpectedNumber,
		CurrentNumber:  activity.CurrentNumber,
		UserId:         activity.UserId,
		UserName:       activity.UserName,
		UserAvatar:     activity.UserAvatar,
	}

	return a
}

func BuildActivitys(items []*model.Activity) (activitys []*Activity) {
	for _, item := range items {
		activity := BuildActivity(item)
		activitys = append(activitys, activity)
	}
	return activitys
}
