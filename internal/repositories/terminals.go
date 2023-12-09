package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/wildanfaz/e-ticket-terminal/internal/constants"
	"github.com/wildanfaz/e-ticket-terminal/internal/models"
)

type ImplementTerminalsRepository struct {
	dbMySql *sql.DB
}

type TerminalsRepository interface {
	AddTerminal(ctx context.Context, payload models.Terminal) error
	AddTransaction(ctx context.Context, payload models.Transaction) error
	UpdateTransaction(ctx context.Context, payload models.UpdateTransaction, user models.User) error
}

func NewTerminalsRepository(dbMySql *sql.DB) TerminalsRepository {
	return &ImplementTerminalsRepository{
		dbMySql: dbMySql,
	}
}

func (r *ImplementTerminalsRepository) AddTerminal(ctx context.Context, payload models.Terminal) error {
	query := `
	INSERT INTO terminals (name, location_id)
	VALUES (?, ?)
	`

	_, err := r.dbMySql.ExecContext(ctx, query, payload.Name, payload.LocationId)
	if err != nil {
		return err
	}

	return nil
}

func (r *ImplementTerminalsRepository) AddTransaction(ctx context.Context, payload models.Transaction) error {
	query := `
	INSERT INTO transactions (user_id, from_terminal_id)
	VALUES (?, ?)
	`

	_, err := r.dbMySql.ExecContext(ctx, query, payload.UserId, payload.FromTerminalId)
	if err != nil {
		return err
	}

	return nil
}

func (r *ImplementTerminalsRepository) UpdateTransaction(ctx context.Context, payload models.UpdateTransaction, user models.User) error {
	var (
		transaction models.Transaction
		price       int
	)

	tx, err := r.dbMySql.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = tx.QueryRowContext(ctx, `
	SELECT id, from_terminal_id FROM transactions
	WHERE user_id = ? AND is_success = false
	ORDER BY created_at DESC
	`, user.Id).Scan(&transaction.Id, &transaction.FromTerminalId)
	if err != nil {
		tx.Rollback()
		return err
	}

	if transaction.FromTerminalId == payload.ToTerminalId {
		tx.Rollback()
		return errors.New("Cannot transaction to same terminal")
	}

	err = tx.QueryRowContext(ctx, `
	SELECT price FROM routes
	WHERE from_terminal_id = ? AND to_terminal_id = ?
	`, transaction.FromTerminalId, payload.ToTerminalId).Scan(&price)
	if err != nil {
		return err
	}

	if user.Balance-price < 0 {
		return errors.New(constants.InsufficientBalance)
	}

	currentBalance := user.Balance - price

	_, err = tx.ExecContext(ctx, `
	UPDATE users
	SET balance = ?
	WHERE id = ?
	`, currentBalance, user.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, `
	UPDATE transactions
	SET to_terminal_id = ?, is_success = true
	WHERE id = ?
	`, payload.ToTerminalId, transaction.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
