// nome do pacote (está sendo utilizado o nome da referida pasta)
package uow

// dependências
import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// definindo interface de repositories
type RepositoryFactory func(tx *sql.Tx) interface{}

// definindo a interface do padrão unity of work
type UowInterface interface {
	Register(name string, fc RepositoryFactory)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, fn func(uow *Uow) error) error
	CommitOrRollback() error
	Rollback() error
	UnRegister(name string)
}

// definindo a estrutura (similar à classe)
type Uow struct {
	// definindo os atributos e seus tipos
	Db           *sql.DB
	Tx           *sql.Tx
	Repositories map[string]RepositoryFactory
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewUow(ctx context.Context, db *sql.DB) *Uow {
	// criando
	return &Uow{
		Db:           db,
		Repositories: make(map[string]RepositoryFactory),
	}
}

// método para registrar um repository
func (u *Uow) Register(name string, fc RepositoryFactory) {
	u.Repositories[name] = fc
}

// método para desregistrar um repository
func (u *Uow) UnRegister(name string) {
	delete(u.Repositories, name)
}

// método para chamada dos repositories, associados em uma transação
func (u *Uow) GetRepository(ctx context.Context, name string) (interface{}, error) {
	// caso não exista transação associada ao uow
	if u.Tx == nil {
		// retorna o erro
		return nil, errors.New("no transaction created")
	}
	// obtém o repository associado à transação
	repo := u.Repositories[name](u.Tx)
	// retorna o repository
	return repo, nil
}

// método para execução de uma função de forma atômica por transação
func (u *Uow) Do(ctx context.Context, fn func(Uow *Uow) error) error {
	// se já existir transação associada
	if u.Tx != nil {
		// retorna o erro
		return fmt.Errorf("transaction already started")
	}
	// inicia a transação
	tx, err := u.Db.BeginTx(ctx, nil)
	// em caso de erro
	if err != nil {
		// não atribui a transação ao uow e retorna o erro
		return err
	}
	// atribui a transação ao uow
	u.Tx = tx
	// executa a função passada como argumento
	err = fn(u)
	// em caso de erro
	if err != nil {
		// realiza o rollback
		errRb := u.Rollback()
		// em caso de erro
		if errRb != nil {
			// retorna o erro
			return fmt.Errorf(fmt.Sprintf("original error: %s, rollback error: %s", err.Error(), errRb.Error()))
		}
		// retorna o erro
		return err
	}
	// comita as operações ou, em caso de erro, aciona o rollback das operações
	return u.CommitOrRollback()
}

// método para efetuar o rollback das operações
func (u *Uow) Rollback() error {
	// caso não exista transação associada ao uow
	if u.Tx == nil {
		// retorna o erro
		return errors.New("no transaction to rollback")
	}
	// efetuando o rollback
	err := u.Tx.Rollback()
	// em caso de erro
	if err != nil {
		// retorna o erro
		return err
	}
	// retira a referencia da transação no uow
	u.Tx = nil
	// não retorna erro
	return nil
}

// método que commita as operações ou, em caso de erro, aciona o rollback das operações
func (u *Uow) CommitOrRollback() error {
	// efetua o commit
	err := u.Tx.Commit()
	// em caso de erro
	if err != nil {
		// realiza o rollback
		errRb := u.Rollback()
		// em caso de erro
		if errRb != nil {
			// retorna o erro
			return fmt.Errorf(fmt.Sprintf("original error: %s, rollback error: %s", err.Error(), errRb.Error()))
		}
		// retorna o erro
		return err
	}
	// retira a referencia da transação no uow
	u.Tx = nil
	// não retorna erro
	return nil
}
