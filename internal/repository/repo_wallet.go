package repository

import (
	"context"
	"database/sql"
	"wallet/internal/domain"
)

type WalletPostgres struct {
	db *sql.DB
}

func NewWalletPostgres(db *sql.DB) *WalletPostgres {
	return &WalletPostgres{db: db}
}

func (r *WalletPostgres) GetByID(ctx context.Context, id string) (*domain.Wallet, error) {
	row := r.db.QueryRowContext(
		ctx,
		`SELECT id,balance FROM wallets WHERE id = $1`,
		id,
	)
	var wallet domain.Wallet
	if err := row.Scan(&wallet.ID, &wallet.Balance); err != nil {
		return nil, domain.ErrWalletNotFound
	}
	return &wallet, nil
}

func (r *WalletPostgres) UpdateBalance(ctx context.Context, id string, amount int64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var balance int64
	err = tx.QueryRowContext(ctx, `SELECT balance FROM wallets WHERE id = $1 FOR UPDATE`, id).Scan(&balance)
	if err != nil {
		return domain.ErrWalletNotFound
	}
	if balance+amount < 0 {
		return domain.ErrClientNotFound
	}
	_, err = tx.ExecContext(ctx, `UPDATE wallets SET balance = balance +$1 WHERE id = $2`, amount, id)
	if err != nil {
		return err
	}
	return tx.Commit()
}
