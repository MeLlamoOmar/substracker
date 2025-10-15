package repository

import (
	"SubsTracker/internal/models"
	"database/sql"
)

type SubsRepository interface {
	Create(s *models.Subscription) error
	FindByID(id string) (*models.Subscription, error)
	Delete(id string) error
	Update(id string, s *models.Subscription) (*models.Subscription, error)
}

type subsRepository struct {
	db *sql.DB
}

func NewSubsRepository(db *sql.DB) SubsRepository {
	return &subsRepository{
		db: db,
	}
}

func (r *subsRepository) Create(s *models.Subscription) error {
	var id string
	q := `INSERT INTO subscriptions (service_name, price, category, payment_method, billing_cycle, is_pinned, notes, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;`
	err := r.db.QueryRow(q, s.ServiceName, s.Price, s.Category, s.PaymentMethod, s.BillingCycle, s.IsPinned, s.Notes, s.UserID).Scan(&id)
	if err != nil {
		return err
	}

	s.ID = id
	return nil
}
func (r *subsRepository) FindByID(id string) (*models.Subscription, error) {
	var s models.Subscription
	q := `SELECT * FROM subscriptions WHERE id = $1;`
	err := r.db.QueryRow(q, id).Scan(&s.ID, &s.ServiceName, &s.Price, &s.Category, &s.PaymentMethod, &s.BillingCycle, &s.IsPinned, &s.Notes, &s.CreatedAt, &s.UpdatedAt, &s.UserID)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
func (r *subsRepository) Delete(id string) error {
	q := `DELETE FROM subscriptions WHERE id = $1;`
	if _, err := r.db.Exec(q, id); err != nil {
		return err
	}
	
	return nil
}
func (r *subsRepository) Update(id string, s *models.Subscription) (*models.Subscription, error) {
	q := `UPDATE subscriptions SET service_name = $1, price = $2, category = $3, payment_method = $4, billing_cycle = $5, is_pinned = $6, notes = $7, updated_at = $8 WHERE id = $9;`
	_, err := r.db.Exec(q, s.ServiceName, s.Price, s.Category, s.PaymentMethod, s.BillingCycle, s.IsPinned, s.Notes, s.UpdatedAt, id)
	if err != nil {
		return nil, err
	}
	
	return s, nil
}

