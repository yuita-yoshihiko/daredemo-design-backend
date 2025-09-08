package middlewares

import (
	"net/http"
	"strings"

	"github.com/go-chi/cors"
	"github.com/yuita-yoshihiko/daredemo-design-backend/config"
)

func NewCorsMiddleware() func(http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins: strings.Split(config.Conf.CORSAllowOrigins, ","),
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})
}
