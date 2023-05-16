package v1

import (
	"ji/consts"
	"ji/pkg/utils"
	"ji/serializer"
	"ji/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateActivity godoc
// @Summary 创建活动
// @Description  创建活动接口
// @Tags activity
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization header parameter"
// @Param CreateActivityInfo body serializer.CreateActivityInfo true "activity create info"
// @Success 200 {object} serializer.Response{data=serializer.Activity}
// @Router /api/v1/activity [post]
func CreateActivity(c *gin.Context) {
	var createActivityService service.ActivityService
	var createActivityInfo serializer.CreateActivityInfo
	claims := utils.GetClaimsFromContext(c)
	if err := c.ShouldBind(&createActivityInfo); err != nil {
		res := createActivityService.CreateActivity(c.Request.Context(), claims.UserID, createActivityInfo)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

// ShowActivity godoc
// @Summary 查看活动详情
// @Description  查看用户详情接口
// @Tags activity
// @Accept  json
// @Produce  json
// @Param aid path int true "activity ID"
// @Success 200 {object} serializer.Response{data=serializer.Activity}
// @Router /api/v1/activity/{aid} [get]
func ShowActivity(c *gin.Context) {
	var showActivityService service.ActivityService
	aIdStr := c.Param("aid")
	if aId, err := strconv.ParseUint(aIdStr, 10, 32); err == nil {
		res := showActivityService.GetActivityById(c.Request.Context(), uint(aId))
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

// ListActivity godoc
// @Summary 查看指定用户创建的所有活动
// @Description  根据用户ID查看活动接口
// @Tags activity
// @Accept  json
// @Produce  json
// @Param page_num query int true "page num"
// @Param page_size query int true "page size"
// @Param uid path int true "user ID"
// @Success 200 {object} serializer.Response{data=serializer.DataList{item=[]serializer.Activity}}
// @Router /api/v1/user/{uid}/activity [get]
func ListUserActivity(c *gin.Context) {
	var listUserActivityService service.ActivityService
	var basePage serializer.BasePage
	uIdStr := c.Param("uid")
	c.ShouldBindQuery(basePage)

	if uId, err := strconv.ParseUint(uIdStr, 10, 32); err == nil {
		res := listUserActivityService.ListActivityByUserId(c.Request.Context(), uint(uId), basePage)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}
