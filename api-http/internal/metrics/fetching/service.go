package fetchingmetric

import (
	"context"
	"example.com/gotraining/go-hexagonal_http_api-course/internal/metrics"
)

// FetchingService is the default FetchingService interface
// implementation returned by fetching.NewCourseFetchingService.
type FetchingService struct {
	metricRepository metrics.MetricRepository
}

// NewCourseFetchingService returns the default Service interface implementation.
func NewCourseFetchingService(metricRepository metrics.MetricRepository) FetchingService {
	return FetchingService{
		metricRepository: metricRepository,
	}
}

// GetAll fetching all courses.
func (s FetchingService) GetAll(ctx context.Context) ([]metrics.Metric, error) {
	return s.metricRepository.GetAll(ctx)
}
