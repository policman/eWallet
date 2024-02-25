package operation

import (
	"context"
	"time"
)

type Storage interface {
	Create(ctx context.Context, time time.Time, fromId, toId string, amount int) (isCreated bool, err error)
	FindFifteen(ctx context.Context, forId string, count int) ([]Operation, error)
	//Delete(ctx context.Context, id string) error
}
