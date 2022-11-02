package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Service represents the account http service
type Service struct{}

func Router(r *gin.RouterGroup) {
	a := Service{}
	pr := r.Group("/profile")
	pr.GET("", a.profile)

}

func (s *Service) profile(c *gin.Context) {
	fmt.Println("service 1")
}
