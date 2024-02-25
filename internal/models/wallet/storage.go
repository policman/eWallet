package wallet

import (
	"context"
)

type Storage interface {
	Create(ctx context.Context) (Wallet, error)
	GetOne(ctx context.Context, wal Wallet) error
	UpdateBalance(ctx context.Context, balance int, id string) (bool, error)
}
