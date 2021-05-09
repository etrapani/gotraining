package fetchingcourse

import (
	"context"
	"example.com/gotraining/go-hexagonal_http_api-course/internal/courses"
)

// FetchingService is the default FetchingService interface
// implementation returned by fetching.NewCourseFetchingService.
type FetchingService struct {
	courseRepository courses.CourseRepository
}

// NewCourseFetchingService returns the default Service interface implementation.
func NewCourseFetchingService(courseRepository courses.CourseRepository) FetchingService {
	return FetchingService{
		courseRepository: courseRepository,
	}
}

// GetAll fetching all courses.
func (s FetchingService) GetAll(ctx context.Context) ([]courses.Course, error) {
	return s.courseRepository.GetAll(ctx)
}
