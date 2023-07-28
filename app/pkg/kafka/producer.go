// nome do pacote (está sendo utilizado o nome da referida pasta)
package kafka

// dependências
import (
	"encoding/json"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// definindo a estrutura (similar à classe)
type Producer struct {
	// definindo os atributos e seus tipos
	ConfigMap *ckafka.ConfigMap
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewKafkaProducer(configMap *ckafka.ConfigMap) *Producer {
	// criando
	return &Producer{ConfigMap: configMap}
}

// função de publicação de mensagens nos tópicos
func (p *Producer) Publish(msg interface{}, key []byte, topic string) error {
	// criando o producer
	producer, err := ckafka.NewProducer(p.ConfigMap)
	// em caso de erro, retorna o erro
	if err != nil {
		return err
	}
	// convertendo a mensagem para json
	msgJson, err := json.Marshal(msg)
	// em caso de erro, retorna o erro
	if err != nil {
		return err
	}
	// preenchendo a mensagem no padrão do kafka
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          msgJson,
		Key:            key,
	}
	// produz a mensagem
	err = producer.Produce(message, nil)
	// em caso de erro, encerra a execução
	if err != nil {
		panic(err)
	}
	// não retorna erro
	return nil
}
