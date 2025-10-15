package repository

import (
	"SubsTracker/internal/models"
	"database/sql"
)

type UserRepository interface {
	Create(u *models.User) error
	FindByID(id string) (*models.User, error)
	Delete(id string) error
	Update(id string, user *models.User) (*models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(u *models.User) error {
	q := `INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id;`
	var id string
	if err := r.db.QueryRow(q, u.Email, u.PasswordHash).Scan(&id); err != nil {
		return err
	}

	u.ID = id
	return nil
}

func (r *userRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	q := `SELECT id, email, password_hash FROM users WHERE id = $1;`
	if err := r.db.QueryRow(q, id).Scan(&user.ID, &user.Email, &user.PasswordHash); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Delete(id string) error {
	q := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(q, id)
	return err
}

func (r *userRepository) Update(id string, user *models.User) (*models.User, error) {
	q := `UPDATE users SET email = $1, password_hash = $2 updated_at = $3 WHERE id = $4 RETURNING id`
	_, err := r.db.Exec(q, user.Email, user.PasswordHash, user.UpdatedAt, id)
	if err != nil {
		return nil, err
	}

	user.ID = id
	return user, nil
}