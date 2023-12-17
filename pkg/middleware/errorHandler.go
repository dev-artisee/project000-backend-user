package middleware

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"project000-backend-user/pkg/cerror"

	"project000-backend-user/pkg/response"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Errors) == 0 {
			return
		}

		// custom error
		for _, err := range c.Errors {
			var customError *cerror.CustomError
			if errors.As(err.Err, &customError) {
				// logging stack trace
				if customError.EnableLogging {
					log.Printf("%+v", customError.Err)
				}

				// TODO: sentry alert
				if customError.EnableAlert {
				}

				c.JSON(customError.Status, response.Fail(customError.Code, customError.Message))
				return
			}
		}

		// unexpected error
		for _, err := range c.Errors {
			// logging stack trace
			log.Printf("%+v", err.Err)
		}

		c.JSON(
			http.StatusInternalServerError,
			response.Fail(response.CodeInternalServerError, "Internal Server Error"),
		)
	}
}
