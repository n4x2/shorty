package controller

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/n4x2/shorty/internal/util"
	"github.com/n4x2/shorty/pkg/domain/model"
	"github.com/n4x2/shorty/pkg/usecase/usecase"
)

// urlController represents the controller implementation for URLs.
type urlController struct {
	urlUsecase usecase.Url
}

// Url defines the interface for URL controllers.
type Url interface {
	GetUrls(c *gin.Context) error
	CreateUrl(c *gin.Context) error
	FindByShortUrl(c *gin.Context) error
}

// NewUrlController creates a new URL controller with the given use case.
func NewUrlController(us usecase.Url) Url {
	return &urlController{us}
}

// GetUrls retrieves all URLs and returns them as a JSON response.
func (uc *urlController) GetUrls(c *gin.Context) error {
	var u []*model.Url

	u, err := uc.urlUsecase.List(u)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"urls": u})
	return nil
}

// CreateUrl creates a new URL based on the provided parameters and returns the created URL as a JSON response.
func (uc *urlController) CreateUrl(c *gin.Context) error {
	var params model.Url

	if err := c.ShouldBindJSON(&params); err != nil {
		return err
	}

	params.ShortUrl = util.ShortenURL(params.LongUrl)

	_, err := uc.urlUsecase.Create(&params)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"shortUrl": os.Getenv("baseURL") + "/" + params.ShortUrl})
	return nil
}

func (uc *urlController) FindByShortUrl(c *gin.Context) error {
	shortURL := c.Param("shortURL")

	u, err := uc.urlUsecase.FindByShortUrl(shortURL)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "URL not found"})
		return err
	}

	c.Redirect(http.StatusMovedPermanently, u.LongUrl)
	return nil
}
