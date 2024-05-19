package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"user/config"
	"user/internal/app"
)

var serveCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts server and runs application",

	Run: startServer,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func startServer(cmd *cobra.Command, args []string) {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = app.RunServer(config.AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
