// nome do pacote (está sendo utilizado o nome da referida pasta)
package events

// dependências
import (
	"errors"
	"sync"
)

// criando um erro customizado para handler já registrado
var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

// definindo a estrutura (similar à classe) que implemente a EventDispatcherInterface
type EventDispatcher struct {
	// definindo os atributos e seus tipos
	// definindo um map com chave=string e valor=array de handlers
	handlers map[string][]EventHandlerInterface
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewEventDispatcher() *EventDispatcher {
	// criando
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

// função de registro de handler
// devem ser descritos a estrutura associada, os argumentos e retornos
func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	// se o evento já estiver registrado no map
	if _, ok := ed.handlers[eventName]; ok {
		// itera sobre o array de handlers deste evento
		for _, h := range ed.handlers[eventName] {
			// se o handler também já estiver registrado
			if h == handler {
				// retorna o erro customizado
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	// senão, adiciona o handler ao evento
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	// não retorna erro
	return nil
}

// função de disparo de evento
// devem ser descritos a estrutura associada, os argumentos e retornos
func (ev *EventDispatcher) Dispatch(event EventInterface) error {
	// se existirem handlers registrados para o evento
	if handlers, ok := ev.handlers[event.GetName()]; ok {
		// ????
		wg := &sync.WaitGroup{}
		// itera sobre o array de handlers deste evento
		for _, handler := range handlers {
			// ?????
			wg.Add(1)
			// executa uma nova thread com o processo do handler
			go handler.Handle(event, wg)
		}
		// ????
		wg.Wait()
	}
	// não retorna erro
	return nil
}

// função de verificação de existência de um handler registrado
// devem ser descritos a estrutura associada, os argumentos e retornos
func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	// se o evento já estiver registrado no map
	if _, ok := ed.handlers[eventName]; ok {
		// itera sobre o array de handlers deste evento
		for _, h := range ed.handlers[eventName] {
			// se o handler também já estiver registrado
			if h == handler {
				// retorna true
				return true
			}
		}
	}
	// retorna false
	return false
}

// função de remoção de handler
// devem ser descritos a estrutura associada, os argumentos e retornos
func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	// se o evento já estiver registrado no map
	if _, ok := ed.handlers[eventName]; ok {
		// itera sobre o array de handlers deste evento
		for i, h := range ed.handlers[eventName] {
			// se o handler também já estiver registrado
			if h == handler {
				// remove o handler e atualiza o array
				ed.handlers[eventName] = append(ed.handlers[eventName][:i], ed.handlers[eventName][i+1:]...)
				// não retorna erro
				return nil
			}
		}
	}
	// não retorna erro
	return nil
}

// função de remoção de todos os handlers e eventos
// devem ser descritos a estrutura associada, os argumentos e retornos
func (ed *EventDispatcher) Clear() {
	// reseta o map
	ed.handlers = make(map[string][]EventHandlerInterface)
}
