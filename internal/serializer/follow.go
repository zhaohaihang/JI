package serializer

import (
	"ji/internal/model"
)

type RequestFollowInfo struct {
	FollowId uint `form:"followed_id" json:"followed_id" `
	Followed      int8 `form:"followded" json:"followded" binding:"lte=3"`
}

type Follow struct {
	UserId     uint `json:"user_id"`
	FollowId uint `json:"followed_id"`
	Followed      int8 `json:"followded"`
}

// BuildUser 序列化
func BuildFollow(follow *model.Follow) *Follow {
	return &Follow{
		UserId:     follow.UserId,
		FollowId: follow.FollowId,
		Followed:      follow.Followed,
	}
}