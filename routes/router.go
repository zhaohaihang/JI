package routes

import (
	"net/http"

	api "ji/api/v1"
	"ji/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "ji/docs"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("../static"))
	 
	v1 := r.Group("api/v1")
	{
		v1.GET("/swagger/*any", func(c *gin.Context) {
			ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "SWAGGER")(c)
		})
		// 用户操作
		v1.POST("user/login", api.UserLogin)

		authed := v1.Group("/") // 需要登陆保护
		authed.Use(middleware.JWT())
		{
			authed.PUT("user",api.UserUpdate)
			authed.GET("user/:uid",api.ViewUser)
			authed.POST("user/avatar",api.UploadUserAvatar)
		}
	}
	return r
}
