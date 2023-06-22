// nome do pacote (está sendo utilizado o nome da referida pasta)
package create_client

// dependências
import (
	"time"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/entity"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/gateway"
)

// definindo os dados de input
type CreateClientInputDTO struct {
	Name  string
	Email string
}

// definindo os dados de output
type CreateClientOutputDTO struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// definindo o usecase
type CreateClientUseCase struct {
	ClientGateway gateway.ClientGateway
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewCreateClientUseCase(clientGateway gateway.ClientGateway) *CreateClientUseCase {
	return &CreateClientUseCase{
		ClientGateway: clientGateway,
	}
}

// função de execução do usecase
// devem ser descritos a estrutura associada, os argumentos e retornos
func (uc *CreateClientUseCase) Execute(input *CreateClientInputDTO) (*CreateClientOutputDTO, error) {
	// criando o client
	client, err := entity.NewClient(input.Name, input.Email)
	// se existirem erros, retorna somente o erro
	if err != nil {
		return nil, err
	}
	// salvando no BD
	err = uc.ClientGateway.Save(client)
	// se existirem erros, retorna somente o erro
	if err != nil {
		return nil, err
	}
	// organizando o output
	output := &CreateClientOutputDTO{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}
	// retornando o output, com erro nulo
	return output, nil
}
