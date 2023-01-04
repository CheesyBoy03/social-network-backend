package repository

import (
	"fmt"

	"github.com/CheesyBoy03/social-media-backend/models"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, first_name, last_name, date_of_birth, user_id) values ($1, $2, $3, $4, $5, $6) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Username, user.Password, user.FirstName, user.LastName, user.DateOfBirth, user.UserID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
