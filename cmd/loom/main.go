package main

import (
	"github.com/go-resty/resty/v2"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
	"loom-api/api/application/sportlink/team/usecases"
	iteam "loom-api/api/infrastructure/persistence/sportlink/team"
	"loom-api/api/infrastructure/rest"
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
	createTeamUC := usecases.NewCreateTeamUC(teamRepository)

	// SportLink Workflows
	creationWorkflow := team.CreationWorkflow{
		CreateTeamActivity: createTeamUC,
	}
	sW := worker.New(temporalClient, "sportlink-task-queue", worker.Options{})
	sW.RegisterWorkflow(creationWorkflow.InvokeCreationWorkflow)
	sW.RegisterActivity(creationWorkflow.CreateTeamActivity.Invoke)

	go func() {
		if err := sW.Run(worker.InterruptCh()); err != nil {
			log.Fatalf("Worker terminó con error: %v", err)
		}
	}()

	teamHandler := rteam.NewHandler(temporalClient)
	router := rest.NewRouter(teamHandler)

	if err := rest.StartServer(":8081", router); err != nil {
		log.Fatalf("Error iniciando servidor HTTP: %v", err)
	}
}
