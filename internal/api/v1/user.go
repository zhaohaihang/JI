package v1

import (
	"ji/pkg/logger"
	"ji/pkg/utils/tokenutil.go"

	"ji/internal/serializer"
	"ji/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type UserController struct {
	log             *logger.Logger
	activityService *service.ActivityService
	userService     *service.UserService
}

func NewUserContrller(log *logger.Logger, as *service.ActivityService, us *service.UserService) *UserController {
	return &UserController{
		log:             log,
		activityService: as,
		userService:     us,
	}
}

var UserControllerProviderSet = wire.NewSet(NewUserContrller)

// UserLogin godoc
// @Summary 用户登录
// @Description  用户登录接口，如果用户不存在则创建用户
// @Tags user
// @Accept  json
// @Produce  json
// @Param LoginUserInfo body serializer.LoginUserInfo true "login user info"
// @Success 200 {object} serializer.Response{data=serializer.TokenData{user=serializer.User}}
// @Router /api/v1/user/login [post]
func (uc *UserController) UserLogin(c *gin.Context) {
	var loginUserInfo serializer.LoginUserInfo
	if err := c.ShouldBind(&loginUserInfo); err == nil {
		res := uc.userService.Login(c.Request.Context(), loginUserInfo)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		uc.log.Logrus.Infoln(err)
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
func (uc *UserController) UserUpdate(c *gin.Context) {
	var updateUserInfo serializer.UpdateUserInfo
	claims := tokenutil.GetTokenClaimsFromContext(c)
	if err := c.ShouldBind(&updateUserInfo); err == nil {
		res := uc.userService.UpdateUserById(c.Request.Context(), claims.UserID, updateUserInfo)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		uc.log.Logrus.Infoln(err)
	}
}

// ViewUser godoc
// @Summary 查看用户信息
// @Description  查看用户信息接口
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization header parameter"
// @Param uid path int true "user ID"
// @Success 200 {object} serializer.Response{data=serializer.User}
// @Router /api/v1/user/{uid} [get]
func (uc *UserController) ViewUser(c *gin.Context) {
	uIdStr := c.Param("uid")

	if uId, err := strconv.ParseUint(uIdStr, 10, 32); err == nil {
		res := uc.userService.GetUserById(c.Request.Context(), uint(uId))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		uc.log.Logrus.Infoln(err)
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
func (uc *UserController) UploadUserAvatar(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		uc.log.Logrus.Infoln(err)
	} else {
		fileSize := fileHeader.Size
		claims := tokenutil.GetTokenClaimsFromContext(c)
		res := uc.userService.UploadUserAvatar(c.Request.Context(), claims.UserID, file, fileSize)
		c.JSON(http.StatusOK, res)
	}
}

// ListUserActivity godoc
// @Summary 查看指定用户创建的所有活动
// @Description  根据用户ID查看该用户创建的活动
// @Tags user
// @Accept  json
// @Produce  json
// @Param page_num query int true "page num"
// @Param page_size query int true "page size"
// @Param uid path int true "user ID"
// @Success 200 {object} serializer.Response{data=serializer.DataList{item=[]serializer.Activity}}
// @Router /api/v1/user/{uid}/activity [get]
func (uc *ActivityController) ListUserActivity(c *gin.Context) {
	var basePage serializer.BasePage
	uIdStr := c.Param("uid")
	c.ShouldBindQuery(basePage)
	if uId, err := strconv.ParseUint(uIdStr, 10, 32); err == nil {
		res := uc.activityService.ListActivityByUserId(c.Request.Context(), uint(uId), basePage)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		uc.log.Logrus.Infoln(err)
	}
}
