package db

import (
	"context"
	"ewallet/internal/models/wallet"
	"ewallet/internal/storage/postgresql"
	"ewallet/utils/scanErr"

	"log/slog"
)

type repository struct {
	client postgresql.Client
	logger *slog.Logger
}

func (r *repository) Create(ctx context.Context) (newWallet wallet.Wallet, err error) {
	q := `INSERT INTO wallet DEFAULT VALUES RETURNING id, balance`

	if err = r.client.QueryRow(ctx, q).Scan(&newWallet.ID, &newWallet.Balance); err != nil {
		msg, err := scanErr.IdentifyErr(err)
		r.logger.Error(msg)
		return wallet.Wallet{}, err
	}
	return
}

func (r *repository) GetOne(ctx context.Context, wal wallet.Wallet) (err error) {
	q := `SELECT id, balance FROM wallet WHERE id = $1`

	if err = r.client.QueryRow(ctx, q, wal.ID).Scan(&wal.Balance); err != nil {
		msg, err := scanErr.IdentifyErr(err)
		r.logger.Error(msg)
		return err
	}
	return err
}

func (r *repository) UpdateBalance(ctx context.Context, amount int, id string) (err error) {
	q := `UPDATE wallet SET balance = $1 WHERE id = $2`
	comTag, err := r.client.Exec(ctx, q, amount, id)
	if err != nil {
		msg, err := scanErr.IdentifyErr(err)
		r.logger.Error(msg)
		return err
	}

	if comTag.RowsAffected() == 0 {
		r.logger.Info("row dont affected")
	}
	return nil
}

func NewRepository(client postgresql.Client, logger *slog.Logger) *repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
