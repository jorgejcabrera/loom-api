package application

import (
	"loom-api/api/domain/poc"
)

type LoadTestUseCase struct{}

// Método que representa la lógica principal, independiente de Temporal

func (u *LoadTestUseCase) RunLoadTest(httpClient HttpClient, p poc.Poc) ([]*poc.HttpResponse, error) {
	var results []*poc.HttpResponse
	for i := 0; i < p.Count; i++ {
		resp, err := httpClient.Get(nil, p.URL)
		if err != nil {
			return nil, err
		}
		results = append(results, resp)
	}
	return results, nil
}
