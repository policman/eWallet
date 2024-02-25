package operation

import "ewallet/internal/models/wallet"

type Operation struct {
	Time   string        `json:"time"`
	FromID wallet.Wallet `json:"fromId"`
	ToID   wallet.Wallet `json:"toId"`
	Amount int           `json:"amount"`
}
