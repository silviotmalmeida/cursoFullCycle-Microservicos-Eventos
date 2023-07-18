// nome do pacote (está sendo utilizado o nome da referida pasta)
package create_account

import (
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/entity"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/gateway"
)

// definindo os dados de input
// foi incluído a customização do nome do atributo ao converter para json
type CreateAccountInputDTO struct {
	ClientID string `json:"client_id"`
}

// definindo os dados de output
type CreateAccountOutputDTO struct {
	ID string
}

// definindo o usecase
type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewCreateAccountUseCase(accountGateway gateway.AccountGateway, clientGateway gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: accountGateway,
		ClientGateway:  clientGateway,
	}
}

// função de execução do usecase
// devem ser descritos a estrutura associada, os argumentos e retornos
func (uc *CreateAccountUseCase) Execute(input *CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	// criando o client
	client, err := uc.ClientGateway.Get(input.ClientID)
	// se existirem erros, retorna somente o erro
	if err != nil {
		return nil, err
	}
	// criando a account
	account := entity.NewAccount(client)
	// salvando no BD
	err = uc.AccountGateway.Save(account)
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
