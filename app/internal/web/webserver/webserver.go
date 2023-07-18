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
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewWebServer(webServerPort string) *WebServer {
	// criando
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: webServerPort,
	}
}

// função para registro dos handlers e endpoints
func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	// definindo o handler para o endpoint informado
	s.Handlers[path] = handler
}

// função para inicializar o webserver
func (s *WebServer) Start() {
	// utilizando o middleware de log
	s.Router.Use(middleware.Logger)
	// carregando os handlers e endpoints registrados e definindo o método como post
	for path, handler := range s.Handlers {
		s.Router.Post(path, handler)
	}
	// inicializando o webserver na porta informada
	http.ListenAndServe(s.WebServerPort, s.Router)
}
