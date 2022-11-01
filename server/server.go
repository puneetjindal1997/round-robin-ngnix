package server

import (
	"os"

	"github.com/gin-gonic/gin"
	mw "github.com/puneetjindal1997/round-robin-ngnix/middleware"
	"github.com/puneetjindal1997/round-robin-ngnix/route"
	"go.uber.org/zap"
)

// Server holds all the routes and their services
type Server struct {
	RouteServices []route.ServicesI
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Run runs our API server
func (server *Server) Run(env string) error {

	r := gin.Default()

	// middleware
	mw.Add(r, CORSMiddleware())
	log, _ := zap.NewDevelopment()
	defer log.Sync()

	// setup default routes
	rsDefault := &route.Services{
		Log:    log,
		R:      r}
	rsDefault.SetupV1Routes()

	// setup all custom/user-defined route services
	for _, rs := range server.RouteServices {
		rs.SetupRoutes()
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	// run with port from config
	return r.Run(":" + port)
}
