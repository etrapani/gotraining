package creatingmetric

import (
	"context"
	"errors"
	"example.com/gotraining/go-hexagonal_http_api-course/kit/bus"
)

const MetricCommandType bus.Type = "bus.creating.metric"

// MetricCommand is the bus dispatched to create a new metric.
type MetricCommand struct {
	url            string
	responseStatus int
	requestSize    int
	responseSize   int
	responseTime   int64
}

// NewMetricCommand creates a new MetricCommand.
func NewMetricCommand(url string, responseStatus, requestSize, responseSize int, responseTime int64) MetricCommand {
	return MetricCommand{
		url:            url,
		responseStatus: responseStatus,
		requestSize:    requestSize,
		responseSize:   responseSize,
		responseTime:   responseTime,
	}
}

func (c MetricCommand) Type() bus.Type {
	return MetricCommandType
}

// MetricCommandHandler is the bus handler
// responsible for creating metric.
type CreateMetricCommandHandler struct {
	service CreateMetricService
}

// NewMetricCommandHandler initializes a new CreateMetricCommandHandler.
func NewMetricCommandHandler(service CreateMetricService) CreateMetricCommandHandler {
	return CreateMetricCommandHandler{
		service: service,
	}
}

// Handle implements the bus.CommandHandler interface.
func (h CreateMetricCommandHandler) Handle(ctx context.Context, cmd bus.Command) error {
	metricCommand, ok := cmd.(MetricCommand)
	if !ok {
		return errors.New("unexpected bus")
	}

	return h.service.CreateMetric(
		ctx,
		metricCommand.url,
		metricCommand.responseStatus,
		metricCommand.requestSize,
		metricCommand.responseSize,
		metricCommand.responseTime,
	)
}
