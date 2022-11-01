package router

import (
	"github.com/puneetjindal1997/round-robin-ngnix/cmd"
	"github.com/puneetjindal1997/round-robin-ngnix/route"
)

// New creates a new Alpaca instance
func New() *Routes {
	return &Routes{}
}

// Alpaca allows us to specify customizations, such as custom route services
type Routes struct {
	RouteServices []route.ServicesI
}

// WithRoutes is the builder method for us to add in custom route services
func (g *Routes) WithRoutes(RouteServices ...route.ServicesI) *Routes {
	return &Routes{RouteServices}
}

// Run executes our alpaca functions or servers
func (g *Routes) Run() {
	cmd.Execute(g.RouteServices)
}
