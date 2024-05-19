package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"product/config"
	"product/internal/db"
)

var healthCheckCmd = &cobra.Command{
	Use:   "health",
	Short: "Check health of application",

	Run: checkHealth,
}

func init() {
	rootCmd.AddCommand(healthCheckCmd)
}

func checkHealth(cmd *cobra.Command, args []string) {
	hasError := false
	err := checkMysqlHealth()
	if err != nil {
		hasError = true
		log.Printf("Mysql healthcheck failed:%v\n", err)
	}
	log.Println("Mysql healthcheck passed!")

	if hasError {
		return

	}

	log.Printf("OK.\n")
}

func checkMysqlHealth() error {
	conn, err := db.Connect(config.AppConfig)
	if err != nil {
		return err
	}
	if err := conn.Close(); err != nil {
		return err
	}
	return nil
}
