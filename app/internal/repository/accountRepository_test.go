// nome do pacote (está sendo utilizado o nome da referida pasta)
package repository

// dependências
import (
	"database/sql"
	"testing"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/entity"
	"github.com/stretchr/testify/suite"
)

// definindo a estrutura (similar à classe)
type AccountRepositoryTestSuite struct {
	// definindo os atributos e seus tipos
	suite.Suite
	db                *sql.DB
	accountRepository *AccountRepository
	client            *entity.Client
}

// função de criação da suíte
// será executado antes de cada teste da suíte
func (s *AccountRepositoryTestSuite) SetupSuite() {
	// definindo o db como sqlite em memória
	db, err := sql.Open("sqlite3", ":memory:")
	// não deve retornar erro
	s.Nil(err)
	// setando o db
	s.db = db
	// criando as tabelas
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("Create table accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	// inicializando o repository
	s.accountRepository = NewAccountRepository(db)
	// criando um client
	s.client, _ = entity.NewClient("John", "j@j.com")
}

// função de encerramento da suíte
// será executado depois de cada teste da suíte
func (s *AccountRepositoryTestSuite) TearDownSuite() {
	// deve-se fechar a conexão ao fim da função
	defer s.db.Close()
	// removendo as tabelas
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
}

// inicializando a suíte como um teste geral
func TestAccountRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AccountRepositoryTestSuite))
}

// testes de unidade

// teste de criação com sucesso
func (s *AccountRepositoryTestSuite) TestSave() {
	// criando a account
	account := entity.NewAccount(s.client)
	// salvando no db
	err := s.accountRepository.Save(account)
	// não deve retornar erro
	s.Nil(err)
}

// teste de busca por id com sucesso
func (s *AccountRepositoryTestSuite) TestFindByID() {
	// inserindo o client no db
	s.db.Exec("Insert into clients (id, name, email, created_at) values (?, ?, ?, ?)",
		s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt,
	)
	// criando a account
	account := entity.NewAccount(s.client)
	// salvando no db
	err := s.accountRepository.Save(account)
	// não deve retornar erro
	s.Nil(err)
	// consultando no db
	accountDB, err := s.accountRepository.FindByID(account.ID)
	// não deve retornar erro
	s.Nil(err)
	// os atributos devem estar consistentes com a entrada
	s.Equal(account.ID, accountDB.ID)
	s.Equal(account.ClientID, accountDB.ClientID)
	s.Equal(account.Balance, accountDB.Balance)
	s.Equal(account.Client.ID, accountDB.Client.ID)
	s.Equal(account.Client.Name, accountDB.Client.Name)
	s.Equal(account.Client.Email, accountDB.Client.Email)
}

// teste de atualização de Balance com sucesso
func (s *AccountRepositoryTestSuite) TestUpdateBalance() {
	// criando a account
	account := entity.NewAccount(s.client)
	// salvando no db
	err := s.accountRepository.Save(account)
	// não deve retornar erro
	s.Nil(err)
	// creditando no balance
	account.Credit(100)
	// salvando no db
	err = s.accountRepository.UpdateBalance(account)
	// não deve retornar erro
	s.Nil(err)
	// consultando no db
	accountDB, err := s.accountRepository.FindByID(account.ID)
	// não deve retornar erro
	s.Nil(err)
	// os atributos devem estar consistentes com a entrada
	s.Equal(account.Balance, accountDB.Balance)
}
