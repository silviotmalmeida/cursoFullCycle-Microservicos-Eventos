// nome do pacote (está sendo utilizado o nome da referida pasta)
package events

// dependências
import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// criando um struct que implemente a EventInterface
// ----------------------------------------------------
type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

func (e *TestEvent) SetPayload(payload interface{}) {
	e.Payload = payload
}

// ----------------------------------------------------

// criando um struct que implemente a EventHandlerInterface
// ----------------------------------------------------
type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
}

// ----------------------------------------------------

// definindo a suíte de testes
type EventDispatcherTestSuite struct {
	// definindo os atributos e seus tipos
	suite.Suite
	event1          TestEvent
	event2          TestEvent
	handler1        TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcher
}

// função de criação da suíte
// será executado antes de cada teste da suíte
func (suite *EventDispatcherTestSuite) SetupTest() {
	// setando o dispatcher
	suite.eventDispatcher = NewEventDispatcher()
	// setando os handlers
	suite.handler1 = TestEventHandler{
		ID: 1,
	}
	suite.handler2 = TestEventHandler{
		ID: 2,
	}
	suite.handler3 = TestEventHandler{
		ID: 3,
	}
	// setando os eventos
	suite.event1 = TestEvent{Name: "test", Payload: "test"}
	suite.event2 = TestEvent{Name: "test2", Payload: "test2"}
}

// inicializando a suíte como um teste geral
func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}

// testes de unidade do dispatcher

// teste de registro com sucesso
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	// registrando o event 1 com o handler 1
	err := suite.eventDispatcher.Register(suite.event1.GetName(), &suite.handler1)
	// não deve retornar erro
	suite.Nil(err)
	// a quantidade de handlers para o event 1 deve ser 1
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	// registrando o event 1 com o handler 2
	err = suite.eventDispatcher.Register(suite.event1.GetName(), &suite.handler2)
	// não deve retornar erro
	suite.Nil(err)
	// a quantidade de handlers para o event 1 deve ser 2
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	// no array de handlers do event 1, o handler 1 deve estar na posição 0
	suite.Equal(&suite.handler1, suite.eventDispatcher.handlers[suite.event1.GetName()][0])
	// no array de handlers do event 1, o handler 1 deve estar na posição 1
	suite.Equal(&suite.handler2, suite.eventDispatcher.handlers[suite.event1.GetName()][1])
}

// teste de registro sem sucesso, com handler repetido
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	// registrando o event 1 com o handler 1
	err := suite.eventDispatcher.Register(suite.event1.GetName(), &suite.handler1)
	// não deve retornar erro
	suite.Nil(err)
	// a quantidade de handlers para o event 1 deve ser 1
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	// registrando o event 1 com o handler 1 novamente
	err = suite.eventDispatcher.Register(suite.event1.GetName(), &suite.handler1)
	// deve retornar um erro do tipo ErrHandlerAlreadyRegistered
	suite.Equal(ErrHandlerAlreadyRegistered, err)
	// a quantidade de handlers para o event 1 deve permanecer em 1
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))
}

// teste de remoção total dos registros com sucesso
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	// registrando o event 1 com o handler 1
	err := suite.eventDispatcher.Register(suite.event1.GetName(), &suite.handler1)
	// não deve retornar erro
	suite.Nil(err)
	// a quantidade de handlers para o event 1 deve ser 1
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	// registrando o event 1 com o handler 2
	err = suite.eventDispatcher.Register(suite.event1.GetName(), &suite.handler2)
	// não deve retornar erro
	suite.Nil(err)
	// a quantidade de handlers para o event 1 deve ser 2
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	// registrando o event 2 com o handler 3
	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	// não deve retornar erro
	suite.Nil(err)
	// a quantidade de handlers para o event 2 deve ser 1
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	// removendo todos os registros
	suite.eventDispatcher.Clear()
	// a quantidade total de registros deve ser 0
	suite.Equal(0, len(suite.eventDispatcher.handlers))
}

// teste de remoção total dos registros com sucesso
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	// registrando o event 1 com o handler 1
	err := suite.eventDispatcher.Register(suite.event1.GetName(), &suite.handler1)
	// não deve retornar erro
	suite.Nil(err)
	// a quantidade de handlers para o event 1 deve ser 1
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	// registrando o event 1 com o handler 2
	err = suite.eventDispatcher.Register(suite.event1.GetName(), &suite.handler2)
	// não deve retornar erro
	suite.Nil(err)
	// a quantidade de handlers para o event 1 deve ser 2
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	// o event 1 deve estar registrado com o handler 1
	suite.True(suite.eventDispatcher.Has(suite.event1.GetName(), &suite.handler1))
	// o event 1 deve estar registrado com o handler 2
	suite.True(suite.eventDispatcher.Has(suite.event1.GetName(), &suite.handler2))
	// o event 1 não deve estar registrado com o handler 3
	suite.False(suite.eventDispatcher.Has(suite.event1.GetName(), &suite.handler3))
}

// teste de remoção individual dos registros com sucesso
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	// registrando o event 1 com o handler 1
	err := suite.eventDispatcher.Register(suite.event1.GetName(), &suite.handler1)
	// não deve retornar erro
	suite.Nil(err)
	// a quantidade de handlers para o event 1 deve ser 1
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	// registrando o event 1 com o handler 2
	err = suite.eventDispatcher.Register(suite.event1.GetName(), &suite.handler2)
	// não deve retornar erro
	suite.Nil(err)
	// a quantidade de handlers para o event 1 deve ser 2
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	// registrando o event 2 com o handler 3
	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	// não deve retornar erro
	suite.Nil(err)
	// a quantidade de handlers para o event 2 deve ser 1
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	// removendo o handler 1 do event 1
	suite.eventDispatcher.Remove(suite.event1.GetName(), &suite.handler1)
	// a quantidade de handlers para o event 1 deve ser 1
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))
	// no array de handlers do event 1, o handler 1 deve estar na posição 0
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event1.GetName()][0])

	// removendo o handler 2 do event 1
	suite.eventDispatcher.Remove(suite.event1.GetName(), &suite.handler2)
	// a quantidade de handlers para o event 1 deve ser 0
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event1.GetName()]))

	// removendo o handler 3 do event 2
	suite.eventDispatcher.Remove(suite.event2.GetName(), &suite.handler3)
	// a quantidade de handlers para o event 2 deve ser 0
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))
}

// criando um mock para o handler
type MockHandler struct {
	mock.Mock
}

// definindo o comportamento do método handle do mock
func (m *MockHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

// teste de disparo das ações registradas com sucesso
func (suite *EventDispatcherTestSuite) TestEventDispatch_Dispatch() {
	// criando o handler 1 mockado
	handler1 := &MockHandler{}
	handler1.On("Handle", &suite.event1)

	// criando o handler 2 mockado
	handler2 := &MockHandler{}
	handler2.On("Handle", &suite.event1)

	// registrando os events e handlers
	suite.eventDispatcher.Register(suite.event1.GetName(), handler1)
	suite.eventDispatcher.Register(suite.event1.GetName(), handler2)

	// disparando as ações relacionadas ao event1
	suite.eventDispatcher.Dispatch(&suite.event1)
	// o método handle do handler1 deve ter sido executado 1 vez
	handler1.AssertNumberOfCalls(suite.T(), "Handle", 1)
	// o método handle do handler2 deve ter sido executado 1 vez
	handler2.AssertNumberOfCalls(suite.T(), "Handle", 1)
}
