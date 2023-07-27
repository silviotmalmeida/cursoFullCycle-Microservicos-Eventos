// nome do pacote (está sendo utilizado o nome da referida pasta)
package mocks

// dependências
import (
	"context"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/pkg/uow"
	"github.com/stretchr/testify/mock"
)

// definindo o mock
type UowMock struct {
	mock.Mock
}

// definindo os comportamentos do mock
// -----------------------------------------------------------------
func (m *UowMock) Register(name string, fc uow.RepositoryFactory) {
	m.Called(name, fc)
}

func (m *UowMock) GetRepository(ctx context.Context, name string) (interface{}, error) {
	args := m.Called(name)
	return args.Get(0), args.Error(1)
}

func (m *UowMock) Do(ctx context.Context, fn func(uow *uow.Uow) error) error {
	args := m.Called(fn)
	return args.Error(0)
}

func (m *UowMock) CommitOrRollback() error {
	args := m.Called()
	return args.Error(0)
}

func (m *UowMock) Rollback() error {
	args := m.Called()
	return args.Error(0)
}

func (m *UowMock) UnRegister(name string) {
	m.Called(name)
}

// -----------------------------------------------------------------
