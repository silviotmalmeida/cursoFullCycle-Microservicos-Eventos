// nome do pacote (está sendo utilizado o nome da referida pasta)
package update_balance

// dependências
import (
	"testing"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// testes de unidade

// teste de execução com sucesso
func TestCreateAccountUseCase_Execute(t *testing.T) {
	// criando o mock do accountGateway
	accountGatewayMock := &mocks.AccountGatewayMock{}
	// definindo o retorno do médoto UpdateBalance como null
	accountGatewayMock.On("UpdateBalance", mock.Anything).Return(nil)
	// criando o usecase
	uc := NewUpdateBalanceUseCase(accountGatewayMock)
	// definindo o input
	input := &UpdateBalanceInputDTO{
		ID: "1234567890",
		Balance: 100.0,
	}
	// executando o usecase
	output, err := uc.Execute(input)
	// não deve retornar erro
	assert.Nil(t, err)
	// deve retornar um output válido
	assert.NotNil(t, output.ID)
	accountGatewayMock.AssertExpectations(t)
	// o método UpdateBalance do accountMock deve ter sido chamado 1 vez
	accountGatewayMock.AssertNumberOfCalls(t, "UpdateBalance", 1)
}
