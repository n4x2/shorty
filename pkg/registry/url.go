package registry

import (
	"github.com/n4x2/shorty/pkg/adapter/controller"
	"github.com/n4x2/shorty/pkg/adapter/repository"
	"github.com/n4x2/shorty/pkg/usecase/usecase"
)

// NewUrlController creates a new instance of the URL controller using the registry.
func (r *registry) NewUrlController() controller.Url {
	// Create instances of the URL use case, repository, and DB repository.
	u := usecase.NewUrlUsecase(
		repository.NewUrlRepository(r.db),
		repository.NewDBRepository(r.db),
	)

	// Create a new URL controller with the URL use case.
	return controller.NewUrlController(u)
}
