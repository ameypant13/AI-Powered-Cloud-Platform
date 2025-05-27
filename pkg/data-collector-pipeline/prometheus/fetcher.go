package prometheus

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

type Fetcher struct {
	Client v1.API
}

func NewPrometheusFetcher(endpoint string) (*Fetcher, error) {
	client, err := api.NewClient(api.Config{Address: endpoint})
	if err != nil {
		return nil, err
	}
	return &Fetcher{Client: v1.NewAPI(client)}, nil
}

func (pf *Fetcher) FetchRange(query string, start, end time.Time, step time.Duration) (model.Matrix, error) {
	r := v1.Range{Start: start, End: end, Step: step}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, warning, err := pf.Client.QueryRange(ctx, query, r)
	if err != nil {
		return nil, fmt.Errorf("Prometheus query error: %w", err)
	}
	if len(warning) > 0 {
		fmt.Println("Prometheus warnings:", warning)
	}
	matrix, ok := result.(model.Matrix)
	if !ok {
		return nil, fmt.Errorf("Unexpected result format")
	}
	return matrix, nil
}
