// nome do pacote (está sendo utilizado o nome da referida pasta)
package database

// dependências
import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/entity"
	"github.com/stretchr/testify/suite"
)

// criando a suíte de testes
type ClientDBTestSuite struct {
	// definindo os atributos e seus tipos
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

// função de criação da suíte
// será executado antes de cada teste da suíte
func (s *ClientDBTestSuite) SetupSuite() {
	// definindo o db como sqlite em memória
	db, err := sql.Open("sqlite3", ":memory:")
	// não deve retornar erro
	s.Nil(err)
	// setando o db
	s.db = db
	// criando a tabela
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	// inicializando o repository
	s.clientDB = NewClientDB(db)
}

// função de encerramento da suíte
// será executado depois de cada teste da suíte
func (s *ClientDBTestSuite) TearDownSuite() {
	// deve-se fechar a conexão ao fim da função
	defer s.db.Close()
	// removendo a tabela
	s.db.Exec("DROP TABLE clients")
}

// inicializando a suíte como um teste geral
func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

// testes de unidade

// teste de criação com sucesso
func (s *ClientDBTestSuite) TestSave() {
	// criando um client
	client := &entity.Client{
		ID:    "1",
		Name:  "Test",
		Email: "j@j.com",
	}
	// salvando no db
	err := s.clientDB.Save(client)
	// não deve retornar erro
	s.Nil(err)
}

// teste de busca por id com sucesso
func (s *ClientDBTestSuite) TestGet() {
	// criando um client, desconsiderando o retorno do erro
	client, _ := entity.NewClient("John", "j@j.com")
	// salvando no db
	s.clientDB.Save(client)
	// consultando no db
	clientDB, err := s.clientDB.Get(client.ID)
	// não deve retornar erro
	s.Nil(err)
	// os atributos devem estar consistentes com a entrada
	s.Equal(client.ID, clientDB.ID)
	s.Equal(client.Name, clientDB.Name)
	s.Equal(client.Email, clientDB.Email)
}
