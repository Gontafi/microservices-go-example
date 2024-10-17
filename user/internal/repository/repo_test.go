package repository

import (
	"context"
	"database/sql"
	"testing"
	"time"
	"user/internal/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewRepository(db)

	user := models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "passwordhash",
	}

	userProfile := models.UserProfile{
		FirstName: "First",
		LastName:  "Last",
	}

	ctx := context.Background()

	mock.ExpectBegin()
	mock.ExpectExec(`insert into users`).WithArgs(user.Username, user.Email, user.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`insert into user_profiles`).WithArgs(1, userProfile.FirstName, userProfile.LastName).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.CreateUser(ctx, user, userProfile)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewRepository(db)

	user := models.User{
		ID:       1,
		Username: "testuser",
		Email:    "test@example.com",
	}

	ctx := context.Background()

	rows := sqlmock.NewRows([]string{"id", "username", "email", "created_at"}).
		AddRow(user.ID, user.Username, user.Email, sql.NullTime{Time: time.Now(), Valid: true})

	mock.ExpectQuery(`select id, username, email, created_at from users where id = ?`).
		WithArgs(user.ID).
		WillReturnRows(rows)

	result, err := repo.GetUserByID(ctx, user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.ID, result.ID)
	assert.Equal(t, user.Username, result.Username)
	assert.Equal(t, user.Email, result.Email)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewRepository(db)

	user := models.User{
		ID:       1,
		Username: "testuser",
		Email:    "test@example.com",
		Password: "passwordhash",
	}

	ctx := context.Background()

	rows := sqlmock.NewRows([]string{"id", "username", "email", "passwordHash", "created_at"}).
		AddRow(user.ID, user.Username, user.Email, user.Password, sql.NullTime{Time: time.Now(), Valid: true})

	mock.ExpectQuery(`select id, username, email, passwordHash, created_at from users where username = ?`).
		WithArgs(user.Username).
		WillReturnRows(rows)

	result, err := repo.GetUserByUsername(ctx, user.Username)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.ID, result.ID)
	assert.Equal(t, user.Username, result.Username)
	assert.Equal(t, user.Email, result.Email)
	assert.Equal(t, user.Password, result.Password)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserProfileByUserID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewRepository(db)

	userProfile := models.UserProfile{
		ID:        1,
		UserID:    1,
		FirstName: "First",
		LastName:  "Last",
	}

	ctx := context.Background()

	rows := sqlmock.NewRows([]string{"id", "user_id", "first_name", "last_name"}).
		AddRow(userProfile.ID, userProfile.UserID, userProfile.FirstName, userProfile.LastName)

	mock.ExpectQuery(`select id, user_id, first_name, last_name from user_profiles where user_id = ?`).
		WithArgs(userProfile.UserID).
		WillReturnRows(rows)

	result, err := repo.GetUserProfileByUserID(ctx, userProfile.UserID)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, userProfile.ID, result.ID)
	assert.Equal(t, userProfile.UserID, result.UserID)
	assert.Equal(t, userProfile.FirstName, result.FirstName)
	assert.Equal(t, userProfile.LastName, result.LastName)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateUserProfileByUserID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewRepository(db)

	userProfile := models.UserProfile{
		FirstName: "UpdatedFirst",
		LastName:  "UpdatedLast",
	}

	ctx := context.Background()
	query := `update user_profiles
			set first_name  = ?, last_name  = ?
			where user_id = ?`

	mock.ExpectExec(query).
		WithArgs(userProfile.FirstName, userProfile.LastName, int64(1)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateUserProfileByUserID(ctx, 1, userProfile)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewRepository(db)

	user := models.User{
		ID:       1,
		Email:    "updated@example.com",
		Password: "updatedpasswordhash",
	}

	ctx := context.Background()
	query := `update users
			set email  = ?, passwordHash  = ?
			where id = ?`

	mock.ExpectExec(query).
		WithArgs(user.Email, user.Password, user.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateUserByID(ctx, &user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
