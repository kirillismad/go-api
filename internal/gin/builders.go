package gin

import (
	"go-api/internal/apps/entities"
	"log/slog"
	"os"

	"github.com/caarlos0/env/v10"
	g "github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	App struct {
		Debug bool `env:"DEBUG" envDefault:"false"`
	} `envPrefix:"APP_"`
	Http struct {
		Host         string `env:"HOST,notEmpty"`
		Port         int    `env:"PORT,notEmpty"`
		ReadTimeout  int    `env:"READ_TIMEOUT" envDefault:"5"`
		WriteTimeout int    `env:"WRITE_TIMEOUT" envDefault:"5"`
		IdleTimeout  int    `env:"IDLE_TIMEOUT" envDefault:"20"`
	} `envPrefix:"HTTP_"`
	Db struct {
		ConnString string `env:"CONN_STRING,notEmpty"`
	} `envPrefix:"DB_"`
}

func BuildLogger(config Config) (*slog.Logger, error) {
	if config.App.Debug {
		return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})), nil
	}
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})), nil
}

func BuildDb(config Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.Db.ConnString), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Build(config Config, logger *slog.Logger, db *gorm.DB) (*g.Engine, error) {
	if config.App.Debug {
		g.SetMode(g.DebugMode)
	} else {
		g.SetMode(g.ReleaseMode)
	}

	// providers
	entityService := entities.NewEntityService(db)

	// router config
	router := g.New()
	router.Use(HandleRequestID())
	router.Use(Logging(logger))
	router.Use(HandleErrors())
	router.Use(g.Recovery())

	router.GET("/ping", func(ctx *g.Context) {
		ctx.Status(200)
	})
	router.POST("/entities", CreateEntityHandler(entityService))

	return router, nil
}

func GetConfig() (Config, error) {
	var config Config
	if err := env.Parse(&config); err != nil {
		return Config{}, err
	}
	return config, nil
}
