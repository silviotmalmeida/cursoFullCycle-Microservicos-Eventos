// nome do pacote (está sendo utilizado o nome da referida pasta)
package repository

// dependências
import (
	"database/sql"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/entity"
)

// definindo a estrutura (similar à classe)
type AccountRepository struct {
	// definindo os atributos e seus tipos
	DB *sql.DB
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewAccountRepository(db *sql.DB) *AccountRepository {
	// criando
	return &AccountRepository{
		DB: db,
	}
}

// função de busca por id
// devem ser descritos a estrutura associada, os argumentos e retornos
func (a *AccountRepository) FindByID(id string) (*entity.Account, error) {
	// criando um account vazio
	account := &entity.Account{}
	// abrindo a conexão e preparando a query
	stmt, err := a.DB.Prepare("SELECT id, balance FROM accounts WHERE id = ?")
	// em caso de erro, retorna somente o erro
	if err != nil {
		return nil, err
	}
	// deve-se fechar a conexão ao final da função
	defer stmt.Close()
	// realizando a query
	row := stmt.QueryRow(id)
	// populando os dados retornados na account
	err = row.Scan(
		&account.ID,
		&account.Balance)
	// em caso de erro, retorna somente o erro
	if err != nil {
		return nil, err
	}
	// senão retorna a account populada
	return account, nil
}

// função de criação
// devem ser descritos a estrutura associada, os argumentos e retornos
func (a *AccountRepository) Save(account *entity.Account) error {
	// abrindo a conexão e preparando a query
	stmt, err := a.DB.Prepare("INSERT INTO accounts (id, balance) VALUES (?, ?)")
	// em caso de erro, retorna somente o erro
	if err != nil {
		return err
	}
	// deve-se fechar a conexão ao final da função
	defer stmt.Close()
	// realizando a query
	_, err = stmt.Exec(account.ID, account.Balance)
	// em caso de erro, retorna somente o erro
	if err != nil {
		return err
	}
	// não retorna erro
	return nil
}

// função de atualização do balance
// devem ser descritos a estrutura associada, os argumentos e retornos
func (a *AccountRepository) UpdateBalance(account *entity.Account) error {
	// abrindo a conexão e preparando a query
	stmt, err := a.DB.Prepare("UPDATE accounts SET balance = ? WHERE id = ?")
	// em caso de erro, retorna somente o erro
	if err != nil {
		return err
	}
	// deve-se fechar a conexão ao final da função
	defer stmt.Close()
	// realizando a query
	_, err = stmt.Exec(account.Balance, account.ID)
	// em caso de erro, retorna somente o erro
	if err != nil {
		return err
	}
	// não retorna erro
	return nil
}
