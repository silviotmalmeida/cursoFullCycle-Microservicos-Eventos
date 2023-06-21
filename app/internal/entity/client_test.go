// nome do pacote (está sendo utilizado o nome da referida pasta)
package entity

// dependências
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// testes de unidade

// teste de criação com sucesso
func TestCreateNewClient(t *testing.T) {
	// criando o client
	client, err := NewClient("John Doe", "j@j.com")
	// não deve retornar erro
	assert.Nil(t, err)
	// deve retornar um client válido
	assert.NotNil(t, client)
	// os atributos do client devem ser iguais aos argumentos passados
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

// teste de criação sem sucesso, argumentos inválidos
func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	// criando o client
	client, err := NewClient("", "")
	// deve retornar erro
	assert.EqualError(t, err, "name is required")
	// não deve retornar o client
	assert.Nil(t, client)
	// criando o client
	client, err = NewClient("John Doe", "")
	// deve retornar erro
	assert.EqualError(t, err, "email is required")
	// não deve retornar o client
	assert.Nil(t, client)
}

// teste de atualização com sucesso
func TestUpdateClient(t *testing.T) {
	// criando o client
	// como o retorno do erro não será utilizado, foi omitido
	client, _ := NewClient("John Doe", "j@j.com")
	// atualizando
	err := client.Update("John Doe Update", "j@j.com")
	// não deve retornar erro
	assert.Nil(t, err)
	// os atributos do client devem ser iguais aos argumentos passados
	assert.Equal(t, "John Doe Update", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

// teste de atualização sem sucesso, argumentos inválidos
func TestUpdateClientWithInvalidArgs(t *testing.T) {
	// criando o client
	// como o retorno do erro não será utilizado, foi omitido
	client, _ := NewClient("John Doe", "j@j.com")
	// atualizando
	err := client.Update("", "j@j.com")
	// deve retornar erro
	assert.EqualError(t, err, "name is required")
	// atualizando
	err = client.Update("John Doe", "")
	// deve retornar erro
	assert.EqualError(t, err, "email is required")
}

// teste de inclusão de account
func TestAddAccountToClient(t *testing.T) {
	// criando o client
	// como o retorno do erro não será utilizado, foi omitido
	client, _ := NewClient("John Doe", "j@j")
	// a lista de accounts associada deve estar vazia
	assert.Equal(t, 0, len(client.Accounts))
	// criando o account 1
	account1 := NewAccount(client)
	// adicionando o account
	err := client.AddAccount(account1)
	// não deve retornar erro
	assert.Nil(t, err)
	// a lista de accounts associada deve conter 1 elemento
	assert.Equal(t, 1, len(client.Accounts))
	// criando o account 2
	account2 := NewAccount(client)
	// adicionando o account
	err = client.AddAccount(account2)
	// não deve retornar erro
	assert.Nil(t, err)
	// a lista de accounts associada deve conter 2 elementos
	assert.Equal(t, 2, len(client.Accounts))
}

// teste de inclusão de account sem sucesso, outro ID de client
func TestAddAccountToClientWithOtherCLientID(t *testing.T) {
	// criando o client 1
	// como o retorno do erro não será utilizado, foi omitido
	client1, _ := NewClient("John Doe", "j@j")
	// criando o client 2
	// como o retorno do erro não será utilizado, foi omitido
	client2, _ := NewClient("Jane Doe", "j@n")
	// criando o account 1
	account1 := NewAccount(client1)
	// criando o account 2
	account2 := NewAccount(client2)
	// adicionando o account 1
	err := client1.AddAccount(account1)
	// não deve retornar erro
	assert.Nil(t, err)
	// adicionando o account 2
	err = client1.AddAccount(account2)
	// deve retornar erro
	assert.EqualError(t, err, "account does not belong to client")
}
