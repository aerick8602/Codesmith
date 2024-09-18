package repository

import (
	"database/sql"
	"instagram-api/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	_, err := r.DB.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
	return err
}

// GetUser retrieves a user by ID
func (r *UserRepository) GetUser(id int) (*models.User, error) {
	row := r.DB.QueryRow("SELECT id, name, email, password FROM users WHERE id = $1", id)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil // No user found
	} else if err != nil {
		return nil, err // Error querying the database
	}

	return &user, nil
}
