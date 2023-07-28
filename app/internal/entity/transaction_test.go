// nome do pacote (está sendo utilizado o nome da referida pasta)
package entity

// dependências
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// testes de unidade

// teste de criação com sucesso
func TestCreateTransaction(t *testing.T) {
	// criando o client 1
	client1, _ := NewClient("John Doe", "j@j")
	// criando o account 1
	account1 := NewAccount(client1)
	// criando o client 2
	client2, _ := NewClient("John Doe 2", "j@j2")
	// criando o account 2
	account2 := NewAccount(client2)
	// inserindo créditos
	account1.Credit(1000)
	account2.Credit(1000)
	// realizando a transaction
	transaction, err := NewTransaction(account1, account2, 100)
	// não deve retornar erro
	assert.Nil(t, err)
	// deve retornar uma transaction válida
	assert.NotNil(t, transaction)
	// validando os balances após a transaction
	assert.Equal(t, 2100.0, account2.Balance)
	assert.Equal(t, 1900.0, account1.Balance)
}

// teste de criação sem sucesso, balance insuficiente
func TestCreateTransactionWithInsuficientBalance(t *testing.T) {
	// criando o client 1
	client1, _ := NewClient("John Doe", "j@j")
	// criando o account 1
	account1 := NewAccount(client1)
	// criando o client 2
	client2, _ := NewClient("John Doe 2", "j@j2")
	// criando o account 2
	account2 := NewAccount(client2)
	// inserindo créditos
	account1.Credit(1000)
	account2.Credit(1000)
	// realizando a transaction
	transaction, err := NewTransaction(account1, account2, 3000)
	// deve retornar erro
	assert.NotNil(t, err)
	assert.EqualError(t, err, "insufficient funds")
	// não deve retornar a transaction
	assert.Nil(t, transaction)
	// validando os balances após a transaction (inalterados)
	assert.Equal(t, 2000.0, account2.Balance)
	assert.Equal(t, 2000.0, account1.Balance)
}
