package postgres

import "fmt"

type PersistenceError struct {
	Message string `json:"message"`
}

func NewPreparationErr(queryName string, repository string, err error) *PersistenceError {
	preparationErr := fmt.Errorf(
		"unable to prepare the query:`%s` on %s repository, original err: %s",
		queryName,
		repository,
		err.Error(),
	)

	return newPersistenceError(preparationErr, "prepare", "postgres")
}

func NewStatementNotPreparedErr(queryName string, repository string) *PersistenceError {
	preparationErr := fmt.Errorf("query `%s` is not prepared on %s repository", queryName, repository)
	return newPersistenceError(preparationErr, "statement not prepared", "postgres")
}

func (e *PersistenceError) Error() string {
	return e.Message
}

func newPersistenceError(originalErr error, action, datasource string) *PersistenceError {
	return &PersistenceError{
		Message: fmt.Sprintf("%s persistence error on `%s`: %s", datasource, action, originalErr.Error()),
	}
}
