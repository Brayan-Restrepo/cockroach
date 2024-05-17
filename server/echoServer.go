package server

import (
	cockroachHandlers "cockroach/cockroach/handlers"
	"cockroach/cockroach/repositories"
	cockroachUsecases "cockroach/cockroach/usecases"
	"fmt"

	"cockroach/config"
	"cockroach/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type echoServer struct {
	app  *echo.Echo
	db   database.Database
	conf *config.Config
}

func NewEchoServer(conf *config.Config, db database.Database) Server {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	return &echoServer{
		app:  echoApp,
		db:   db,
		conf: conf,
	}
}

func (s *echoServer) Start() {
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())

	s.app.GET("v1/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	s.initializeCockroachHttpHandler()

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (s *echoServer) initializeCockroachHttpHandler() {
	// Initialize all layers
	cockroachPostgresRepository := repositories.NewCockroachPostgresRepository(s.db)
	cockroachFCMMessaging := repositories.NewCockroachFCMMessaging()

	cockroachUsecase := cockroachUsecases.NewCockroachUsecaseImpl(
		cockroachPostgresRepository,
		cockroachFCMMessaging,
	)

	cockroachHttpHandler := cockroachHandlers.NewCockroachHttpHandler(cockroachUsecase)

	// Routers
	cockroachRouters := s.app.Group("v1/cockroach")
	cockroachRouters.POST("", cockroachHttpHandler.DetectCockroach)
	// Routers Animals
	animalRouters := s.app.Group("v1/animal")
	animalRouters.POST("", cockroachHttpHandler.DetectCockroach)
}
