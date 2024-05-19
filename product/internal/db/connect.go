package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"product/config"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	dbConn, err := sql.Open(cfg.DBDriver, fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	),
	)
	log.Println(fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	))
	if err != nil {
		return nil, err
	}

	dbConn.SetMaxOpenConns(cfg.DB_MAX_CONN)
	dbConn.SetMaxIdleConns(cfg.DB_MAX_IDLE_CONN)
	dbConn.SetConnMaxIdleTime(cfg.DB_MAX_IDLE_TIME)

	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}

	return dbConn, err
}
