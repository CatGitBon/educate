package repository

import (
	"database/sql"

	"github.com/vctrl/currency-service/auth/internal/dto"
)

type Auth struct {
	DB *sql.DB
}

func NewAuth(db *sql.DB) (Auth, error) {
	return Auth{
		DB: db,
	}, nil
}

func (a Auth) GetUserByID(id int64) (*dto.User, error) {
	query := "SELECT id, username, email, created_at, updated_at FROM auth.users WHERE id = $1"
	row := a.DB.QueryRow(query, id)

	var user dto.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}
