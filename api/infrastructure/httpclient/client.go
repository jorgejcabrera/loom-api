package httpclient

import (
	"context"
	"github.com/go-resty/resty/v2"
	"loom-api/api/application"
	"loom-api/api/domain/poc"
)

var _ application.HttpClient = (*RestyHttpClient)(nil)

type RestyHttpClient struct{}

func (c *RestyHttpClient) Get(ctx context.Context, url string) (*poc.HttpResponse, error) {
	client := resty.New()
	resp, err := client.R().Get(url)
	if err != nil {
		return nil, err
	}

	return &poc.HttpResponse{
		StatusCode: resp.StatusCode(),
		Body:       string(resp.Body()),
	}, nil
}
