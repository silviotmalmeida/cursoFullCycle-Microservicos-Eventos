// nome do pacote (está sendo utilizado o nome da referida pasta)
package create_account

// dependências
import (
	"testing"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/entity"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// testes de unidade

// teste de execução com sucesso
func TestCreateAccountUseCase_Execute(t *testing.T) {
	// criando o client, desconsiderando o retorno do erro
	client, _ := entity.NewClient("John Doe", "j@j")
	// criando o mock do clientGateway
	clientGatewayMock := &mocks.ClientGatewayMock{}
	// definindo o retorno do médoto Get como o client criado
	clientGatewayMock.On("Get", client.ID).Return(client, nil)
	// criando o mock do accountGateway
	accountGatewayMock := &mocks.AccountGatewayMock{}
	// definindo o retorno do médoto Save como null
	accountGatewayMock.On("Save", mock.Anything).Return(nil)
	// criando o usecase
	uc := NewCreateAccountUseCase(accountGatewayMock, clientGatewayMock)
	// definindo o input
	input := &CreateAccountInputDTO{
		ClientID: client.ID,
	}
	// executando o usecase
	output, err := uc.Execute(input)
	// não deve retornar erro
	assert.Nil(t, err)
	// deve retornar um output válido
	assert.NotNil(t, output.ID)
	clientGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertExpectations(t)
	// o método Get do clientMock deve ter sido chamado 1 vez
	clientGatewayMock.AssertNumberOfCalls(t, "Get", 1)
	// o método Save do accountMock deve ter sido chamado 1 vez
	accountGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}
