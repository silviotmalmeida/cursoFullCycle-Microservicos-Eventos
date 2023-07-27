// nome do pacote (está sendo utilizado o nome da referida pasta)
package web

// dependências
import (
	"encoding/json"
	"net/http"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/usecase/create_transaction_uow"
)

// definindo a estrutura (similar à classe)
// responsável por tratar as requisições do endpoint "/transactions"
type WebTransactionUowHandler struct {
	// definindo os atributos e seus tipos
	CreateTransactionUowUseCase create_transaction_uow.CreateTransactionUowUseCase
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewWebTransactionUowHandler(createTransactionUowUseCase create_transaction_uow.CreateTransactionUowUseCase) *WebTransactionUowHandler {
	// criando
	return &WebTransactionUowHandler{
		CreateTransactionUowUseCase: createTransactionUowUseCase,
	}
}

// função responsável por criar uma transaction
func (h *WebTransactionUowHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	// inicializando a variável de input
	var dto create_transaction_uow.CreateTransactionInputDTO
	// populando o input a partir dos dados do body do request
	err := json.NewDecoder(r.Body).Decode(&dto)
	// em caso de erro
	if err != nil {
		// preenche o header do response com o erro
		w.WriteHeader(http.StatusBadRequest)
		// encerra
		return
	}
	// criando o context
	ctx := r.Context()
	// executando o usecase
	output, err := h.CreateTransactionUowUseCase.Execute(ctx, dto)
	// em caso de erro
	if err != nil {
		// preenche o header do response com o erro
		w.WriteHeader(http.StatusBadRequest)
		// detalha o erro
		w.Write([]byte(err.Error()))
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
