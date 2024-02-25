package db

import (
	"context"
	"ewallet/internal/models/operation"
	"ewallet/internal/storage/postgresql"
	"ewallet/utils/scanErr"
	"log/slog"
	"time"
)

type repository struct {
	client postgresql.Client
	logger *slog.Logger
}

func (r *repository) Create(ctx context.Context, time time.Time, fromId, toId string, amount int) (isCreated bool, err error) {
	q := `INSERT INTO operation (time, fromid, toid, amount) 
			VALUES ($1, $2, $3, $4)`

	comTag, err := r.client.Exec(ctx, q, time, fromId, toId, amount)
	if err != nil {
		msg, err := scanErr.IdentifyErr(err)
		r.logger.Error(msg, err)
		return false, err
	}

	if comTag.RowsAffected() == 1 {
		isCreated = true
		return
	} else {
		r.logger.Error("New operation is not created")
		isCreated = false
		return
	}

}

func (r *repository) FindFifteen(ctx context.Context, forId string, count int) ([]operation.Operation, error) {
	q := `SELECT time, fromid, toid, amount 
			FROM operation 
			WHERE fromid = $1 OR toid = $1 AND fromid != toid 
			ORDER BY time DESC
			LIMIT $2`

	rows, err := r.client.Query(ctx, q, forId, count)
	if err != nil {
		msg, err := scanErr.IdentifyErr(err)
		r.logger.Error(msg, err)
		return nil, err
	}
	rows.Close()

	var operations []operation.Operation
	for rows.Next() {
		var op operation.Operation
		if err := rows.Scan(&op.Time, &op.FromID, &op.ToID, &op.Amount); err != nil {
			r.logger.Error("Error with read some operation")
			return nil, err
		}
		operations = append(operations, op)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Error in iterations of read history")
		return nil, err
	}

	return operations, nil
}

func NewRepository(client postgresql.Client, logger *slog.Logger) operation.Storage {
	return &repository{
		client: client,
		logger: logger,
	}
}
