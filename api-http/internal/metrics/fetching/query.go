package fetchingmetric

import (
	"context"
	"errors"
	"example.com/gotraining/go-hexagonal_http_api-course/kit/bus"
)

const MetricQueryType bus.Type = "bus.fetching.metrics"

// MetricQuery is the bus dispatched to create a new Metric.
type MetricQuery struct {
}

// NewFetchMetricQuery creates a new MetricQuery.
func NewFetchMetricQuery() MetricQuery {
	return MetricQuery{}
}

func (c MetricQuery) Type() bus.Type {
	return MetricQueryType
}

// MetricQueryHandler is the bus handler
// responsible for fetching Metrics.
type MetricQueryHandler struct {
	service FetchingService
}

// NewMetricQueryHandler initializes a new NewMetricQueryHandler.
func NewMetricQueryHandler(service FetchingService) MetricQueryHandler {
	return MetricQueryHandler{
		service: service,
	}
}

// Handle implements the bus.QueryHandler interface.
func (h MetricQueryHandler) Handle(ctx context.Context, query bus.Query) (bus.QueryResponse, error) {
	_, ok := query.(MetricQuery)
	if !ok {
		return nil, errors.New("unexpected bus")
	}

	return h.service.GetAll(ctx)
}
