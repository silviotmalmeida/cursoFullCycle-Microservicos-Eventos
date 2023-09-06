// nome do pacote (está sendo utilizado o nome da referida pasta)
package list_accounts

import (
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/entity"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/gateway"
)

// definindo os dados de output
type ListAccountsOutputDTO struct {
	Accounts []*entity.Account
}

// definindo o usecase
type ListAccountsUseCase struct {
	AccountGateway gateway.AccountGateway
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewListAccountsUseCase(accountGateway gateway.AccountGateway) *ListAccountsUseCase {
	return &ListAccountsUseCase{
		AccountGateway: accountGateway,
	}
}

// função de execução do usecase
// devem ser descritos a estrutura associada, os argumentos e retornos
func (uc *ListAccountsUseCase) Execute() (*ListAccountsOutputDTO, error) {
	// consultando no BD
	accounts, err := uc.AccountGateway.List()
	// se existirem erros, retorna somente o erro
	if err != nil {
		return nil, err
	}
	// organizando o output
	output := &ListAccountsOutputDTO{
		Accounts: accounts,
	}
	// retornando o output, com erro nulo
	return output, nil
}
