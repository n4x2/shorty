package main

import (
	"github.com/gin-gonic/gin"
	"github.com/n4x2/shorty/pkg/infrastructure/datastore"
	"github.com/n4x2/shorty/pkg/infrastructure/router"
	"github.com/n4x2/shorty/pkg/registry"
)

func main() {
	// Create a new database connection
	db := datastore.NewDB()

	// Create a new registry with the database connection
	r := registry.NewRegistry(db)

	// Create a new Gin router
	g := gin.New()

	// Initialize the router with the registry's app controller
	g = router.NewRouter(g, r.NewAppController())

	// Start the server on port
	g.Run(":8081")
}
