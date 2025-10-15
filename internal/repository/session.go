package repository

import (
	"SubsTracker/internal/models"
	"database/sql"
)

type SessionRepository interface {
	Create(s *models.Session) (*models.Session, error)
	FindByID(id string) (*models.Session, error)
	Delete(id string) error
	Update(id string, s *models.Session) (*models.Session, error)
}

type sessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) SessionRepository {
	return &sessionRepository{
		db: db,
	}
}

func (r *sessionRepository) Create(s *models.Session) (*models.Session, error) {
	var idSession string
	q := `INSERT INTO sessions (token, expired_at, user_agent, ip_address, user_id) VALUES ($1, $2, $3, $4, $5) RETURNING id_session;`
	if err := r.db.QueryRow(q, s.Token, s.ExpiredAt, s.UserAgent, s.IPAddress, s.UserID).Scan(&idSession); err != nil {
		return nil, err
	}

	s.IDSession = idSession
	return s, nil
}
func (r *sessionRepository) FindByID(id string) (*models.Session, error) {
	var session models.Session
	q := `SELECT id_session, token, expired_at, user_agent, ip_address FROM sessions WHERE id_session = $1;`
	if err := r.db.QueryRow(q, id).Scan(&session.IDSession, &session.Token, &session.ExpiredAt, &session.UserAgent, &session.IPAddress); err != nil {
		return nil, err
	}
	
	return &session, nil
}
func (r *sessionRepository) Delete(id string) error {
	q := `DELETE FROM sessions WHERE id_session = &1;`
	_, err := r.db.Exec(q, id)
	return err
}
func (r *sessionRepository) Update(id string, s *models.Session) (*models.Session, error) {
	q := `UPDATE sessions SET token = $1, expired_at = $2, updated_at = $3, user_agent = $4, ip_address = $5 WHERE id_session = $6 RETURNING id_session;`
	if _, err := r.db.Exec(q, s.Token, s.ExpiredAt, s.UpdatedAt, s.UserAgent, s.IPAddress, id); err != nil {
		return nil, err
	}

	s.IDSession = id
	return s, nil
}