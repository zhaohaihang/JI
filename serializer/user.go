package serializer

import (
	"ji/config"
	"ji/repository/db/model"
	"time"
	// geo "github.com/kellydunn/golang-geo"
)

type User struct {
	ID        uint        `json:"id"`
	UserName  string      `json:"username"`
	Email     string      `json:"email"`
	Status    string      `json:"status"`
	Avatar    string      `json:"avatar"`
	LastLogin time.Time   `json:"last_login"`
	Address   string      `json:"address"`
	Biography string      `json:"biography"`
	Phone     string      `json:"phone"`
	Location  model.Point `json:"location"`
	Extra     string      `json:"xxtra"`
}

// BuildUser 序列化用户
func BuildUser(user *model.User) *User {
	u := &User{
		ID:        user.ID,
		UserName:  user.UserName,  
		Email:     user.Email,
		Status:    user.Status,
		LastLogin: user.LastLogin,
		Avatar:    config.Conf.Static.StaticHost + config.Conf.Static.StaticPort + config.Conf.Static.AvatarPath + user.AvatarURL(),
		Address:   user.Address,
		Biography: user.Address,
		Phone:     user.Phone,
		Location:  user.Location,
		Extra: 	   user.Extra,
	}

	return u
}

func BuildUsers(items []*model.User) (users []*User) {
	for _, item := range items {
		user := BuildUser(item)
		users = append(users, user)
	}
	return users
}
