// nome do pacote (está sendo utilizado o nome da referida pasta)
package main

// dependências
import (
	"database/sql"
	"fmt"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/event"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/repository"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/usecase/create_account"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/usecase/create_client"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/usecase/create_transaction"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/pkg/events"

	_ "github.com/go-sql-driver/mysql"
)

// função responsável pela criação de um servidor web
// versão sem o gerenciamento da transação com unity of work - uow
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

	// configMap := ckafka.ConfigMap{
	// 	"bootstrap.servers": "kafka:29092",
	// 	"group.id":          "wallet",
	// }
	// kafkaProducer := kafka.NewKafkaProducer(&configMap)

	eventDispatcher := events.NewEventDispatcher()
	// eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	// eventDispatcher.Register("BalanceUpdated", handler.NewUpdateBalanceKafkaHandler(kafkaProducer))
	// transactionCreatedEvent := event.NewTransactionCreated()
	transactionCreatedEvent := event.NewTransactionCreatedEvent()
	// balanceUpdatedEvent := event.NewBalanceUpdated()

	clientDb := repository.NewClientRepository(db)
	accountDb := repository.NewAccountRepository(db)
	transactionDb := repository.NewTransactionRepository(db)

	// ctx := context.Background()
	// uow := uow.NewUow(ctx, db)

	// uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
	// 	return repository.NewAccountRepository(db)
	// })

	// uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
	// 	return repository.NewTransactionRepository(db)
	// })
	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(transactionDb, accountDb, eventDispatcher, transactionCreatedEvent)
	// createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent, balanceUpdatedEvent)

	webserver := webserver.NewWebServer(":8080")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Server is running")
	webserver.Start()
}
