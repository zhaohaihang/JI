package serializer

import (
	"ji/config"
	"ji/repository/db/model"
	"time"
	// geo "github.com/kellydunn/golang-geo"
)

type User struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"user_name"`
	NickName  string    `json:"nickname"`
	Type      int       `json:"type"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	Avatar    string    `json:"avatar"`
	LastLogin time.Time `json:"last_login"`
	CreateAt  int64     `json:"create_at"`
	Address   string    `json:"address"`
	Biography string    `json:"biography"`
	Phone     string    `json:"phone"`
	Location  model.Point     `json:"location"`
}

// BuildUser 序列化用户
func BuildUser(user *model.User) *User {
	u := &User{
		ID:        user.ID,
		UserName:  user.UserName,
		NickName:  user.NickName,
		Email:     user.Email,
		Status:    user.Status,
		LastLogin: user.LastLogin,
		Avatar:    config.Conf.Static.StaticHost + config.Conf.Static.StaticPort + config.Conf.Static.AvatarPath + user.AvatarURL(),
		CreateAt:  user.CreatedAt.Unix(),
		Address:   user.Address,
		Biography: user.Address,
		Phone:     user.Phone,
		Location:  user.Location,
	}

	// if conf.UploadModel == consts.UploadModelOss {
	// 	u.Avatar = user.Avatar
	// }

	return u
}

func BuildUsers(items []*model.User) (users []*User) {
	for _, item := range items {
		user := BuildUser(item)
		users = append(users, user)
	}
	return users
}
