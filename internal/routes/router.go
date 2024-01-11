package routes

import (
	"net/http"

	v1 "ji/internal/api/v1"

	"ji/pkg/middleware"

	"ji/internal/valid"

	_ "ji/docs"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(
	ac *v1.ActivityController, 
	uc *v1.UserController, 
	ec *v1.EngageController,
	lc *v1.LikeController,
	) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("../static"))

	valid.Init()

	v1 := r.Group("api/v1")
	{
		v1.GET("/swagger/*any", func(c *gin.Context) {
			ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "SWAGGER")(c)
		})

		v1.POST("user/login", uc.UserLogin)
		v1.GET("user/:uid/activitys", ac.ListUserActivity)

		v1.GET("activity/:aid", ac.ShowActivity)
		v1.GET("activity/near", ac.ListNearActivity)
		v1.GET("user/:uid", uc.ViewUser)
		v1.GET("engage/activity/:aid/user/list", ec.EngageUserList)
		v1.GET("engage/user/:uid/activity/list", ec.EngageActivityList)

		authed := v1.Group("/") // 需要登陆保护
		authed.Use(middleware.JWT())
		{
			authed.PUT("user", uc.UserUpdate)
			authed.PUT("user/changepasswd", uc.ChangePasswd)
			authed.POST("user/avatar", uc.UploadUserAvatar)

			authed.POST("activity", ac.CreateActivity)
			authed.PUT("activity/cover", ac.UploadActivityCover)

			authed.PUT("activity/:aid", ac.UpdateActivity)
			authed.DELETE("activity/cover", ac.DeleteActivity)

			authed.POST("engage/:aid", ec.EngageActivity)
			authed.DELETE("engage/:aid", ec.CancelEngageActivity)

			authed.POST("/api/v1/like",lc.LikeActivity)
		}
	}
	return r
}

var RouterProviderSet = wire.NewSet(NewRouter)
