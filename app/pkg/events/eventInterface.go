// nome do pacote (está sendo utilizado o nome da referida pasta)
package events

// dependências
import (
	"sync"
	"time"
)

// definindo a interface do event
type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
	SetPayload(payload interface{})
}

// definindo a interface do event handler
// o handler é responsável por realizar alguma operação a partir do evento
type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

// definindo a interface do event dispatcher
// responsável por registrar/remover os eventos e suas operações
// e disparar os eventos
type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear()
}
