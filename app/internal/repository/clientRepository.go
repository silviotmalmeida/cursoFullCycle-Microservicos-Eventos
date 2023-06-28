// nome do pacote (está sendo utilizado o nome da referida pasta)
package repository

// dependências
import (
	"database/sql"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/entity"
)

// definindo a estrutura (similar à classe)
type ClientRepository struct {
	// definindo os atributos e seus tipos
	DB *sql.DB
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewClientRepository(db *sql.DB) *ClientRepository {
	// criando
	return &ClientRepository{
		DB: db,
	}
}

// função de busca por id
// devem ser descritos a estrutura associada, os argumentos e retornos
func (c *ClientRepository) Get(id string) (*entity.Client, error) {
	// criando um client vazio
	client := &entity.Client{}
	// abrindo a conexão e preparando a query
	stmt, err := c.DB.Prepare("SELECT id, name, email, created_at FROM clients WHERE id = ?")
	// em caso de erro, retorna somente o erro
	if err != nil {
		return nil, err
	}
	// deve-se fechar a conexão ao final da função
	defer stmt.Close()
	// realizando a query
	row := stmt.QueryRow(id)
	// populando os dados no client
	err = row.Scan(&client.ID, &client.Name, &client.Email, &client.CreatedAt)
	// em caso de erro, retorna somente o erro
	if err != nil {
		return nil, err
	}
	// senão retorna o client populado
	return client, nil
}

// função de criação
// devem ser descritos a estrutura associada, os argumentos e retornos
func (c *ClientRepository) Save(client *entity.Client) error {
	// abrindo a conexão e preparando a query
	stmt, err := c.DB.Prepare("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)")
	// em caso de erro, retorna somente o erro
	if err != nil {
		return err
	}
	// deve-se fechar a conexão ao final da função
	defer stmt.Close()
	// realizando a query
	_, err = stmt.Exec(client.ID, client.Name, client.Email, client.CreatedAt)
	// em caso de erro, retorna somente o erro
	if err != nil {
		return err
	}
	// não retorna erro
	return nil
}
