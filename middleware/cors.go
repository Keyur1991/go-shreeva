package middleware

import (
	"github.com/Keyur1991/go-shreeva/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors(conf *config.Cors) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     conf.AllowedOrigins,
		AllowCredentials: conf.AllowCredentials,
		AllowHeaders:     conf.AllowedHeaders,
		ExposeHeaders:    conf.ExposedHeaders,
		AllowMethods:     conf.AllowedMethods,
	})
}
