package app

import (
	"context"
	"start/internal/config"
	repository_event "start/internal/repository/event"
	repository_user "start/internal/repository/user"
	service_event "start/internal/service/event"
	service_user "start/internal/service/user"
	usecase_event "start/internal/usecase/event"
	usecase_user "start/internal/usecase/user"
	"start/pkg/postgres"
)

func init() {
	// Initialize User Service
	ctx := context.Background() // Assuming you have a context available
	// Load configuration and initialize database connection
	conf, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}
	postgresPool, err := postgres.NewPool(ctx, conf.DatabaseURL)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	userRepo := repository_user.New(postgresPool) // Assuming a NewUserRepository function exists
	userUsecase := usecase_user.New(userRepo)
	userService := service_user.New(userUsecase)

	// Initialize Event Service
	eventRepo := repository_event.New() // Assuming a NewEventRepository function exists
	eventUsecase := usecase_event.New(eventRepo)
	eventService := service_event.New(eventUsecase)

	// Register services in the application context or DI container
	RegisterUserService(userService)
	RegisterEventService(eventService)
}
0