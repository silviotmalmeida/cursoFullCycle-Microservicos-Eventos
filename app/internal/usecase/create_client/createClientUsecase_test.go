package create_client

import (
	"testing"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/entity"
	"github.com/stretchr/testify/assert"
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

// testes de unidade

// teste de execução com sucesso
func TestCreateClientUseCase_Execute(t *testing.T) {
	// criando o mock
	m := &ClientGatewayMock{}
	// definindo o retorno do médoto Save como null
	m.On("Save", mock.Anything).Return(nil)
	// criando o usecase
	uc := NewCreateClientUseCase(m)
	// definindo o input
	input := &CreateClientInputDTO{
		Name:  "John Doe",
		Email: "j@j",
	}
	// executando o usecase
	output, err := uc.Execute(input)
	// não deve retornar erro
	assert.Nil(t, err)
	// deve retornar um output válido
	assert.NotNil(t, output)
	// o ID do do output deve estar preenchido
	assert.NotEmpty(t, output.ID)
	// o Name do do output deve correspender ao input
	assert.Equal(t, input.Name, output.Name)
	// o Email do do output deve correspender ao input
	assert.Equal(t, input.Email, output.Email)
	// o método Save do mock deve ter sido chamado 1 vez
	m.AssertNumberOfCalls(t, "Save", 1)
	// o método Get do mock não deve ter sido chamado
	m.AssertNumberOfCalls(t, "Get", 0)
}
