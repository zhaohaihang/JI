package routes

import (
	"net/http"

	api "ji/api/v1"
	"ji/middleware"

	_ "ji/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

		v1.POST("user/login", api.UserLogin)
		v1.GET("user/:uid/activitys", api.ListUserActivity)

		v1.GET("activity/:aid", api.ShowActivity)
		v1.GET("activity/near", api.ListNearActivity)

		authed := v1.Group("/") // 需要登陆保护
		authed.Use(middleware.JWT())
		{
			authed.PUT("user", api.UserUpdate)
			authed.GET("user/:uid", api.ViewUser)
			authed.POST("user/avatar", api.UploadUserAvatar)

			authed.POST("activity", api.CreateActivity)
		}
	}
	return r
}
