package v1

import (
	"ji/internal/serializer"
	"ji/internal/service"
	"ji/pkg/logger"
	"ji/pkg/utils/tokenutil.go"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type ActivityController struct {
	log             *logger.Logger
	activityService *service.ActivityService
	userService     *service.UserService
}

func NewActivityContrller(log *logger.Logger, as *service.ActivityService, us *service.UserService) *ActivityController {
	return &ActivityController{
		log:             log,
		activityService: as,
		userService:     us,
	}
}

var ActivityControllerProviderSet = wire.NewSet(NewActivityContrller)

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
func (ac *ActivityController) CreateActivity(c *gin.Context) {
	var createActivityInfo serializer.CreateActivityInfo
	claims := tokenutil.GetTokenClaimsFromContext(c)
	if err := c.ShouldBind(&createActivityInfo); err == nil {
		res := ac.activityService.CreateActivity(c.Request.Context(), claims.UserID, createActivityInfo)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		ac.log.Logrus.Infoln(err)
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
func (ac *ActivityController) ShowActivity(c *gin.Context) {
	aIdStr := c.Param("aid")
	if aId, err := strconv.ParseUint(aIdStr, 10, 32); err == nil {
		res := ac.activityService.GetActivityById(c.Request.Context(), uint(aId))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		ac.log.Logrus.Infoln(err)
	}
}

// ListNearActivity godoc
// @Summary 查看指定点周围的所有活动
// @Description  根据位置和半径查看活动接口
// @Tags activity
// @Accept  json
// @Produce  json
// @Param lat query int true "lat"
// @Param lng query int true "lng"
// @Param rad query int true "rad"
// @Success 200 {object} serializer.Response{data=serializer.DataList{item=[]serializer.Activity}}
// @Router /api/v1/activity/near [get]
func (ac *ActivityController) ListNearActivity(c *gin.Context) {
	var nearInfo serializer.NearInfo
	if err := c.ShouldBindQuery(nearInfo); err == nil {
		ac.activityService.ListNearActivity(c.Request.Context(), nearInfo)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		ac.log.Logrus.Infoln(err)
	}
}
