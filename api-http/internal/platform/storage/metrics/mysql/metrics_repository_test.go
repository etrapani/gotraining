package mysql

import (
	"context"
	"errors"
	"example.com/gotraining/go-hexagonal_http_api-course/internal/metrics"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_MetricRepository_Save_RepositoryError(t *testing.T) {
	metricID, metricUrl, metricResponseStatus, metricRequestSize, metricResponseSize, metricResponseTime := "37a0f027-15e6-47cc-a5d2-64183281087e", "/prueba/test", 500, 50, 100, int64(16)
	metric, err := metrics.NewMetric(metricID, metricUrl, metricResponseStatus, metricRequestSize, metricResponseSize, metricResponseTime)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO metrics (id, url, response_status, request_size, response_size, response_time) VALUES (?, ?, ?, ?, ?, ?)").
		WithArgs(metricID, metricUrl, metricResponseStatus, metricRequestSize, metricResponseSize, metricResponseTime).
		WillReturnError(errors.New("something-failed"))

	repo := NewMetricRepository(db)

	err = repo.Save(context.Background(), metric)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_MetricRepository_Save_Succeed(t *testing.T) {
	metricID, metricUrl, metricResponseStatus, metricRequestSize, metricResponseSize, metricResponseTime := "37a0f027-15e6-47cc-a5d2-64183281087e", "/prueba/test", 500, 50, 100, int64(16)
	metric, err := metrics.NewMetric(metricID, metricUrl, metricResponseStatus, metricRequestSize, metricResponseSize, metricResponseTime)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO metrics (id, url, response_status, request_size, response_size, response_time) VALUES (?, ?, ?, ?, ?, ?)").
		WithArgs(metricID, metricUrl, metricResponseStatus, metricRequestSize, metricResponseSize, metricResponseTime).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewMetricRepository(db)

	err = repo.Save(context.Background(), metric)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}
