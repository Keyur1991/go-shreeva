package request

import (
	mw "github.com/Keyur1991/go-shreeva/middleware"

	"github.com/gin-gonic/gin"
)

func Router(config *Config) *gin.Engine {
	r := gin.Default()

	r.Use(mw.Cors(config.Cors))
	// r := mux.NewRouter()

	// r.Use(mw.Cors(config.Cors))
	// r.Use(mw.RateLimit(config.RateLimit))
	// r.Use(mw.Logger)

	return r
}
