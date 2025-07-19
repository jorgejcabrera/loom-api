package team

import (
	"go.temporal.io/sdk/workflow"
	"loom-api/api/application"
	"loom-api/api/domain/sportlink/team"
	"time"
)

type CreationWorkflow struct {
	CreateTeamActivity application.UseCase[team.Entity, team.Entity]
}

func (c *CreationWorkflow) InvokeCreationWorkflow(ctx workflow.Context) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	err := workflow.ExecuteActivity(ctx, c.CreateTeamActivity.Invoke, team.Entity{
		ID:       "Boca",
		Sport:    team.Football,
		Category: team.L1,
	}).Get(ctx, nil)

	if err != nil {
		return err
	}

	return nil
}
