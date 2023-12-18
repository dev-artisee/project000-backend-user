package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"project000-backend-user/config"
	"project000-backend-user/pkg/datasource"
)

type App struct {
	Gin    *gin.Engine
	DB     *gorm.DB
	Config config.Config
}

func NewApp(cfg config.Config) (*App, error) {
	db, err := datasource.NewGORMDatabase(cfg)
	if err != nil {
		return nil, err
	}

	app := &App{
		Gin:    gin.Default(),
		DB:     db,
		Config: cfg,
	}

	return app, nil
}

func (app *App) Run() error {
	server := http.Server{
		Addr:         "",
		Handler:      app.Gin,
		ReadTimeout:  app.Config.Server.ReadTimeout,
		WriteTimeout: app.Config.Server.WriteTimeout,
	}

	return server.ListenAndServe()
}
