package loading

import (
	"ji/repository/db/dao"
	"ji/valid"
)

func Init() {
	valid.Init()
	dao.InitMySQL()
}
