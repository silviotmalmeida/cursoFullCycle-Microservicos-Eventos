// nome do pacote (está sendo utilizado o nome da referida pasta)
package event

// dependências
import "time"

// definindo a estrutura (similar à classe) que implemente a EventInterface
type TransactionCreatedEvent struct {
	// definindo os atributos e seus tipos
	Name    string
	Payload interface{}
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewTransactionCreatedEvent() *TransactionCreatedEvent {
	// criando
	return &TransactionCreatedEvent{
		Name: "TransactionCreated",
	}
}

// getters e setters
// --------------------------------------------------
func (e *TransactionCreatedEvent) GetName() string {
	return e.Name
}

func (e *TransactionCreatedEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TransactionCreatedEvent) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *TransactionCreatedEvent) GetDateTime() time.Time {
	return time.Now()
}

// --------------------------------------------------
