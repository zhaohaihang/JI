package v1

import (
	"ji/internal/service"

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
