package usecase

import (
	"errors"

	"github.com/n4x2/shorty/pkg/domain/model"
	"github.com/n4x2/shorty/pkg/usecase/repository"
)

// urlUsecase represents the use case implementation for URLs.
type urlUsecase struct {
	urlRepository repository.UrlRepository
	dbRepository  repository.DBRepository
}

// Url defines the interface for URL use cases.
type Url interface {
	List(u []*model.Url) ([]*model.Url, error)
	Create(u *model.Url) (*model.Url, error)
	FindByShortUrl(shortURL string) (*model.Url, error)
}

// NewUrlUsecase creates a new URL use case with the given repositories.
func NewUrlUsecase(r repository.UrlRepository, d repository.DBRepository) Url {
	return &urlUsecase{r, d}
}

// List retrieves all URLs and populates the provided slice.
func (uu *urlUsecase) List(u []*model.Url) ([]*model.Url, error) {
	u, err := uu.urlRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// Create creates a new URL record and returns the created URL.
func (uu *urlUsecase) Create(u *model.Url) (*model.Url, error) {
	data, err := uu.dbRepository.Transaction(func(i interface{}) (interface{}, error) {
		u, err := uu.urlRepository.Create(u)

		return u, err
	})

	url, ok := data.(*model.Url)
	if !ok {
		return nil, errors.New("cast error")
	}

	if err != nil {
		return nil, err
	}

	return url, nil
}

func (uu *urlUsecase) FindByShortUrl(shortUrl string) (*model.Url, error) {
	url, err := uu.urlRepository.FindByShortUrl(shortUrl)
	if err != nil {
		return nil, err
	}

	return url, nil
}
