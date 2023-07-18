// nome do pacote (está sendo utilizado o nome da referida pasta)
package web

// dependências
import (
	"encoding/json"
	"net/http"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/usecase/create_transaction"
)

// definindo a estrutura (similar à classe)
// responsável por tratar as requisições do endpoint "/transactions"
type WebTransactionHandler struct {
	// definindo os atributos e seus tipos
	CreateTransactionUseCase create_transaction.CreateTransactionUseCase
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewWebTransactionHandler(createTransactionUseCase create_transaction.CreateTransactionUseCase) *WebTransactionHandler {
	// criando
	return &WebTransactionHandler{
		CreateTransactionUseCase: createTransactionUseCase,
	}
}

// função responsável por criar uma transaction
func (h *WebTransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	// inicializando a variável de input
	var dto create_transaction.CreateTransactionInputDTO
	// populando o input a partir dos dados do body do request
	err := json.NewDecoder(r.Body).Decode(&dto)
	// em caso de erro
	if err != nil {
		// preenche o header do response com o erro
		w.WriteHeader(http.StatusBadRequest)
		// encerra
		return
	}
	// ctx := r.Context()
	// output, err := h.CreateTransactionUseCase.Execute(ctx, dto)
	// executando o usecase
	output, err := h.CreateTransactionUseCase.Execute(&dto)
	// em caso de erro
	if err != nil {
		// preenche o header do response com o erro
		w.WriteHeader(http.StatusBadRequest)
		// w.Write([]byte(err.Error()))
		// encerra
		return
	}
	// preenche o header do response com o content type
	w.Header().Set("Content-Type", "application/json")
	// preenche o response com o output
	err = json.NewEncoder(w).Encode(output)
	// em caso de erro
	if err != nil {
		// preenche o header do response com o erro
		w.WriteHeader(http.StatusInternalServerError)
		// encerra
		return
	}
	// preenche o header do response com sucesso
	w.WriteHeader(http.StatusCreated)
}
