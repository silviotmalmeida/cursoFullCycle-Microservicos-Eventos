// nome do pacote (está sendo utilizado o nome da referida pasta)
package create_transaction_uow

// dependências
import (
	"context"
	"testing"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/entity"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/event"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/usecase/mocks"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// testes de unidade

// teste de execução com sucesso
func TestCreateTransactionUowUseCase_Execute(t *testing.T) {
	// criando o client 1
	client1, _ := entity.NewClient("client1", "j@j.com")
	// criando o account 1
	account1 := entity.NewAccount(client1)
	// creditando no account 1
	account1.Credit(1000)
	// criando o client 2
	client2, _ := entity.NewClient("client2", "j@j2.com")
	// criando o account 2
	account2 := entity.NewAccount(client2)
	// creditando no account 2
	account2.Credit(1000)
	// criando o mock do uow
	mockUow := &mocks.UowMock{}
	// definindo o retorno do médoto Do como null
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)
	// definindo o input
	inputDto := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}
	// criando o dispatcher
	eventDispatcher := events.NewEventDispatcher()
	// criando o event
	eventTransaction := event.NewTransactionCreatedEvent()
	// criando o event
	eventBalance := event.NewBalanceUpdatedEvent()
	// criando o contexto
	ctx := context.Background()
	// criando o usecase
	uc := NewCreateTransactionUowUseCase(mockUow, eventDispatcher, eventTransaction, eventBalance)
	// executando o usecase
	output, err := uc.Execute(ctx, inputDto)
	// não deve retornar erro
	assert.Nil(t, err)
	// deve retornar um output válido
	assert.NotNil(t, output)
	mockUow.AssertExpectations(t)
	// o método Do do mock deve ter sido chamado 1 vez
	mockUow.AssertNumberOfCalls(t, "Do", 1)
}
