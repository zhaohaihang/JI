package routes

import (
	"net/http"

	v1 "ji/internal/api/v1"

	"ji/pkg/cors"
	"ji/pkg/jwt"

	"ji/pkg/valid"

	_ "ji/docs"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(ac *v1.ActivityController, uc *v1.UserController) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Cors())
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

		authed := v1.Group("/") // 需要登陆保护
		authed.Use(jwt.JWT())
		{
			authed.PUT("user", uc.UserUpdate)
			authed.GET("user/:uid", uc.ViewUser)
			authed.POST("user/avatar", uc.UploadUserAvatar)

			authed.POST("activity", ac.CreateActivity)
		}
	}
	return r
}

var RouterProviderSet = wire.NewSet(NewRouter)
