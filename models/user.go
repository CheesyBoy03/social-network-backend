package models

type User struct {
	ID          int    `json:"-" db:"id"`
	UserID      string `json:"user_id" db:"user_id"`
	FirstName   string `json:"first_name" db:"first_name" binding:"required"`
	LastName    string `json:"last_name" db:"last_name" binding:"required"`
	DateOfBirth string `json:"date_of_birth" db:"date_of_birth" binding:"required"`
	Username    string `json:"username" db:"username" binding:"required"`
	Password    string `json:"password" db:"password" binding:"required"`
}
