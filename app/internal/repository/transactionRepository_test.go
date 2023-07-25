// nome do pacote (está sendo utilizado o nome da referida pasta)
package repository

// dependências
import (
	"database/sql"
	"testing"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/entity"
	"github.com/stretchr/testify/suite"
)

// definindo a suíte de testes
type TransactionRepositoryTestSuite struct {
	// definindo os atributos e seus tipos
	suite.Suite
	db                    *sql.DB
	client                *entity.Client
	client2               *entity.Client
	accountFrom           *entity.Account
	accountTo             *entity.Account
	transactionRepository *TransactionRepository
}

// função de criação da suíte
// será executado antes de cada teste da suíte
func (s *TransactionRepositoryTestSuite) SetupSuite() {
	// definindo o db como sqlite em memória
	db, err := sql.Open("sqlite3", ":memory:")
	// não deve retornar erro
	s.Nil(err)
	// setando o db
	s.db = db
	// criando as tabelas
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("Create table accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	db.Exec("Create table transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date)")
	// criando e atribuindo os clients
	client, err := entity.NewClient("John", "j@j.com")
	s.Nil(err)
	s.client = client
	client2, err := entity.NewClient("John2", "jj@j.com")
	s.Nil(err)
	s.client2 = client2
	// inserindo os clients no db
	s.db.Exec("Insert into clients (id, name, email, created_at) values (?, ?, ?, ?)",
		s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt,
	)
	s.db.Exec("Insert into clients (id, name, email, created_at) values (?, ?, ?, ?)",
		s.client2.ID, s.client2.Name, s.client2.Email, s.client2.CreatedAt,
	)
	// criando e atribuindo os accounts
	accountFrom := entity.NewAccount(s.client)
	accountFrom.Balance = 1000
	s.accountFrom = accountFrom
	accountTo := entity.NewAccount(s.client2)
	accountTo.Balance = 1000
	s.accountTo = accountTo
	// inserindo os accounts no db
	s.db.Exec("Insert into accounts (id, client_id, balance, created_at) values (?, ?, ?, ?)",
		s.accountFrom.ID, s.accountFrom.ClientID, s.accountFrom.Balance, s.accountFrom.CreatedAt,
	)
	s.db.Exec("Insert into accounts (id, client_id, balance, created_at) values (?, ?, ?, ?)",
		s.accountTo.ID, s.accountTo.ClientID, s.accountTo.Balance, s.accountTo.CreatedAt,
	)
	// inicializando o repository
	s.transactionRepository = NewTransactionRepository(db)
}

// função de encerramento da suíte
// será executado depois de cada teste da suíte
func (s *TransactionRepositoryTestSuite) TearDownSuite() {
	// deve-se fechar a conexão ao fim da função
	defer s.db.Close()
	// removendo as tabelas
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

// inicializando a suíte como um teste geral
func TestTransactionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionRepositoryTestSuite))
}

// testes de unidade

// teste de criação com sucesso
func (s *TransactionRepositoryTestSuite) TestCreate() {
	// criando a transaction
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	// não deve retornar erro
	s.Nil(err)
	// salvando no db
	err = s.transactionRepository.Create(transaction)
	// não deve retornar erro
	s.Nil(err)
}
