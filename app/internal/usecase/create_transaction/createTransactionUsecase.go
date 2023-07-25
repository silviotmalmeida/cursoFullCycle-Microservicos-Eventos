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
	// salvando no BD
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

// // definindo os dados de output
// // foram incluídas as customizações dos nomes dos atributos ao converter para json
// type BalanceUpdatedOutputDTO struct {
// 	AccountIDFrom        string  `json:"account_id_from"`
// 	AccountIDTo          string  `json:"account_id_to"`
// 	BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
// 	BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
// }

// type CreateTransactionUseCase struct {
// 	Uow                uow.UowInterface
// 	EventDispatcher    events.EventDispatcherInterface
// 	TransactionCreated events.EventInterface
// 	BalanceUpdated     events.EventInterface
// }

// func NewCreateTransactionUseCase(
// 	Uow uow.UowInterface,
// 	eventDispatcher events.EventDispatcherInterface,
// 	transactionCreated events.EventInterface,
// 	balanceUpdated events.EventInterface,
// ) *CreateTransactionUseCase {
// 	return &CreateTransactionUseCase{
// 		Uow:                Uow,
// 		EventDispatcher:    eventDispatcher,
// 		TransactionCreated: transactionCreated,
// 		BalanceUpdated:     balanceUpdated,
// 	}
// }

// func (uc *CreateTransactionUseCase) Execute(ctx context.Context, input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
// 	output := &CreateTransactionOutputDTO{}
// 	balanceUpdatedOutput := &BalanceUpdatedOutputDTO{}
// 	err := uc.Uow.Do(ctx, func(_ *uow.Uow) error {
// 		accountRepository := uc.getAccountRepository(ctx)
// 		transactionRepository := uc.getTransactionRepository(ctx)

// 		accountFrom, err := accountRepository.FindByID(input.AccountIDFrom)
// 		if err != nil {
// 			return err
// 		}
// 		accountTo, err := accountRepository.FindByID(input.AccountIDTo)
// 		if err != nil {
// 			return err
// 		}
// 		transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
// 		if err != nil {
// 			return err
// 		}

// 		err = accountRepository.UpdateBalance(accountFrom)
// 		if err != nil {
// 			return err
// 		}

// 		err = accountRepository.UpdateBalance(accountTo)
// 		if err != nil {
// 			return err
// 		}

// 		err = transactionRepository.Create(transaction)
// 		if err != nil {
// 			return err
// 		}
// 		output.ID = transaction.ID
// 		output.AccountIDFrom = input.AccountIDFrom
// 		output.AccountIDTo = input.AccountIDTo
// 		output.Amount = input.Amount

// 		balanceUpdatedOutput.AccountIDFrom = input.AccountIDFrom
// 		balanceUpdatedOutput.AccountIDTo = input.AccountIDTo
// 		balanceUpdatedOutput.BalanceAccountIDFrom = accountFrom.Balance
// 		balanceUpdatedOutput.BalanceAccountIDTo = accountTo.Balance
// 		return nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	uc.TransactionCreated.SetPayload(output)
// 	uc.EventDispatcher.Dispatch(uc.TransactionCreated)

// 	uc.BalanceUpdated.SetPayload(balanceUpdatedOutput)
// 	uc.EventDispatcher.Dispatch(uc.BalanceUpdated)
// 	return output, nil
// }

// func (uc *CreateTransactionUseCase) getAccountRepository(ctx context.Context) gateway.AccountGateway {
// 	repo, err := uc.Uow.GetRepository(ctx, "AccountDB")
// 	if err != nil {
// 		panic(err)
// 	}
// 	return repo.(gateway.AccountGateway)
// }

// func (uc *CreateTransactionUseCase) getTransactionRepository(ctx context.Context) gateway.TransactionGateway {
// 	repo, err := uc.Uow.GetRepository(ctx, "TransactionDB")
// 	if err != nil {
// 		panic(err)
// 	}
// 	return repo.(gateway.TransactionGateway)
// }
