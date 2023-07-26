// nome do pacote (está sendo utilizado o nome da referida pasta)
package event

// dependências
import "time"

// definindo a estrutura (similar à classe) que implemente a EventInterface
type BalanceUpdatedEvent struct {
	// definindo os atributos e seus tipos
	Name    string
	Payload interface{}
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewBalanceUpdatedEvent() *BalanceUpdatedEvent {
	// criando
	return &BalanceUpdatedEvent{
		Name: "BalanceUpdated",
	}
}

// getters e setters
// --------------------------------------------------
func (e *BalanceUpdatedEvent) GetName() string {
	return e.Name
}

func (e *BalanceUpdatedEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *BalanceUpdatedEvent) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *BalanceUpdatedEvent) GetDateTime() time.Time {
	return time.Now()
}

// --------------------------------------------------
