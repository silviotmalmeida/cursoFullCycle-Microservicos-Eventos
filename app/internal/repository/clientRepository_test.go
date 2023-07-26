// nome do pacote (está sendo utilizado o nome da referida pasta)
package repository

// dependências
import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/entity"
	"github.com/stretchr/testify/suite"
)

// criando a suíte de testes
type ClientRepositoryTestSuite struct {
	// definindo os atributos e seus tipos
	suite.Suite
	db               *sql.DB
	clientRepository *ClientRepository
}

// função de criação da suíte
// será executado antes de cada teste da suíte
func (s *ClientRepositoryTestSuite) SetupSuite() {
	// definindo o db como sqlite em memória, com as restrições de chave estrangeira ativadas
	db, err := sql.Open("sqlite3", "file::memory:?_foreign_keys=on")
	// não deve retornar erro
	s.Nil(err)
	// setando o db
	s.db = db
	// criando a tabela
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date, primary key (id))")
	// inicializando o repository
	s.clientRepository = NewClientRepository(db)
}

// função de encerramento da suíte
// será executado depois de cada teste da suíte
func (s *ClientRepositoryTestSuite) TearDownSuite() {
	// deve-se fechar a conexão ao fim da função
	defer s.db.Close()
	// removendo a tabela
	s.db.Exec("DROP TABLE clients")
}

// inicializando a suíte como um teste geral
func TestClientRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ClientRepositoryTestSuite))
}

// testes de unidade

// teste de criação com sucesso
func (s *ClientRepositoryTestSuite) TestSave() {
	// criando um client
	client := &entity.Client{
		ID:    "1",
		Name:  "Test",
		Email: "j@j.com",
	}
	// salvando no db
	err := s.clientRepository.Save(client)
	// não deve retornar erro
	s.Nil(err)
}

// teste de busca por id com sucesso
func (s *ClientRepositoryTestSuite) TestGet() {
	// criando um client, desconsiderando o retorno do erro
	client, _ := entity.NewClient("John", "j@j.com")
	// salvando no db
	s.clientRepository.Save(client)
	// consultando no db
	clientDB, err := s.clientRepository.Get(client.ID)
	// não deve retornar erro
	s.Nil(err)
	// os atributos devem estar consistentes com a entrada
	s.Equal(client.ID, clientDB.ID)
	s.Equal(client.Name, clientDB.Name)
	s.Equal(client.Email, clientDB.Email)
}
