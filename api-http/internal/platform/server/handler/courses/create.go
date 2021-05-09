package courses

import (
	"errors"
	"example.com/gotraining/go-hexagonal_http_api-course/internal/courses"
	creatingcourse "example.com/gotraining/go-hexagonal_http_api-course/internal/courses/creating"
	"example.com/gotraining/go-hexagonal_http_api-course/kit/bus"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation.
func CreateHandler(bus bus.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := bus.DispatchCommand(ctx, creatingcourse.NewCourseCommand(
			req.ID,
			req.Name,
			req.Duration,
		))

		if err != nil {
			switch {
			case errors.Is(err, courses.ErrInvalidCourseID),
				errors.Is(err, courses.ErrEmptyCourseName), errors.Is(err, courses.ErrInvalidCourseID):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
