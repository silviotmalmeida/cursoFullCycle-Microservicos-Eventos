// nome do pacote (está sendo utilizado o nome da referida pasta)
package mocks

// dependências
import (
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/entity"
	"github.com/stretchr/testify/mock"
)

// definindo o mock do gateway
type ClientGatewayMock struct {
	mock.Mock
}

// definindo o comportamento do Save do mock
func (m *ClientGatewayMock) Save(client *entity.Client) error {
	// recebe o client
	args := m.Called(client)
	// retorna como argumento 0 um erro
	return args.Error(0)
}

// definindo o comportamento do Get do mock
func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	// recebe o id
	args := m.Called(id)
	// retorna como argumento 0 um client
	// retorna como argumento 1 um erro
	return args.Get(0).(*entity.Client), args.Error(1)
}
