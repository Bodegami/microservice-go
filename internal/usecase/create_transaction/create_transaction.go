package createtransaction

import (
	"github.com.br/devbodegami/microservice-go/fc-ms-wallet/internal/entity"
	"github.com.br/devbodegami/microservice-go/fc-ms-wallet/internal/gateway"
)

type CreateTransactionInputDTO struct {
	AccountIdFrom string
	AccountIdTo   string
	Amount        float64
}

type CreateTransactionOutputDTO struct {
	ID string
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway     gateway.AccountGateway
}

func NewCreateTransactionUseCase(transactionGateway gateway.TransactionGateway, accountGateway gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: transactionGateway,
		AccountGateway:     accountGateway,
	}
}

func (c *CreateTransactionUseCase) Execute(input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	accountFrom, err := c.AccountGateway.FindByID(input.AccountIdFrom)
	if err != nil {
		return nil, err
	}

	accountTo, err := c.AccountGateway.FindByID(input.AccountIdTo)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}
	err = c.TransactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}

	return &CreateTransactionOutputDTO{
		ID: transaction.ID,
	}, nil
}
