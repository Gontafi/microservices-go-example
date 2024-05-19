package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"orders/config"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "orders application",
}

func Execute() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not configure application: %s\n", err.Error())
	}

	err = rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
