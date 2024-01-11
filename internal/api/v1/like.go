package v1

import (
	"ji/internal/serializer"
	"ji/internal/service"
	"ji/pkg/utils/tokenutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LikeController struct {
	logger      *logrus.Logger
	likeService *service.LikeService
}

func NewLikeContrller(l *logrus.Logger, ls *service.LikeService) *LikeController {
	return &LikeController{
		logger:      l,
		likeService: ls,
	}
}

// LikeActivity godoc
// @Summary 用户为活动点赞
// @Description  用户为活动点赞
// @Tags like
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization header parameter"
// @Param aid path int true "activity ID"
// @Success 200 {object} serializer.Response{}
// @Router /api/v1/like [post]
func (lc *LikeController) LikeActivity(c *gin.Context) {
	var likeInfo serializer.RequestLikeInfo
	claims := tokenutil.GetTokenClaimsFromContext(c)
	if err := c.ShouldBind(&likeInfo); err == nil {
		res := lc.likeService.UserLikeActivity(claims.UserID, likeInfo.AcitivtyId, likeInfo.Liked)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		lc.logger.Infoln(err)
	}
}
