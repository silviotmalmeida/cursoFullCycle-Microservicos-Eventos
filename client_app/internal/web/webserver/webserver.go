// nome do pacote (está sendo utilizado o nome da referida pasta)
package webserver

// dependências
import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

// definindo a estrutura (similar à classe)
type WebServer struct {
	// definindo os atributos e seus tipos
	Router        chi.Router
	HandlersGET      map[string]http.HandlerFunc
	HandlersPOST      map[string]http.HandlerFunc
	WebServerPort string
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewWebServer(webServerPort string) *WebServer {
	// criando
	return &WebServer{
		Router:        chi.NewRouter(),
		HandlersGET:      make(map[string]http.HandlerFunc),
		HandlersPOST:      make(map[string]http.HandlerFunc),
		WebServerPort: webServerPort,
	}
}

// função para registro dos handlers e endpoints
func (s *WebServer) AddHandler(path string, handler http.HandlerFunc, method string) {

	// populando o array conforme o verbo http utilizado
	switch method {

		case "GET":
			s.HandlersGET[path] = handler

		case "POST":
			s.HandlersPOST[path] = handler
	}
}

// função para inicializar o webserver
func (s *WebServer) Start() {
	// utilizando o middleware de log
	s.Router.Use(middleware.Logger)
	// carregando os handlers e endpoints registrados
	for path, handler := range s.HandlersGET {
		s.Router.Get(path, handler)
	}
	for path, handler := range s.HandlersPOST {
		s.Router.Post(path, handler)
	}
	// inicializando o webserver na porta informada
	http.ListenAndServe(s.WebServerPort, s.Router)
}
