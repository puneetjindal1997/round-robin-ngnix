package main

import (
	"fmt"

	"github.com/puneetjindal1997/round-robin-ngnix/router"
)

func main() {
	router.New().
		WithRoutes(&MyServices{}).
		Run()
}

type MyServices struct{}

// SetupRoutes is our implementation of custom routes
func (s *MyServices) SetupRoutes() {
	fmt.Println("set up our custom routes!")
}
