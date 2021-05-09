package bootstrap

import (
	"context"
	"database/sql"
	creatingcourse "example.com/gotraining/go-hexagonal_http_api-course/internal/courses/creating"
	fetchingcourse "example.com/gotraining/go-hexagonal_http_api-course/internal/courses/fetching"
	creatingmetric "example.com/gotraining/go-hexagonal_http_api-course/internal/metrics/creating"
	"example.com/gotraining/go-hexagonal_http_api-course/internal/platform/bus/inmemory"
	"example.com/gotraining/go-hexagonal_http_api-course/internal/platform/server"
	coursemysql "example.com/gotraining/go-hexagonal_http_api-course/internal/platform/storage/courses/mysql"
	metricmysql "example.com/gotraining/go-hexagonal_http_api-course/internal/platform/storage/metrics/mysql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	host = "localhost"
	port = 8080
	shutdownTimeout = 10 * time.Second

	dbUser = "root"
	dbPass = "password"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "codelytv"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	var (
		bus = inmemory.NewCommandBus()
	)

	courseRepository := coursemysql.NewCourseRepository(db)
	metricRepository := metricmysql.NewMetricRepository(db)

	creatingCourseService := creatingcourse.NewCourseService(courseRepository)
	fetchingCourseService := fetchingcourse.NewCourseFetchingService(courseRepository)
	creatingMetricService := creatingmetric.NewCreateMetricService(metricRepository)

	createCourseCommandHandler := creatingcourse.NewCourseCommandHandler(creatingCourseService)
	fetchingCourseQueryHandler := fetchingcourse.NewCourseQueryHandler(fetchingCourseService)
	creatingMetricQueryHandler := creatingmetric.NewMetricCommandHandler(creatingMetricService)

	bus.RegisterCommandHandler(creatingcourse.CourseCommandType, createCourseCommandHandler)
	bus.RegisterQueryHandler(fetchingcourse.CourseQueryType, fetchingCourseQueryHandler)
	bus.RegisterCommandHandler(creatingmetric.MetricCommandType, creatingMetricQueryHandler)


	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, bus)
	return srv.Run(ctx)
}
