package temporal

import (
	"context"
	"log"
	"loom-api/api/application"
	"loom-api/api/domain/poc"
)

// Actividades reciben las dependencias necesarias via struct

type Activities struct {
	HttpClient application.HttpClient
}

func (a *Activities) HttpGetActivity(ctx context.Context, url string) (*poc.HttpResponse, error) {
	return a.HttpClient.Get(ctx, url)
}

func (a *Activities) ProcessResponseActivity(ctx context.Context, resp *poc.HttpResponse) error {
	log.Printf("Procesando respuesta: Status=%d, Body length=%d\n", resp.StatusCode, len(resp.Body))
	return nil
}
