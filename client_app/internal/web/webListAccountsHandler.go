// nome do pacote (está sendo utilizado o nome da referida pasta)
package web

// dependências
import (
	"encoding/json"
	"net/http"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/usecase/list_accounts"
)

// definindo a estrutura (similar à classe)
// responsável por tratar as requisições do endpoint "/accounts"
type WebListAccountsHandler struct {
	// definindo os atributos e seus tipos
	ListAccountsUseCase list_accounts.ListAccountsUseCase
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewWebListAccountsHandler(createAccountUseCase list_accounts.ListAccountsUseCase) *WebListAccountsHandler {
	// criando
	return &WebListAccountsHandler{
		ListAccountsUseCase: createAccountUseCase,
	}
}

// função responsável por listar os accounts
func (h *WebListAccountsHandler) ListAccounts(w http.ResponseWriter, r *http.Request) {
	// executando o usecase
	output, err := h.ListAccountsUseCase.Execute()
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
