package serializer

import (
	"ji/internal/model"
)

type RequestLikeInfo struct {
	AcitivtyId uint `form:"activity_id" json:"activity_id" `
	Liked      int8 `form:"liked" json:"liked" binding:"lte=3"`
}

type Like struct {
	UserId     uint `json:"user_id"`
	ActivityId uint `json:"activity_id"`
	Liked      int8 `json:"liked"`
}

// BuildUser 序列化用户
func BuildLike(like *model.Like) *Like {
	l := &Like{
		UserId:     like.UserId,
		ActivityId: like.AcitivtyId,
		Liked:      like.Liked,
	}
	return l
}
