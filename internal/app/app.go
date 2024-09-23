package app

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/prajnasatryass/tic-be/config"
	authV1 "github.com/prajnasatryass/tic-be/internal/auth/delivery/http/v1"
	authRepository "github.com/prajnasatryass/tic-be/internal/auth/repository"
	authUsecase "github.com/prajnasatryass/tic-be/internal/auth/usecase"
	"github.com/prajnasatryass/tic-be/internal/database"
	"github.com/prajnasatryass/tic-be/internal/middleware"
	userV1 "github.com/prajnasatryass/tic-be/internal/user/delivery/http/v1"
	userRepository "github.com/prajnasatryass/tic-be/internal/user/repository"
	userUsecase "github.com/prajnasatryass/tic-be/internal/user/usecase"
	"github.com/samber/lo"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	cfg  config.Config
	echo *echo.Echo
	db   *sqlx.DB
}

func NewApp(cfg config.Config) *App {
	return &App{
		cfg:  cfg,
		echo: middleware.NewEcho(cfg),
		db:   lo.Must(database.NewDatabase(database.DriverPostgres, cfg.Database)),
	}
}

func (app *App) init() {
	app.echo.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	app.echo.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	app.echo.GET("/protected", func(c echo.Context) error {
		return c.String(http.StatusOK, "authenticated")
	}, middleware.JWTAuth(app.cfg))

	egAPI := app.echo.Group("/api")
	egAPIV1 := egAPI.Group("/v1")

	authRepo := authRepository.NewAuthRepository(app.db)
	userRepo := userRepository.NewUserRepository(app.db)

	authUC := authUsecase.NewAuthUsecase(authRepo, userRepo, app.cfg.JWT)
	authV1.NewAuthController(egAPIV1, authUC)

	userUC := userUsecase.NewUserUsecase(userRepo)
	userV1.NewUserController(egAPIV1, userUC)
}

func (app *App) Run() error {
	app.init()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-quit
		log.Info("Shutting down gracefully ✨")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		app.db.Close()
		app.echo.Shutdown(ctx)
	}()

	return app.echo.Start(":" + app.cfg.Server.Port)
}
