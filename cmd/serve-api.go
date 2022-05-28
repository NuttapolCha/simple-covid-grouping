package cmd

import (
	"fmt"

	"github.com/NuttapolCha/simple-covid-grouping/api"
	"github.com/NuttapolCha/simple-covid-grouping/app"
	"github.com/NuttapolCha/simple-covid-grouping/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveApiCommand = &cobra.Command{
	Use:   "serve-api",
	Short: "Start API Server",
	RunE: func(cmd *cobra.Command, args []string) error {
		port := fmt.Sprintf(":%d", viper.GetInt("API.ServeHTTPPort"))

		logger, err := log.NewLogger()
		if err != nil {
			return err
		}
		logger.Debugf("logger initialized with level: %s", logger.Level())

		// initiailize application
		application, err := app.New(logger)
		if err != nil {
			return err
		}

		// initialize API router
		router := api.Init(application)
		return router.Run(port)
	},
}

func init() {
	rootCmd.AddCommand(serveApiCommand)
}
