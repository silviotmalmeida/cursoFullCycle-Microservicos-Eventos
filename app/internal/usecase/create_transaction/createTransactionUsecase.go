// nome do pacote (está sendo utilizado o nome da referida pasta)
package create_transaction

// dependências
import (
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/entity"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/gateway"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/pkg/events"
)

// definindo os dados de input
// foram incluídas as customizações dos nomes dos atributos ao converter para json
type CreateTransactionInputDTO struct {
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

// definindo os dados de output
// foram incluídas as customizações dos nomes dos atributos ao converter para json
type CreateTransactionOutputDTO struct {
	ID            string  `json:"id"`
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

// definindo o usecase (sem o gerenciamento da transação com unity of work)
type CreateTransactionUseCase struct {
	TransactionGateway      gateway.TransactionGateway
	AccountGateway          gateway.AccountGateway
	EventDispatcher         events.EventDispatcherInterface
	TransactionCreatedEvent events.EventInterface
}

// definindo o método contrutor (sem o gerenciamento da transação com unity of work)
// devem ser descritos os argumentos e retornos
func NewCreateTransactionUseCase(
	transactionGateway gateway.TransactionGateway,
	accountGateway gateway.AccountGateway,
	eventDispatcher events.EventDispatcherInterface,
	transactionCreatedEvent events.EventInterface) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway:      transactionGateway,
		AccountGateway:          accountGateway,
		EventDispatcher:         eventDispatcher,
		TransactionCreatedEvent: transactionCreatedEvent,
	}
}

// função de execução do usecase (sem o gerenciamento da transação com unity of work - uow)
// devem ser descritos a estrutura associada, os argumentos e retornos
func (uc *CreateTransactionUseCase) Execute(input *CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	// consultando o accountFrom
	accountFrom, err := uc.AccountGateway.FindByID(input.AccountIDFrom)
	// se existirem erros, retorna somente o erro
	if err != nil {
		return nil, err
	}
	// consultando o accountTo
	accountTo, err := uc.AccountGateway.FindByID(input.AccountIDTo)
	// se existirem erros, retorna somente o erro
	if err != nil {
		return nil, err
	}
	// criando a transaction
	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	// se existirem erros, retorna somente o erro
	if err != nil {
		return nil, err
	}
	// atualizando o balance da accountFrom
	err = uc.AccountGateway.UpdateBalance(accountFrom)
	// se existirem erros, retorna somente o erro
	if err != nil {
		return nil, err
	}
	// atualizando o balance da accountTo
	err = uc.AccountGateway.UpdateBalance(accountTo)
	// se existirem erros, retorna somente o erro
	if err != nil {
		return nil, err
	}
	// salvando a transaction no BD
	err = uc.TransactionGateway.Create(transaction)
	// se existirem erros, retorna somente o erro
	if err != nil {
		return nil, err
	}
	// organizando o output
	output := &CreateTransactionOutputDTO{
		ID:            transaction.ID,
		AccountIDFrom: transaction.AccountFromID,
		AccountIDTo:   transaction.AccountToID,
		Amount:        transaction.Amount,
	}
	// populando o evento com o output do usecase
	uc.TransactionCreatedEvent.SetPayload(output)
	// disparando as ações associadas ao evento TransactionCreatedEvent
	uc.EventDispatcher.Dispatch(uc.TransactionCreatedEvent)
	// retornando o output, com erro nulo
	return output, nil
}
