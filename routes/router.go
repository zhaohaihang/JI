package routes

import (
	"net/http"

	api "ji/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.StaticFS("/static", http.Dir("./static"))

	v1 := r.Group("api/v1")
	{
		// 用户操作
		v1.POST("user/login", api.UserLogin)
	}
	return r
}
