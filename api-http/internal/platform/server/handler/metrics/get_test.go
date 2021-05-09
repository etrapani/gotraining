package metrics

import (
	"encoding/json"
	"errors"
	"example.com/gotraining/go-hexagonal_http_api-course/internal/metrics"
	fetchingmetric "example.com/gotraining/go-hexagonal_http_api-course/internal/metrics/fetching"
	"example.com/gotraining/go-hexagonal_http_api-course/internal/platform/bus/inmemory"
	"example.com/gotraining/go-hexagonal_http_api-course/internal/platform/storage/metrics/storagemocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHandler(t *testing.T) {
	tests := map[string]struct {
		mockData []metrics.Metric
		mockError error
		expectedResponse []getResponse
		expectedStatus int
	}{
		"Empty repo return 200 with empty list": {
			mockData: []metrics.Metric{},
			mockError: nil,
			expectedStatus:http.StatusOK,
			expectedResponse: []getResponse{}},
		"Repo error return 500": {
			mockData: []metrics.Metric{},
			mockError: errors.New("the field Duration can not be empty"),
			expectedStatus:http.StatusInternalServerError,
			expectedResponse: []getResponse{}},
		"Fully repo return 200 with list of courses":{
			mockData: []metrics.Metric{mockMetric("8a1c5cdc-ba57-445a-994d-aa412d23723f", "/prueba/test", 500, 50,75, 100)},
			mockError: nil,
			expectedStatus:http.StatusOK,
			expectedResponse: []getResponse{{Id: "8a1c5cdc-ba57-445a-994d-aa412d23723f", Url: "/prueba/test", ResponseStatus: "500", RequestSize: 50, ResponseSize: 75, ResponseTime: 100}}},
	}
	for key, value := range tests {
		t.Run(key, func(t *testing.T) {
			metricRepository := storagemocks.MetricRepository{}
			bus := inmemory.NewCommandBus()
			fetchingMetricService := fetchingmetric.NewCourseFetchingService(&metricRepository)
			fetchingMetricQueryHandler := fetchingmetric.NewMetricQueryHandler(fetchingMetricService)
			bus.RegisterQueryHandler(fetchingmetric.MetricQueryType, fetchingMetricQueryHandler)

			metricRepository.On(
				"GetAll",
				mock.Anything,
			).Return(value.mockData, value.mockError)
			gin.SetMode(gin.TestMode)
			r := gin.New()
			r.GET("/metrics", GetHandler(bus))

			req, err := http.NewRequest(http.MethodGet, "/metrics", nil)
			require.NoError(t, err)

			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			assert.Equal(t, value.expectedStatus, res.StatusCode)
			var response []getResponse
			if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
				log.Fatalln(err)
			}

			assert.Equal(t, value.expectedResponse, response)
		})
	}
}

func mockMetric(id, url string, responseStatus, requestSize, responseSize int, responseTime int64) metrics.Metric {
	metric, err := metrics.NewMetric(id, url, responseStatus, requestSize, responseSize, responseTime)
	if err != nil{
		log.Fatalln(err)
	}
	return metric
}
