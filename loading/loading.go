package loading

import "ji/repository/db/dao"

func Init() {
	dao.InitMySQL()
}
