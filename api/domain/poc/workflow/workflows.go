package workflow

import (
	"go.temporal.io/sdk/workflow"
	"loom-api/api/domain/poc/workflow/activity"
	"time"
)

func LoadTestWorkflow(ctx workflow.Context, url string, count int) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	for i := 0; i < count; i++ {
		var resp *activity.HttpResponse
		// Ejecutar Activity 1 y obtener struct
		err := workflow.ExecuteActivity(ctx, activity.HttpGetActivity, url).Get(ctx, &resp)
		if err != nil {
			return err
		}

		// Pasar struct a Activity 2
		err = workflow.ExecuteActivity(ctx, activity.ProcessResponseActivity, resp).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
