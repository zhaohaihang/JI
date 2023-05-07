package model

import (
	"time"

	"github.com/DanPlayer/randomname"
	geo "github.com/kellydunn/golang-geo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//User 用户模型
type User struct {
	gorm.Model
	UserName		string `gorm:"unique"`
	PasswordDigest	string
	Biography		string `gorm:"size:1000"`
	Address			string 
	Email			string
	Phone			string 
	NickName		string
	Status			string
	Avatar			string `gorm:"size:1000"`
	LastLogin		time.Time
	Latitude		float64
	Lngitude		float64
	Location		*geo.Point `gorm:"type:point"`
}

const (
	PassWordCost = 12       //密码加密难度
	Active  = "active" //激活用户
)

func (user *User) BeforeSave(db *gorm.DB) error {
	user.LastLogin = time.Now()
	return nil
}

//SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

//CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

func (user *User) GenerateRandomNickName() {
	user.NickName = randomname.GenerateName()
}

// AvatarUrl 头像地址
func (user *User) AvatarURL() string {
	signedGetURL := user.Avatar
	return signedGetURL
}
