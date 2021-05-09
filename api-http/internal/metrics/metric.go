package metrics

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
)


var ErrInvalidMetricID = errors.New("invalid Metric ID")

// MetricID represents the course unique identifier.
type MetricID struct {
	value string
}

// NewMetricID instantiate the VO for MetricID
func NewMetricID(value string) (MetricID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return MetricID{}, fmt.Errorf("%w: %s", ErrInvalidMetricID, value)
	}

	return MetricID{
		value: v.String(),
	}, nil
}

// String type converts the CourseID into string.
func (id MetricID) String() string {
	return id.value
}


var ErrInvalidUrl = errors.New("url mandatory")

type Metric struct {
	id MetricID
	url string
	responseStatus int
	requestSize int
	responseSize int
	responseTime int64
}

// MetricRepository defines the expected behaviour from a course storage.
type MetricRepository interface {
	Save(ctx context.Context, metric Metric) error
	GetAll(ctx context.Context) ([]Metric, error)
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=../platform/storage/metrics/storagemocks --name=MetricRepository

// NewMetric creates a new metric.
func NewMetric(id, url string, responseStatus, requestSize, responseSize int, responseTime int64) (Metric, error) {
	idVO, err := NewMetricID(id)
	if err != nil {
		return Metric{}, err
	}

	if url == "" {
		return Metric{}, ErrInvalidUrl
	}

	return Metric{
		id: 			idVO,
		url:       		url,
		responseStatus: responseStatus,
		requestSize:    requestSize,
		responseSize: 	responseSize,
		responseTime: 	responseTime,
	}, nil
}


// ID returns the metric unique identifier.
func (m Metric) ID() MetricID {
	return m.id
}

// URL returns the request url.
func (m Metric) URL() string {
	return m.url
}

// ResponseStatus returns the responnse status.
func (m Metric) ResponseStatus() int {
	return m.responseStatus
}

// RequestSize returns the request size.
func (m Metric) RequestSize() int {
	return m.requestSize
}

// ResponseSize returns the response size.
func (m Metric) ResponseSize() int {
	return m.responseSize
}

// ResponseTime returns the response time.
func (m Metric) ResponseTime() int64 {
	return m.responseTime
}
