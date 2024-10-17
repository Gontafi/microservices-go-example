package db

import (
	"database/sql"
	"fmt"
	"testing"
	"user/config"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestConnect_Success(t *testing.T) {
	cfg := &config.Config{
		DBDriver:         "mysql",
		DBUser:           "testuser",
		DBPassword:       "testpass",
		DBHost:           "localhost",
		DBPort:           "3306",
		DBName:           "testdb",
		DB_MAX_CONN:      10,
		DB_MAX_IDLE_CONN: 5,
		DB_MAX_IDLE_TIME: 100,
	}

	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	assert.NoError(t, err)
	defer db.Close()

	open := func(driverName, dataSourceName string) (*sql.DB, error) {
		return db, nil
	}

	mock.ExpectPing()

	conn, err := Connect(cfg, open)
	assert.NoError(t, err)
	assert.NotNil(t, conn)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestConnect_FailToOpen(t *testing.T) {
	cfg := &config.Config{
		DBDriver:         "mysql",
		DBUser:           "testuser",
		DBPassword:       "testpass",
		DBHost:           "localhost",
		DBPort:           "3306",
		DBName:           "testdb",
		DB_MAX_CONN:      10,
		DB_MAX_IDLE_CONN: 5,
		DB_MAX_IDLE_TIME: 100,
	}

	expectedErr := fmt.Errorf("failed to open database")

	open := func(driverName, dataSourceName string) (*sql.DB, error) {
		return nil, expectedErr
	}

	conn, err := Connect(cfg, open)
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Nil(t, conn)
}

func TestConnect_FailToPing(t *testing.T) {
	cfg := &config.Config{
		DBDriver:         "mysql",
		DBUser:           "testuser",
		DBPassword:       "testpass",
		DBHost:           "localhost",
		DBPort:           "3306",
		DBName:           "testdb",
		DB_MAX_CONN:      10,
		DB_MAX_IDLE_CONN: 5,
		DB_MAX_IDLE_TIME: 100,
	}

	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	assert.NoError(t, err)
	defer db.Close()

	open := func(driverName, dataSourceName string) (*sql.DB, error) {
		return db, nil
	}

	mock.ExpectPing().WillReturnError(fmt.Errorf("ping failed"))

	conn, err := Connect(cfg, open)
	assert.Error(t, err)
	assert.Nil(t, conn)
	assert.NoError(t, mock.ExpectationsWereMet())
}
