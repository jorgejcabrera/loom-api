package usecases

import (
	"loom-api/api/application"
	"loom-api/api/domain/sportlink/team"
)

type RetrieveTeamUC struct {
	repository team.SearchRepository
}

func NewRetrieveTeamUC(repository team.SearchRepository) application.UseCase[team.ID, team.Entity] {
	return &RetrieveTeamUC{
		repository: repository,
	}
}

func (r RetrieveTeamUC) Invoke(id team.ID) (*team.Entity, error) {
	entity, err := r.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}
