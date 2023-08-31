// nome do pacote (está sendo utilizado o nome da referida pasta)
package create_account

import (
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/entity"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/gateway"
)

// definindo os dados de input
type CreateAccountInputDTO struct {
	ID string
	Balance float64
}

// definindo os dados de output
type CreateAccountOutputDTO struct {
	ID string
}

// definindo o usecase
type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewCreateAccountUseCase(accountGateway gateway.AccountGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: accountGateway,
	}
}

// função de execução do usecase
// devem ser descritos a estrutura associada, os argumentos e retornos
func (uc *CreateAccountUseCase) Execute(input *CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	// criando a account
	account := entity.NewAccount(input.ID, input.Balance)
	// salvando no BD
	err := uc.AccountGateway.Save(account)
	// se existirem erros, retorna somente o erro
	if err != nil {
		return nil, err
	}
	// organizando o output
	output := &CreateAccountOutputDTO{
		ID: account.ID,
	}
	// retornando o output, com erro nulo
	return output, nil
}
