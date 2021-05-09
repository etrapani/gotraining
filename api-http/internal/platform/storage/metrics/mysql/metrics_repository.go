package mysql

import (
	"context"
	"database/sql"
	"example.com/gotraining/go-hexagonal_http_api-course/internal/metrics"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
)

// MetricRepository is a MySQL mooc.CourseRepository implementation.
type MetricRepository struct {
	db *sql.DB
}

// NewMetricRepository initializes a MySQL-based implementation of mooc.CourseRepository.
func NewMetricRepository(db *sql.DB) *MetricRepository {
	return &MetricRepository{
		db: db,
	}
}

// Save implements the mooc.CourseRepository interface.
func (r *MetricRepository) Save(ctx context.Context, metric metrics.Metric) error {
	metricSQLStruct := sqlbuilder.NewStruct(new(sqlMetric))
	query, args := metricSQLStruct.InsertInto(sqlMetricTable, sqlMetric{
		ID:             metric.ID().String(),
		Url:            metric.URL(),
		ResponseStatus: metric.ResponseStatus(),
		RequestSize:    metric.RequestSize(),
		ResponseSize:   metric.ResponseSize(),
		ResponseTime:   metric.ResponseTime(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	return nil
}

// GetAll implements the mooc.CourseRepository interface.
func (r *MetricRepository) GetAll(ctx context.Context) (response []metrics.Metric, err error) {
	metricSQLStruct := sqlbuilder.NewSelectBuilder()
	metricSQLStruct.Select("id", "url", "response_status", "request_size", "response_size", "response_time")
	metricSQLStruct.From(sqlMetricTable)

	sqlQuery, args := metricSQLStruct.Build()

	rows, err := r.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("error trying to get course on database: %v", err)
	}
	defer rows.Close()
	response = []metrics.Metric{}
	for rows.Next() {
		var sqlMetric sqlMetric
		err := rows.Scan(sqlMetric.ID, sqlMetric.Url, sqlMetric.ResponseStatus, sqlMetric.RequestSize, sqlMetric.ResponseSize, sqlMetric.ResponseTime)
		if err != nil {
			return nil, err
		}
		metric, err := metrics.NewMetric(sqlMetric.ID, sqlMetric.Url, sqlMetric.ResponseStatus, sqlMetric.RequestSize, sqlMetric.ResponseSize, sqlMetric.ResponseTime)
		if err != nil {
			return nil, err
		}
		response = append(response, metric)
	}
	return response, nil
}
