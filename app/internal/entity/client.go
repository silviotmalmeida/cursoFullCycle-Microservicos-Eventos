// nome do pacote (está sendo utilizado o nome da referida pasta)
package entity

// dependências
import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// definindo a estrutura (similar à classe)
type Client struct {
	// definindo os atributos e seus tipos
	ID        string
	Name      string
	Email     string
	Accounts  []*Account
	CreatedAt time.Time
	UpdatedAt time.Time
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewClient(name string, email string) (*Client, error) {
	// criando
	client := &Client{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// validando
	err := client.Validate()
	// se existirem erros de validação, retorna somente o erro
	if err != nil {
		return nil, err
	}
	// retorna somente a estrutura
	return client, nil
}

// função de autovalidação
// devem ser descritos a estrutura associada, os argumentos e retornos
func (c *Client) Validate() error {
	// se o atributo name for vazio, retorna um erro
	if c.Name == "" {
		return errors.New("name is required")
	}
	// se o atributo email for vazio, retorna um erro
	if c.Email == "" {
		return errors.New("email is required")
	}
	// senão, não retorna nada
	return nil
}

// função de atualização
// devem ser descritos a estrutura associada, os argumentos e retornos
func (c *Client) Update(name string, email string) error {
	// atualizando os atributos
	c.Name = name
	c.Email = email
	c.UpdatedAt = time.Now()
	// validando
	err := c.Validate()
	// se existirem erros de validação, retorna somente o erro
	if err != nil {
		return err
	}
	// senão, não retorna nada
	return nil
}

// função de inclusão de account
// devem ser descritos a estrutura associada, os argumentos e retornos
func (c *Client) AddAccount(account *Account) error {
	// se a account fornecida pertencer a outro client, retorna um erro
	if account.Client.ID != c.ID {
		return errors.New("account does not belong to client")
	}
	// incrementa a lista de accounts do client
	c.Accounts = append(c.Accounts, account)
	// não retorna nada
	return nil
}
