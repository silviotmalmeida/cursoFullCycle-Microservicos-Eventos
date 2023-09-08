// nome do pacote (está sendo utilizado o nome da referida pasta)
package web

// dependências
import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos-Desafio/internal/usecase/update_balance"
)

// definindo a estrutura (similar à classe)
// responsável por tratar as requisições do endpoint "/update-balance"
type WebUpdateBalanceHandler struct {
	// definindo os atributos e seus tipos
	UpdateBalanceUseCase update_balance.UpdateBalanceUseCase
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewWebUpdateBalanceHandler(createAccountUseCase update_balance.UpdateBalanceUseCase) *WebUpdateBalanceHandler {
	// criando
	return &WebUpdateBalanceHandler{
		UpdateBalanceUseCase: createAccountUseCase,
	}
}

// função responsável por atualizar o balance de um account
func (h *WebUpdateBalanceHandler) UpdateBalance(w http.ResponseWriter, r *http.Request) {
	// inicializando a variável de input
	var dto update_balance.UpdateBalanceInputDTO
	// populando o input a partir dos dados do body do request
	err := json.NewDecoder(r.Body).Decode(&dto)
	// em caso de erro
	if err != nil {
		// preenche o header do response com o erro
		w.WriteHeader(http.StatusBadRequest)
		// encerra
		return
	}

	fmt.Println(&dto)
	// executando o usecase
	output, err := h.UpdateBalanceUseCase.Execute(&dto)
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
