// nome do pacote (está sendo utilizado o nome da referida pasta)
package mocks

// dependências
import (
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/entity"
	"github.com/stretchr/testify/mock"
)

// definindo o mock do gateway
type AccountGatewayMock struct {
	mock.Mock
}

// definindo o comportamento do Save do mock
func (m *AccountGatewayMock) Save(account *entity.Account) error {
	// recebe o account
	args := m.Called(account)
	// retorna como argumento 0 um erro
	return args.Error(0)
}

// definindo o comportamento do FindByID do mock
func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	// recebe o id
	args := m.Called(id)
	// retorna como argumento 0 um account
	// retorna como argumento 1 um erro
	return args.Get(0).(*entity.Account), args.Error(1)
}

// definindo o comportamento do UpdateBalance do mock
func (m *AccountGatewayMock) UpdateBalance(account *entity.Account) error {
	// recebe o account
	args := m.Called(account)
	// retorna como argumento 0 um erro
	return args.Error(0)
}
