package team

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"loom-api/api/domain/sportlink/team"
)

type SearchRepositoryAdapter struct {
	client *resty.Client
}

func NewSearchRepository(client *resty.Client) team.SearchRepository {
	return &SearchRepositoryAdapter{
		client: client,
	}
}

func (s SearchRepositoryAdapter) FindByID(id team.ID) (team.Entity, error) {
	url := fmt.Sprintf("http://localhost:8080/sport/%s/team/%s", id.Sport, id.Name)
	var dto Dto
	resp, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&dto).
		Get(url)

	if err != nil {
		return team.Entity{}, fmt.Errorf("error haciendo POST a /team: %w", err)
	}

	if resp.IsError() {
		return team.Entity{}, fmt.Errorf("respuesta con error del servidor: %s", resp.Status())
	}

	return team.Entity{
		ID:       dto.Name,
		Category: team.Category(dto.Category),
		Sport:    team.Sport(dto.Sport),
	}, nil
}
