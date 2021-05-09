package creatingmetric

import (
	"context"
	"example.com/gotraining/go-hexagonal_http_api-course/internal/metrics"
	"github.com/google/uuid"
)

// CreateMetricService is the default CreateMetricService interface
// implementation returned by creatingmetric.NewCourseService.
type CreateMetricService struct {
	metricRepository metrics.MetricRepository
}

// NewCreateMetricService returns the default Service interface implementation.
func NewCreateMetricService(metricRepository metrics.MetricRepository) CreateMetricService {
	return CreateMetricService{
		metricRepository: metricRepository,
	}
}

// CreateMetric implements the creatingmetric.CreateMetricService interface.
func (s CreateMetricService) CreateMetric(ctx context.Context, url string, responseStatus, requestSize, responseSize int, responseTime int64) error {
	uuid, err := uuid.NewRandom()
	metric, err := metrics.NewMetric(uuid.String(), url, responseStatus, requestSize, responseSize, responseTime)
	if err != nil {
		return err
	}
	return s.metricRepository.Save(ctx, metric)
}

