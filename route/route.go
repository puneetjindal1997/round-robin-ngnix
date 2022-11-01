package route

import (
	"github.com/alpacahq/ribbit-backend/service"

	"github.com/gin-gonic/gin" // gin-swagger middleware
	// swagger embed files
	"go.uber.org/zap"
)

// NewServices creates a new router services
func NewServices(Log *zap.Logger, R *gin.Engine) *Services {
	return &Services{Log, R}
}

// Services lets us bind specific services when setting up routes
type Services struct {
	Log *zap.Logger
	R   *gin.Engine
}

// SetupV1Routes instances various repos and services and sets up the routers
func (s *Services) SetupV1Routes() {

	// prefixed with /v1 and protected by jwt
	v1Router := s.R.Group("/v1")
	service.AccountRouter(accountService, v1Router)
}
