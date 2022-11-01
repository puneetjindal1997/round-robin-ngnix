package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/puneetjindal1997/round-robin-ngnix/route"
	"github.com/puneetjindal1997/round-robin-ngnix/server"
	"github.com/spf13/cobra"
)

// routes will be attached to s
var s server.Server

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "load balancing",
	Short: "API middleware",
	Long:  `load balancing using ngnix and later will do CI/CD with jenikins`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var env string
		// var ok bool
		// if env, ok = os.LookupEnv("ALPACA_ENV"); !ok {
		// 	env = "dev"
		// 	fmt.Printf("Run server in %s mode\n", env)
		// }
		err := s.Run(env)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(customRouteServices []route.ServicesI) {
	s.RouteServices = customRouteServices
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err, "sdjsbhfjbhsfb")
		os.Exit(1)
	}
}
