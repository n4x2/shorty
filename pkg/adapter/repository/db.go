package repository

import (
	"log"

	"github.com/n4x2/shorty/pkg/usecase/repository"
	"gorm.io/gorm"
)

// dbRepository represents the repository implementation for database operations.
type dbRepository struct {
	db *gorm.DB
}

// NewDBRepository creates a new database repository with the given database connection.
func NewDBRepository(db *gorm.DB) repository.DBRepository {
	return &dbRepository{db}
}

// Transaction starts a new transaction and executes the provided function within the transaction.
// It handles transaction rollback in case of error and commits the transaction if there are no errors.
func (r *dbRepository) Transaction(txFunc func(interface{}) (interface{}, error)) (data interface{}, err error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if p := recover(); p != nil {
			log.Print("recover")
			tx.Rollback()
			panic(p)
		} else if err != nil {
			log.Print("rollback")
			tx.Rollback()
			panic("error")
		} else {
			err = tx.Commit().Error
		}
	}()

	data, err = txFunc(tx)
	return data, err
}
