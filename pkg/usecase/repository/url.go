package repository

import "github.com/n4x2/shorty/pkg/domain/model"

// UrlRepository defines the interface for URL repositories.
type UrlRepository interface {
	// FindAll retrieves all URLs and populates the provided slice.
	FindAll(u []*model.Url) ([]*model.Url, error)

	// Create creates a new URL record and returns the created URL.
	Create(u *model.Url) (*model.Url, error)
}
