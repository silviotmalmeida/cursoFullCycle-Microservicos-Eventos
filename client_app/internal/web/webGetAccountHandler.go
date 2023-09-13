// nome do pacote (está sendo utilizado o nome da referida pasta)
package web

// dependências
import (
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/usecase/get_account"
)

// definindo a estrutura (similar à classe)
// responsável por tratar as requisições do endpoint "/accounts/{id}"
type WebGetAccountHandler struct {
	// definindo os atributos e seus tipos
	GetAccountUseCase get_account.GetAccountUseCase
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewWebGetAccountHandler(createAccountUseCase get_account.GetAccountUseCase) *WebGetAccountHandler {
	// criando
	return &WebGetAccountHandler{
		GetAccountUseCase: createAccountUseCase,
	}
}

// função responsável por listar os accounts
func (h *WebGetAccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	// obtendo o valor do  id passado na URL
	id := chi.URLParam(r, "id")
	// definindo o input do usecase
	input := &get_account.GetAccountInputDTO{
		ID: id,
	}
	// executando o usecase
	output, err := h.GetAccountUseCase.Execute(input)
	// em caso de erro
	if err != nil {
		// preenche o header do response com o erro
		w.WriteHeader(http.StatusInternalServerError)
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
