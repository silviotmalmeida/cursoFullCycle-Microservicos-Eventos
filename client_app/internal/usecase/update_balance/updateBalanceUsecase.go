// nome do pacote (está sendo utilizado o nome da referida pasta)
package update_balance

import (
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/entity"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/gateway"
)

// definindo os dados de input
type UpdateBalanceInputDTO struct {
	ID string
	Balance float64
}

// definindo os dados de output
type UpdateBalanceOutputDTO struct {
	ID string
}

// definindo o usecase
type UpdateBalanceUseCase struct {
	AccountGateway gateway.AccountGateway
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewUpdateBalanceUseCase(accountGateway gateway.AccountGateway) *UpdateBalanceUseCase {
	return &UpdateBalanceUseCase{
		AccountGateway: accountGateway,
	}
}

// função de execução do usecase
// devem ser descritos a estrutura associada, os argumentos e retornos
func (uc *UpdateBalanceUseCase) Execute(input *UpdateBalanceInputDTO) (*UpdateBalanceOutputDTO, error) {
	// criando a account
	account := entity.NewAccount(input.ID, input.Balance)
	// salvando no BD
	err := uc.AccountGateway.UpdateBalance(account)
	// se existirem erros, retorna somente o erro
	if err != nil {
		return nil, err
	}
	// organizando o output
	output := &UpdateBalanceOutputDTO{
		ID: account.ID,
	}
	// retornando o output, com erro nulo
	return output, nil
}
