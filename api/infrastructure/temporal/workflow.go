package temporal

import (
	"go.temporal.io/sdk/workflow"
	"loom-api/api/domain/poc"
	"time"
)

func (a *Activities) LoadTestWorkflow(ctx workflow.Context, p poc.Poc) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	for i := 0; i < p.Count; i++ {
		var resp *poc.HttpResponse
		err := workflow.ExecuteActivity(ctx, a.HttpGetActivity, p.URL).Get(ctx, &resp)
		if err != nil {
			return err
		}
		err = workflow.ExecuteActivity(ctx, a.ProcessResponseActivity, resp).Get(ctx, nil)
		if err != nil {
			return err
		}
	}
	return nil
}
