package main

import (
	"github.com/go-resty/resty/v2"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
	"loom-api/api/application/sportlink/team/usecases"
	iteam "loom-api/api/infrastructure/persistence/sportlink/team"
	"loom-api/api/infrastructure/rest"
	"loom-api/api/infrastructure/rest/doc"
	rteam "loom-api/api/infrastructure/rest/sportlink/team"
	"loom-api/api/infrastructure/temporal/sportlink/team"
)

func main() {
	temporalClient, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalf("No se pudo conectar a Temporal: %v", err)
	}
	defer temporalClient.Close()

	// SportLink
	teamRepository := iteam.NewRepository(resty.New())
	searchRepository := iteam.NewSearchRepository(resty.New())
	createTeamUC := usecases.NewCreateTeamUC(teamRepository)
	retrieveTeamUc := usecases.NewRetrieveTeamUC(searchRepository)

	// SportLink Workflows
	creationWorkflow := team.CreationWorkflow{
		CreateTeamActivity:   createTeamUC,
		RetrieveTeamActivity: retrieveTeamUc,
	}
	sW := worker.New(temporalClient, "sportlink-task-queue", worker.Options{})
	sW.RegisterWorkflow(creationWorkflow.InvokeCreationWorkflow)
	sW.RegisterActivityWithOptions(creationWorkflow.CreateTeamActivity.Invoke, activity.RegisterOptions{
		Name: "CreateTeamActivity",
	})
	sW.RegisterActivityWithOptions(creationWorkflow.RetrieveTeamActivity.Invoke, activity.RegisterOptions{
		Name: "RetrieveTeamActivity",
	})

	go func() {
		if err := sW.Run(worker.InterruptCh()); err != nil {
			log.Fatalf("Worker terminó con error: %v", err)
		}
	}()

	teamHandler := rteam.NewHandler(temporalClient)
	docHandler := doc.NewHandler("docs/")
	router := rest.NewRouter(teamHandler, docHandler)

	if err := rest.StartServer(":8081", router); err != nil {
		log.Fatalf("Error iniciando servidor HTTP: %v", err)
	}
}
