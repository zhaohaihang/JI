package v1

import (
	"ji/consts"
	"ji/pkg/utils"
	"ji/serializer"
	"ji/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserLogin godoc
// @Summary 用户登录
// @Description  用户登录接口，如果用户不存在则创建用户
// @Tags user
// @Accept  json
// @Produce  json
// @Param LoginUserInfo body serializer.LoginUserInfo true "login user info"
// @Success 200 {object} serializer.Response{data=serializer.TokenData{user=serializer.User}}
// @Router /api/v1/user/login [post]
func UserLogin(c *gin.Context) {
	var userLoginService service.UserService
	var loginUserInfo serializer.LoginUserInfo
	if err := c.ShouldBind(&loginUserInfo); err == nil {
		res := userLoginService.Login(c.Request.Context(), loginUserInfo)
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
// @Param Authorization header string true "Authorization header parameter"
// @Param UpdateUserInfo body serializer.UpdateUserInfo true "user update info"
// @Success 200 {object} serializer.Response{data=serializer.User}
// @Router /api/v1/user [put]
func UserUpdate(c *gin.Context) {
	var userUpdateService service.UserService
	var updateUserInfo serializer.UpdateUserInfo
	claims := utils.GetClaimsFromContext(c)
	if err := c.ShouldBind(&updateUserInfo); err == nil {
		res := userUpdateService.UpdateUserById(c.Request.Context(), claims.UserID, updateUserInfo)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

// ViewUser godoc
// @Summary 查看用户信息
// @Description  查看用户信息接口
// @Tags user
// @Accept  json
// @Produce  json
// @Param uid path int true "user ID"
// @Success 200 {object} serializer.Response{data=serializer.User}
// @Router /api/v1/user/{uid} [get]
func ViewUser(c *gin.Context) {
	var viewUserService service.UserService
	uIdStr := c.Param("uid")

	if uId, err := strconv.ParseUint(uIdStr, 10, 32); err == nil {
		res := viewUserService.GetUserById(c.Request.Context(), uint(uId))
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

// UploadUserAvatar godoc
// @Summary 上传用户头像
// @Description  上传用户头像接口
// @Tags user
// @Accept  multipart/form-data
// @Produce  json
// @Param file formData file true "图片文件"
// @Param Authorization header string true "Authorization header parameter"
// @Success 200 {object} serializer.Response{data=serializer.User}
// @Router /api/v1/user/avatar [post]
func UploadUserAvatar(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}else {
		fileSize := fileHeader.Size
		var uploadUserAvatarService  service.UserService
		claims := utils.GetClaimsFromContext(c)
		res := uploadUserAvatarService.UploadUserAvatar(c.Request.Context(), claims.UserID, file, fileSize)
		c.JSON(consts.StatusOK, res)
	}
}