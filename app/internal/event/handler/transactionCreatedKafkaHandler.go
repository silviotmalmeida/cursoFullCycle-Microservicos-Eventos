// nome do pacote (está sendo utilizado o nome da referida pasta)
package handler

// dependências
import (
	"fmt"
	"sync"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/pkg/events"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/pkg/kafka"
)

// definindo a estrutura (similar à classe)
type TransactionCreatedKafkaHandler struct {
	// definindo os atributos e seus tipos
	Kafka *kafka.Producer
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewTransactionCreatedKafkaHandler(kafka *kafka.Producer) *TransactionCreatedKafkaHandler {
	// criando
	return &TransactionCreatedKafkaHandler{
		Kafka: kafka,
	}
}

// função de publicação no kafka
func (h *TransactionCreatedKafkaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	// encerra o waitgroup
	defer wg.Done()
	// publica a mensagem no tópico
	h.Kafka.Publish(message, nil, "transactions")
	// imprime no terminal (somente para visualização)
	fmt.Println("TransactionCreatedKafkaHandler: ", message.GetPayload())
}
