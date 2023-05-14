package v1

import (
	"ji/consts"
	"ji/pkg/utils"
	"ji/serializer"
	"ji/service"

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
