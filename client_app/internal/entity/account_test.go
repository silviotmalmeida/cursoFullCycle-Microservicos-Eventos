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
	// parâmetros de entrada
	id := "1234567890qwerty"
	balance := 100.0
	// criando o account
	account := NewAccount(id, balance)
	// deve retornar um account válido
	assert.NotNil(t, account)
	// os atributos devem corresponder aos recebidos por argumento
	assert.Equal(t, id, account.ID)
	assert.Equal(t, balance, account.Balance)
}

// teste de atualização de balance com sucesso
func TestSetBalance(t *testing.T) {
	// parâmetros de entrada
	id := "1234567890qwerty"
	balance := 100.0
	// criando o account
	account := NewAccount(id, balance)
	// deve retornar um account válido
	assert.NotNil(t, account)
	// os atributos devem corresponder aos recebidos por argumento
	assert.Equal(t, id, account.ID)
	assert.Equal(t, balance, account.Balance)
	// atualizando o balance
	account.SetBalance(150.0)
	// verificando a atualização
	assert.Equal(t, 150.0, account.Balance)
}