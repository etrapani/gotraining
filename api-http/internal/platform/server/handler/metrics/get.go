package metrics

import (
	"example.com/gotraining/go-hexagonal_http_api-course/internal/metrics"
	fetchingmetric "example.com/gotraining/go-hexagonal_http_api-course/internal/metrics/fetching"
	"example.com/gotraining/go-hexagonal_http_api-course/kit/bus"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getResponse struct {
	Id             string `json:"id"`
	Url            string `json:"url"`
	ResponseStatus int `json:"responseStatus"`
	RequestSize    int    `json:"requestSize"`
	ResponseSize   int    `json:"responseSize"`
	ResponseTime   int64    `json:"responseTime"`
}

// GetHandler returns an HTTP handler for courses.
func GetHandler(queryBus bus.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var queryResponse, err = queryBus.DispatchQuery(ctx, fetchingmetric.NewFetchMetricQuery())

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, []getResponse{})
			return
		}
		metrics, ok := queryResponse.([]metrics.Metric)
		if ok {
			var response = make([]getResponse, 0, len(metrics))
			for _, metric := range metrics {
				response = append(response, getResponse{
					Id:             metric.ID().String(),
					Url:            metric.URL(),
					ResponseStatus: metric.ResponseStatus(),
					RequestSize:    metric.RequestSize(),
					ResponseSize:   metric.ResponseSize(),
					ResponseTime:   metric.ResponseTime(),
				})
			}
			ctx.JSON(http.StatusOK, response)
		}
		ctx.JSON(http.StatusInternalServerError, []getResponse{})
	}
}
