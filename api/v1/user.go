package v1

import (
	"ji/consts"
	"ji/pkg/utils"
	"ji/serializer"
	"ji/service"

	"github.com/gin-gonic/gin"
)

// UserLogin godoc
// @Summary 用户登录
// @Description  用户登录接口，如果用户不存在则创建用户
// @Tags user
// @Accept  json
// @Produce  json
// @Param post body serializer.LoginUserInfo true "user info"
// @Success 200 {object} serializer.Response{data=serializer.TokenData{user=serializer.User}}
// @Router /api/v1/user/login [post]
func UserLogin(c *gin.Context) {
	var userLoginService service.UserService
	var loginUserInfo serializer.LoginUserInfo
	if err := c.ShouldBind(&loginUserInfo); err == nil {
		res := userLoginService.Login(c.Request.Context(),loginUserInfo)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

// UserUpdate godoc
// @Summary 用户更新信息
// @Description  用户更新信息接口
// @Tags user
// @Accept  json
// @Produce  json
// @Param id   path  int  true  "user id"
// @Param post body serializer.UpdateUserInfo true "user update info"
// @Success 200 {object} serializer.Response{data=serializer.User}
// @Router /api/v1/user/{id} [put]
func UserUpdate(c *gin.Context) {
	var userUpdateService service.UserService
	// var updateUserInfo serializer.UpdateUserInfo
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdateService); err == nil {
		res := userUpdateService.UpdateUserById(c.Request.Context(), claims.UserID)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}