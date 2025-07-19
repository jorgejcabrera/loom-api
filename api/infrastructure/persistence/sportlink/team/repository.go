package team

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"loom-api/api/domain/sportlink/team"
)

type RepositoryAdapter struct {
	client *resty.Client
}

func NewRepository(client *resty.Client) team.Repository {
	return &RepositoryAdapter{
		client: client,
	}
}

func (r RepositoryAdapter) Save(entity team.Entity) error {
	resp, err := r.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(mapRequest(entity)).
		Post("http://localhost:8080/team")

	if err != nil {
		return fmt.Errorf("error haciendo POST a /team: %w", err)
	}

	if resp.IsError() {
		return fmt.Errorf("respuesta con error del servidor: %s", resp.Status())
	}

	return nil
}

func mapRequest(entity team.Entity) map[string]interface{} {
	return map[string]interface{}{
		"name":     entity.ID,
		"category": entity.Category,
		"sport":    string(entity.Sport),
	}
}
