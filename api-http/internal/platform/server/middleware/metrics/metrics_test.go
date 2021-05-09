package metrics

import (
	"context"
	creatingmetric "example.com/gotraining/go-hexagonal_http_api-course/internal/metrics/creating"
	"example.com/gotraining/go-hexagonal_http_api-course/internal/platform/bus/inmemory"
	"example.com/gotraining/go-hexagonal_http_api-course/internal/platform/storage/metrics/storagemocks"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMiddleware(t *testing.T) {
	bus := inmemory.NewCommandBus()
	metricRepository := storagemocks.MetricRepository{}
	createMetricService := creatingmetric.NewCreateMetricService(&metricRepository)
	createMetricQueryHandler := creatingmetric.NewMetricCommandHandler(createMetricService)
	bus.RegisterCommandHandler(creatingmetric.MetricCommandType, createMetricQueryHandler)

	metricRepository.On(
		"Save",
		mock.Anything,
	).Return(nil)
	// Setting up the Gin server
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	engine.Use(Middleware(bus))

	// Setting up the HTTP recorder and the request
	httpRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/test-middleware", nil)
	require.NoError(t, err)

	// Performing the request
	engine.ServeHTTP(httpRecorder, req)

	response, err := metricRepository.GetAll(context.Background())
	// Asserting the output contains some expected values
	assert.Equal(t, 1, len(response))
}
