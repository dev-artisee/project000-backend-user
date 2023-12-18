package app

import (
	userHandler "project000-backend-user/internal/handler/user"
	userRepo "project000-backend-user/internal/repository/user"
	userUseCase "project000-backend-user/internal/usecase/user"
)

func (app *App) StartService() {
	userRW := userRepo.NewRepository(app.DB)
	userUC := userUseCase.NewService(userRW)
	userCtrl := userHandler.NewController(userUC)

	userGroup := app.Gin.Group("/user")
	userHandler.Route(userGroup, userCtrl, app.Config)
}
