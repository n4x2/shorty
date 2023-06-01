package router

import (
	"github.com/gin-gonic/gin"
	"github.com/n4x2/shorty/pkg/adapter/controller"
)

// NewRouter creates a new router with the provided Gin engine and app controller.
func NewRouter(r *gin.Engine, c controller.AppController) *gin.Engine {
	// Use Gin middleware for logging and recovery.
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Define the URL routes and their corresponding controller methods.
	r.POST("/url", func(ctx *gin.Context) { c.Url.CreateUrl(ctx) })
	r.GET("/urls", func(ctx *gin.Context) { c.Url.GetUrls(ctx) })

	return r
}
