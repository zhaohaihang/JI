package loading

import (
	"ji/repository/db/dao"
	"ji/pkg/valid"
)

func Init() {
	valid.Init()
	dao.InitMySQL()
}
