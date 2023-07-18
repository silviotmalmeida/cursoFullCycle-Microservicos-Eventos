// nome do pacote (está sendo utilizado o nome da referida pasta)
package web

// dependências
import (
	"encoding/json"
	"net/http"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/usecase/create_client"
)

// definindo a estrutura (similar à classe)
// responsável por tratar as requisições do endpoint "/clients"
type WebClientHandler struct {
	// definindo os atributos e seus tipos
	CreateClientUseCase create_client.CreateClientUseCase
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewWebClientHandler(createClientUseCase create_client.CreateClientUseCase) *WebClientHandler {
	// criando
	return &WebClientHandler{
		CreateClientUseCase: createClientUseCase,
	}
}

// função responsável por criar um client
func (h *WebClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	// inicializando a variável de input
	var dto create_client.CreateClientInputDTO
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
	output, err := h.CreateClientUseCase.Execute(&dto)
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
