package v1

import (
	"ji/internal/serializer"
	"ji/internal/service"
	"ji/pkg/utils/tokenutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type FollowController struct {
	logger        *logrus.Logger
	followService *service.FollowService
}

func NewFollowContrller(l *logrus.Logger, fs *service.FollowService) *FollowController {
	return &FollowController{
		logger:        l,
		followService: fs,
	}
}

// UserFollow godoc
// @Summary 用户关注或取消关注
// @Description  用户关注其他用户
// @Tags follow
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization header parameter"
// @Param RequestFollowInfo body serializer.RequestFollowInfo true "user follow info"
// @Success 200 {object} serializer.Response{}
// @Router /api/v1/follow [post]
func (fc FollowController) UserFollow(c *gin.Context) {
	var followInfo serializer.RequestFollowInfo
	claims := tokenutil.GetTokenClaimsFromContext(c)
	if err := c.ShouldBind(&followInfo); err == nil {
		res := fc.followService.UserFollow(claims.UserID, followInfo.FollowId,followInfo.Followed)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		fc.logger.Infoln(err)
	}
}
