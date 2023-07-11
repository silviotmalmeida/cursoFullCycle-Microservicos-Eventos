// nome do pacote (está sendo utilizado o nome da referida pasta)
package repository

// dependências
import (
	"database/sql"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/entity"
)

// definindo a estrutura (similar à classe)
type TransactionRepository struct {
	// definindo os atributos e seus tipos
	DB *sql.DB
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	// criando
	return &TransactionRepository{
		DB: db,
	}
}

// função de criação
// devem ser descritos a estrutura associada, os argumentos e retornos
func (t *TransactionRepository) Create(transaction *entity.Transaction) error {
	// abrindo a conexão e preparando a query
	stmt, err := t.DB.Prepare("INSERT INTO transactions (id, account_id_from, account_id_to, amount, created_at) VALUES (?, ?, ?, ?, ?)")
	// em caso de erro, retorna somente o erro
	if err != nil {
		return err
	}
	// deve-se fechar a conexão ao final da função
	defer stmt.Close()
	// realizando a query
	_, err = stmt.Exec(transaction.ID, transaction.AccountFrom.ID, transaction.AccountTo.ID, transaction.Amount, transaction.CreatedAt)
	// em caso de erro, retorna somente o erro
	if err != nil {
		return err
	}
	// não retorna erro
	return nil
}
