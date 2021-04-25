package bootstrap

import (
	"context"
	"database/sql"
	"example.com/gotraining/go-hexagonal_http_api-course/internal/creating"
	"example.com/gotraining/go-hexagonal_http_api-course/internal/fetching"
	"example.com/gotraining/go-hexagonal_http_api-course/internal/platform/bus/inmemory"
	"example.com/gotraining/go-hexagonal_http_api-course/internal/platform/server"
	"example.com/gotraining/go-hexagonal_http_api-course/internal/platform/storage/mysql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	host = "localhost"
	port = 8080
	shutdownTimeout = 10 * time.Second

	dbUser = "codely"
	dbPass = "codely"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "codely"
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

	courseRepository := mysql.NewCourseRepository(db)

	creatingCourseService := creating.NewCourseService(courseRepository)
	fetchingCourseService := fetching.NewCourseFetchingService(courseRepository)

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)
	fetchingCourseQueryHandler := fetching.NewCourseQueryHandler(fetchingCourseService)

	bus.RegisterCommandHandler(creating.CourseCommandType, createCourseCommandHandler)
	bus.RegisterQueryHandler(fetching.CourseQueryType, fetchingCourseQueryHandler)

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, bus)
	return srv.Run(ctx)
}
