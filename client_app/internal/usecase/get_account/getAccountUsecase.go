// nome do pacote (está sendo utilizado o nome da referida pasta)
package get_account

import (
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/gateway"
)

// definindo os dados de input
type GetAccountInputDTO struct {
	ID string
}

// definindo os dados de output
type GetAccountOutputDTO struct {
	ID string
	Balance float64
}

// definindo o usecase
type GetAccountUseCase struct {
	AccountGateway gateway.AccountGateway
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewGetAccountUseCase(accountGateway gateway.AccountGateway) *GetAccountUseCase {
	return &GetAccountUseCase{
		AccountGateway: accountGateway,
	}
}

// função de execução do usecase
// devem ser descritos a estrutura associada, os argumentos e retornos
func (uc *GetAccountUseCase) Execute(input *GetAccountInputDTO) (*GetAccountOutputDTO, error) {
	// buscando no BD
	account, err := uc.AccountGateway.FindByID(input.ID)
	// se existirem erros, retorna somente o erro
	if err != nil {
		return nil, err
	}
	// organizando o output
	output := &GetAccountOutputDTO{
		ID: account.ID,
		Balance: account.Balance,
	}
	// retornando o output, com erro nulo
	return output, nil
}
