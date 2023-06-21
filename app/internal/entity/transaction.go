// nome do pacote (está sendo utilizado o nome da referida pasta)
package entity

// dependências
import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// definindo a estrutura (similar à classe)
// foi incluído o detalhamento da associação com a tabela Account para ser utilizado pelo gorm
type Transaction struct {
	ID            string
	AccountFrom   *Account `gorm:"foreignKey:AccountFromID"`
	AccountFromID string
	AccountTo     *Account `gorm:"foreignKey:AccountToID"`
	AccountToID   string
	Amount        float64
	CreatedAt     time.Time
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewTransaction(accountFrom *Account, accountTo *Account, amount float64) (*Transaction, error) {
	// criando
	transaction := &Transaction{
		ID:          uuid.New().String(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}
	// validando
	err := transaction.Validate()
	// se existirem erros de validação, retorna somente o erro
	if err != nil {
		return nil, err
	}
	// realiza a transação nas accounts, atualizando os respectivos balances
	transaction.Commit()
	// retorna somente a estrutura
	return transaction, nil
}

// função de autovalidação
// devem ser descritos a estrutura associada, os argumentos e retornos
func (t *Transaction) Validate() error {
	// se o atributo amount for menor ou igual a zero, retorna um erro
	if t.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	// se o atributo balance da accountFrom for menor que o amount, retorna um erro
	if t.AccountFrom.Balance < t.Amount {
		return errors.New("insufficient funds")
	}
	// senão, não retorna nada
	return nil
}

// função de execução da transação
// devem ser descritos a estrutura associada, os argumentos e retornos
func (t *Transaction) Commit() {
	// atualizando os saldos das accounts
	t.AccountFrom.Debit(t.Amount)
	t.AccountTo.Credit(t.Amount)
}
