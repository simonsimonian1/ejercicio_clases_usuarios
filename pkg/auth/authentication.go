package auth

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/simonsimonian1/modulo7_GoWeb/clase3/C3-tm-estructuraDeCapas/EjerciciosPersonales/grupo4_clase7_tm/pkg/web"
)

func respondWithError(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, web.NewResponse(code, nil, message))
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Para continuar se necesita setear un token como variable de entorno")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			respondWithError(c, http.StatusUnauthorized, "API Token requerido")
			return
		}
		if token != requiredToken {
			respondWithError(c, http.StatusUnauthorized, "Token inválido")
			return
		}

		c.Next()
	}
}
