package cmd

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/spf13/cobra"
	"log"
	"user/config"
	"user/internal/db"
)

var version int

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Manipulate database migrations",
	Long: `Manipulate with database migrations with options" +
		Usage example migrate COMMAND [args...]
	
	Commands:
  	up       Apply all up migrations
  	down -v [N]   Apply all or N down migrations
  	drop         Drop everything inside database
  	force -v [N]  Set version V but don't run migration (ignores dirty state)
  	version      Print current migration version`,
	Args: cobra.ExactArgs(1),
	Run:  migrateDatabase,
}

func init() {
	migrateCmd.Flags().IntVarP(&version, "version", "v", 0, "Version of migration")
	rootCmd.AddCommand(migrateCmd)
}

func migrateDatabase(cmd *cobra.Command, args []string) {

	dbConn, err := db.Connect(config.AppConfig, sql.Open)
	if err != nil {
		log.Println(config.AppConfig.DBHost, config.AppConfig.DBPassword, config.AppConfig.DBPort, config.AppConfig.DBUser)
		log.Printf("Could not connect to db: %s\n", err.Error())
		return
	}

	defer dbConn.Close()
	m, err := NewMigrator(dbConn)
	if err != nil {
		log.Printf("Could not create instance of migrator: %s\n", err.Error())
		return
	}

	defer func() {
		if err := dbConn.Close(); err != nil {
			log.Printf("mysql conncetion close error: %s", err.Error())
		}
	}()

	switch args[0] {

	case "up":
		err := m.Up()
		if err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Printf("up migration error error: %s\n", err)
			return
		}
		log.Printf("up migration finished.\n")
	case "down":
		if version == 0 {
			err := m.Down()
			if err != nil {
				log.Printf("down migration error: %s\n", err)
			} else {
				log.Println("full down migration finished.")
			}
			return
		}

		err := m.Steps(-1 * version)
		if err != nil {
			log.Printf("down migration error: %s\n", err)
			return
		}
		log.Printf("down migration finished.\n")

	case "drop":
		err := m.Drop()
		if err != nil {
			log.Printf("drop migration error: %s\n", err)
		}
		log.Printf("drop migration finished \n")
	case "force":
		if version == 0 {
			log.Println("Please specify the version")
			return
		}
		err := m.Force(version)
		if err != nil {
			log.Printf("force migration error: %s\n", err)
			return
		}
		log.Printf("force migration finished \n")

	case "version":
		ver, _, _ := m.Version()
		log.Printf("Current db version: %d\n", ver)
		return
	}

}

func NewMigrator(conn *sql.DB) (*migrate.Migrate, error) {
	driver, err := mysql.WithInstance(conn, &mysql.Config{})
	if err != nil {
		return nil, err
	}
	return migrate.NewWithDatabaseInstance(
		"file://internal/migrations",
		config.AppConfig.DBName, driver)
}
