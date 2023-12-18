package docs

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SwaggerConfig(engine *gin.Engine) {
	SwaggerInfo.Title = "project000-user-backend"
	SwaggerInfo.BasePath = ""
	SwaggerInfo.Schemes = []string{"http", "https"}

	engine.GET("/user-swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler,
			ginSwagger.PersistAuthorization(true),
		),
	)
}
