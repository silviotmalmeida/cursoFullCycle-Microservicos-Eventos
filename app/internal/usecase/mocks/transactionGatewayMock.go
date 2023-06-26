// nome do pacote (está sendo utilizado o nome da referida pasta)
package mocks

// dependências
import (
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/entity"
	"github.com/stretchr/testify/mock"
)

// definindo o mock do gateway
type TransactionGatewayMock struct {
	mock.Mock
}

// definindo o comportamento do Create do mock
func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	// recebe a transaction
	args := m.Called(transaction)
	// retorna como argumento 0 um erro
	return args.Error(0)
}
