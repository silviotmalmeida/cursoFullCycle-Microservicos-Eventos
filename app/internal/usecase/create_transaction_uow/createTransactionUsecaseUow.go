// nome do pacote (está sendo utilizado o nome da referida pasta)
package create_transaction_uow

// dependências
import (
	"context"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/entity"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/gateway"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/pkg/events"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/pkg/uow"
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

// definindo os dados de output
// foram incluídas as customizações dos nomes dos atributos ao converter para json
type BalanceUpdatedOutputDTO struct {
	AccountIDFrom        string  `json:"account_id_from"`
	AccountIDTo          string  `json:"account_id_to"`
	BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
	BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
}

// definindo o usecase (com o gerenciamento da transação com unity of work)
type CreateTransactionUseCaseUow struct {
	Uow                     uow.UowInterface
	EventDispatcher         events.EventDispatcherInterface
	TransactionCreatedEvent events.EventInterface
	BalanceUpdatedEvent     events.EventInterface
}

// definindo o método contrutor (com o gerenciamento da transação com unity of work)
// devem ser descritos os argumentos e retornos
func NewCreateTransactionUseCaseUow(
	uow uow.UowInterface,
	eventDispatcher events.EventDispatcherInterface,
	transactionCreatedEvent events.EventInterface,
	balanceUpdatedEvent events.EventInterface,
) *CreateTransactionUseCaseUow {
	return &CreateTransactionUseCaseUow{
		Uow:                     uow,
		EventDispatcher:         eventDispatcher,
		TransactionCreatedEvent: transactionCreatedEvent,
		BalanceUpdatedEvent:     balanceUpdatedEvent,
	}
}

// função de execução do usecase (com o gerenciamento da transação com unity of work - uow)
// devem ser descritos a estrutura associada, os argumentos e retornos
func (uc *CreateTransactionUseCaseUow) Execute(ctx context.Context, input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	// inicializando os outputs
	output := &CreateTransactionOutputDTO{}
	balanceUpdatedOutput := &BalanceUpdatedOutputDTO{}
	// realizando a transaçao do Uow
	err := uc.Uow.Do(ctx,
		// função a ser executada com atomicidade
		func(_ *uow.Uow) error {
			// obtendo os repositories
			accountRepository := uc.getAccountRepository(ctx)
			transactionRepository := uc.getTransactionRepository(ctx)
			// consultando o accountFrom
			accountFrom, err := accountRepository.FindByID(input.AccountIDFrom)
			// se existirem erros, retorna o erro
			if err != nil {
				return err
			}
			// consultando o accountTo
			accountTo, err := accountRepository.FindByID(input.AccountIDTo)
			// se existirem erros, retorna o erro
			if err != nil {
				return err
			}
			// criando a transaction
			transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
			// se existirem erros, retorna o erro
			if err != nil {
				return err
			}
			// atualizando o balance da accountFrom
			err = accountRepository.UpdateBalance(accountFrom)
			// se existirem erros, retorna o erro
			if err != nil {
				return err
			}
			// atualizando o balance da accountTo
			err = accountRepository.UpdateBalance(accountTo)
			// se existirem erros, retorna o erro
			if err != nil {
				return err
			}
			// salvando a transaction no BD
			err = transactionRepository.Create(transaction)
			// se existirem erros, retorna o erro
			if err != nil {
				return err
			}
			// organizando o output
			output.ID = transaction.ID
			output.AccountIDFrom = input.AccountIDFrom
			output.AccountIDTo = input.AccountIDTo
			output.Amount = input.Amount
			balanceUpdatedOutput.AccountIDFrom = input.AccountIDFrom
			balanceUpdatedOutput.AccountIDTo = input.AccountIDTo
			balanceUpdatedOutput.BalanceAccountIDFrom = accountFrom.Balance
			balanceUpdatedOutput.BalanceAccountIDTo = accountTo.Balance
			// não retorna erro
			return nil
		})
	// se existirem erros, retorna somente o erro
	if err != nil {
		return nil, err
	}
	// populando o evento TransactionCreatedEvent com o output do usecase
	uc.TransactionCreatedEvent.SetPayload(output)
	// disparando as ações associadas ao evento TransactionCreatedEvent
	uc.EventDispatcher.Dispatch(uc.TransactionCreatedEvent)
	// populando o evento BalanceUpdatedEvent com o output do usecase
	uc.BalanceUpdatedEvent.SetPayload(balanceUpdatedOutput)
	// disparando as ações associadas ao evento BalanceUpdatedEvent
	uc.EventDispatcher.Dispatch(uc.BalanceUpdatedEvent)
	// retornando o output, com erro nulo
	return output, nil
}

// método para obtenção de um AccountRepository registrado no Uow
func (uc *CreateTransactionUseCaseUow) getAccountRepository(ctx context.Context) gateway.AccountGateway {
	// obtendo o repository
	repo, err := uc.Uow.GetRepository(ctx, "AccountDB")
	// em caso de erro
	if err != nil {
		// encerra a execução
		panic(err)
	}
	// retorna o repository
	return repo.(gateway.AccountGateway)
}

// método para obtenção de um TransactionRepository registrado no Uow
func (uc *CreateTransactionUseCaseUow) getTransactionRepository(ctx context.Context) gateway.TransactionGateway {
	// obtendo o repository
	repo, err := uc.Uow.GetRepository(ctx, "TransactionDB")
	// em caso de erro
	if err != nil {
		// encerra a execução
		panic(err)
	}
	// retorna o repository
	return repo.(gateway.TransactionGateway)
}
