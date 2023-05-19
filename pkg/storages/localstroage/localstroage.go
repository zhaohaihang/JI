package localstroage

import (
	"io/ioutil"
	"ji/config"
	"log"
	"mime/multipart"
	"os"
	"strconv"

	"github.com/google/wire"
)

type LocalStroage struct {
	config *config.Config
}

func NewLocalStroage(c *config.Config) *LocalStroage {
	return &LocalStroage{
		config: c,
	}
}

var LocalStroageProviderSet = wire.NewSet(NewLocalStroage)

// UploadAvatarToLocalStatic 上传头像
func (l *LocalStroage)UploadAvatarToLocalStatic(file multipart.File, userId uint, userName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId))
	basePath := "." + l.config.Static.AvatarPath + "user" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + userName + ".jpg"
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return "", err
	}
	return "user" + bId + "/" + userName + ".jpg", err
}

// DirExistOrNot 判断文件是否存在
func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir()
}

// CreateDir 创建文件夹
func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 7550)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
