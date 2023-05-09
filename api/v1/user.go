package v1

import (
	"ji/consts"
	"ji/service"
	"ji/pkg/utils"

	"github.com/gin-gonic/gin"
)

// UserLogin godoc
// @Summary 用户登录
// @Description  用户登录接口，如果用户不存在则创建用户
// @Tags user
// @Accept  json
// @Produce  json
// @Param post body service.LoginUserInfo true "user info"
// @Success 200 {object} serializer.Response{data=serializer.TokenData{user=serializer.User}}
// @Router /api/v1/user/login [post]
func UserLogin(c *gin.Context) {
	var userLoginService service.UserService
	var loginUserInfo service.LoginUserInfo
	if err := c.ShouldBind(&loginUserInfo); err == nil {
		res := userLoginService.Login(c.Request.Context(),loginUserInfo)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

func UserUpdate(c *gin.Context) {
	var userUpdateService service.UserService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdateService); err == nil {
		res := userUpdateService.UpdateUserById(c.Request.Context(), claims.UserID)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}