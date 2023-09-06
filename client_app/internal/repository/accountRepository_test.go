// nome do pacote (está sendo utilizado o nome da referida pasta)
package repository

// dependências
import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/entity"
	"github.com/stretchr/testify/suite"
)

// definindo a suíte de testes
type AccountRepositoryTestSuite struct {
	// definindo os atributos e seus tipos
	suite.Suite
	db                *sql.DB
	accountRepository *AccountRepository
}

// função de criação da suíte
// será executado antes de cada teste da suíte
func (s *AccountRepositoryTestSuite) SetupSuite() {
	// definindo o db como sqlite em memória, com as restrições de chave estrangeira ativadas
	db, err := sql.Open("sqlite3", "file::memory:?_foreign_keys=on")
	// não deve retornar erro
	s.Nil(err)
	// setando o db
	s.db = db
	// criando a tabela
	db.Exec("Create table accounts (id varchar(255), balance int, primary key (id))")
	// inicializando o repository
	s.accountRepository = NewAccountRepository(db)
}

// função de encerramento da suíte
// será executado depois de cada teste da suíte
func (s *AccountRepositoryTestSuite) TearDownSuite() {
	// deve-se fechar a conexão ao fim da função
	defer s.db.Close()
	// removendo a tabela
	s.db.Exec("DROP TABLE accounts")
}

// inicializando a suíte como um teste geral
func TestAccountRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AccountRepositoryTestSuite))
}

// testes de unidade

// teste de criação com sucesso
func (s *AccountRepositoryTestSuite) Test1Save() {
	// criando a account
	account := entity.NewAccount("123", 100.0)
	// salvando no db
	err := s.accountRepository.Save(account)
	// não deve retornar erro
	s.Nil(err)	
}

// teste de busca por id com sucesso
func (s *AccountRepositoryTestSuite) Test2FindByID() {
	// criando a account
	account := entity.NewAccount("456", 200.0)
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
	s.Equal(account.Balance, accountDB.Balance)
}

// teste de atualização de Balance com sucesso
func (s *AccountRepositoryTestSuite) Test3UpdateBalance() {
	// criando a account
	account := entity.NewAccount("789", 300.0)
	// salvando no db
	err := s.accountRepository.Save(account)
	// não deve retornar erro
	s.Nil(err)
	// atualizando o balance
	account.SetBalance(150.0)
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

// teste de listagem com sucesso
func (s *AccountRepositoryTestSuite) Test4List() {	
	// consultando no db
	accountDB, err := s.accountRepository.List()
	// não deve retornar erro
	s.Nil(err)
	// os atributos devem estar consistentes com a entrada
	s.Equal("123", accountDB[0].ID)
	s.Equal(100.0, accountDB[0].Balance)
	s.Equal("456", accountDB[1].ID)
	s.Equal(200.0, accountDB[1].Balance)
	s.Equal("789", accountDB[2].ID)
	s.Equal(150.0, accountDB[2].Balance)
}
