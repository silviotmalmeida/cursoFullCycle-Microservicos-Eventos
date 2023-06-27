// nome do pacote (está sendo utilizado o nome da referida pasta)
package create_client

// dependências
import (
	"testing"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// testes de unidade

// teste de execução com sucesso
func TestCreateClientUseCase_Execute(t *testing.T) {
	// criando o mock
	clientGatewayMock := &mocks.ClientGatewayMock{}
	// definindo o retorno do médoto Save como null
	clientGatewayMock.On("Save", mock.Anything).Return(nil)
	// criando o usecase
	uc := NewCreateClientUseCase(clientGatewayMock)
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
	// o Name do do output deve corresponder ao input
	assert.Equal(t, input.Name, output.Name)
	// o Email do do output deve corresponder ao input
	assert.Equal(t, input.Email, output.Email)
	clientGatewayMock.AssertExpectations(t)
	// o método Save do mock deve ter sido chamado 1 vez
	clientGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}
