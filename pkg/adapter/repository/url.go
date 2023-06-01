package repository

import (
	"github.com/n4x2/shorty/pkg/domain/model"
	"github.com/n4x2/shorty/pkg/usecase/repository"
	"gorm.io/gorm"
)

// urlRepository represents the repository implementation for URLs.
type urlRepository struct {
	db *gorm.DB
}

// NewUrlRepository creates a new URL repository with the given database connection.
func NewUrlRepository(db *gorm.DB) repository.UrlRepository {
	return &urlRepository{db}
}

// FindAll retrieves all URLs and populates the provided slice.
func (ur *urlRepository) FindAll(u []*model.Url) ([]*model.Url, error) {
	if err := ur.db.Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

// Create creates a new URL record.
func (ur *urlRepository) Create(u *model.Url) (*model.Url, error) {
	if err := ur.db.Create(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}
