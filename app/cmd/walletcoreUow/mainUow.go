// nome do pacote (está sendo usado main pois trata-se de entrypoint da aplicação)
package main

// dependências
import (
	"context"
	"database/sql"
	"fmt"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/event"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/event/handler"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/repository"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/usecase/create_account"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/usecase/create_client"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/usecase/create_transaction_uow"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/web"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/web/webserver"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/pkg/events"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/pkg/kafka"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/pkg/uow"

	_ "github.com/go-sql-driver/mysql"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// função responsável pela criação de um servidor web
// versão com o gerenciamento da transação com unity of work - uow
func main() {
	// criando a conexão com o bd
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	// em caso de erro
	if err != nil {
		// encerra a conexão
		panic(err)
	}
	// no fim da execução, fecha a conexão
	defer db.Close()

	// iniciando o kafka
	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	// criando o producer
	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	// criando o gerenciador de eventos
	eventDispatcher := events.NewEventDispatcher()

	// criando os eventos
	transactionCreatedEvent := event.NewTransactionCreatedEvent()
	balanceUpdatedEvent := event.NewBalanceUpdatedEvent()

	// registrando os eventos e respectivos handlers
	eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	eventDispatcher.Register("BalanceUpdated", handler.NewUpdateBalanceKafkaHandler(kafkaProducer))

	// criando os repositories
	clientDb := repository.NewClientRepository(db)
	accountDb := repository.NewAccountRepository(db)

	// criando o centext
	ctx := context.Background()
	// criando o gerenciador da transação uow
	uow := uow.NewUow(ctx, db)

	// registrando os repositories no uow
	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return repository.NewAccountRepository(db)
	})
	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return repository.NewTransactionRepository(db)
	})

	// criando os usecases
	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUowUseCase := create_transaction_uow.NewCreateTransactionUowUseCase(uow, eventDispatcher, transactionCreatedEvent, balanceUpdatedEvent)

	// criando o webserver e definindo a porta a ser utilizada
	port := "8080"
	webserver := webserver.NewWebServer(":" + port)

	// criando os handlers dos endpoints
	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionUowHandler(*createTransactionUowUseCase)

	// adicionando os handlers ao webserver e configurando os endpoints
	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	// inicializando o webserver
	fmt.Println("Server is running on port", port)
	webserver.Start()
}
