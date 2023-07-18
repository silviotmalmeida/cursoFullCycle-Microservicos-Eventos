// nome do pacote (está sendo utilizado o nome da referida pasta)
package web

// dependências
import (
	"encoding/json"
	"net/http"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/usecase/create_account"
)

// definindo a estrutura (similar à classe)
// responsável por tratar as requisições do endpoint "/accounts"
type WebAccountHandler struct {
	// definindo os atributos e seus tipos
	CreateAccountUseCase create_account.CreateAccountUseCase
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewWebAccountHandler(createAccountUseCase create_account.CreateAccountUseCase) *WebAccountHandler {
	// criando
	return &WebAccountHandler{
		CreateAccountUseCase: createAccountUseCase,
	}
}

// função responsável por criar um account
func (h *WebAccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	// inicializando a variável de input
	var dto create_account.CreateAccountInputDTO
	// populando o input a partir dos dados do body do request
	err := json.NewDecoder(r.Body).Decode(&dto)
	// em caso de erro
	if err != nil {
		// preenche o header do response com o erro
		w.WriteHeader(http.StatusBadRequest)
		// encerra
		return
	}
	// executando o usecase
	output, err := h.CreateAccountUseCase.Execute(&dto)
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
