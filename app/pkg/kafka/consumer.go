// nome do pacote (está sendo utilizado o nome da referida pasta)
package kafka

// dependências
import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

// definindo a estrutura (similar à classe)
type Consumer struct {
	// definindo os atributos e seus tipos
	ConfigMap *ckafka.ConfigMap
	Topics    []string
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	// criando
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

// ????
func (c *Consumer) Consume(msgChan chan *ckafka.Message) error {
	consumer, err := ckafka.NewConsumer(c.ConfigMap)
	if err != nil {
		panic(err)
	}
	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		panic(err)
	}
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			msgChan <- msg
		}
	}
}
