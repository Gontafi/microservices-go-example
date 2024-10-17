package db

import (
	"database/sql"
	"fmt"
	"user/config"
)

type DBConnector func(driverName, dataSourceName string) (*sql.DB, error)

func Connect(cfg *config.Config, open DBConnector) (*sql.DB, error) {
	dbConn, err := open(cfg.DBDriver, fmt.Sprintf(
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

	return dbConn, nil
}
