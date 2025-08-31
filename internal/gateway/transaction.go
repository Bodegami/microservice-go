package gateway

import "github.com.br/devbodegami/microservice-go/fc-ms-wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
