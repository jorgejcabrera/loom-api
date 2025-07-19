package activity

import (
	"context"
	"github.com/go-resty/resty/v2"
	"log"
)

type HttpResponse struct {
	StatusCode int
	Body       string
}

func HttpGetActivity(ctx context.Context, url string) (*HttpResponse, error) {
	client := resty.New()
	resp, err := client.R().Get(url)
	if err != nil {
		return nil, err
	}

	return &HttpResponse{
		StatusCode: resp.StatusCode(),
		Body:       string(resp.Body()),
	}, nil
}

func ProcessResponseActivity(ctx context.Context, resp *HttpResponse) error {
	// Por ejemplo, imprimimos info o hacemos algo con la respuesta
	log.Printf("Procesando response - Status: %d, Body length: %d\n", resp.StatusCode, len(resp.Body))
	return nil
}
