package repository

// DBRepository defines the interface for a database repository.
type DBRepository interface {
	// Transaction starts a new transaction and executes the provided function within the transaction.
	// It returns the result of the function and any error encountered during the transaction.
	Transaction(func(interface{}) (interface{}, error)) (interface{}, error)
}
