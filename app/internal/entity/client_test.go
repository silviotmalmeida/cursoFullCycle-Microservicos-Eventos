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
	// criando a estrutura
	client, err := NewClient("John Doe", "j@j.com")
	// não deve retornar erro
	assert.Nil(t, err)
	// deve retornar uma estrutura válida
	assert.NotNil(t, client)
	// os atributos da estrutura devem ser iguais aos argumentos passados
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

// teste de criação sem sucesso, argumentos inválidos
func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	// criando a estrutura
	client, err := NewClient("", "")
	// deve retornar erro
	assert.EqualError(t, err, "name is required")
	// não deve retornar a estrutura
	assert.Nil(t, client)

	// criando a estrutura
	client, err = NewClient("John Doe", "")
	// deve retornar erro
	assert.EqualError(t, err, "email is required")
	// não deve retornar a estrutura
	assert.Nil(t, client)
}

// teste de atualização com sucesso
func TestUpdateClient(t *testing.T) {
	// criando a estrutura
	// como o retorno do erro não será utilizado, foi omitido
	client, _ := NewClient("John Doe", "j@j.com")
	// atualizando
	err := client.Update("John Doe Update", "j@j.com")
	// não deve retornar erro
	assert.Nil(t, err)
	// os atributos da estrutura devem ser iguais aos argumentos passados
	assert.Equal(t, "John Doe Update", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

// teste de atualização sem sucesso, argumentos inválidos
func TestUpdateClientWithInvalidArgs(t *testing.T) {
	// criando a estrutura
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

// func TestAddAccountToClient(t *testing.T) {
// 	client, _ := NewClient("John Doe", "j@j")
// 	account := NewAccount(client)
// 	err := client.AddAccount(account)
// 	assert.Nil(t, err)
// 	assert.Equal(t, 1, len(client.Accounts))
// }
