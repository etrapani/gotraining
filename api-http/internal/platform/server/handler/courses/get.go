package courses

import (
	courses2 "example.com/gotraining/go-hexagonal_http_api-course/internal/courses"
	fetchingcourse "example.com/gotraining/go-hexagonal_http_api-course/internal/courses/fetching"
	"example.com/gotraining/go-hexagonal_http_api-course/kit/bus"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Duration string `json:"duration"`
}

// GetHandler returns an HTTP handler for courses.
func GetHandler(queryBus bus.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var queryResponse, err = queryBus.DispatchQuery(ctx, fetchingcourse.NewFetchCourseQuery())

		if err != nil {
			// Si quiero devolver error en ves de la lista se rompe me genera un error de unmarshal
			ctx.JSON(http.StatusInternalServerError, []getResponse{})
			return
		}
		courses, ok := queryResponse.([]courses2.Course)
		if ok {
			var response = make([]getResponse, 0, len(courses))
			for _, course := range courses {
				response = append(response, getResponse{
					Id:       course.ID().String(),
					Name:     course.Name().String(),
					Duration: course.Duration().String(),
				})
			}
			ctx.JSON(http.StatusOK, response)
		}
		ctx.JSON(http.StatusInternalServerError, []getResponse{})
	}
}
