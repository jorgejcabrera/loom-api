package application

import (
	"context"
	"loom-api/api/domain/poc"
)

// Definimos interfaces que nuestra app necesita

type HttpClient interface {
	Get(ctx context.Context, url string) (*poc.HttpResponse, error)
}
