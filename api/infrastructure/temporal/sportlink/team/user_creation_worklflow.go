package team

import (
	"go.temporal.io/sdk/workflow"
	"loom-api/api/application"
	"loom-api/api/domain/sportlink/team"
	"time"
)

type CreationWorkflow struct {
	CreateTeamActivity   application.UseCase[team.Entity, team.Entity]
	RetrieveTeamActivity application.UseCase[team.ID, team.Entity]
}

func (c *CreationWorkflow) InvokeCreationWorkflow(ctx workflow.Context) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	err := workflow.ExecuteActivity(ctx, "CreateTeamActivity", team.Entity{
		ID:       "Boca",
		Sport:    team.Football,
		Category: team.L1,
	}).Get(ctx, nil)

	workflow.GetLogger(ctx).Info("Team created. Now searching....")

	if err != nil {
		return err
	}

	err = workflow.ExecuteActivity(ctx, "RetrieveTeamActivity", team.ID{
		Name:  "Boca",
		Sport: team.Football,
	}).Get(ctx, nil)

	if err != nil {
		return err
	}

	workflow.GetLogger(ctx).Info("Team found.")

	return nil
}
