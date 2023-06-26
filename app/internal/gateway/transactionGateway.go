// nome do pacote (está sendo utilizado o nome da referida pasta)
package gateway

// dependências
import "github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/entity"

// definindo a interface do gateway
type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
