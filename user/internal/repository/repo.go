package repository

import (
	"context"
	"database/sql"
	"user/internal/models"
)

type Repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) CreateUser(ctx context.Context, user models.User, userProfile models.UserProfile) error {
	queryUserCreate := `insert into users(username, email, passwordHash) 
				values (?, ?, ?)`
	queryUserProfileCreate := `insert into user_profiles(user_id, first_name, last_name) 
				values (?, ?, ?)`

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	res, err := tx.ExecContext(ctx, queryUserCreate,
		user.Username,
		user.Email,
		user.Password,
	)

	if err != nil {
		return err
	}

	lastIndertedID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, queryUserProfileCreate,
		lastIndertedID,
		userProfile.FirstName,
		userProfile.LastName,
	)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetUserByID(ctx context.Context, userID int64) (*models.User, error) {
	query := `select id, username, email, created_at
			from users
			where id = ?`

	user := models.User{}
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repo) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `select id, username, email, passwordHash, created_at
			from users
			where username = ?`

	user := models.User{}
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repo) GetUserProfileByUserID(ctx context.Context, userID int64) (*models.UserProfile, error) {
	query := `select id, user_id, first_name, last_name
			from user_profiles
			where user_id = ?`

	userProfile := models.UserProfile{}

	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&userProfile.ID,
		&userProfile.UserID,
		&userProfile.FirstName,
		&userProfile.LastName,
	)
	if err != nil {
		return nil, err
	}

	return &userProfile, nil
}

func (r *Repo) UpdateUserProfileByUserID(ctx context.Context, userID int64, profile models.UserProfile) error {
	query := `update user_profiles
			set first_name  = ?, last_name  = ?
			where user_id = ?`

	_, err := r.db.ExecContext(ctx, query,
		profile.FirstName,
		profile.LastName,
		userID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) UpdateUserByID(ctx context.Context, user *models.User) error {
	query := `update users
			set email  = ?, passwordHash  = ?
			where id = ?`

	_, err := r.db.ExecContext(ctx, query,
		user.Email,
		user.Password,
		user.ID,
	)
	if err != nil {
		return err
	}

	return nil
}
