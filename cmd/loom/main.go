// main.go
package main

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
	"loom-api/api/infrastructure/httpclient"
	"loom-api/api/infrastructure/rest"
	restPoc "loom-api/api/infrastructure/rest/poc"
	"loom-api/api/infrastructure/temporal"
)

func main() {
	// Crear cliente Temporal
	temporalClient, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalf("No se pudo conectar a Temporal: %v", err)
	}
	defer temporalClient.Close()

	// Instanciar activities con sus dependencias
	httpClient := &httpclient.RestyHttpClient{}
	activities := &temporal.Activities{
		HttpClient: httpClient,
	}

	// Crear y arrancar worker (workflow + activities)
	w := worker.New(temporalClient, "load-test-task-queue", worker.Options{})

	w.RegisterWorkflow(activities.LoadTestWorkflow)
	w.RegisterActivity(activities.HttpGetActivity)
	w.RegisterActivity(activities.ProcessResponseActivity)

	go func() {
		if err := w.Run(worker.InterruptCh()); err != nil {
			log.Fatalf("Worker terminó con error: %v", err)
		}
	}()

	// Crear handler HTTP con cliente Temporal para disparar workflows
	handler := restPoc.NewHandler(temporalClient)

	// Levantar servidor HTTP usando función StartServer de rest
	if err := rest.StartServer(":8081", handler); err != nil {
		log.Fatalf("Error iniciando servidor HTTP: %v", err)
	}
}
