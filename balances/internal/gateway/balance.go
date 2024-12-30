package gateway

import "github.com.br/devfullcycle/fc-ms-balances/internal/entity"

type BalanceGateway interface {
	Save(balance *entity.Balance) error
	FindByAccountID(accountID string) (*entity.Balance, error)
}