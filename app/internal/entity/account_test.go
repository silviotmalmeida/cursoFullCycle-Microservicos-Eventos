// nome do pacote (está sendo utilizado o nome da referida pasta)
package entity

// dependências
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// testes de unidade

// teste de criação com sucesso
func TestCreateAccount(t *testing.T) {
	// criando o client
	client, _ := NewClient("John Doe", "j@j")
	// criando o account
	account := NewAccount(client)
	// deve retornar um account válido
	assert.NotNil(t, account)
	// o Client ID deve corresponder ao client recebido por argumento
	assert.Equal(t, client.ID, account.Client.ID)
}

// teste de criação sem sucesso, cliente nulo
func TestCreateAccountWithNilClient(t *testing.T) {
	// criando o account
	account := NewAccount(nil)
	// não deve retornar o account
	assert.Nil(t, account)
}

// teste de crédito com sucesso
func TestCreditAccount(t *testing.T) {
	// criando o client
	client, _ := NewClient("John Doe", "j@j")
	// criando o account
	account := NewAccount(client)
	// o balance inicial deve ser 1000
	assert.Equal(t, float64(1000), account.Balance)
	// creditando o balance
	account.Credit(100)
	// o valor incrementado deve corresponder ao que foi creditado
	assert.Equal(t, float64(1100), account.Balance)
	// creditando o balance
	account.Credit(50)
	// o valor incrementado deve corresponder ao que foi creditado
	assert.Equal(t, float64(1150), account.Balance)
}

// teste de débito com sucesso
func TestDebitAccount(t *testing.T) {
	// criando o client
	client, _ := NewClient("John Doe", "j@j")
	// criando o account
	account := NewAccount(client)
	// o balance inicial deve ser 1000
	assert.Equal(t, float64(1000), account.Balance)
	// creditando o balance
	account.Credit(100)
	// o valor incrementado deve corresponder ao que foi creditado
	assert.Equal(t, float64(1100), account.Balance)
	// debitando o balance
	account.Debit(50)
	// o valor decrementado deve corresponder ao que foi debitado
	assert.Equal(t, float64(1050), account.Balance)
	// debitando o balance
	account.Debit(20)
	// o valor decrementado deve corresponder ao que foi debitado
	assert.Equal(t, float64(1030), account.Balance)
}
