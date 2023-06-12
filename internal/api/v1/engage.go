package v1

import (
	"ji/internal/service"
	"ji/pkg/utils/tokenutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type EngageController struct {
	logger        *logrus.Logger
	engageService *service.EngageService
}

func NewEngageController(
	l *logrus.Logger,
	es *service.EngageService,
) *EngageController {
	return &EngageController{
		logger:        l,
		engageService: es,
	}
}

// EngageActivity godoc
// @Summary 加入活动
// @Description  加入活动
// @Tags engage
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization header parameter"
// @Param aid path int true "activity ID"
// @Success 200 {object} serializer.Response{}
// @Router /api/v1/engage/{aid} [post]
func (ec *EngageController) EngageActivity(c *gin.Context) {
	claims := tokenutil.GetTokenClaimsFromContext(c)
	aidStr := c.Param("aid")
	if aId, err := strconv.ParseUint(aidStr, 10, 32); err == nil {
		res := ec.engageService.EngageActivity(claims.UserID, uint(aId))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		ec.logger.Infoln(err)
	}
}

// CancelEngageActivity godoc
// @Summary 取消加入活动
// @Description  取消加入活动
// @Tags engage
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization header parameter"
// @Param aid path int true "activity ID"
// @Success 200 {object} serializer.Response{}
// @Router /api/v1/engage/{aid} [delete]
func (ec *EngageController) CancelEngageActivity(c *gin.Context) {
	claims := tokenutil.GetTokenClaimsFromContext(c)
	aidStr := c.Param("aid")
	if aId, err := strconv.ParseUint(aidStr, 10, 32); err == nil {
		res := ec.engageService.DelEngageActivity(claims.UserID, uint(aId))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		ec.logger.Infoln(err)
	}
}

// EngageUserList godoc
// @Summary 获取参加某一个活动的所有用户
// @Description  获取参加某一个活动的所有用户
// @Tags engage
// @Accept  json
// @Produce  json
// @Param aid path int true "activity ID"
// @Success 200 {object} serializer.Response{data=serializer.DataList{item=[]serializer.User}}
// @Router /api/v1/engage/activity/{aid}/user/list [get]
func (ec *EngageController)EngageUserList(c *gin.Context) {
	aidStr := c.Param("aid")
	if aId, err := strconv.ParseUint(aidStr, 10, 32); err == nil {
		res := ec.engageService.ListUsersByActivityId(uint(aId))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		ec.logger.Infoln(err)
	}
}

// EngageActivityList godoc
// @Summary 获取某一个用户参加的所有活动
// @Description  获取某一个用户参加的所有活动
// @Tags engage
// @Accept  json
// @Produce  json
// @Param uid path int true "user ID"
// @Success 200 {object} serializer.Response{data=serializer.DataList{item=[]serializer.Activity}}
// @Router /api/v1/engage/user/{uid}/activity/list [get]
func (ec *EngageController)EngageActivityList(c *gin.Context) {
	uidStr := c.Param("uid")
	if aId, err := strconv.ParseUint(uidStr, 10, 32); err == nil {
		res := ec.engageService.ListActivitysByUserId(uint(aId))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		ec.logger.Infoln(err)
	}
}