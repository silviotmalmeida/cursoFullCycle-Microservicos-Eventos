// nome do pacote (está sendo usado main pois trata-se de entrypoint da aplicação)
package main

// dependências
import (
	"database/sql"
	"fmt"

	// "github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/event"
	// "github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/event/handler"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/repository"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/usecase/create_account"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/usecase/list_accounts"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/usecase/update_balance"

	// "github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/usecase/create_client"
	// "github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/usecase/create_transaction"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/web"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/web/webserver"

	// "github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/pkg/events"
	// "github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/pkg/kafka"

	_ "github.com/go-sql-driver/mysql"
)

// função responsável pela criação de um servidor web
// versão sem o gerenciamento da transação com unity of work - uow
func main() {
	// criando a conexão com o bd
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql-client", "3306", "wallet"))
	// em caso de erro
	if err != nil {
		// encerra a conexão
		panic(err)
	}
	// no fim da execução, fecha a conexão
	defer db.Close()

	// // iniciando o kafka
	// configMap := ckafka.ConfigMap{
	// 	"bootstrap.servers": "kafka:29092",
	// 	"group.id":          "wallet",
	// }
	// // criando o producer
	// kafkaProducer := kafka.NewKafkaProducer(&configMap)

	// // criando o gerenciador de eventos
	// eventDispatcher := events.NewEventDispatcher()

	// // criando o evento de transactionCreated
	// transactionCreatedEvent := event.NewTransactionCreatedEvent()
	// balanceUpdatedEvent := event.NewBalanceUpdatedEvent()

	// // registrando os eventos e respectivos handlers
	// eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	// eventDispatcher.Register("BalanceUpdated", handler.NewUpdateBalanceKafkaHandler(kafkaProducer))

	// criando os repositories
	accountDb := repository.NewAccountRepository(db)

	// criando os usecases
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb)
	updateBalanceUseCase := update_balance.NewUpdateBalanceUseCase(accountDb)
	listAccountsUseCase := list_accounts.NewListAccountsUseCase(accountDb)

	// criando o webserver e definindo a porta a ser utilizada
	port := "3003"
	webserver := webserver.NewWebServer(":" + port)

	// criando os handlers dos endpoints
	createAccountHandler := web.NewWebCreateAccountHandler(*createAccountUseCase)
	updateBalanceHandler := web.NewWebUpdateBalanceHandler(*updateBalanceUseCase)
	listAccountsHandler := web.NewWebListAccountsHandler(*listAccountsUseCase)

	// adicionando os handlers ao webserver e configurando os endpoints
	webserver.AddHandler("/create-account", createAccountHandler.CreateAccount)
	webserver.AddHandler("/update-balance", updateBalanceHandler.UpdateBalance)
	webserver.AddHandler("/accounts", listAccountsHandler.ListAccounts)

	// inicializando o webserver
	fmt.Println("Server is running on port", port)
	webserver.Start()
}
