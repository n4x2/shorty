package registry

import (
	"github.com/n4x2/shorty/pkg/adapter/controller"
	"gorm.io/gorm"
)

// registry represents the registry implementation.
type registry struct {
	db *gorm.DB
}

// Registry defines the interface for the registry.
type Registry interface {
	NewAppController() controller.AppController
}

// NewRegistry creates a new registry with the provided database connection.
func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

// NewAppController creates a new instance of the app controller using the registry.
func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Url: r.NewUrlController(),
	}
}
