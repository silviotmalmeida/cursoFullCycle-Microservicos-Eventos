// nome do pacote (está sendo usado main pois trata-se de entrypoint da aplicação)
package main

// dependências
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/repository"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/usecase/create_account"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/usecase/get_account"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/usecase/list_accounts"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/usecase/update_balance"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/web"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/web/webserver"

	_ "github.com/go-sql-driver/mysql"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// definindo a estrutura para recebimento da mensagem do tópico
type Response struct {
	Name    string
	Payload Payload
}
type Payload struct {
	AccountIdFrom        string  `json:"account_id_from"`
	AccountIdTo          string  `json:"account_id_to"`
	BalanceAccountIdFrom float64 `json:"balance_account_id_from"`
	BalanceAccountIdTo   float64 `json:"balance_account_id_to"`
}

// função responsável pela criação de um servidor web
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

	// criando os repositories
	accountDb := repository.NewAccountRepository(db)

	// criando os usecases
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb)
	updateBalanceUseCase := update_balance.NewUpdateBalanceUseCase(accountDb)
	listAccountsUseCase := list_accounts.NewListAccountsUseCase(accountDb)
	getAccountUseCase := get_account.NewGetAccountUseCase(accountDb)

	// criando o webserver e definindo a porta a ser utilizada
	port := "3003"
	webserver := webserver.NewWebServer(":" + port)

	// criando os handlers dos endpoints
	createAccountHandler := web.NewWebCreateAccountHandler(*createAccountUseCase)
	updateBalanceHandler := web.NewWebUpdateBalanceHandler(*updateBalanceUseCase)
	listAccountsHandler := web.NewWebListAccountsHandler(*listAccountsUseCase)
	getAccountHandler := web.NewWebGetAccountHandler(*getAccountUseCase)

	// adicionando os handlers ao webserver e configurando os endpoints
	webserver.AddHandler("/create-account", createAccountHandler.CreateAccount, "POST")
	webserver.AddHandler("/update-balance", updateBalanceHandler.UpdateBalance, "POST")
	webserver.AddHandler("/balances", listAccountsHandler.ListAccounts, "GET")
	webserver.AddHandler("/balances/{id}", getAccountHandler.GetAccount, "GET")

	// inicializando o webserver
	fmt.Println("Server is running on port", port)
	go webserver.Start()

	// configurando o acesso ao kafka
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "client_app-group-" + time.Now().UTC().String(),
		"auto.offset.reset": "earliest",
	}
	// lista de tópicos a serem lidos
	topics := []string{"balances"}
	// criando o consumer
	fmt.Println("Initializing consumer...")
	consumer, err := ckafka.NewConsumer(configMap)
	if err != nil {
		panic(err)
	}
	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		panic(err)
	}
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {

			var response Response
			json.Unmarshal(msg.Value, &response)

			// verificando se a account_from já está cadastrada
			// definindo o input do usecase do get_account
			inputGetFrom := &get_account.GetAccountInputDTO{
				ID: response.Payload.AccountIdFrom,
			}
			// executando o usecase
			_, err := getAccountUseCase.Execute(inputGetFrom)
			// caso a account ainda não exista, deve-se criá-la
			if err != nil {
				// definindo o input do usecase do create_account
				inputCreateFrom := &create_account.CreateAccountInputDTO{
					ID:      response.Payload.AccountIdFrom,
					Balance: response.Payload.BalanceAccountIdFrom,
				}
				// executando o usecase
				outputCreateFrom, err := createAccountUseCase.Execute(inputCreateFrom)
				if err != nil {
					panic(err)
				}
				fmt.Println("criado account_from", outputCreateFrom)
			} else {
				// definindo o input do usecase do update_balance
				inputUpdateFrom := &update_balance.UpdateBalanceInputDTO{
					ID:      response.Payload.AccountIdFrom,
					Balance: response.Payload.BalanceAccountIdFrom,
				}
				// executando o usecase
				outputUpdateFrom, err := updateBalanceUseCase.Execute(inputUpdateFrom)
				if err != nil {
					panic(err)
				}
				fmt.Println("atualizado account_from", outputUpdateFrom)
			}

			// verificando se a account_to já está cadastrada
			// definindo o input do usecase do get_account
			inputGetTo := &get_account.GetAccountInputDTO{
				ID: response.Payload.AccountIdTo,
			}
			// executando o usecase
			_, err = getAccountUseCase.Execute(inputGetTo)
			// caso a account ainda não exista, deve-se criá-la
			if err != nil {
				// definindo o input do usecase do create_account
				inputCreateTo := &create_account.CreateAccountInputDTO{
					ID:      response.Payload.AccountIdTo,
					Balance: response.Payload.BalanceAccountIdTo,
				}
				// executando o usecase
				outputCreateTo, err := createAccountUseCase.Execute(inputCreateTo)
				if err != nil {
					panic(err)
				}
				fmt.Println("criado account_to", outputCreateTo)
			} else {
				// definindo o input do usecase do update_balance
				inputUpdateTo := &update_balance.UpdateBalanceInputDTO{
					ID:      response.Payload.AccountIdTo,
					Balance: response.Payload.BalanceAccountIdTo,
				}
				// executando o usecase
				outputUpdateTo, err := updateBalanceUseCase.Execute(inputUpdateTo)
				if err != nil {
					panic(err)
				}
				fmt.Println("atualizado account_to", outputUpdateTo)
			}
		}
	}
}
