package usecases

import (
	"fmt"
	"loom-api/api/application"
	"loom-api/api/domain/sportlink/team"
)

type CreateTeamUC struct {
	teamRepository team.Repository
}

func NewCreateTeamUC(teamRepository team.Repository) application.UseCase[team.Entity, team.Entity] {
	return &CreateTeamUC{
		teamRepository: teamRepository,
	}
}

func (uc *CreateTeamUC) Invoke(input team.Entity) (*team.Entity, error) {
	err := uc.teamRepository.Save(input)
	if err != nil {
		return nil, fmt.Errorf("error while inserting team in database: %w", err)
	}
	return &input, nil
}
