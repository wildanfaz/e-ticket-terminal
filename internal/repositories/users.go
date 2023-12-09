package repositories

import (
	"context"
	"database/sql"

	"github.com/wildanfaz/e-ticket-terminal/internal/models"
)

type ImplementUsersRepository struct {
	dbMySql *sql.DB
}

type UsersRepository interface {
	Register(ctx context.Context, payload models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	TopUp(ctx context.Context, payload models.User) error
}

func NewUsersRepository(dbMySql *sql.DB) UsersRepository {
	return &ImplementUsersRepository{
		dbMySql: dbMySql,
	}
}

func (r *ImplementUsersRepository) Register(ctx context.Context, payload models.User) error {
	query := `
	INSERT INTO users (email, password)
	VALUES (?, ?)
	`

	_, err := r.dbMySql.ExecContext(ctx, query, payload.Email, payload.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r ImplementUsersRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var (
		user models.User
	)

	query := `
	SELECT id, email, password, balance, created_at, updated_at
	FROM users
	WHERE email = ?
	`

	err := r.dbMySql.QueryRowContext(ctx, query, email).Scan(
		&user.Id, &user.Email, &user.Password,
		&user.Balance, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &user, nil
}

func (r *ImplementUsersRepository) TopUp(ctx context.Context, payload models.User) error {
	query := `
	UPDATE users
	SET balance = balance + ?
	WHERE email = ?
	`

	_, err := r.dbMySql.ExecContext(ctx, query, payload.Balance, payload.Email)
	if err != nil {
		return err
	}

	return nil
}
