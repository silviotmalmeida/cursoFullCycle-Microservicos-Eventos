// nome do pacote (está sendo utilizado o nome da referida pasta)
package gateway

// dependências
import "github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/entity"

// definindo a interface do gateway
type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
	UpdateBalance(account *entity.Account) error
}
