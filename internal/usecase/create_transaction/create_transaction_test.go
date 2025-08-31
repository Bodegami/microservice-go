package createtransaction

import (
	"testing"

	"github.com.br/devbodegami/microservice-go/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("Client1", "c1@test.com")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entity.NewClient("Client2", "c2@test.com")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000)

	mockAccountGateway := &AccountGatewayMock{}
	mockAccountGateway.On("FindByID", account1.ID).Return(account1, nil)
	mockAccountGateway.On("FindByID", account2.ID).Return(account2, nil)

	mockTransactionGateway := &TransactionGatewayMock{}
	mockTransactionGateway.On("Create", mock.Anything).Return(nil)

	inputDto := CreateTransactionInputDTO{
		AccountIdFrom: account1.ID,
		AccountIdTo:   account2.ID,
		Amount:        100,
	}

	uc := NewCreateTransactionUseCase(mockTransactionGateway, mockAccountGateway)

	output, err := uc.Execute(inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	mockAccountGateway.AssertExpectations(t)
	mockAccountGateway.AssertNumberOfCalls(t, "FindByID", 2)
	mockTransactionGateway.AssertExpectations(t)
	mockTransactionGateway.AssertNumberOfCalls(t, "Create", 1)
}
