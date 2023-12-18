package main

import (
	"os"

	"project000-backend-user/config"
	"project000-backend-user/internal/app"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	env := os.Getenv("env")
	if env == "" {
		env = "local"
	}

	cfg, err := config.LoadConfig(env)
	if err != nil {
		panic(err)
	}

	userApp, err := app.NewApp(cfg)
	if err != nil {
		panic(err)
	}

	if err = userApp.Run(); err != nil {
		panic(err)
	}
}
